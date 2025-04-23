package frp

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/frp"
	frpReq "github.com/flipped-aurora/gin-vue-admin/server/model/frp/request"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/frp_node"
)

type PubService struct{}

func (*PubService) Auth(c context.Context, req *frpReq.Auth) error {
	var o frp.FrpOrder
	if err := global.GVA_DB.Where("uuid=?", req.Uuid).First(&o).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return errors.New("KEY不正确或者未拥有次隧道")
	}

	if *o.Flow < 0 {
		return errors.New("流量不足")
	}

	var n frp.FrpNode
	if err := global.GVA_DB.Where("id=?", o.NodeId).First(&n).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return errors.New("KEY不正确或者未拥有次隧道")
	}

	v := map[string]string{
		"uuid":  req.Uuid,
		"runId": req.RunId,
		"t":     fmt.Sprintf("%d", req.T),
	}
	sign := frp_node.GenerateSign(v, *n.Key)
	if sign != req.Sign {
		return errors.New("签名错误")
	}

	o.RunId = &req.RunId
	global.GVA_DB.Save(&o)
	return nil
}

func (*PubService) Auth2(c context.Context, req *frpReq.Auth2) error {
	var o frp.FrpOrder
	if err := global.GVA_DB.Where("uuid=?", req.Uuid).First(&o).Error; err != nil {
		return errors.New("KEY不正确或者未拥有次隧道")
	}

	if fmt.Sprintf(":%d", *o.Port) != req.Port {
		return errors.New("请检查端口是否拥有")
	}

	var n frp.FrpNode
	if err := global.GVA_DB.Where("id=?", o.NodeId).First(&n).Error; err != nil {
		return errors.New("KEY不正确或者未拥有次隧道")
	}

	v := map[string]string{
		"uuid": req.Uuid,
		"port": req.Port,
		"t":    fmt.Sprintf("%d", req.T),
	}

	sign := frp_node.GenerateSign(v, *n.Key)
	if sign != req.Sign {
		return errors.New("签名错误")
	}

	return nil
}

func (*PubService) Speed(c context.Context, req *frpReq.Speed) (int, error) {
	var o frp.FrpOrder
	if err := global.GVA_DB.Where("run_id=?", req.RunId).First(&o).Error; err != nil {
		return 1, errors.New("KEY不正确或者未拥有次隧道")
	}

	return *o.Speed, nil
}
