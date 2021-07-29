package model

import (
	"github.com/mats9693/utils/uuid"
	"time"
)

type Common struct {
	ID          string        `pg:",pk"`
	CreatedTime time.Duration `pg:",use_zero"`
	UpdateTime  time.Duration `pg:",use_zero"`
	// todo: 乐观锁
}

type User struct {
	UserID     string `pg:",unique"`
	Name       string
	Password   string
	IsLocked   bool  `pg:",use_zero"`
	Permission uint8 `pg:",use_zero"`

	Common
}

func NewCommon() Common {
	return Common{
		ID:          uuid.New(),
		CreatedTime: time.Duration(time.Now().Unix()),
		UpdateTime:  time.Duration(time.Now().Unix()),
	}
}