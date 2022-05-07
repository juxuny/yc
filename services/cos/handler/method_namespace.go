package handler

import (
	"context"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
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
		_, err = db.TableNamespace.UpdateById(ctx, *req.Id, orm.H{db.TableNamespace.Namespace: req.Namespace})
	} else {
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
	total, err := db.TableNamespace.CountByCreatorId(ctx, currentId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	modelNamespaceList, err := db.TableNamespace.PageByCreatorId(ctx, req.Pagination.PageNum, req.Pagination.PageSize, currentId, orm.DESC(db.TableNamespace.CreateTime))
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
	return &cos.DeleteNamespaceResponse{}, nil
}
