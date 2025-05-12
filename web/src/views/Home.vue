<template>
  <div class="home">
    <Header />
    <div class="main-content">
      <aside class="sidebar">
        <nav>
          <router-link to="/allergy/records">过敏记录</router-link>
          <router-link to="/allergy/analysis">过敏分析</router-link>
          <router-link to="/allergy/prevention">预防建议</router-link>
        </nav>
      </aside>
      <main class="content">
        <h1>欢迎使用过敏管理系统</h1>
        <div class="quick-actions">
          <div class="action-card" @click="$router.push('/allergy/records/new')">
            <i class="icon-add"></i>
            <span>添加过敏记录</span>
          </div>
          <div class="action-card" @click="$router.push('/allergy/analysis')">
            <i class="icon-analysis"></i>
            <span>查看分析报告</span>
          </div>
          <div class="action-card" @click="$router.push('/allergy/prevention')">
            <i class="icon-prevention"></i>
            <span>获取预防建议</span>
          </div>
        </div>
        <div class="recent-records" v-if="recentRecords.length">
          <h2>最近记录</h2>
          <ul>
            <li v-for="record in recentRecords" :key="record.id">
              <span class="date">{{ formatDate(record.date) }}</span>
              <span class="type">{{ record.allergyType }}</span>
              <span class="severity">严重程度: {{ record.severity }}</span>
            </li>
          </ul>
        </div>
      </main>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import Header from '@/components/layout/Header.vue'
import { allergyApi } from '@/api/allergy'

export default {
  name: 'Home',
  components: {
    Header
  },
  setup() {
    const recentRecords = ref([])

    const formatDate = (date) => {
      return new Date(date).toLocaleDateString()
    }

    const fetchRecentRecords = async () => {
      try {
        const res = await allergyApi.getRecentRecords()
        recentRecords.value = res.data
      } catch (error) {
        console.error('获取最近记录失败:', error)
      }
    }

    onMounted(() => {
      fetchRecentRecords()
    })

    return {
      recentRecords,
      formatDate
    }
  }
}
</script>

<style scoped>
.home {
  min-height: 100vh;
  background-color: #f5f5f5;
}

.main-content {
  display: flex;
  padding: 20px;
  gap: 20px;
}

.sidebar {
  width: 200px;
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.sidebar nav {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.sidebar a {
  padding: 10px;
  color: #333;
  text-decoration: none;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.sidebar a:hover {
  background-color: #f0f0f0;
}

.content {
  flex: 1;
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

h1 {
  margin-bottom: 30px;
  color: #333;
}

.quick-actions {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.action-card {
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.3s, box-shadow 0.3s;
  text-align: center;
}

.action-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.action-card i {
  font-size: 24px;
  margin-bottom: 10px;
  display: block;
}

.recent-records {
  margin-top: 30px;
}

.recent-records h2 {
  margin-bottom: 20px;
  color: #333;
}

.recent-records ul {
  list-style: none;
  padding: 0;
}

.recent-records li {
  padding: 15px;
  border-bottom: 1px solid #eee;
  display: flex;
  gap: 20px;
}

.recent-records .date {
  color: #666;
  min-width: 100px;
}

.recent-records .type {
  font-weight: bold;
  color: #333;
  flex: 1;
}

.recent-records .severity {
  color: #666;
}
</style> 