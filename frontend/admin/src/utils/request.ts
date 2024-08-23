import axios from "axios"
import { netConfig } from "../config/net.config"

const request = axios.create({
  baseURL: netConfig.baseURL,
  timeout: netConfig.timeout,
})

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    if (config.url !== "user/login") {
      // 判断请求是否是登录接口
      config.headers.Authorization = sessionStorage.getItem("token") // 如果不是登录接口，就给请求头里面设置token
    }
    return config // 返回这个配置对象，如果没有返回，这个请求就不会发送出去
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (res) => {
    // 响应数据处理
    if (res.data.status !== 200) {
      ElMessage({
        message: res.data.message,
        grouping: true,
        type: "error",
      })
    }
    return res
  },
  (error) => {
    // 响应错误处理
    if (error.response) {
      // 请求已发出，但服务器响应的状态码不在2xx范围内
      ElMessage({
        message: `Error ${error.response.status}: ${error.response.data.message}`,
        grouping: true,
        type: "error",
      })
    } else {
      // 请求超时或网络错误
      ElMessage({
        message: error.message,
        grouping: true,
        type: "error",
      })
    }
    // 将未处理的异常往外抛
    return Promise.reject(error)
  }
)

export default request
