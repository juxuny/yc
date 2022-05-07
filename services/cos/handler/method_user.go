package handler

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
	"github.com/juxuny/yc/utils"
)

func (t *handler) UserInfo(ctx context.Context, req *cos.UserInfoRequest) (resp *cos.UserInfoResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	if req.UserId != nil && req.UserId.Valid && !req.UserId.Equal(currentId) {
		return nil, cos.Error.NoPermissionAccessUserInfo.WithField("userId", req.UserId)
	}
	modelAccount, found, err := db.TableAccount.FindOneById(ctx, currentId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.AccountNotFound
	}
	return modelAccount.ToUserInfoResponseAsPointer(), nil
}

func (t *handler) UpdateInfo(ctx context.Context, req *cos.UpdateInfoRequest) (resp *cos.UpdateInfoResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	if req.UserId != nil && req.UserId.Valid && !req.UserId.Equal(currentId) {
		return nil, cos.Error.NoPermissionAccessUserInfo.WithField("userId", req.UserId)
	}
	modelAccount, found, err := db.TableAccount.FindOneById(ctx, currentId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.AccountNotFound
	}
	if modelAccount.IsDisabled {
		return nil, cos.Error.AccountForbidden
	}
	_, err = db.TableAccount.UpdateById(ctx, currentId, orm.H{
		db.TableAccount.Nick: req.Nick,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cos.UpdateInfoResponse{}, nil
}

func (t *handler) ModifyPassword(ctx context.Context, req *cos.ModifyPasswordRequest) (resp *cos.ModifyPasswordResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	if req.UserId != nil && req.UserId.Valid && !req.UserId.Equal(currentId) {
		return nil, cos.Error.NoPermissionAccessUserInfo.WithField("userId", req.UserId)
	}
	modelAccount, found, err := db.TableAccount.FindOneById(ctx, currentId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.AccountNotFound
	}
	if modelAccount.IsDisabled {
		return nil, cos.Error.AccountForbidden
	}
	if !utils.IsBcryptMatched(req.OldPassword, modelAccount.Credential) {
		return nil, cos.Error.IncorrectPassword
	}
	_, err = db.TableAccount.UpdateById(ctx, currentId, orm.H{
		db.TableAccount.Credential: utils.Bcrypt(req.NewPassword),
	})
	return &cos.ModifyPasswordResponse{}, nil
}

func (t *handler) SaveOrCreateUser(ctx context.Context, req *cos.SaveOrCreateUserRequest) (resp *cos.SaveOrCreateUserResponse, err error) {
	if req.UserId != nil && req.GetUserId().Valid && req.GetUserId().Uint64 > 0 {
		_, found, err := db.TableAccount.FindOneById(ctx, *req.UserId)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if !found {
			return nil, cos.Error.AccountNotFound.Wrap(fmt.Errorf("userId=%v", req.UserId.Uint64))
		}
		rowsAffected, err := db.TableAccount.UpdateById(ctx, *req.UserId, orm.H{
			db.TableAccount.Nick: req.Nick,
		})
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if rowsAffected == 0 {
			return nil, cos.Error.NoDataModified
		}
	} else {
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
			CreatorId:   dt.NewIDPointer(0),
		}
		_, err = db.TableAccount.Create(ctx, account)
		if err != nil {
			return nil, err
		}
	}
	return &cos.SaveOrCreateUserResponse{}, nil
}
