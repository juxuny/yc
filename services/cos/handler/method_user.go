package handler

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
	"github.com/juxuny/yc/utils"
)

func (t *handler) UserInfo(ctx context.Context, req *cos.UserInfoRequest) (resp *cos.UserInfoResponse, err error) {
	return &cos.UserInfoResponse{}, nil
}

func (t *handler) UpdateInfo(ctx context.Context, req *cos.UpdateInfoRequest) (resp *cos.UpdateInfoResponse, err error) {
	return &cos.UpdateInfoResponse{}, nil
}

func (t *handler) ModifyPassword(ctx context.Context, req *cos.ModifyPasswordRequest) (resp *cos.ModifyPasswordResponse, err error) {
	return &cos.ModifyPasswordResponse{}, nil
}

func (t *handler) CreateUser(ctx context.Context, req *cos.CreateUserRequest) (resp *cos.CreateUserResponse, err error) {
	count, err := db.TableAccount.Count(ctx, orm.NewAndWhereWrapper().Eq(db.TableAccount.Identifier, req.Identifier))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, cos.Error.AccountExists
	}
	account := db.ModelAccount{
		Identifier:  req.Identifier,
		Credential:  utils.Bcrypt(req.Credential),
		AccountType: req.AccountType,
		CreateTime:  orm.Now(),
		UpdateTime:  orm.Now(),
		DeletedAt:   0,
		IsDisabled:  false,
		CreatorId:   dt.NewID(0),
	}
	_, err = db.TableAccount.Create(ctx, account)
	if err != nil {
		return nil, err
	}
	return &cos.CreateUserResponse{}, nil
}
