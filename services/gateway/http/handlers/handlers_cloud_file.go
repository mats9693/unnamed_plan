package handlers

import (
    "context"
    "github.com/mats9693/unnamed_plan/services/gateway/http/structure_defination"
    "github.com/mats9693/unnamed_plan/services/gateway/rpc"
    "github.com/mats9693/unnamed_plan/services/shared/proto/impl"
    "github.com/mats9693/utils/toy_server/const"
    "github.com/mats9693/utils/toy_server/http"
    "github.com/mats9693/utils/toy_server/log"
    "go.uber.org/zap"
    "net/http"
    "time"
)

func ListCloudFileByUploader(r *http.Request) *mhttp.ResponseData {
    params := &structure.ListCloudFileByUploaderReqParams{}
    if errMsg := params.Decode(r); len(errMsg) > 0 {
        mlog.Logger().Error("parse request params failed", zap.String("err msg", errMsg))
        return mhttp.ResponseWithError(errMsg)
    }

    res, err := rpc.GetRPCClient().CloudFileClient.ListByUploader(context.Background(), &rpc_impl.CloudFile_ListByUploaderReq{
        OperatorId: params.OperatorID,
        Page: &rpc_impl.Pagination{
            PageSize: uint32(params.PageSize),
            PageNum:  uint32(params.PageNum),
        },
    })
    if err != nil {
        mlog.Logger().Error("list cloud file by uploader failed", zap.Error(err))
        return mhttp.ResponseWithError(err.Error())
    }

    return mhttp.Response(structure.MakeListCloudFileByUploaderRes(res.Total, filesRPCToHTTP(res.Files...)))
}

func ListPublicCloudFile(r *http.Request) *mhttp.ResponseData {
    params := &structure.ListPublicCloudFileReqParams{}
    if errMsg := params.Decode(r); len(errMsg) > 0 {
        mlog.Logger().Error("parse request params failed", zap.String("err msg", errMsg))
        return mhttp.ResponseWithError(errMsg)
    }

    res, err := rpc.GetRPCClient().CloudFileClient.ListPublic(context.Background(), &rpc_impl.CloudFile_ListPublicReq{
        OperatorId: params.OperatorID,
        Page: &rpc_impl.Pagination{
            PageSize: uint32(params.PageSize),
            PageNum:  uint32(params.PageNum),
        },
    })
    if err != nil {
        mlog.Logger().Error("list public cloud file failed", zap.Error(err))
        return mhttp.ResponseWithError(err.Error())
    }

    return mhttp.Response(structure.MakeListPublicCloudFileRes(res.Total, filesRPCToHTTP(res.Files...)))
}

func UploadCloudFile(r *http.Request) *mhttp.ResponseData {
    params := &structure.UploadCloudFileReqParams{}
    if errMsg := params.Decode(r); len(errMsg) > 0 {
        mlog.Logger().Error("parse request params failed", zap.String("err msg", errMsg))
        return mhttp.ResponseWithError(errMsg)
    }

    _, err := rpc.GetRPCClient().CloudFileClient.Upload(context.Background(), &rpc_impl.CloudFile_UploadReq{
        OperatorId:       params.OperatorID,
        File:             params.File,
        FileName:         params.FileName,
        ExtensionName:    params.ExtensionName,
        FileSize:         params.FileHeader.Size,
        LastModifiedTime: int64(params.LastModifiedTime),
        IsPublic:         params.IsPublic,
    })
    if err != nil {
        mlog.Logger().Error("upload cloud file failed", zap.Error(err))
        return mhttp.ResponseWithError(err.Error())
    }

    return mhttp.Response(mconst.EmptyHTTPRes)
}

func ModifyCloudFile(r *http.Request) *mhttp.ResponseData {
    params := &structure.ModifyCloudFileReqParams{}
    if errMsg := params.Decode(r); len(errMsg) > 0 {
        mlog.Logger().Error("parse request params failed", zap.String("err msg", errMsg))
        return mhttp.ResponseWithError(errMsg)
    }

    _, err := rpc.GetRPCClient().CloudFileClient.Modify(context.Background(), &rpc_impl.CloudFile_ModifyReq{
        OperatorId:       params.OperatorID,
        FileId:           params.FileID,
        Password:         params.Password,
        FileName:         params.FileName,
        ExtensionName:    params.ExtensionName,
        IsPublic:         params.IsPublic,
        File:             params.File,
        FileSize:         params.FileHeader.Size,
        LastModifiedTime: int64(params.LastModifiedTime),
    })
    if err != nil {
        mlog.Logger().Error("modify cloud file failed", zap.Error(err))
        return mhttp.ResponseWithError(err.Error())
    }

    return mhttp.Response(mconst.EmptyHTTPRes)
}

func DeleteCloudFile(r *http.Request) *mhttp.ResponseData {
    params := &structure.DeleteCloudFileReqParams{}
    if errMsg := params.Decode(r); len(errMsg) > 0 {
        mlog.Logger().Error("parse request params failed", zap.String("err msg", errMsg))
        return mhttp.ResponseWithError(errMsg)
    }

    _, err := rpc.GetRPCClient().CloudFileClient.Delete(context.Background(), &rpc_impl.CloudFile_DeleteReq{
        OperatorId: params.OperatorID,
        Password:   params.Password,
        FileId:     params.FileID,
    })
    if err != nil {
        mlog.Logger().Error("delete cloud file failed", zap.Error(err))
        return mhttp.ResponseWithError(err.Error())
    }

    return mhttp.Response(mconst.EmptyHTTPRes)
}

func filesRPCToHTTP(data ...*rpc_impl.CloudFile_Data) []*structure.FileRes {
    res := make([]*structure.FileRes, 0, len(data))
    for i := range data {
        res = append(res, &structure.FileRes{
            FileID:           data[i].FileId,
            FileName:         data[i].FileName,
            LastModifiedTime: time.Duration(data[i].LastModifiedTime),
            FileURL:          data[i].FileUrl,
            IsPublic:         data[i].IsPublic,
            UpdateTime:       time.Duration(data[i].UpdateTime),
            CreatedTime:      time.Duration(data[i].CreatedTime),
        })
    }

    return res
}
