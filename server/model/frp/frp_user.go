
// 自动生成模板FrpUser
package frp
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Frp用户 结构体  FrpUser
type FrpUser struct {
    global.GVA_MODEL
  UpId  *string `json:"upId" form:"upId" gorm:"column:up_id;" binding:"required"`  //上游ID
  Uuid  *string `json:"uuid" form:"uuid" gorm:"column:uuid;"`  //UUID
  Desc  *string `json:"desc" form:"desc" gorm:"column:desc;"`  //描述
}


// TableName Frp用户 FrpUser自定义表名 frp_user
func (FrpUser) TableName() string {
    return "frp_user"
}





