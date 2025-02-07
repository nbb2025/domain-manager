import axios from 'axios'
import router from '../router'
import { ErrorCode, getErrorMessage } from './errorCode'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
api.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => {
    console.error('请求配置错误:', error)
    return Promise.reject(new Error('请求配置错误'))
  }
)

// 响应拦截器
api.interceptors.response.use(
  response => {
    const { code, message, data } = response.data
    
    if (code === 0) {
      return data
    }
    
    // 处理特定错误码
    if (code === ErrorCode.Unauthorized) {
      localStorage.removeItem('token')
      router.push({ name: 'login' })
    }
    
    throw new Error(message || getErrorMessage(code) || '请求失败，请稍后重试')
  },

  error => {
    let errorMessage = '网络连接失败，请检查网络设置'

    if (error.code === 'ECONNABORTED') {
      errorMessage = '请求超时，请稍后重试'
    }

    // 记录错误信息
    console.error('API请求错误:', {
      url: error.config?.url,
      method: error.config?.method,
      message: errorMessage
    })

    return Promise.reject(new Error(errorMessage))
  }
)

// 封装 HTTP 请求方法
export const get = async (url, params) => {
  try {
    return await api.get(url, { params })
  } catch (error) {
    throw error
  }
}

export const post = async (url, data) => {
  try {
    return await api.post(url, data)
  } catch (error) {
    throw error
  }
}

export const put = async (url, data) => {
  try {
    return await api.put(url, data)
  } catch (error) {
    throw error
  }
}

export const del = async (url) => {
  try {
    return await api.delete(url)
  } catch (error) {
    throw error
  }
}

export default api