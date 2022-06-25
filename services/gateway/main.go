package main

import (
	"github.com/mats9693/unnamed_plan/services/gateway/http"
	"github.com/mats9693/unnamed_plan/services/shared/config"
	"github.com/mats9693/unnamed_plan/services/shared/const"
	"github.com/mats9693/unnamed_plan/services/shared/http"
	"github.com/mats9693/unnamed_plan/services/shared/init"
	"github.com/mats9693/unnamed_plan/services/shared/log"
	"github.com/mats9693/unnamed_plan/services/shared/registration_center_embedded"
	"go.uber.org/zap"
)

func main() {
	err := initialize.InitFromConfigCenter(mconst.UID_Service_Gateway, http.Init)
	if err != nil {
		mlog.Logger().Error("init failed", zap.Error(err))
		return
	}

	err = rc_embedded.Init(mconfig.GetRegistrationCenterTarget())
	if err != nil {
		mlog.Logger().Error("register server to rc failed", zap.Error(err))
		return
	}

	mhttp.StartServer(http.GetHandler())
}
