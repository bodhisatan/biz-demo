// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package handler

import (
	"context"
	"github.com/cloudwego/biz-demo/book-shop/app/item/common/converter"
	"github.com/cloudwego/biz-demo/book-shop/app/item/domain/service"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
)

type EditHandler struct {
	ctx   context.Context
	param *item.EditReq
}

func NewEditHandler(ctx context.Context, req *item.EditReq) *EditHandler {
	return &EditHandler{
		ctx:   ctx,
		param: req,
	}
}

func (h *EditHandler) Edit() (*item.EditResp, error) {
	resp := &item.EditResp{
		BaseResp: errno.BuildBaseResp(errno.Success),
	}

	updateService := service.GetProductUpdateServiceInstance()
	queryService := service.GetProductQueryServiceInstance()

	originEntity, err := queryService.GetProduct(h.ctx, h.param.ProductId)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	targetEntity, err := converter.ConvertEditReq2Entity(originEntity, h.param)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	err = updateService.EditProduct(h.ctx, originEntity, targetEntity)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	return resp, nil
}
