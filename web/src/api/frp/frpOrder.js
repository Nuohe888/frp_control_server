import service from '@/utils/request'
// @Tags FrpOrder
// @Summary 创建Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FrpOrder true "创建Frp订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Fo/createFrpOrder [post]
export const createFrpOrder = (data) => {
  return service({
    url: '/Fo/createFrpOrder',
    method: 'post',
    data
  })
}

// @Tags FrpOrder
// @Summary 删除Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FrpOrder true "删除Frp订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Fo/deleteFrpOrder [delete]
export const deleteFrpOrder = (params) => {
  return service({
    url: '/Fo/deleteFrpOrder',
    method: 'delete',
    params
  })
}

// @Tags FrpOrder
// @Summary 批量删除Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Frp订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Fo/deleteFrpOrder [delete]
export const deleteFrpOrderByIds = (params) => {
  return service({
    url: '/Fo/deleteFrpOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags FrpOrder
// @Summary 更新Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FrpOrder true "更新Frp订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Fo/updateFrpOrder [put]
export const updateFrpOrder = (data) => {
  return service({
    url: '/Fo/updateFrpOrder',
    method: 'put',
    data
  })
}

// @Tags FrpOrder
// @Summary 用id查询Frp订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.FrpOrder true "用id查询Frp订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Fo/findFrpOrder [get]
export const findFrpOrder = (params) => {
  return service({
    url: '/Fo/findFrpOrder',
    method: 'get',
    params
  })
}

// @Tags FrpOrder
// @Summary 分页获取Frp订单列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Frp订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Fo/getFrpOrderList [get]
export const getFrpOrderList = (params) => {
  return service({
    url: '/Fo/getFrpOrderList',
    method: 'get',
    params
  })
}
// @Tags FrpOrder
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Fo/findFrpOrderDataSource [get]
export const getFrpOrderDataSource = () => {
  return service({
    url: '/Fo/getFrpOrderDataSource',
    method: 'get',
  })
}

// @Tags FrpOrder
// @Summary 不需要鉴权的Frp订单接口
// @Accept application/json
// @Produce application/json
// @Param data query frpReq.FrpOrderSearch true "分页获取Frp订单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Fo/getFrpOrderPublic [get]
export const getFrpOrderPublic = () => {
  return service({
    url: '/Fo/getFrpOrderPublic',
    method: 'get',
  })
}
