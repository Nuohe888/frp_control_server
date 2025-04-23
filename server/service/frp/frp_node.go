package frp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/frp"
	frpReq "github.com/flipped-aurora/gin-vue-admin/server/model/frp/request"
)

type NodeService struct{}

// CreateFrpNode 创建frp节点记录
// Author [yourname](https://github.com/yourname)
func (FnService *NodeService) CreateFrpNode(ctx context.Context, Fn *frp.FrpNode) (err error) {
	err = global.GVA_DB.Create(Fn).Error
	return err
}

// DeleteFrpNode 删除frp节点记录
// Author [yourname](https://github.com/yourname)
func (FnService *NodeService) DeleteFrpNode(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&frp.FrpNode{}, "id = ?", ID).Error
	return err
}

// DeleteFrpNodeByIds 批量删除frp节点记录
// Author [yourname](https://github.com/yourname)
func (FnService *NodeService) DeleteFrpNodeByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]frp.FrpNode{}, "id in ?", IDs).Error
	return err
}

// UpdateFrpNode 更新frp节点记录
// Author [yourname](https://github.com/yourname)
func (FnService *NodeService) UpdateFrpNode(ctx context.Context, Fn frp.FrpNode) (err error) {
	err = global.GVA_DB.Model(&frp.FrpNode{}).Where("id = ?", Fn.ID).Updates(&Fn).Error
	return err
}

// GetFrpNode 根据ID获取frp节点记录
// Author [yourname](https://github.com/yourname)
func (FnService *NodeService) GetFrpNode(ctx context.Context, ID string) (Fn frp.FrpNode, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Fn).Error
	return
}

// GetFrpNodeInfoList 分页获取frp节点记录
// Author [yourname](https://github.com/yourname)
func (FnService *NodeService) GetFrpNodeInfoList(ctx context.Context, info frpReq.FrpNodeSearch) (list []frp.FrpNode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&frp.FrpNode{})
	var Fns []frp.FrpNode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Fns).Error
	return Fns, total, err
}
func (FnService *NodeService) GetFrpNodePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
