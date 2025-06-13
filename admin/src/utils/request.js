import axios from 'axios'
import { router } from '@/router'
import config from '@/config/config'

const request = axios.create({
  baseURL: config.api_base_url,
  timeout: 60000,
})

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    // 配置请求头
    config.headers.Authorization = localStorage.getItem('token') || ''
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

// 响应拦截器
request.interceptors.response.use(
  (res) => {
    // 响应数据处理
    if (res.data.code === 401) {
      ElMessageBox.confirm(`${res.data.msg}，请重新登录`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(() => {
          localStorage.removeItem('token')
          router.push('/login')
        })
        .catch(() => {})
    } else if (res.data.code !== 200) {
      ElMessage({
        message: res.data.msg,
        grouping: true,
        type: 'error',
      })
    }
    return res
  },
  (error) => {
    if (error.code === 'ECONNABORTED') {
      // 超时错误
      ElMessage.error('请求超时，请稍后重试')
    } else if (!error.response) {
      // 网络错误（比如断网、DNS 无响应等）
      ElMessage.error('网络错误，请检查您的网络连接')
    } else {
      // 其他响应错误
      ElMessage.error(`${error.response.status}: ${error.response.data.error || '请求失败'}`)
    }
    return Promise.reject(error)
  },
)

export default request
