package dao

import (
	"github.com/mats9693/unnamed_plan/services/shared/const"
	"github.com/mats9693/unnamed_plan/services/shared/db"
	"github.com/mats9693/unnamed_plan/services/shared/db/dal"
	"github.com/mats9693/unnamed_plan/services/shared/db/model"
	"time"
)

type TaskPostgresql model.Task

var _ TaskDao = (*TaskPostgresql)(nil)

func (t *TaskPostgresql) Insert(task *model.Task) error {
	if len(task.ID) < 1 {
		task.Common = model.NewCommon()
	}

	return mdb.DB().WithTx(func(conn mdal.Conn) error {
		_, err := conn.PostgresqlConn.Model(task).Insert()
		return err
	})
}

func (t *TaskPostgresql) QueryByPoster(userID string) (tasks []*model.Task, count int, err error) {
	err = mdb.DB().WithNoTx(func(conn mdal.Conn) error {
		count, err = conn.PostgresqlConn.Model(&tasks).
			Where(model.Task_PostedBy+" = ?", userID).
			Where(model.Task_Status+" != ?", mconst.TaskStatus_History).
			Order(model.Common_UpdateTime + " ASC").
			SelectAndCount()

		return err
	})
	if err != nil {
		tasks = nil
		count = 0
	}

	return
}

func (t *TaskPostgresql) QueryOne(taskID string) (note *model.Task, err error) {
	note = &model.Task{}

	err = mdb.DB().WithNoTx(func(conn mdal.Conn) error {
		return conn.PostgresqlConn.Model(note).Where(model.Common_ID+" = ?", taskID).Select()
	})
	if err != nil {
		note = nil
	}

	return
}

func (t *TaskPostgresql) UpdateColumnsByTaskID(task *model.Task, columns ...string) error {
	task.UpdateTime = time.Duration(time.Now().Unix())
	task.OptimisticLock++

	return mdb.DB().WithTx(func(conn mdal.Conn) error {
		query := conn.PostgresqlConn.Model(task).Column(model.Common_UpdateTime, model.Common_OptimisticLock)
		for i := range columns {
			query.Column(columns[i])
		}

		_, err := query.Where(model.Common_ID+" = ?"+model.Common_ID).
			Where(model.Common_OptimisticLock+" = ?", task.OptimisticLock-1).Update()
		return err
	})
}