package handler

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
)

func (t *handler) SaveConfig(ctx context.Context, req *cos.SaveConfigRequest) (resp *cos.SaveConfigResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if req.NamespaceId == nil || !req.NamespaceId.Valid {
		log.Error("namespace ID is nil")
		return nil, cos.Error.NamespaceNotFound
	}
	modelNamespace, found, err := db.TableNamespace.FindOneById(ctx, *req.NamespaceId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.NamespaceNotFound
	}
	if modelNamespace.CreatorId == nil || !currentId.Equal(*modelNamespace.CreatorId) {
		return nil, cos.Error.NoPermissionToAssessConfig
	}
	if req.BaseId != nil {
		_, found, err := db.TableConfig.FindOne(ctx, orm.NewAndWhereWrapper().Eq(db.TableConfig.Id, req.BaseId).Eq(db.TableConfig.IsDisabled, false))
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if !found {
			return nil, cos.Error.ConfigNotFound.WithField(db.TableConfig.Id.LowerFirstHump(), req.BaseId)
		}
	}
	if req.Id != nil && req.Id.Valid && req.Id.Uint64 > 0 {
		count, err := db.TableConfig.Count(ctx, orm.NewAndWhereWrapper().Eq(db.TableConfig.ConfigId, req.ConfigId).Neq(db.TableConfig.Id, req.Id))
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if count > 0 {
			return nil, cos.Error.ConfigIdDuplicated.WithField(db.TableConfig.ConfigId.LowerFirstHump(), req.ConfigId)
		}
		update := orm.H{
			db.TableConfig.ConfigId:   req.ConfigId,
			db.TableConfig.UpdateTime: orm.Now(),
		}
		_, err = db.TableConfig.UpdateById(ctx, *req.Id, update)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	} else {
		count, err := db.TableConfig.Count(ctx, orm.NewAndWhereWrapper().Eq(db.TableConfig.ConfigId, req.ConfigId))
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if count > 0 {
			return nil, cos.Error.ConfigIdDuplicated.WithField(db.TableConfig.ConfigId.LowerFirstHump(), req.ConfigId)
		}
		_, err = db.TableConfig.Create(ctx, db.ModelConfig{
			CreateTime:  orm.Now(),
			UpdateTime:  orm.Now(),
			ConfigId:    req.ConfigId,
			IsDisabled:  false,
			CreatorId:   &currentId,
			BaseId:      req.BaseId,
			NamespaceId: req.NamespaceId,
		})
		if err != nil {
			return nil, err
		}
	}
	return &cos.SaveConfigResponse{}, nil
}

func (t *handler) DeleteConfig(ctx context.Context, req *cos.DeleteConfigRequest) (resp *cos.DeleteConfigResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	modelConfig, found, err := db.TableConfig.FindOne(ctx, orm.NewAndWhereWrapper().Eq(db.TableConfig.Id, req.Id).Eq(db.TableConfig.CreatorId, currentId))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		log.Error("not found config:", req.Id)
		return nil, cos.Error.ConfigNotFound
	}
	_, err = db.TableConfig.SoftDelete(ctx, orm.NewAndWhereWrapper().Eq(db.TableConfig.DeletedAt, modelConfig.DeletedAt).Eq(db.TableConfig.Id, req.Id))
	return nil, err
}

func (t *handler) ListConfig(ctx context.Context, req *cos.ListConfigRequest) (resp *cos.ListConfigResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	where := orm.NewAndWhereWrapper().Eq(db.TableConfig.CreatorId, currentId).Eq(db.TableConfig.NamespaceId, req.NamespaceId)
	if req.IsDisabled != nil {
		where.Eq(db.TableConfig.IsDisabled, req.IsDisabled)
	}
	if req.SearchKey != "" {
		where.Like(db.TableConfig.ConfigId, fmt.Sprintf("%%%s%%", req.SearchKey))
	}
	count, err := db.TableConfig.Count(ctx, where)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if count == 0 {
		return &cos.ListConfigResponse{
			Pagination: req.Pagination.ToRespPointer(count),
			List:       []*cos.ListConfigItem{},
		}, nil
	}
	list, err := db.TableConfig.Page(ctx, req.Pagination.PageNum, req.Pagination.PageSize, where, orm.DESC(db.TableConfig.CreateTime))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cos.ListConfigResponse{
		Pagination: req.Pagination.ToRespPointer(count),
		List:       list.MapToListConfigItemList(),
	}, nil
}
