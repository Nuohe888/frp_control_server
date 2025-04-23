package frp

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	FrpNodeRouter
	FrpUserRouter
	FrpOrderRouter
	ApiRouter
}

var (
	FnApi = api.ApiGroupApp.FrpApiGroup.NodeApi
	FuApi = api.ApiGroupApp.FrpApiGroup.UserApi
	FoApi = api.ApiGroupApp.FrpApiGroup.OrderApi
	FaApi = api.ApiGroupApp.FrpApiGroup.PubApi
)
