package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/mats9693/unnamed_plan/shared/go/config"
	. "github.com/mats9693/unnamed_plan/shared/go/const"
	"time"
)

type dbConfig struct {
	Addr     string        `json:"addr"`
	User     string        `json:"user"`
	Password string        `json:"password"`
	Database string        `json:"dbName"`
	Timeout  time.Duration `json:"timeout"` // db read and write timeout, second
	ShowSQL  bool          `json:"showSQL"` // show sql before query
}

var db *pg.DB

func init() {
	conf := getDBConfig()

	db = pg.Connect(&pg.Options{
		Addr:         conf.Addr,
		User:         conf.User,
		Password:     conf.Password,
		Database:     conf.Database,
		ReadTimeout:  conf.Timeout * time.Second, // default behavior is blocking
		WriteTimeout: conf.Timeout * time.Second, // default behavior is blocking
	})

	if conf.ShowSQL {
		//db.AddQueryHook(&dbConfig{}) // todo: try official method to print sql.
		db.AddQueryHook(pgdebug.DebugHook{
			Verbose: true,
		})
	}

	fmt.Println("> Database init finish.")
}

func getConn() *pg.Conn {
	return db.Conn()
}

func WithTx(task func(conn orm.DB) error) error {
	if task == nil {
		return nil // todo: return special error
	}

	conn := getConn()
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = tx.Commit()
		} else {
			err = tx.Rollback()
		}
	}()

	err = task(tx)

	return err
}

func WithNoTx(task func(conn orm.DB) error) error {
	if task == nil {
		return nil
	}

	conn := getConn()
	defer conn.Close()

	err := task(conn)

	return err
}

func getDBConfig() *dbConfig {
	byteSlice := config.GetConfig(UID_DB)

	conf := &dbConfig{}
	err := json.Unmarshal(byteSlice, conf)
	if err != nil {
		fmt.Printf("json unmarshal failed, uid: %s, error: %v\n", UID_DB, err)
		return nil
	}

	return conf
}

func (d *dbConfig) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	sqlBytes, err := q.FormattedQuery()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sqlBytes))
	}

	return c, nil
}

func (d *dbConfig) AfterQuery(_ context.Context, _ *pg.QueryEvent) error {
	return nil
}