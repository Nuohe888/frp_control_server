
// 自动生成模板FrpNode
package frp
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// frp节点 结构体  FrpNode
type FrpNode struct {
    global.GVA_MODEL
  Desc  *string `json:"desc" form:"desc" gorm:"column:desc;"`  //描述
  Name  *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`  //名称标识
  Key  *string `json:"key" form:"key" gorm:"column:key;" binding:"required"`  //节点密钥
  Status  *string `json:"status" form:"status" gorm:"default:0;column:status;" binding:"required"`  //状态
}


// TableName frp节点 FrpNode自定义表名 frp_node
func (FrpNode) TableName() string {
    return "frp_node"
}





