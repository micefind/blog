import request from '@/utils/request'

// 文章标签列表
export const getArticleTagList = async (req) => {
  const { data: res } = await request.post('/article/tag/list', req)
  return res
}

// 文章标签详情
export const getArticleTagDetail = async (req) => {
  const { data: res } = await request.post('/article/tag/detail', req)
  return res
}

// 创建文章标签
export const createArticleTag = async (req) => {
  const { data: res } = await request.post('/article/tag/create', req)
  return res
}

// 更新文章标签
export const updateArticleTag = async (req) => {
  const { data: res } = await request.post('/article/tag/update', req)
  return res
}

// 删除文章标签
export const deleteArticleTag = async (req) => {
  const { data: res } = await request.post('/article/tag/delete', req)
  return res
}
