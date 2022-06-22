package handler

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
	"github.com/juxuny/yc/services/cos/impl"
	"github.com/juxuny/yc/utils"
	"google.golang.org/protobuf/proto"
	"strings"
)

func (t *handler) SaveValue(ctx context.Context, req *cos.SaveValueRequest) (resp *cos.SaveValueResponse, err error) {
	currentId, err := yc.GetIncomingUserId(ctx)
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
		ConfigKey:   req.ConfigKey,
		ConfigValue: req.ConfigValue,
		ValueType:   req.ValueType,
		ConfigId:    req.ConfigId,
		CreatorId:   currentId,
		IsHot:       req.IsHot,
	}); err != nil {
		log.Error(err)
		return nil, err
	}

	return &cos.SaveValueResponse{}, nil
}

func (t *handler) DeleteValue(ctx context.Context, req *cos.DeleteValueRequest) (resp *cos.DeleteValueRequest, err error) {
	userId, err := yc.GetIncomingUserId(ctx)
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
	_, err = db.TableKeyValue.SoftDeleteById(ctx, modelKeyValue.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cos.DeleteValueRequest{}, nil
}

func (t *handler) ListValue(ctx context.Context, req *cos.ListValueRequest) (resp *cos.ListValueResponse, err error) {
	userId, err := yc.GetIncomingUserId(ctx)
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
	for i := range resp.List {
		resp.List[i].ConfigId = proto.Clone(req.ConfigId).(*dt.ID)
	}
	return resp, nil
}

func (t *handler) DisableValue(ctx context.Context, req *cos.DisableValueRequest) (resp *cos.DisableValueResponse, err error) {
	return nil, errors.SystemError.DisabledMethod
}

func (t *handler) ListAllValue(ctx context.Context, req *cos.ListAllValueRequest) (resp *cos.ListAllValueResponse, err error) {
	userId, err := yc.GetIncomingUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if req.ConfigId == nil || !req.ConfigId.Valid {
		return nil, cos.Error.MissingConfigId
	}
	modelConfig, found, err := db.TableConfig.FindOneById(ctx, req.ConfigId)
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
	where := orm.NewAndWhereWrapper().Eq(db.TableKeyValue.CreatorId, userId).Eq(db.TableKeyValue.IsDisabled, req.IsDisabled).Eq(db.TableKeyValue.IsHot, req.IsHot)
	req.SearchKey = strings.TrimSpace(req.SearchKey)
	if req.SearchKey != "" {
		where.Nested(orm.NewOrWhereWrapper().Like(db.TableKeyValue.ConfigKey, fmt.Sprintf("%%%s%%", req.SearchKey)))
	}
	where.Nested(orm.NewOrWhereWrapper().Eq(db.TableKeyValue.ConfigId, req.ConfigId).Eq(db.TableKeyValue.ConfigId, modelConfig.BaseId))

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
	items, err := db.TableKeyValue.Find(ctx, where, orm.ASC(db.TableKeyValue.ConfigKey))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if modelConfig.BaseId != nil && modelConfig.BaseId.Valid {
		baseItems := make(db.ModelKeyValueList, 0)
		currentItems := make(db.ModelKeyValueList, 0)
		baseKeys := utils.NewStringSet()
		currentKeys := utils.NewStringSet()
		for _, item := range items {
			if item.ConfigId.Equal(modelConfig.Id) {
				currentKeys.Add(item.ConfigKey)
				currentItems = append(currentItems, item)
			} else {
				baseKeys.Add(item.ConfigKey)
				baseItems = append(baseItems, item)
			}
		}
		interactKeys := baseKeys.Intersect(currentKeys)
		baseItems = baseItems.Filter(func(index int, item db.ModelKeyValue) bool {
			return !interactKeys.Exists(item.ConfigKey)
		})
		items = currentItems.MergeSort(baseItems, func(a, b db.ModelKeyValue) bool {
			return a.ConfigKey < b.ConfigKey
		})
	}
	resp.List = items.MapToKeyValueRespList()
	for i := range resp.List {
		resp.List[i].ConfigId = req.ConfigId.Clone()
	}
	return resp, nil
}

func (t *handler) UpdateStatusValue(ctx context.Context, req *cos.UpdateStatusValueRequest) (resp *cos.UpdateStatusValueResponse, err error) {
	userId, _ := yc.GetIncomingUserId(ctx)
	if req.Id == nil || !req.Id.Valid {
		return nil, cos.Error.MissingArguments.Wrap(fmt.Errorf("missing: id"))
	}
	modelKeyValue, found, err := db.TableKeyValue.FindOne(ctx, orm.NewAndWhereWrapper().Eq(db.TableKeyValue.CreatorId, userId).Eq(db.TableKeyValue.Id, req.Id))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.KeyNotFound
	}
	_, err = db.TableKeyValue.UpdateById(ctx, modelKeyValue.Id, orm.H{
		db.TableKeyValue.IsDisabled: req.IsDisabled,
		db.TableKeyValue.UpdateTime: orm.Now(),
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cos.UpdateStatusValueResponse{}, nil
}

func (t *handler) ListAllValueByConfigId(ctx context.Context, req *cos.ListAllValueByConfigIdRequest) (resp *cos.ListAllValueByConfigIdResponse, err error) {
	resp, err = &cos.ListAllValueByConfigIdResponse{}, nil
	userId, err := yc.GetIncomingUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	modelConfig, found, err := db.TableConfig.FindOneByConfigId(ctx, req.ConfigId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.ConfigNotFound.WithField(db.TableConfig.ConfigId.LowerFirstHump(), req.ConfigId)
	}
	if modelConfig.CreatorId == nil || !modelConfig.CreatorId.Equal(userId) {
		return nil, cos.Error.NoPermissionToAssessConfig.WithField(db.TableConfig.ConfigId.LowerFirstHump(), req.ConfigId)
	}
	values, err := t.ListAllValue(ctx, &cos.ListAllValueRequest{
		ConfigId:   modelConfig.Id.Clone(),
		IsDisabled: req.IsDisabled,
		IsHot:      req.IsHot,
		SearchKey:  req.SearchKey,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	resp.List = values.List
	return
}
