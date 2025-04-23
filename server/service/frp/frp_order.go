package frp

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/frp"
	frpReq "github.com/flipped-aurora/gin-vue-admin/server/model/frp/request"
)

type OrderService struct{}

// CreateFrpOrder 创建Frp订单记录
// Author [yourname](https://github.com/yourname)
func (FoService *OrderService) CreateFrpOrder(ctx context.Context, Fo *frp.FrpOrder) (err error) {
	var c int64

	global.GVA_DB.Where("port=?", Fo.Port).Find(&frp.FrpOrder{}).Count(&c)

	if c > 0 {
		return errors.New("端口重复")
	}

	err = global.GVA_DB.Create(Fo).Error
	return err
}

// DeleteFrpOrder 删除Frp订单记录
// Author [yourname](https://github.com/yourname)
func (FoService *OrderService) DeleteFrpOrder(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&frp.FrpOrder{}, "id = ?", ID).Error
	return err
}

// DeleteFrpOrderByIds 批量删除Frp订单记录
// Author [yourname](https://github.com/yourname)
func (FoService *OrderService) DeleteFrpOrderByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]frp.FrpOrder{}, "id in ?", IDs).Error
	return err
}

// UpdateFrpOrder 更新Frp订单记录
// Author [yourname](https://github.com/yourname)
func (FoService *OrderService) UpdateFrpOrder(ctx context.Context, Fo frp.FrpOrder) (err error) {

	var c int64
	global.GVA_DB.Where("port=?", Fo.Port).Find(&frp.FrpOrder{}).Count(&c)
	if c > 0 {
		return errors.New("端口重复")
	}
	err = global.GVA_DB.Model(&frp.FrpOrder{}).Where("id = ?", Fo.ID).Updates(&Fo).Error
	return err
}

// GetFrpOrder 根据ID获取Frp订单记录
// Author [yourname](https://github.com/yourname)
func (FoService *OrderService) GetFrpOrder(ctx context.Context, ID string) (Fo frp.FrpOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Fo).Error
	return
}

// GetFrpOrderInfoList 分页获取Frp订单记录
// Author [yourname](https://github.com/yourname)
func (FoService *OrderService) GetFrpOrderInfoList(ctx context.Context, info frpReq.FrpOrderSearch) (list []frp.FrpOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&frp.FrpOrder{})
	var Fos []frp.FrpOrder
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

	err = db.Find(&Fos).Error
	return Fos, total, err
}
func (FoService *OrderService) GetFrpOrderDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	nodeId := make([]map[string]any, 0)

	global.GVA_DB.Table("frp_node").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&nodeId)
	res["nodeId"] = nodeId
	return
}
func (FoService *OrderService) GetFrpOrderPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
