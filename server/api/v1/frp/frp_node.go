package frp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/frp"
	frpReq "github.com/flipped-aurora/gin-vue-admin/server/model/frp/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NodeApi struct{}

// CreateFrpNode 创建frp节点
// @Tags FrpNode
// @Summary 创建frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body frp.FrpNode true "创建frp节点"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Fn/createFrpNode [post]
func (FnApi *NodeApi) CreateFrpNode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var Fn frp.FrpNode
	err := c.ShouldBindJSON(&Fn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FnService.CreateFrpNode(ctx, &Fn)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteFrpNode 删除frp节点
// @Tags FrpNode
// @Summary 删除frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body frp.FrpNode true "删除frp节点"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Fn/deleteFrpNode [delete]
func (FnApi *NodeApi) DeleteFrpNode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := FnService.DeleteFrpNode(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteFrpNodeByIds 批量删除frp节点
// @Tags FrpNode
// @Summary 批量删除frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Fn/deleteFrpNodeByIds [delete]
func (FnApi *NodeApi) DeleteFrpNodeByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := FnService.DeleteFrpNodeByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateFrpNode 更新frp节点
// @Tags FrpNode
// @Summary 更新frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body frp.FrpNode true "更新frp节点"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Fn/updateFrpNode [put]
func (FnApi *NodeApi) UpdateFrpNode(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var Fn frp.FrpNode
	err := c.ShouldBindJSON(&Fn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FnService.UpdateFrpNode(ctx, Fn)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindFrpNode 用id查询frp节点
// @Tags FrpNode
// @Summary 用id查询frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询frp节点"
// @Success 200 {object} response.Response{data=frp.FrpNode,msg=string} "查询成功"
// @Router /Fn/findFrpNode [get]
func (FnApi *NodeApi) FindFrpNode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reFn, err := FnService.GetFrpNode(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reFn, c)
}

// GetFrpNodeList 分页获取frp节点列表
// @Tags FrpNode
// @Summary 分页获取frp节点列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query frpReq.FrpNodeSearch true "分页获取frp节点列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Fn/getFrpNodeList [get]
func (FnApi *NodeApi) GetFrpNodeList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo frpReq.FrpNodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := FnService.GetFrpNodeInfoList(ctx, pageInfo)
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

// GetFrpNodePublic 不需要鉴权的frp节点接口
// @Tags FrpNode
// @Summary 不需要鉴权的frp节点接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Fn/getFrpNodePublic [get]
func (FnApi *NodeApi) GetFrpNodePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	FnService.GetFrpNodePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的frp节点接口信息",
	}, "获取成功", c)
}
