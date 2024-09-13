import axios, {
  AxiosInstance,
  InternalAxiosRequestConfig,
  AxiosResponse,
} from "axios"

// 定义接口响应数据的类型
interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

// 创建 axios 实例
const service: AxiosInstance = axios.create({
  baseURL: "https://api.example.com", // 替换为你的 API 根路径
  timeout: 5000, // 请求超时时间
  headers: {
    "Content-Type": "application/json;charset=utf-8",
  },
})

// 请求拦截器
service.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // 在请求发送之前做些什么，比如设置token
    const token = uni.getStorageSync("token") // 从本地存储获取token
    if (token && config.headers) {
      config.headers["Authorization"] = "Bearer " + token // 在请求头中带上token
    }
    return config
  },
  (error) => {
    // 处理请求错误
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse<ApiResponse>) => {
    // 处理响应数据
    const res = response.data
    if (res.code !== 200) {
      uni.showToast({
        title: res.message || "Error",
        icon: "none",
        duration: 2000,
      })
      return Promise.reject(new Error(res.message || "Error"))
    } else {
      return res.data // 直接返回数据部分
    }
  },
  (error) => {
    // 处理响应错误
    uni.showToast({
      title: error.message,
      icon: "none",
      duration: 2000,
    })
    return Promise.reject(error)
  }
)

export default service
