package mconfig

import (
	"encoding/json"
	"github.com/mats9693/unnamed_plan/services/shared/const"
	"log"
)

type servicePublicConfig struct {
	init bool

	CoreTarget    string `json:"coreTarget"`
	RetryTimes    int    `json:"retryTimes"`
	RetryInterval int64  `json:"retryInterval"` // unit: second
}

var servicePublicConfigIns = &servicePublicConfig{}

func initServicePublicConfig() error {
	if servicePublicConfigIns.init { // have initialized
		log.Println("already initialized")
		return nil
	}

	byteSlice := GetConfig(mconst.UID_Public_Config)

	err := json.Unmarshal(byteSlice, servicePublicConfigIns)
	if err != nil {
		log.Printf("json unmarshal failed, uid: %s, error: %v\n", mconst.UID_Public_Config, err)
		return err
	}

	servicePublicConfigIns.init = true

	return nil
}

func GetCoreTarget() string {
	return servicePublicConfigIns.CoreTarget
}
