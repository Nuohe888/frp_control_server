import service from '@/utils/request'
// @Tags FrpNode
// @Summary 创建frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FrpNode true "创建frp节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Fn/createFrpNode [post]
export const createFrpNode = (data) => {
  return service({
    url: '/Fn/createFrpNode',
    method: 'post',
    data
  })
}

// @Tags FrpNode
// @Summary 删除frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FrpNode true "删除frp节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Fn/deleteFrpNode [delete]
export const deleteFrpNode = (params) => {
  return service({
    url: '/Fn/deleteFrpNode',
    method: 'delete',
    params
  })
}

// @Tags FrpNode
// @Summary 批量删除frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除frp节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Fn/deleteFrpNode [delete]
export const deleteFrpNodeByIds = (params) => {
  return service({
    url: '/Fn/deleteFrpNodeByIds',
    method: 'delete',
    params
  })
}

// @Tags FrpNode
// @Summary 更新frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FrpNode true "更新frp节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Fn/updateFrpNode [put]
export const updateFrpNode = (data) => {
  return service({
    url: '/Fn/updateFrpNode',
    method: 'put',
    data
  })
}

// @Tags FrpNode
// @Summary 用id查询frp节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.FrpNode true "用id查询frp节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Fn/findFrpNode [get]
export const findFrpNode = (params) => {
  return service({
    url: '/Fn/findFrpNode',
    method: 'get',
    params
  })
}

// @Tags FrpNode
// @Summary 分页获取frp节点列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取frp节点列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Fn/getFrpNodeList [get]
export const getFrpNodeList = (params) => {
  return service({
    url: '/Fn/getFrpNodeList',
    method: 'get',
    params
  })
}

// @Tags FrpNode
// @Summary 不需要鉴权的frp节点接口
// @Accept application/json
// @Produce application/json
// @Param data query frpReq.FrpNodeSearch true "分页获取frp节点列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Fn/getFrpNodePublic [get]
export const getFrpNodePublic = () => {
  return service({
    url: '/Fn/getFrpNodePublic',
    method: 'get',
  })
}
