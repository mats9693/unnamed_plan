package rpc

import (
	"context"
	"fmt"
	"github.com/mats9693/unnamed_plan/services/shared/const"
	"github.com/mats9693/unnamed_plan/services/shared/db/model"
	"github.com/mats9693/unnamed_plan/services/shared/proto/impl"
	"github.com/mats9693/unnamed_plan/services/shared/utils"
	"github.com/mats9693/unnamed_plan/services/user/config"
	"github.com/mats9693/unnamed_plan/services/user/db"
	"github.com/pkg/errors"
)

type userServerImpl struct {
	rpc_impl.UnimplementedIUserServer
}

var _ rpc_impl.IUserServer = (*userServerImpl)(nil)

var userServerImplIns = &userServerImpl{}

func GetUserServer() *userServerImpl {
	return userServerImplIns
}

func (s *userServerImpl) Login(_ context.Context, req *rpc_impl.User_LoginReq) (*rpc_impl.User_LoginRes, error) {
	if len(req.UserName) < 1 || len(req.Password) < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	user, err := db.GetUserDao().QueryOneByUserName(req.UserName)
	if err != nil {
		return nil, err
	}

	if user.Password != utils.CalcSHA256(req.Password, user.Salt) {
		return nil, utils.NewError(mconst.Error_InvalidAccountOrPassword)
	}

	return &rpc_impl.User_LoginRes{
		UserId:     user.ID,
		Nickname:   user.Nickname,
		Permission: uint32(user.Permission),
	}, nil
}

func (s *userServerImpl) List(_ context.Context, req *rpc_impl.User_ListReq) (*rpc_impl.User_ListRes, error) {
	if len(req.OperatorId) < 1 || req.GetPage() == nil || req.GetPage().PageSize < 1 || req.GetPage().PageNum < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	pageSize := int(req.GetPage().PageSize)
	pageNum := int(req.GetPage().PageNum)

	users, count, err := db.GetUserDao().QueryPageLEPermission(pageSize, pageNum, req.OperatorId)
	if err != nil {
		return nil, err
	}

	return &rpc_impl.User_ListRes{
		Total: uint32(count),
		Users: usersDBToRPC(users...),
	}, nil
}

func (s *userServerImpl) Create(_ context.Context, req *rpc_impl.User_CreateReq) (*rpc_impl.User_CreateRes, error) {
	if len(req.OperatorId) < 1 || len(req.UserName) < 1 || len(req.Password) < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	operator, err := db.GetUserDao().QueryOne(req.OperatorId)
	if err != nil {
		return nil, err
	}

	permission := uint8(req.Permission)
	if operator.Permission < config.GetConfig().ARankAdminPermission || operator.Permission <= permission {
		return nil, utils.NewError(mconst.Error_PermissionDenied)
	}

	salt := utils.RandomHexString(10)
	err = db.GetUserDao().Insert(&model.User{
		UserName:   req.UserName,
		Nickname:   req.UserName,
		Password:   utils.CalcSHA256(req.Password, salt),
		Salt:       salt,
		Permission: permission,
		CreatedBy:  req.OperatorId,
	})
	if err != nil {
		return nil, err
	}

	return &rpc_impl.User_CreateRes{}, nil
}

func (s *userServerImpl) Lock(_ context.Context, req *rpc_impl.User_LockReq) (*rpc_impl.User_LockRes, error) {
	if len(req.OperatorId) < 1 || len(req.UserId) < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	users, err := db.GetUserDao().Query([]string{req.OperatorId, req.UserId})
	if err != nil {
		return nil, err
	}

	users, err = sortUsersByUserID(users, []string{req.OperatorId, req.UserId})
	if err != nil {
		return nil, err
	}

	if users[1].IsLocked {
		return nil, utils.NewError(mconst.Error_UserAlreadyLocked)
	}
	if users[0].Permission < config.GetConfig().ARankAdminPermission || users[0].Permission <= users[1].Permission {
		return nil, utils.NewError(mconst.Error_PermissionDenied)
	}

	users[1].IsLocked = true

	err = db.GetUserDao().UpdateColumnsByUserID(users[1], model.User_IsLocked)
	if err != nil {
		return nil, err
	}

	return &rpc_impl.User_LockRes{}, nil
}

func (s *userServerImpl) Unlock(_ context.Context, req *rpc_impl.User_UnlockReq) (*rpc_impl.User_UnlockRes, error) {
	if len(req.OperatorId) < 1 || len(req.UserId) < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	users, err := db.GetUserDao().Query([]string{req.OperatorId, req.UserId})
	if err != nil {
		return nil, err
	}

	users, err = sortUsersByUserID(users, []string{req.OperatorId, req.UserId})
	if err != nil {
		return nil, err
	}

	if !users[1].IsLocked {
		return nil, utils.NewError(mconst.Error_UserAlreadyUnlocked)
	}
	if users[0].Permission < config.GetConfig().ARankAdminPermission || users[0].Permission <= users[1].Permission {
		return nil, utils.NewError(mconst.Error_PermissionDenied)
	}

	users[1].IsLocked = false

	err = db.GetUserDao().UpdateColumnsByUserID(users[1], model.User_IsLocked)
	if err != nil {
		return nil, err
	}

	return &rpc_impl.User_UnlockRes{}, nil
}

