package frp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/frp"
	frpReq "github.com/flipped-aurora/gin-vue-admin/server/model/frp/request"
	"github.com/google/uuid"
)

type UserService struct{}

// CreateFrpUser 创建Frp用户记录
// Author [yourname](https://github.com/yourname)
func (FuService *UserService) CreateFrpUser(ctx context.Context, Fu *frp.FrpUser) (err error) {
	_uuid := uuid.NewString()
	Fu.Uuid = &_uuid
	err = global.GVA_DB.Create(Fu).Error
	return err
}

// DeleteFrpUser 删除Frp用户记录
// Author [yourname](https://github.com/yourname)
func (FuService *UserService) DeleteFrpUser(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&frp.FrpUser{}, "id = ?", ID).Error
	return err
}

// DeleteFrpUserByIds 批量删除Frp用户记录
// Author [yourname](https://github.com/yourname)
func (FuService *UserService) DeleteFrpUserByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]frp.FrpUser{}, "id in ?", IDs).Error
	return err
}

// UpdateFrpUser 更新Frp用户记录
// Author [yourname](https://github.com/yourname)
func (FuService *UserService) UpdateFrpUser(ctx context.Context, Fu frp.FrpUser) (err error) {
	err = global.GVA_DB.Model(&frp.FrpUser{}).Where("id = ?", Fu.ID).Updates(&Fu).Error
	return err
}

// GetFrpUser 根据ID获取Frp用户记录
// Author [yourname](https://github.com/yourname)
func (FuService *UserService) GetFrpUser(ctx context.Context, ID string) (Fu frp.FrpUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Fu).Error
	return
}

// GetFrpUserInfoList 分页获取Frp用户记录
// Author [yourname](https://github.com/yourname)
func (FuService *UserService) GetFrpUserInfoList(ctx context.Context, info frpReq.FrpUserSearch) (list []frp.FrpUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&frp.FrpUser{})
	var Fus []frp.FrpUser
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

	err = db.Find(&Fus).Error
	return Fus, total, err
}
func (FuService *UserService) GetFrpUserPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
