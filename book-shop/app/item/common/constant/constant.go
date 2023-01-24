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

package constant

type ProductStatus = int64

const (
	ProductStatusOnline  ProductStatus = 0 // 商品上架
	ProductStatusOffline ProductStatus = 1 // 商品下架
	ProductStatusDelete  ProductStatus = 2 // 商品删除
)

var ProductStatusDescMap = map[ProductStatus]string{
	ProductStatusOnline:  "上架",
	ProductStatusOffline: "下架",
	ProductStatusDelete:  "删除",
}

type StateOperationType = int64

const (
	StateOperationTypeAdd     StateOperationType = 1 //新建商品
	StateOperationTypeSave    StateOperationType = 2 //保存商品
	StateOperationTypeDel     StateOperationType = 3 //删除商品
	StateOperationTypeOffline StateOperationType = 4 //商品下架
	StateOperationTypeOnline  StateOperationType = 5 //商品上架
)
