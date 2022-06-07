package handler

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
	"strings"
)

func (t *handler) SaveNamespace(ctx context.Context, req *cos.SaveNamespaceRequest) (resp *cos.SaveNamespaceResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if req.Id != nil && req.Id.Valid && req.Id.Uint64 > 0 {
		where := orm.NewAndWhereWrapper().Eq(db.TableNamespace.CreatorId, currentId).Eq(db.TableNamespace.Id, req.Id)
		_, found, err := db.TableNamespace.FindOne(ctx, where)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if !found {
			return nil, cos.Error.NamespaceNotFound
		}
		if modelNamespace, found, err := db.TableNamespace.FindOneByNamespace(ctx, req.Namespace); err != nil {
			log.Error(err)
			return nil, err
		} else if found && !modelNamespace.Id.Equal(*req.Id) {
			return nil, cos.Error.NamespaceDuplicated
		}
		_, err = db.TableNamespace.UpdateById(ctx, *req.Id, orm.H{db.TableNamespace.Namespace: req.Namespace})
	} else {
		count, err := db.TableNamespace.CountByNamespace(ctx, req.Namespace)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if count > 0 {
			return nil, cos.Error.NamespaceDuplicated
		}
		_, err = db.TableNamespace.Create(ctx, db.ModelNamespace{
			Namespace:  req.Namespace,
			CreateTime: orm.Now(),
			UpdateTime: orm.Now(),
			DeletedAt:  0,
			IsDisabled: false,
			CreatorId:  currentId.NewPointer(),
		})
	}
	return nil, err
}

func (t *handler) ListNamespace(ctx context.Context, req *cos.ListNamespaceRequest) (resp *cos.ListNamespaceResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	where := orm.NewAndWhereWrapper().Eq(db.TableNamespace.CreatorId, currentId)
	if req.IsDisabled != nil && req.IsDisabled.Valid {
		where.Eq(db.TableNamespace.IsDisabled, req.IsDisabled)
	}
	if strings.TrimSpace(req.SearchKey) != "" {
		where.Like(db.TableNamespace.Namespace, fmt.Sprintf("%%%s%%", strings.TrimSpace(req.SearchKey)))
	}
	total, err := db.TableNamespace.Count(ctx, where)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	modelNamespaceList, err := db.TableNamespace.Page(ctx, req.Pagination.PageNum, req.Pagination.PageSize, where, orm.DESC(db.TableNamespace.CreateTime))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respList := make([]*cos.ListNamespaceItem, 0)
	for _, item := range modelNamespaceList {
		respList = append(respList, item.ToListNamespaceItemAsPointer())
	}
	resp = &cos.ListNamespaceResponse{
		Pagination: req.Pagination.ToRespPointer(total),
		List:       respList,
	}
	return resp, nil
}

func (t *handler) DeleteNamespace(ctx context.Context, req *cos.DeleteNamespaceRequest) (resp *cos.DeleteNamespaceResponse, err error) {
	currentId, err := yc.GetUserId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	modelNamespace, found, err := db.TableNamespace.FindOneById(ctx, *req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		return nil, cos.Error.NamespaceNotFound
	}
	if !modelNamespace.CreatorId.Equal(currentId) {
		return nil, cos.Error.NoPermissionDeleteNamespace
	}
	_, err = db.TableNamespace.SoftDeleteById(ctx, *req.Id)
	return nil, err
}

func (t *handler) UpdateStatusNamespace(ctx context.Context, req *cos.UpdateStatusNamespaceRequest) (resp *cos.UpdateStatusNamespaceResponse, err error) {
	userId, _ := yc.GetUserId(ctx)
	modelNamespace, found, err := db.TableNamespace.FindOneById(ctx, *req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !found {
		log.Error(err)
		return nil, cos.Error.NamespaceNotFound
	}
	if !modelNamespace.CreatorId.Equal(userId) {
		return nil, cos.Error.NoPermissionAccessNamespace
	}
	_, err = db.TableNamespace.UpdateById(ctx, *req.Id, orm.H{
		db.TableNamespace.IsDisabled: req.IsDisabled,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cos.UpdateStatusNamespaceResponse{}, nil
}

func (t *handler) SelectorNamespace(ctx context.Context, req *cos.SelectorRequest) (resp *cos.SelectorResponse, err error) {
	userId, _ := yc.GetUserId(ctx)
	where := orm.NewAndWhereWrapper().Eq(db.TableNamespace.CreatorId, userId)
	if req.IsDisabled != nil && req.IsDisabled.Valid {
		where.Eq(db.TableNamespace.IsDisabled, req.IsDisabled)
	}
	modelNamespaceList, err := db.TableNamespace.Find(ctx, where, orm.DESC(db.TableNamespace.CreateTime))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	list := make([]*cos.SelectorItem, 0)
	for _, item := range modelNamespaceList {
		list = append(list, &cos.SelectorItem{
			Label: item.Namespace,
			Value: item.Id.NumberAsString(),
		})
	}
	resp = &cos.SelectorResponse{
		List: list,
	}
	return resp, nil
}
