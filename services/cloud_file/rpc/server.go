package rpc

import (
	"context"
	"github.com/mats9693/unnamed_plan/services/cloud_file/config"
	"github.com/mats9693/unnamed_plan/services/cloud_file/db"
	"github.com/mats9693/unnamed_plan/services/shared/const"
	"github.com/mats9693/unnamed_plan/services/shared/db/model"
	"github.com/mats9693/unnamed_plan/services/shared/proto/client"
	"github.com/mats9693/unnamed_plan/services/shared/proto/impl"
	"github.com/mats9693/unnamed_plan/services/shared/utils"
	"os"
	"strconv"
	"time"
)

const backupFileSuffix = ".backup"

type cloudFileServerImpl struct {
	rpc_impl.UnimplementedICloudFileServer

	userClient rpc_impl.IUserClient
}

var _ rpc_impl.ICloudFileServer = (*cloudFileServerImpl)(nil)

var cloudFileServerImplIns = &cloudFileServerImpl{}

func GetCloudFileServer(userServerTarget string) (*cloudFileServerImpl, error) {
	userClient, err := client.ConnectUserServer(userServerTarget)
	if err != nil {
		return nil, err
	}

	cloudFileServerImplIns.userClient = userClient

	return cloudFileServerImplIns, nil
}

func (c *cloudFileServerImpl) ListByUploader(_ context.Context, req *rpc_impl.CloudFile_ListByUploaderReq) (*rpc_impl.CloudFile_ListByUploaderRes, error) {
	if len(req.OperatorId) < 1 || req.GetPage() == nil || req.GetPage().PageSize < 1 || req.GetPage().PageNum < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	pageSize := int(req.GetPage().PageSize)
	pageNum := int(req.GetPage().PageNum)

	files, count, err := db.GetCloudFileDao().QueryPageByUploader(pageSize, pageNum, req.OperatorId)
	if err != nil {
		return nil, err
	}

	return &rpc_impl.CloudFile_ListByUploaderRes{
		Total: uint32(count),
		Files: filesDBToRPC(files...),
	}, nil
}

func (c *cloudFileServerImpl) ListPublic(_ context.Context, req *rpc_impl.CloudFile_ListPublicReq) (*rpc_impl.CloudFile_ListPublicRes, error) {
	if len(req.OperatorId) < 1 || req.GetPage() == nil || req.GetPage().PageSize < 1 || req.GetPage().PageNum < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	pageSize := int(req.GetPage().PageSize)
	pageNum := int(req.GetPage().PageNum)

	files, count, err := db.GetCloudFileDao().QueryPageInPublic(pageSize, pageNum, req.OperatorId)
	if err != nil {
		return nil, err
	}

	return &rpc_impl.CloudFile_ListPublicRes{
		Total: uint32(count),
		Files: filesDBToRPC(files...),
	}, nil
}

