package frp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	frpReq "github.com/flipped-aurora/gin-vue-admin/server/model/frp/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PubApi struct{}

func (*PubApi) Auth(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var req frpReq.Auth
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FpService.Auth(ctx, &req)
	if err != nil {
		global.GVA_LOG.Error("验证失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("登录成功", c)
}

func (*PubApi) Auth2(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var req frpReq.Auth2
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FpService.Auth2(ctx, &req)
	if err != nil {
		global.GVA_LOG.Error("验证失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("登录成功", c)
}

func (*PubApi) Speed(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var req frpReq.Speed
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	limit, err := FpService.Speed(ctx, &req)
	if err != nil {
		global.GVA_LOG.Error("验证失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{
		"limit": limit,
	}, c)
}
