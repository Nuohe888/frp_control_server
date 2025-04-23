package frp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/frp"
	frpReq "github.com/flipped-aurora/gin-vue-admin/server/model/frp/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

// CreateFrpUser 创建Frp用户
// @Tags FrpUser
// @Summary 创建Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body frp.FrpUser true "创建Frp用户"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Fu/createFrpUser [post]
func (FuApi *UserApi) CreateFrpUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var Fu frp.FrpUser
	err := c.ShouldBindJSON(&Fu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FuService.CreateFrpUser(ctx, &Fu)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteFrpUser 删除Frp用户
// @Tags FrpUser
// @Summary 删除Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body frp.FrpUser true "删除Frp用户"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Fu/deleteFrpUser [delete]
func (FuApi *UserApi) DeleteFrpUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := FuService.DeleteFrpUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteFrpUserByIds 批量删除Frp用户
// @Tags FrpUser
// @Summary 批量删除Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Fu/deleteFrpUserByIds [delete]
func (FuApi *UserApi) DeleteFrpUserByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := FuService.DeleteFrpUserByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateFrpUser 更新Frp用户
// @Tags FrpUser
// @Summary 更新Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body frp.FrpUser true "更新Frp用户"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Fu/updateFrpUser [put]
func (FuApi *UserApi) UpdateFrpUser(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var Fu frp.FrpUser
	err := c.ShouldBindJSON(&Fu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FuService.UpdateFrpUser(ctx, Fu)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindFrpUser 用id查询Frp用户
// @Tags FrpUser
// @Summary 用id查询Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询Frp用户"
// @Success 200 {object} response.Response{data=frp.FrpUser,msg=string} "查询成功"
// @Router /Fu/findFrpUser [get]
func (FuApi *UserApi) FindFrpUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reFu, err := FuService.GetFrpUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reFu, c)
}

// GetFrpUserList 分页获取Frp用户列表
// @Tags FrpUser
// @Summary 分页获取Frp用户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query frpReq.FrpUserSearch true "分页获取Frp用户列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Fu/getFrpUserList [get]
func (FuApi *UserApi) GetFrpUserList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo frpReq.FrpUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := FuService.GetFrpUserInfoList(ctx, pageInfo)
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

// GetFrpUserPublic 不需要鉴权的Frp用户接口
// @Tags FrpUser
// @Summary 不需要鉴权的Frp用户接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Fu/getFrpUserPublic [get]
func (FuApi *UserApi) GetFrpUserPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	FuService.GetFrpUserPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的Frp用户接口信息",
	}, "获取成功", c)
}
