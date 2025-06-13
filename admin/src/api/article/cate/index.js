import request from '@/utils/request'

// 文章分类列表
export const getArticleCateList = async (req) => {
  const { data: res } = await request.post('/article/cate/list', req)
  return res
}

// 文章分类详情
export const getArticleCateDetail = async (req) => {
  const { data: res } = await request.post('/article/cate/detail', req)
  return res
}

// 创建文章分类
export const createArticleCate = async (req) => {
  const { data: res } = await request.post('/article/cate/create', req)
  return res
}

// 更新文章分类
export const updateArticleCate = async (req) => {
  const { data: res } = await request.post('/article/cate/update', req)
  return res
}

// 删除文章分类
export const deleteArticleCate = async (req) => {
  const { data: res } = await request.post('/article/cate/delete', req)
  return res
}
