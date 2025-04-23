// 自动生成模板FrpOrder
package frp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Frp订单 结构体  FrpOrder
type FrpOrder struct {
	global.GVA_MODEL
	Uuid   *string `json:"uuid" form:"uuid" gorm:"column:uuid;" binding:"required"`        //内部UUID
	Speed  *int    `json:"speed" form:"speed" gorm:"column:speed;" binding:"required"`     //限速
	Flow   *int    `json:"flow" form:"flow" gorm:"column:flow;" binding:"required"`        //流量
	NodeId *int    `json:"nodeId" form:"nodeId" gorm:"column:node_id;" binding:"required"` //节点ID
	Port   *int    `json:"port" form:"port" gorm:"column:port;" binding:"required"`        //端口
	RunId  *string `json:"runId" form:"runId" gorm:"column:run_id;"`
}

// TableName Frp订单 FrpOrder自定义表名 frp_order
func (FrpOrder) TableName() string {
	return "frp_order"
}
