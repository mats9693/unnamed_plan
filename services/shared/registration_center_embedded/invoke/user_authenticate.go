package rce_invoke

import (
	"context"
	"github.com/mats9693/unnamed_plan/services/shared/const"
	"github.com/mats9693/unnamed_plan/services/shared/log"
	"github.com/mats9693/unnamed_plan/services/shared/proto/go"
	"github.com/mats9693/unnamed_plan/services/shared/registration_center_embedded"
	"github.com/mats9693/unnamed_plan/services/shared/utils"
	"go.uber.org/zap"
)

func AuthUserInfo(ctx context.Context, req *rpc_impl.User_AuthenticateReq) *rpc_impl.Error {
	conn, err := rce.GetClientConn(mconst.UID_Service_User)
	if err != nil {
		mlog.Logger().Error("get client conn failed", zap.Error(err))
		return utils.NewExecError(err.Error())
	}

	res, err := rpc_impl.NewIUserClient(conn).Authenticate(ctx, req)
	if err != nil { // grpc connection error
		mlog.Logger().Error(mconst.Error_GrpcConnectionError, zap.Error(err))

		rce.ReportInvalidTarget(mconst.UID_Service_User, conn.Target())

		return utils.NewGrpcConnectionError(err.Error())
	}
	if res != nil && res.Err != nil { // service exec error
		mlog.Logger().Error(mconst.Error_ExecutionError, zap.String("error", res.Err.String()))
		return res.Err
	}

	return nil
}
