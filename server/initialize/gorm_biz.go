package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/frp"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(frp.FrpNode{}, frp.FrpUser{}, frp.FrpOrder{})
	if err != nil {
		return err
	}
	return nil
}
