package frp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/frp"
	frpReq "github.com/flipped-aurora/gin-vue-admin/server/model/frp/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderApi struct{}

// CreateFrpOrder 创建Frp订单
// @Tags FrpOrder
// @Summary 创建Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body frp.FrpOrder true "创建Frp订单"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Fo/createFrpOrder [post]
func (FoApi *OrderApi) CreateFrpOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var Fo frp.FrpOrder
	err := c.ShouldBindJSON(&Fo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FoService.CreateFrpOrder(ctx, &Fo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteFrpOrder 删除Frp订单
// @Tags FrpOrder
// @Summary 删除Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body frp.FrpOrder true "删除Frp订单"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Fo/deleteFrpOrder [delete]
func (FoApi *OrderApi) DeleteFrpOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := FoService.DeleteFrpOrder(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteFrpOrderByIds 批量删除Frp订单
// @Tags FrpOrder
// @Summary 批量删除Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Fo/deleteFrpOrderByIds [delete]
func (FoApi *OrderApi) DeleteFrpOrderByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := FoService.DeleteFrpOrderByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateFrpOrder 更新Frp订单
// @Tags FrpOrder
// @Summary 更新Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body frp.FrpOrder true "更新Frp订单"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Fo/updateFrpOrder [put]
func (FoApi *OrderApi) UpdateFrpOrder(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var Fo frp.FrpOrder
	err := c.ShouldBindJSON(&Fo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FoService.UpdateFrpOrder(ctx, Fo)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindFrpOrder 用id查询Frp订单
// @Tags FrpOrder
// @Summary 用id查询Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询Frp订单"
// @Success 200 {object} response.Response{data=frp.FrpOrder,msg=string} "查询成功"
// @Router /Fo/findFrpOrder [get]
func (FoApi *OrderApi) FindFrpOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reFo, err := FoService.GetFrpOrder(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reFo, c)
}

// GetFrpOrderList 分页获取Frp订单列表
// @Tags FrpOrder
// @Summary 分页获取Frp订单列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query frpReq.FrpOrderSearch true "分页获取Frp订单列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Fo/getFrpOrderList [get]
func (FoApi *OrderApi) GetFrpOrderList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo frpReq.FrpOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := FoService.GetFrpOrderInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetFrpOrderDataSource 获取FrpOrder的数据源
// @Tags FrpOrder
// @Summary 获取FrpOrder的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /Fo/getFrpOrderDataSource [get]
func (FoApi *OrderApi) GetFrpOrderDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := FoService.GetFrpOrderDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetFrpOrderPublic 不需要鉴权的Frp订单接口
// @Tags FrpOrder
// @Summary 不需要鉴权的Frp订单接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Fo/getFrpOrderPublic [get]
func (FoApi *OrderApi) GetFrpOrderPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	FoService.GetFrpOrderPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的Frp订单接口信息",
	}, "获取成功", c)
}
