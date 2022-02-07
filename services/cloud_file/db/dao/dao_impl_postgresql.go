package dao

import (
	"github.com/mats9693/unnamed_plan/services/shared/const"
	mdb "github.com/mats9693/unnamed_plan/services/shared/db/dal"
	"github.com/mats9693/unnamed_plan/services/shared/db/model"
	"github.com/mats9693/unnamed_plan/services/shared/utils"
	"time"
)

type CloudFilePostgresql model.CloudFile

var _ CloudFileDao = (*CloudFilePostgresql)(nil)

func (cf *CloudFilePostgresql) Insert(cloudFile *model.CloudFile) error {
	if len(cloudFile.ID) < 1 {
		cloudFile.Common = model.NewCommon()
	}

	return mdb.DB().WithTx(func(conn mdb.Conn) error {
		_, err := conn.PostgresqlConn.Model(cloudFile).Insert()
		return err
	})
}

func (cf *CloudFilePostgresql) QueryOne(cloudFileID string) (*model.CloudFile, error) {
	file := &model.CloudFile{}

	err := mdb.DB().WithNoTx(func(conn mdb.Conn) error {
		return conn.PostgresqlConn.Model(file).
			Where(model.CloudFile_IsDeleted+" = ?", false).
			Where(model.CloudFile_FileID+" = ?", cloudFileID).
			First()
	})
	if err != nil {
		file = nil
	}

	return file, err
}

func (cf *CloudFilePostgresql) QueryPageByUploader(
	pageSize int,
	pageNum int,
	userID string,
) ([]*model.CloudFile, int, error) {
	files := make([]*model.CloudFile, 0)
	count := 0

	err := mdb.DB().WithNoTx(func(conn mdb.Conn) error {
		var err2 error
		count, err2 = conn.PostgresqlConn.Model(&files).
			Where(model.CloudFile_IsDeleted+" = ?", false).
			Where(model.CloudFile_UploadedBy+" = ?", userID).
			Order(model.Common_UpdateTime + " DESC").
			Offset((pageNum - 1) * pageSize).Limit(pageSize).SelectAndCount()

		return err2
	})
	if err != nil {
		files = nil
		count = 0
	}

	return files, count, err
}

// QueryPageInPublic
/**
Core: sub-query
	select *
	from cloud_files cf
	where cf.uploaded_by in (
		select "id"
		from users u
		where "permission" <= (
			select "permission"
			from users u2
			where id = 'user id'
		)
	);
*/
func (cf *CloudFilePostgresql) QueryPageInPublic(
	pageSize int,
	pageNum int,
	userID string,
) ([]*model.CloudFile, int, error) {
	files := make([]*model.CloudFile, 0)
	count := 0

	err := mdb.DB().WithNoTx(func(conn mdb.Conn) error {
		permission := conn.PostgresqlConn.Model((*model.User)(nil)).Column(model.User_Permission).Where(model.Common_ID+" = ?", userID)
		userIDs := conn.PostgresqlConn.Model((*model.User)(nil)).Column(model.Common_ID).Where(model.User_Permission+" <= (?)", permission)

		var err2 error
		count, err2 = conn.PostgresqlConn.Model(&files).
			Where(model.CloudFile_IsDeleted+" = ?", false).
			Where(model.CloudFile_IsPublic+" = ?", true).
			Where(model.CloudFile_UploadedBy+" in (?)", userIDs).
			Order(model.Common_UpdateTime + " DESC").
			Offset((pageNum - 1) * pageSize).
			Limit(pageSize).
			SelectAndCount()

		return err2
	})
	if err != nil {
		files = nil
		count = 0
	}

	return files, count, err
}

func (cf *CloudFilePostgresql) UpdateColumnsByFileID(data *model.CloudFile, columns ...string) error {
	data.UpdateTime = time.Duration(time.Now().Unix())

	return mdb.DB().WithTx(func(conn mdb.Conn) error {
		query := conn.PostgresqlConn.Model(data).Column(model.Common_UpdateTime)
		for i := range columns {
			query.Column(columns[i])
		}

		res, err := query.Where(model.CloudFile_IsDeleted+" = ?", false).
			Where(model.CloudFile_FileID + " = ?" + model.CloudFile_FileID).Update()
		if err != nil {
			return err
		}

		if res.RowsAffected() < 0 {
			return utils.NewError(mconst.Error_FileAlreadyDeleted)
		}

		return nil
	})
}
