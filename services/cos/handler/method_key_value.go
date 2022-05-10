package handler

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
	"github.com/juxuny/yc/services/cos/impl"
)

func (t *handler) SaveValue(ctx context.Context, req *cos.SaveValueRequest) (resp *cos.SaveValueResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if req.ConfigId == nil || !req.ConfigId.Valid {
		return nil, cos.Error.MissingArguments.Wrap(fmt.Errorf("missing: configId"))
	}
	modelConfig, found, err := db.TableConfig.FindOne(ctx, orm.NewAndWhereWrapper().Eq(db.TableConfig.Id, req.ConfigId).Eq(db.TableConfig.CreatorId, currentId))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.ConfigNotFound
	}
	if modelConfig.IsDisabled {
		return nil, cos.Error.ConfigDisabled
	}
	if err := impl.SaveValue(ctx, db.ModelKeyValue{
		CreateTime:  orm.Now(),
		UpdateTime:  orm.Now(),
		DeletedAt:   0,
		IsDisabled:  false,
		ConfigKey:   req.Key,
		ConfigValue: req.Value,
		ValueType:   req.ValueType,
		ConfigId:    req.ConfigId,
		CreatorId:   &currentId,
		IsHot:       req.IsHot,
	}); err != nil {
		log.Error(err)
		return nil, err
	}

	return &cos.SaveValueResponse{}, nil
}

func (t *handler) DeleteValue(ctx context.Context, req *cos.DeleteValueRequest) (resp *cos.DeleteValueRequest, err error) {
	userId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if req.Id == nil || !req.Id.Valid {
		return nil, cos.Error.MissingArguments.Wrap(fmt.Errorf("missing: configId"))
	}
	modelKeyValue, found, err := db.TableKeyValue.FindOne(ctx, orm.NewAndWhereWrapper().Eq(db.TableKeyValue.CreatorId, userId).Eq(db.TableKeyValue.Id, req.Id))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.KeyNotFound
	}
	_, err = db.TableKeyValue.SoftDeleteById(ctx, *modelKeyValue.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cos.DeleteValueRequest{}, nil
}

func (t *handler) ListValue(ctx context.Context, req *cos.ListValueRequest) (resp *cos.ListValueResponse, err error) {
	userId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if req.ConfigId == nil || !req.ConfigId.Valid {
		return nil, cos.Error.MissingConfigId
	}
	where := orm.NewAndWhereWrapper().Eq(db.TableKeyValue.CreatorId, userId).Eq(db.TableKeyValue.ConfigId, req.ConfigId).Eq(db.TableKeyValue.IsDisabled, req.IsDisabled)
	if req.SearchKey != "" {
		where.Like(db.TableKeyValue.ConfigKey, fmt.Sprintf("%%%s%%", req.SearchKey))
	}
	count, err := db.TableKeyValue.Count(ctx, where)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if count == 0 {
		return &cos.ListValueResponse{
			Pagination: req.Pagination.ToRespPointer(count),
			List:       []*cos.KeyValueResp{},
		}, nil
	}
	resp = &cos.ListValueResponse{Pagination: req.Pagination.ToRespPointer(count)}
	items, err := db.TableKeyValue.Page(ctx, req.Pagination.PageNum, req.Pagination.PageSize, where)
	if err != nil {
		return nil, err
	}
	resp.List = items.MapToKeyValueRespList()
	return resp, nil
}
func (t *handler) DisableValue(ctx context.Context, req *cos.DisableValueRequest) (resp *cos.DisableValueResponse, err error) {
	userId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if req.Id == nil || !req.Id.Valid {
		return nil, cos.Error.MissingArguments.Wrap(fmt.Errorf("missing: configId"))
	}
	modelKeyValue, found, err := db.TableKeyValue.FindOne(ctx, orm.NewAndWhereWrapper().Eq(db.TableKeyValue.CreatorId, userId).Eq(db.TableKeyValue.Id, req.Id))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.KeyNotFound
	}
	_, err = db.TableKeyValue.UpdateById(ctx, *modelKeyValue.Id, orm.H{
		db.TableKeyValue.IsDisabled: req.IsDisabled,
		db.TableKeyValue.UpdateTime: orm.Now(),
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cos.DisableValueResponse{}, nil
}
func (t *handler) ListAllValue(ctx context.Context, req *cos.ListAllValueRequest) (resp *cos.ListAllValueResponse, err error) {
	userId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if req.ConfigId == nil || !req.ConfigId.Valid {
		return nil, cos.Error.MissingConfigId
	}
	where := orm.NewAndWhereWrapper().Eq(db.TableKeyValue.CreatorId, userId).Eq(db.TableKeyValue.ConfigId, req.ConfigId).Eq(db.TableKeyValue.IsDisabled, req.IsDisabled)
	count, err := db.TableKeyValue.Count(ctx, where)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if count == 0 {
		return &cos.ListAllValueResponse{
			List: []*cos.KeyValueResp{},
		}, nil
	}
	resp = &cos.ListAllValueResponse{}
	items, err := db.TableKeyValue.Find(ctx, where)
	if err != nil {
		return nil, err
	}
	resp.List = items.MapToKeyValueRespList()
	return resp, nil
}
