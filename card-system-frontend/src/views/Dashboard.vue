<template>
  <div class="dashboard">
    <div class="stat-card">
      <h2>总用户数</h2>
      <p>{{ stats.totalUsers }}</p>
    </div>
    
    <div class="chart-container">
      <order-trend></order-trend>
      <top-products></top-products>
    </div>
  </div>
</template>

<script>
import OrderTrend from '@/components/OrderTrend.vue'
import TopProducts from '@/components/TopProducts.vue'
import { getDashboardStats } from '@/api/dashboard'

export default {
  components: { OrderTrend, TopProducts },
  data() {
    return {
      stats: {}
    }
  },
  mounted() {
    this.fetchStats()
  },
  methods: {
    async fetchStats() {
      this.stats = await getDashboardStats()
    }
  }
}
</script>