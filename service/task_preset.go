package service

import (
	"context"
	v1 "go-ops/api/v1"
	"go-ops/model/entity"
	"go-ops/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

type (
	sTaskPreset struct{}
)

var (
	insTaskPreset = sTaskPreset{}
)

func TaskPreset() *sTaskPreset {
	return &insTaskPreset
}

func (self *sTaskPreset) Create(ctx context.Context, req *v1.AddTaskPresetReq) (res *v1.TaskPresetItemRes, err error) {

	item := &entity.TaskPreset{
		Created: gtime.Now(),
		Name:    req.Name,
		Content: req.Content,
		Creater: req.Creater,
		Type:    req.Type,
		Uuid:    guid.S(),
	}

	_, err = dao.TaskPreset.Ctx(ctx).Data(item).Insert()

	if err != nil {
		return
	}

	res = &v1.TaskPresetItemRes{
		Name:    req.Name,
		Content: req.Content,
		Creater: req.Creater,
		Type:    req.Type,
		Uuid:    item.Uuid,
		Created: item.Created.String(),
	}

	return
}

func (self *sTaskPreset) Update(ctx context.Context, req *v1.UpdateTaskPresetReq) (res *v1.TaskPresetItemRes, err error) {

	item := &entity.TaskPreset{
		Updated: gtime.Now(),
		Name:    req.Name,
		Content: req.Content,
		Updater: req.Updater,
		Type:    req.Type,
		Uuid:    req.Uuid,
	}

	_, err = dao.TaskPreset.Ctx(ctx).Data(item).Where("uuid = ?", req.Uuid).Update()

	if err != nil {
		return
	}

	res = &v1.TaskPresetItemRes{
		Name:    req.Name,
		Content: req.Content,
		Updater: req.Updater,
		Type:    req.Type,
		Uuid:    item.Uuid,
		Updated: item.Updated.String(),
	}

	return
}

func (self *sTaskPreset) Query(ctx context.Context, req *v1.QueryTaskPresetReq) (res *v1.QueryTaskPresetRes, err error) {

	m := g.Map{}

	if req.Name != "" {
		m["name"] = req.Name

	}

	if req.Uuid != "" {
		m["uuid"] = req.Uuid
	}

	if req.Creater != "" {
		m["creater"] = req.Creater
	}

	list := make([]*entity.TaskPreset, 0)

	err = dao.TaskPreset.Ctx(ctx).Where(m).Scan(&list)

	if err != nil {
		return
	}

	res = new(v1.QueryTaskPresetRes)

	for _, item := range list {
		res.List = append(res.List, &v1.TaskPresetItemRes{
			Name:    item.Name,
			Content: item.Content,
			Updater: item.Updater,
			Type:    item.Type,
			Uuid:    item.Uuid,
			Updated: item.Updated.String(),
			Created: item.Created.String(),
		})
	}

	return
}

func (self *sTaskPreset) Delete(ctx context.Context, req *v1.DeleteTaskPresetReq) (err error) {

	_, err = dao.TaskPreset.Ctx(ctx).WhereIn("uuid", req.Uuids).Delete()
	return
}
