package frp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FrpNodeRouter struct {}

// InitFrpNodeRouter 初始化 frp节点 路由信息
func (s *FrpNodeRouter) InitFrpNodeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	FnRouter := Router.Group("Fn").Use(middleware.OperationRecord())
	FnRouterWithoutRecord := Router.Group("Fn")
	FnRouterWithoutAuth := PublicRouter.Group("Fn")
	{
		FnRouter.POST("createFrpNode", FnApi.CreateFrpNode)   // 新建frp节点
		FnRouter.DELETE("deleteFrpNode", FnApi.DeleteFrpNode) // 删除frp节点
		FnRouter.DELETE("deleteFrpNodeByIds", FnApi.DeleteFrpNodeByIds) // 批量删除frp节点
		FnRouter.PUT("updateFrpNode", FnApi.UpdateFrpNode)    // 更新frp节点
	}
	{
		FnRouterWithoutRecord.GET("findFrpNode", FnApi.FindFrpNode)        // 根据ID获取frp节点
		FnRouterWithoutRecord.GET("getFrpNodeList", FnApi.GetFrpNodeList)  // 获取frp节点列表
	}
	{
	    FnRouterWithoutAuth.GET("getFrpNodePublic", FnApi.GetFrpNodePublic)  // frp节点开放接口
	}
}
