import request from '@/utils/request'

export const userApi = {
  // 用户登录
  login(data) {
    return request({
      url: '/api/v1/user/login',
      method: 'post',
      data
    })
  },

  // 用户注册
  register(data) {
    return request({
      url: '/api/v1/user/register',
      method: 'post',
      data
    })
  },

  // 获取用户信息
  getUserInfo() {
    return request({
      url: '/api/v1/user/info',
      method: 'get'
    })
  },

  // 上传头像
  uploadAvatar(data) {
    return request({
      url: '/api/v1/user/avatar',
      method: 'post',
      data,
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  }
} 