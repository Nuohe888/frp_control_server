package frp

import (
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

// InitFrpNodeRouter 初始化 frp节点 路由信息
func (s *ApiRouter) InitFrpApiRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	PublicRouter.POST("/api/node/auth", FaApi.Auth)
	PublicRouter.POST("/api/node/auth2", FaApi.Auth2)
	PublicRouter.POST("/api/node/speed", FaApi.Speed)
}
