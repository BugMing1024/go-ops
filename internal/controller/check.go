package controller

import (
	"context"
	v1 "go-ops/api/v1"
)

var CheckItem *checkitem = new(checkitem)

type checkitem struct{}

func (self *checkitem) Create(ctx context.Context, req *v1.AddCheckItemReq) (res *v1.CheckItemRes, err error) {
	return
}

func (self *checkitem) Update(ctx context.Context, req *v1.UpdateCheckItemReq) (res *v1.CheckItemRes, err error) {
	return
}

func (self *checkitem) Query(ctx context.Context, req *v1.QueryCheckItemReq) (res *v1.QueryCheckItemRes, err error) {
	return
}