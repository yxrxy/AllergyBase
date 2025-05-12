import axios from 'axios'
import router from '@/router'

const request = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080',
  timeout: 5000
})

// 请求拦截器
request.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  error => {
    console.error('请求错误:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code !== '10000') {
      console.error('接口错误:', res.message)
      return Promise.reject(new Error(res.message || '未知错误'))
    }
    return res
  },
  error => {
    if (error.response) {
      const { status } = error.response
      if (status === 401) {
        // token 过期或无效
        localStorage.removeItem('token')
        router.push('/login')
      }
      console.error('响应错误:', error.response.data.message)
      return Promise.reject(new Error(error.response.data.message || '请求失败'))
    }
    console.error('网络错误:', error)
    return Promise.reject(error)
  }
)

export default request 