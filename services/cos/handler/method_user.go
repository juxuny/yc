package handler

import (
	"context"
	"fmt"
	"github.com/juxuny/yc/services/cos/impl"
	"strings"

	"github.com/juxuny/yc"
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
	if req.UserId != nil && req.UserId.Valid && !req.UserId.Equal(&currentId) {
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
	if req.UserId != nil && req.UserId.Valid && !req.UserId.Equal(&currentId) {
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
	if req.UserId != nil && req.UserId.Valid && !req.UserId.Equal(&currentId) {
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
	userId, _ := yc.GetUserId(ctx)
	if req.UserId != nil && req.GetUserId().Valid && req.GetUserId().Uint64 > 0 {
		userInfo, found, err := db.TableAccount.FindOneById(ctx, *req.UserId)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if !found {
			return nil, cos.Error.AccountNotFound.Wrap(fmt.Errorf("userId=%v", req.UserId.Uint64))
		}
		if userInfo.CreatorId == nil || !userInfo.CreatorId.Valid || !userInfo.CreatorId.Equal(&userId) {
			return nil, cos.Error.NoPermissionUpdateUserInfo
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
		currentAccount, found, err := db.TableAccount.FindOneById(ctx, userId)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if !found {
			log.Warn("current user data not found")
			return nil, cos.Error.IllegalUserData
		}
		if currentAccount.CreatorId != nil && currentAccount.CreatorId.Uint64 != 0 {
			return nil, cos.Error.NotAllowedCreateAccount
		}
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
			CreatorId:   currentAccount.Id,
			Nick:        req.Nick,
		}
		_, err = db.TableAccount.Create(ctx, account)
		if err != nil {
			return nil, err
		}
	}
	return &cos.SaveOrCreateUserResponse{}, nil
}

func (t *handler) UserList(ctx context.Context, req *cos.UserListRequest) (resp *cos.UserListResponse, err error) {
	userId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	where := orm.NewAndWhereWrapper()
	where.Eq(db.TableAccount.CreatorId, userId).Eq(db.TableAccount.IsDisabled, req.IsDisabled)
	if req.SearchKey != "" {
		searchKey := fmt.Sprintf("%%%s%%", req.SearchKey)
		where.Nested(orm.NewOrWhereWrapper().Like(db.TableAccount.Nick, searchKey).Like(db.TableAccount.Identifier, searchKey))
	}
	count, err := db.TableAccount.Count(ctx, where)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if count == 0 {
		return &cos.UserListResponse{
			Pagination: req.Pagination.ToRespPointer(count),
			List:       []*cos.UserListItem{},
		}, nil
	}
	accountItems, err := db.TableAccount.Page(ctx, req.Pagination.PageNum, req.Pagination.PageSize, where)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	resp = &cos.UserListResponse{
		Pagination: req.Pagination.ToRespPointer(count),
	}
	resp.List = accountItems.MapToUserListItemList()
	return resp, nil
}

func (t *handler) UserUpdateStatus(ctx context.Context, req *cos.UserUpdateStatusRequest) (resp *cos.UserUpdateStatusResponse, err error) {
	userId, _ := yc.GetUserId(ctx)
	modelAccount, found, err := db.TableAccount.FindOneById(ctx, *req.UserId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.AccountNotFound
	}
	if !modelAccount.CreatorId.Equal(&userId) {
		log.Error("no permission set user status")
		return nil, cos.Error.NoPermissionAccessUserInfo
	}
	_, err = db.TableAccount.UpdateById(ctx, *req.UserId, orm.H{
		db.TableAccount.IsDisabled: req.IsDisabled,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cos.UserUpdateStatusResponse{}, nil
}

func (t *handler) UserDelete(ctx context.Context, req *cos.UserDeleteRequest) (resp *cos.UserDeleteResponse, err error) {
	userId, _ := yc.GetUserId(ctx)
	modelAccount, found, err := db.TableAccount.FindOneById(ctx, *req.UserId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.AccountNotFound
	}
	if !modelAccount.CreatorId.Equal(&userId) {
		return nil, cos.Error.NoPermissionAccessUserInfo
	}
	_, err = db.TableAccount.SoftDeleteById(ctx, *req.UserId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cos.UserDeleteResponse{}, nil
}

func (t *handler) AccessKeyList(ctx context.Context, req *cos.AccessKeyListRequest) (resp *cos.AccessKeyListResponse, err error) {
	userId, _ := yc.GetUserId(ctx)
	where := orm.NewAndWhereWrapper().Eq(db.TableAccessKey.UserId, userId).Eq(db.TableAccessKey.IsDisabled, req.IsDisabled)
	if strings.TrimSpace(req.SearchKey) != "" {
		where.Like(db.TableAccessKey.Remark, fmt.Sprintf("%%%s%%", req.SearchKey))
	}
	count, err := db.TableAccessKey.Count(ctx, where)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if count == 0 {
		return &cos.AccessKeyListResponse{
			Pagination: req.Pagination.ToRespPointer(count),
			List:       make([]*cos.AccessKeyItem, 0),
		}, nil
	}
	items, err := db.TableAccessKey.Page(ctx, req.Pagination.PageNum, req.Pagination.PageSize, where, orm.DESC(db.TableAccessKey.CreateTime))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	resp = &cos.AccessKeyListResponse{
		Pagination: req.Pagination.ToRespPointer(count),
		List:       items.MapToAccessKeyItemList(),
	}
	return resp, nil
}

func (t *handler) CreateAccessKey(ctx context.Context, req *cos.CreateAccessKeyRequest) (resp *cos.CreateAccessKeyResponse, err error) {
	resp = &cos.CreateAccessKeyResponse{
		AccessKey: "",
		Remark:    req.Remark,
	}
	userId, _ := yc.GetUserId(ctx)
	err = yc.Retry(func() (isEnd bool, err error) {
		resp.AccessKey = utils.StringHelper.RandString(64)
		_, err = db.TableAccessKey.Create(ctx, db.ModelAccessKey{
			CreateTime:     orm.Now(),
			UpdateTime:     orm.Now(),
			IsDisabled:     false,
			DeletedAt:      0,
			UserId:         &userId,
			AccessKey:      resp.AccessKey,
			HasValidity:    req.HasValidity,
			ValidStartTime: req.ValidStartTime,
			ValidEndTime:   req.ValidEndTime,
			Remark:         req.Remark,
		})
		if err != nil {
			log.Error(err)
			return false, err
		}
		return true, nil
	})
	return resp, nil
}

func (t *handler) UpdateStatusAccessKey(ctx context.Context, req *cos.UpdateStatusAccessKeyRequest) (resp *cos.UpdateStatusAccessKeyResponse, err error) {
	if err := impl.CheckIfAllowToAccessAccessKey(ctx, *req.Id); err != nil {
		return nil, err
	}
	_, err = db.TableAccessKey.UpdateById(ctx, *req.Id, orm.H{
		db.TableAccessKey.IsDisabled: req.IsDisabled,
		db.TableAccessKey.UpdateTime: orm.Now(),
	})
	return &cos.UpdateStatusAccessKeyResponse{}, err
}

func (t *handler) DeleteAccessKey(ctx context.Context, req *cos.DeleteAccessKeyRequest) (resp *cos.DeleteAccessKeyResponse, err error) {
	if err := impl.CheckIfAllowToAccessAccessKey(ctx, *req.Id); err != nil {
		return nil, err
	}
	_, err = db.TableAccessKey.SoftDeleteById(ctx, *req.Id)
	return &cos.DeleteAccessKeyResponse{}, err
}

func (t *handler) SetRemarkAccessKey(ctx context.Context, req *cos.SetAccessKeyRemarkRequest) (resp *cos.SetAccessKeyRemarkResponse, err error) {
	if err := impl.CheckIfAllowToAccessAccessKey(ctx, *req.Id); err != nil {
		return nil, err
	}
	_, err = db.TableAccessKey.UpdateById(ctx, *req.Id, orm.H{
		db.TableAccessKey.Remark:     req.Remark,
		db.TableAccessKey.UpdateTime: orm.Now(),
	})
	return &cos.SetAccessKeyRemarkResponse{}, err
}
