// 封装通用请求
import axios from 'axios'
import router from '../router'
axios.defaults.timeout = 5000 // 超时时间：5s
axios.defaults.withCredentials = true// 允许跨域
// Content-Type 响应头
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'
// 访问基础url
axios.defaults.baseURL = 'http://localhost:8083'

// 响应拦截器
axios.interceptors.response.use(
  response => {
    // 如果response里面的status是200，说明访问到接口了，否则失败
    if (response.status === 200) {
      // Promise：异步框架
      return Promise.resolve(response)
    } else {
      return Promise.reject(response)
    }
  },
  error => {
    if (error.response.status) {
      // 根据访问失败返回的状态码，分别做不同的处理
      switch (error.response.status) {
        case 401: // 未登录
          router.replace({
            path: '/',
            query: {
              redirect: router.currentRoute.fullPath // 存之前访问地址
            }
          })
          break
        case 404: // not found
          break
      }
      return Promise.reject(error.response)
    }
  }
)

/**
 * 封装get请求
 */
export function get (url, params = {}) {
  return new Promise((resolve, reject) => {
    axios.get(url, {params: params})
      .then(response => {
        resolve(response.data)
      })
      .catch(err => {
        reject(err)
      })
  })
}

/**
 * 封装post请求
 */
export function post (url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.post(url, data)
      .then(response => {
        resolve(response.data)
      })
      .catch(err => {
        reject(err)
      })
  })
}