func (s *userServerImpl) ModifyInfo(_ context.Context, req *rpc_impl.User_ModifyInfoReq) (*rpc_impl.User_ModifyInfoRes, error) {
	if len(req.OperatorId) < 1 || len(req.UserId) < 1 || req.OperatorId != req.UserId {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	if len(req.Nickname)+len(req.Password) < 1 {
		return nil, utils.NewError(mconst.Error_NoValidModification)
	}

	user, err := verifyPwdByUserID(req.UserId, req.CurrPwd)
	if err != nil {
		return nil, err
	}

	updateColumns := make([]string, 0, 2)
	if len(req.Nickname) > 0 {
		user.Nickname = req.Nickname
		updateColumns = append(updateColumns, model.User_Nickname)
	}
	if len(req.Password) > 0 {
		user.Password = utils.CalcSHA256(req.Password, user.Salt)
		updateColumns = append(updateColumns, model.User_Password)
	}

	err = db.GetUserDao().UpdateColumnsByUserID(user, updateColumns...)
	if err != nil {
		return nil, err
	}

	return &rpc_impl.User_ModifyInfoRes{}, nil
}

func (s *userServerImpl) ModifyPermission(_ context.Context, req *rpc_impl.User_ModifyPermissionReq) (*rpc_impl.User_ModifyPermissionRes, error) {
	if len(req.OperatorId) < 1 || len(req.UserId) < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	users, err := db.GetUserDao().Query([]string{req.OperatorId, req.UserId})
	if err != nil {
		return nil, err
	}

	users, err = sortUsersByUserID(users, []string{req.OperatorId, req.UserId})
	if err != nil {
		return nil, err
	}

	permission := uint8(req.Permission)
	if users[0].Permission < config.GetConfig().SRankAdminPermission ||
		users[1].Permission >= config.GetConfig().SRankAdminPermission ||
		permission >= config.GetConfig().SRankAdminPermission {
		return nil, utils.NewError(mconst.Error_PermissionDenied)
	}

	users[1].Permission = permission

	err = db.GetUserDao().UpdateColumnsByUserID(users[1], model.User_Permission)
	if err != nil {
		return nil, err
	}

	return &rpc_impl.User_ModifyPermissionRes{}, nil
}

func (s *userServerImpl) Authenticate(_ context.Context, req *rpc_impl.User_AuthenticateReq) (*rpc_impl.User_AuthenticateRes, error) {
	if len(req.UserId) < 1 || len(req.Password) < 1 {
		return nil, utils.NewError(mconst.Error_InvalidParams)
	}

	_, err := verifyPwdByUserID(req.UserId, req.Password)
	if err != nil {
		return nil, utils.NewError(mconst.Error_InvalidAccountOrPassword)
	}

	return &rpc_impl.User_AuthenticateRes{}, nil
}

func usersDBToRPC(data ...*model.User) []*rpc_impl.User_Data {
	res := make([]*rpc_impl.User_Data, 0, len(data))
	for i := range data {
		res = append(res, &rpc_impl.User_Data{
			UserId:     data[i].ID,
			UserName:   data[i].UserName,
			Nickname:   data[i].Nickname,
			IsLocked:   data[i].IsLocked,
			Permission: uint32(data[i].Permission),
		})
	}

	return res
}

func verifyPwdByUserID(userID string, password string) (*model.User, error) {
	user, err := db.GetUserDao().QueryOne(userID)
	if err != nil {
		return nil, err
	}

	if user.Password != utils.CalcSHA256(password, user.Salt) {
		return nil, utils.NewError(mconst.Error_InvalidAccountOrPassword)
	}

	return user, nil
}

func sortUsersByUserID(users []*model.User, order []string) ([]*model.User, error) {
	if len(users) != len(order) {
		return nil, errors.New(fmt.Sprintf("unmatched users amount, users %d, orders %d", len(users), len(order)))
	}

	length := len(users)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if order[j] == users[i].ID {
				users[i], users[j] = users[j], users[i]
				break
			}
		}
	}

	unmatchedIndex := -1
	for i := 0; i < length; i++ {
		if users[i].ID != order[i] {
			unmatchedIndex = i
			break
		}
	}

	if unmatchedIndex >= 0 {
		return nil, errors.New(fmt.Sprintf("unmatched user id: %s", users[unmatchedIndex].ID))
	}

	return users, nil
}