func (c *cloudFileServerImpl) Upload(_ context.Context, req *rpc_impl.CloudFile_UploadReq) (*rpc_impl.CloudFile_UploadRes, error) {
	if len(req.OperatorId) < 1 || len(req.FileName) < 1 || len(req.ExtensionName) < 1 ||
		req.FileSize < 1 || req.LastModifiedTime < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	dir := spliceFileDir(req.IsPublic, req.OperatorId)
	// make sure private path exist, public path has checked in init
	if !req.IsPublic {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	// save file
	fileID := utils.CalcSHA256(req.OperatorId + time.Now().GoString())
	absolutePath := spliceFilePath(dir, fileID, req.ExtensionName)

	err := os.WriteFile(absolutePath, req.File, 0755)
	if err != nil {
		return nil, err
	}

	// save db, if failed, remove file
	err = db.GetCloudFileDao().Insert(&model.CloudFile{
		UploadedBy:       req.OperatorId,
		FileID:           fileID,
		FileName:         req.FileName,
		ExtensionName:    req.ExtensionName,
		LastModifiedTime: time.Duration(req.LastModifiedTime),
		FileSize:         req.FileSize,
		IsPublic:         req.IsPublic,
	})
	if err != nil {
		err2 := os.Remove(absolutePath)
		return nil, utils.NewError(utils.ErrorsToString(err, err2))
	}

	return &rpc_impl.CloudFile_UploadRes{}, nil
}

func (c *cloudFileServerImpl) Modify(ctx context.Context, req *rpc_impl.CloudFile_ModifyReq) (*rpc_impl.CloudFile_ModifyRes, error) {
	if len(req.OperatorId) < 1 || len(req.FileId) < 1 || len(req.Password) < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	_, err := c.userClient.Authenticate(ctx, &rpc_impl.User_AuthenticateReq{
		UserId:   req.OperatorId,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	fileRecord, err := db.GetCloudFileDao().QueryOne(req.FileId)
	if err != nil {
		return nil, err
	}

	if len(req.FileName)+len(req.ExtensionName) < 1 && fileRecord.IsPublic == req.IsPublic && len(req.File) < 1 {
		return nil, utils.NewError(mconst.Error_NoValidModification)
	}

	updateColumns := make([]string, 0, 5)

	// update file, backup old file in private folder
	if len(req.File) > 0 {
		// make sure private dir exist, only when old file is public
		if fileRecord.IsPublic {
			if err = os.MkdirAll(spliceFileDir(false, req.OperatorId), 0755); err != nil {
				return nil, err
			}
		}

		// backup old file, may implicitly move file from public path to private path
		oldPath := spliceFilePath(spliceFileDir(fileRecord.IsPublic, req.OperatorId), fileRecord.FileID, req.ExtensionName)
		privPath := spliceFilePath(spliceFileDir(false, req.OperatorId), fileRecord.FileID, req.ExtensionName)

		err = os.Rename(oldPath, getValidBackupFileName(privPath))
		if err != nil {
			return nil, err
		}

		// save new file
		absolutePath := spliceFilePath(spliceFileDir(req.IsPublic, req.OperatorId), fileRecord.FileID, req.ExtensionName)
		err = os.WriteFile(absolutePath, req.File, 0755)
		if err != nil {
			return nil, err
		}

		// update db record: file size and file last modified time
		fileRecord.FileSize = req.FileSize
		updateColumns = append(updateColumns, model.CloudFile_FileSize)

		fileRecord.LastModifiedTime = time.Duration(req.LastModifiedTime)
		updateColumns = append(updateColumns, model.CloudFile_LastModifiedTime)
	}

	if len(req.FileName) > 0 {
		fileRecord.FileName = req.FileName
		updateColumns = append(updateColumns, model.CloudFile_FileName)
	}
	if len(req.ExtensionName) > 0 && fileRecord.ExtensionName != req.ExtensionName {
		fileRecord.ExtensionName = req.ExtensionName
		updateColumns = append(updateColumns, model.CloudFile_ExtensionName)
	}
	if fileRecord.IsPublic != req.IsPublic {
		fileRecord.IsPublic = req.IsPublic
		updateColumns = append(updateColumns, model.CloudFile_IsPublic)
	}

	// update db record, if update failed and have new file, remove file
	err = db.GetCloudFileDao().UpdateColumnsByFileID(fileRecord, updateColumns...)
	if err != nil {
		errMsg := err.Error()
		if len(req.File) > 0 {
			dir := spliceFileDir(req.IsPublic, req.OperatorId)
			absolutePath := spliceFilePath(dir, fileRecord.FileID, req.ExtensionName)
			errMsg += utils.ErrorsToString(os.Remove(absolutePath))
		}
		return nil, utils.NewError(errMsg)
	}

	return &rpc_impl.CloudFile_ModifyRes{}, nil
}

func (c *cloudFileServerImpl) Delete(ctx context.Context, req *rpc_impl.CloudFile_DeleteReq) (*rpc_impl.CloudFile_DeleteRes, error) {
	if len(req.OperatorId) < 1 || len(req.Password) < 1 || len(req.FileId) < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	_, err := c.userClient.Authenticate(ctx, &rpc_impl.User_AuthenticateReq{
		UserId:   req.OperatorId,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	err = db.GetCloudFileDao().UpdateColumnsByFileID(&model.CloudFile{
		FileID:    req.FileId,
		IsDeleted: true,
	}, model.CloudFile_IsDeleted)
	if err != nil {
		return nil, err
	}

	return &rpc_impl.CloudFile_DeleteRes{}, nil
}

func getValidBackupFileName(absolutePath string) string {
	absolutePath += backupFileSuffix

	for i := 0; ; i++ {
		fileInfo, err := os.Stat(absolutePath + strconv.Itoa(i))
		if !(err == nil && !fileInfo.IsDir()) {
			absolutePath += strconv.Itoa(i)
			break
		}
	}

	return absolutePath
}

func filesDBToRPC(data ...*model.CloudFile) []*rpc_impl.CloudFile_Data {
	res := make([]*rpc_impl.CloudFile_Data, 0, len(data))
	for i := range data {
		res = append(res, &rpc_impl.CloudFile_Data{
			FileId:           data[i].FileID,
			FileName:         data[i].FileName,
			LastModifiedTime: int64(data[i].LastModifiedTime),
			FileUrl:          spliceFilePath(spliceFileURL(data[i].IsPublic, data[i].UploadedBy), data[i].FileID, data[i].ExtensionName),
			IsPublic:         data[i].IsPublic,
			UpdateTime:       int64(data[i].UpdateTime),
			CreatedTime:      int64(data[i].CreatedTime),
		})
	}

	return res
}

func spliceFileURL(isPublic bool, operatorID string) string {
	url := ""

	if isPublic {
		url += config.GetConfig().CloudFilePublicDir
	} else {
		url += operatorID
	}

	return utils.FormatDirSuffix(url)
}

func spliceFileDir(isPublic bool, operatorID string) string {
	dir := utils.FormatDirSuffix(config.GetConfig().CloudFileRootPath)
	dir += spliceFileURL(isPublic, operatorID)

	return utils.FormatDirSuffix(dir)
}

// spliceFilePath return absolute path
func spliceFilePath(dir string, fileName string, extensionName string) string {
	return dir + fileName + "." + extensionName
}
