import request from '@/utils/request'

// 文章列表
export const getArticleList = async (req) => {
  const { data: res } = await request.post('/article/list', req)
  return res
}

// 文章详情
export const getArticleDetail = async (req) => {
  const { data: res } = await request.post('/article/detail', req)
  return res
}

// 创建文章
export const createArticle = async (req) => {
  const { data: res } = await request.post('/article/create', req)
  return res
}

// 更新文章
export const updateArticle = async (req) => {
  const { data: res } = await request.post('/article/update', req)
  return res
}

// 删除文章
export const deleteArticle = async (req) => {
  const { data: res } = await request.post('/article/delete', req)
  return res
}

// 恢复文章
export const recoverArticle = async (req) => {
  const { data: res } = await request.post('/article/recover', req)
  return res
}
