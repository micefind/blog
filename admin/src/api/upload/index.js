import request from '@/utils/request'

// 上传图片
export const uploadImage = async (req) => {
  const { data: res } = await request.post('/upload/image', req)
  return res
}
