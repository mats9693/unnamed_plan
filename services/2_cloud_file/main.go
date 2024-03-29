package main

import (
	"fmt"
	"github.com/mats9693/unnamed_plan/services/2_cloud_file/config"
	"github.com/mats9693/unnamed_plan/services/2_cloud_file/db"
	"github.com/mats9693/unnamed_plan/services/2_cloud_file/rpc"
	"github.com/mats9693/unnamed_plan/services/shared/config"
	"github.com/mats9693/unnamed_plan/services/shared/const"
	"github.com/mats9693/unnamed_plan/services/shared/db"
	"github.com/mats9693/unnamed_plan/services/shared/init"
	"github.com/mats9693/unnamed_plan/services/shared/log"
	"github.com/mats9693/unnamed_plan/services/shared/proto/go"
	"github.com/mats9693/unnamed_plan/services/shared/registration_center_embedded"
	"github.com/mats9693/unnamed_plan/services/shared/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	err := initialize.InitFromConfigCenter(mconst.UID_Service_Cloud_File, mdb.Init, config.Init, db.Init, initCloudFileDir)
	if err != nil {
		mlog.Logger().Error("init failed", zap.Error(err))
		return
	}

	ip, port, err := initialize.GetIPAndFreePort()
	if err != nil {
		mlog.Logger().Error("get ip or free port failed", zap.Error(err))
		return
	}

	localAddress := fmt.Sprintf("127.0.0.1:%d", port)
	internetAddress := fmt.Sprintf("%s:%d", ip, port)

	rceServer, err := rce.InitAndRegister(mconfig.GetCoreTarget(), mconst.UID_Service_Cloud_File, internetAddress)
	if err != nil {
		mlog.Logger().Error("register server to rc failed", zap.Error(err))
		return
	}

	startCloudFileServer(localAddress, rceServer)
}

func initCloudFileDir() error {
	root := config.GetConfig().CloudFileRootPath
	if len(root) < 1 {
		root = "./files/"
	}

	cloudFileDir := utils.FormatDirSuffix(root) + config.GetConfig().CloudFilePublicDir

	err := os.MkdirAll(cloudFileDir, 0755)
	if err != nil {
		mlog.Logger().Error("os.MkdirAll failed", zap.Error(err))
		return err
	}

	mlog.Logger().Info("> Cloud file directory init finish.")

	return nil
}

func startCloudFileServer(address string, rceServer rpc_impl.IRegistrationCenterEmbeddedServer) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		mlog.Logger().Error(fmt.Sprintf("listen on %s failed", address), zap.Error(err))
		return
	}

	server := grpc.NewServer()
	rpc_impl.RegisterICloudFileServer(server, rpc.GetCloudFileServer())
	rpc_impl.RegisterIRegistrationCenterEmbeddedServer(server, rceServer)
	//reflection.Register(server) // for grpc ui

	mlog.Logger().Info("> Listening at : " + address)

	// Serve is blocked
	err = server.Serve(listener)
	if err != nil {
		mlog.Logger().Error(fmt.Sprintf("serve on %s failed", address), zap.Error(err))
		return
	}
}
