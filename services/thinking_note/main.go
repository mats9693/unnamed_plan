package main

import (
    "fmt"
    "github.com/mats9693/unnamed_plan/services/shared/log"
    "github.com/mats9693/unnamed_plan/services/shared/proto/impl"
    "github.com/mats9693/unnamed_plan/services/thinking_note/config"
    "github.com/mats9693/unnamed_plan/services/thinking_note/rpc"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", config.GetConfig().Address)
    if err != nil {
        mlog.Logger().Error(fmt.Sprintf("listen on %s failed", config.GetConfig().Address), zap.Error(err))
        return
    }

    server := grpc.NewServer()
    thinkingNoteServer, err := rpc.GetThinkingNoteServer(config.GetConfig().UserServerAddress)
    if err != nil {
        mlog.Logger().Error("init thinking note server failed", zap.Error(err))
        return
    }
    rpc_impl.RegisterIThinkingNoteServer(server, thinkingNoteServer)

    mlog.Logger().Info("> Listening at : "+ config.GetConfig().Address)

    // Serve is blocked
    err = server.Serve(listener)
    if err != nil {
        mlog.Logger().Error(fmt.Sprintf("serve on %s failed", config.GetConfig().Address), zap.Error(err))
        return
    }
}
