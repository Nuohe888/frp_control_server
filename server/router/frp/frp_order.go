package frp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FrpOrderRouter struct {}

// InitFrpOrderRouter 初始化 Frp订单 路由信息
func (s *FrpOrderRouter) InitFrpOrderRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	FoRouter := Router.Group("Fo").Use(middleware.OperationRecord())
	FoRouterWithoutRecord := Router.Group("Fo")
	FoRouterWithoutAuth := PublicRouter.Group("Fo")
	{
		FoRouter.POST("createFrpOrder", FoApi.CreateFrpOrder)   // 新建Frp订单
		FoRouter.DELETE("deleteFrpOrder", FoApi.DeleteFrpOrder) // 删除Frp订单
		FoRouter.DELETE("deleteFrpOrderByIds", FoApi.DeleteFrpOrderByIds) // 批量删除Frp订单
		FoRouter.PUT("updateFrpOrder", FoApi.UpdateFrpOrder)    // 更新Frp订单
	}
	{
		FoRouterWithoutRecord.GET("findFrpOrder", FoApi.FindFrpOrder)        // 根据ID获取Frp订单
		FoRouterWithoutRecord.GET("getFrpOrderList", FoApi.GetFrpOrderList)  // 获取Frp订单列表
	}
	{
	    FoRouterWithoutAuth.GET("getFrpOrderDataSource", FoApi.GetFrpOrderDataSource)  // 获取Frp订单数据源
	    FoRouterWithoutAuth.GET("getFrpOrderPublic", FoApi.GetFrpOrderPublic)  // Frp订单开放接口
	}
}
