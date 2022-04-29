package main

import (
	"fmt"
	"github.com/mats9693/unnamed_plan/services/config_center/config"
	"github.com/mats9693/unnamed_plan/services/config_center/db"
	i "github.com/mats9693/unnamed_plan/services/config_center/init"
	"github.com/mats9693/unnamed_plan/services/config_center/rpc"
	"github.com/mats9693/unnamed_plan/services/shared/db"
	"github.com/mats9693/unnamed_plan/services/shared/init"
	"github.com/mats9693/unnamed_plan/services/shared/log"
	"github.com/mats9693/unnamed_plan/services/shared/proto/impl"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	initialize.InitFromFile("config.json", mdb.Init, db.Init, config.Init, i.InitSupportConfig)

	// todo: start http server for web

	listener, err := net.Listen("tcp", config.GetConfig().Address)
	if err != nil {
		mlog.Logger().Error(fmt.Sprintf("listen on %s failed", config.GetConfig().Address), zap.Error(err))
		return
	}

	server := grpc.NewServer()
	rpc_impl.RegisterIConfigCenterServer(server, rpc.GetConfigCenterServer())

	mlog.Logger().Info("> Listening at : " + config.GetConfig().Address)

	// Serve is blocked
	err = server.Serve(listener)
	if err != nil {
		mlog.Logger().Error(fmt.Sprintf("serve on %s failed", config.GetConfig().Address), zap.Error(err))
		return
	}
}