import request from '@/utils/request'

// 登录
export const login = async (req) => {
  const { data: res } = await request.post('/user/login', req)
  return res
}

// 用户列表
export const getUserList = async (req) => {
  const { data: res } = await request.post('/user/list', req)
  return res
}

// 用户详情
export const getUserDetail = async (req) => {
  const { data: res } = await request.post('/user/detail', req)
  return res
}

// 创建用户
export const createUser = async (req) => {
  const { data: res } = await request.post('/user/create', req)
  return res
}

// 更新用户
export const updateUser = async (req) => {
  const { data: res } = await request.post('/user/update', req)
  return res
}

// 删除用户
export const logoutUser = async (req) => {
  const { data: res } = await request.post('/user/logout', req)
  return res
}

// 恢复用户
export const recoverUser = async (req) => {
  const { data: res } = await request.post('/user/recover', req)
  return res
}

// 重置密码
export const resetUserPassword = async (req) => {
  const { data: res } = await request.post('/user/password/reset', req)
  return res
}

// 修改密码
export const changeUserPassword = async (req) => {
  const { data: res } = await request.post('/user/password/change', req)
  return res
}
