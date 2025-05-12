<template>
  <header class="header">
    <div class="logo">
      <router-link to="/">过敏管理系统</router-link>
    </div>
    <nav class="nav">
      <router-link to="/allergy/records">过敏记录</router-link>
      <router-link to="/allergy/analysis">过敏分析</router-link>
      <router-link to="/allergy/prevention">预防建议</router-link>
    </nav>
    <div class="user-info" v-if="userInfo">
      <div class="avatar" @click="showUserMenu = !showUserMenu">
        <img :src="userInfo.avatar || defaultAvatar" :alt="userInfo.username">
        <div class="user-menu" v-if="showUserMenu">
          <div class="menu-item" @click="$router.push('/user/profile')">个人资料</div>
          <div class="menu-item" @click="handleLogout">退出登录</div>
        </div>
      </div>
    </div>
    <div class="auth-buttons" v-else>
      <router-link to="/login" class="btn">登录</router-link>
      <router-link to="/register" class="btn btn-primary">注册</router-link>
    </div>
  </header>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { userApi } from '@/api/user'

export default {
  name: 'Header',
  setup() {
    const router = useRouter()
    const userInfo = ref(null)
    const showUserMenu = ref(false)
    const defaultAvatar = '/assets/default-avatar.png'

    const fetchUserInfo = async () => {
      try {
        const res = await userApi.getUserInfo()
        userInfo.value = res.data
      } catch (error) {
        console.error('获取用户信息失败:', error)
      }
    }

    const handleLogout = () => {
      localStorage.removeItem('token')
      userInfo.value = null
      router.push('/login')
    }

    onMounted(() => {
      if (localStorage.getItem('token')) {
        fetchUserInfo()
      }
    })

    return {
      userInfo,
      showUserMenu,
      defaultAvatar,
      handleLogout
    }
  }
}
</script>

<style scoped>
.header {
  background: white;
  padding: 15px 30px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.logo a {
  font-size: 20px;
  font-weight: bold;
  color: #333;
  text-decoration: none;
}

.nav {
  display: flex;
  gap: 20px;
}

.nav a {
  color: #666;
  text-decoration: none;
  padding: 5px 10px;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.nav a:hover {
  background-color: #f5f5f5;
}

.nav a.router-link-active {
  color: #409eff;
  font-weight: bold;
}

.user-info {
  position: relative;
}

.avatar {
  cursor: pointer;
  position: relative;
}

.avatar img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.user-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: white;
  border-radius: 4px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  padding: 5px 0;
  margin-top: 10px;
  min-width: 120px;
}

.menu-item {
  padding: 8px 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.menu-item:hover {
  background-color: #f5f5f5;
}

.auth-buttons {
  display: flex;
  gap: 10px;
}

.btn {
  padding: 8px 16px;
  border-radius: 4px;
  text-decoration: none;
  transition: all 0.3s;
}

.btn:not(.btn-primary) {
  color: #666;
}

.btn-primary {
  background-color: #409eff;
  color: white;
}

.btn:hover {
  opacity: 0.8;
}
</style> 