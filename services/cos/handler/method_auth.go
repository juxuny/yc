package handler

import (
	"context"
	"github.com/juxuny/yc/jwt"
	"github.com/juxuny/yc/log"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
	"github.com/juxuny/yc/utils"
)

func (t *handler) Login(ctx context.Context, req *cos.LoginRequest) (resp *cos.LoginResponse, err error) {
	modelAccount, found, err := db.TableAccount.FindOneByIdentifier(ctx, req.UserName)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.AccountNotFound
	}
	log.Debug(utils.ToJson(modelAccount))
	if !utils.IsBcryptMatched(req.Password, modelAccount.Credential) {
		return nil, cos.Error.IncorrectPassword
	}
	token, err := jwt.GenerateToken(modelAccount.Id.Uint64, modelAccount.Identifier)
	if err != nil {
		log.Error(err)
		return nil, cos.Error.AuthorizeFailed.Wrap(err)
	}
	resp = &cos.LoginResponse{
		UserId: modelAccount.Id.Uint64,
		Name:   modelAccount.Identifier,
		Token:  token,
	}
	return resp, nil
}
