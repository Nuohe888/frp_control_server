import service from '@/utils/request'
// @Tags FrpUser
// @Summary 创建Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FrpUser true "创建Frp用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Fu/createFrpUser [post]
export const createFrpUser = (data) => {
  return service({
    url: '/Fu/createFrpUser',
    method: 'post',
    data
  })
}

// @Tags FrpUser
// @Summary 删除Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FrpUser true "删除Frp用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Fu/deleteFrpUser [delete]
export const deleteFrpUser = (params) => {
  return service({
    url: '/Fu/deleteFrpUser',
    method: 'delete',
    params
  })
}

// @Tags FrpUser
// @Summary 批量删除Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Frp用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Fu/deleteFrpUser [delete]
export const deleteFrpUserByIds = (params) => {
  return service({
    url: '/Fu/deleteFrpUserByIds',
    method: 'delete',
    params
  })
}

// @Tags FrpUser
// @Summary 更新Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.FrpUser true "更新Frp用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Fu/updateFrpUser [put]
export const updateFrpUser = (data) => {
  return service({
    url: '/Fu/updateFrpUser',
    method: 'put',
    data
  })
}

// @Tags FrpUser
// @Summary 用id查询Frp用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.FrpUser true "用id查询Frp用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Fu/findFrpUser [get]
export const findFrpUser = (params) => {
  return service({
    url: '/Fu/findFrpUser',
    method: 'get',
    params
  })
}

// @Tags FrpUser
// @Summary 分页获取Frp用户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Frp用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Fu/getFrpUserList [get]
export const getFrpUserList = (params) => {
  return service({
    url: '/Fu/getFrpUserList',
    method: 'get',
    params
  })
}

// @Tags FrpUser
// @Summary 不需要鉴权的Frp用户接口
// @Accept application/json
// @Produce application/json
// @Param data query frpReq.FrpUserSearch true "分页获取Frp用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Fu/getFrpUserPublic [get]
export const getFrpUserPublic = () => {
  return service({
    url: '/Fu/getFrpUserPublic',
    method: 'get',
  })
}
