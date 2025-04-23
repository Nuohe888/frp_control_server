package frp

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	NodeApi
	UserApi
	OrderApi
	PubApi
}

var (
	FnService = service.ServiceGroupApp.FrpServiceGroup.NodeService
	FuService = service.ServiceGroupApp.FrpServiceGroup.UserService
	FoService = service.ServiceGroupApp.FrpServiceGroup.OrderService
	FpService = service.ServiceGroupApp.FrpServiceGroup.PubService
)
