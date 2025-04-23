package frp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FrpUserRouter struct {}

// InitFrpUserRouter 初始化 Frp用户 路由信息
func (s *FrpUserRouter) InitFrpUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	FuRouter := Router.Group("Fu").Use(middleware.OperationRecord())
	FuRouterWithoutRecord := Router.Group("Fu")
	FuRouterWithoutAuth := PublicRouter.Group("Fu")
	{
		FuRouter.POST("createFrpUser", FuApi.CreateFrpUser)   // 新建Frp用户
		FuRouter.DELETE("deleteFrpUser", FuApi.DeleteFrpUser) // 删除Frp用户
		FuRouter.DELETE("deleteFrpUserByIds", FuApi.DeleteFrpUserByIds) // 批量删除Frp用户
		FuRouter.PUT("updateFrpUser", FuApi.UpdateFrpUser)    // 更新Frp用户
	}
	{
		FuRouterWithoutRecord.GET("findFrpUser", FuApi.FindFrpUser)        // 根据ID获取Frp用户
		FuRouterWithoutRecord.GET("getFrpUserList", FuApi.GetFrpUserList)  // 获取Frp用户列表
	}
	{
	    FuRouterWithoutAuth.GET("getFrpUserPublic", FuApi.GetFrpUserPublic)  // Frp用户开放接口
	}
}
