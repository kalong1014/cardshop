<template>
  <div class="stats-panel">
    <el-card>
      <h3>商户总数：{{ stats.merchantCount }}</h3>
      <h3>卡密总数：{{ stats.cardCount }}</h3>
      <h3>今日订单：{{ stats.todayOrders }}</h3>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      stats: {
        merchantCount: 0,
        cardCount: 0,
        todayOrders: 0
      }
    }
  },
  mounted() {
    this.fetchStats()
  },
  methods: {
    async fetchStats() {
      const res = await this.$api.getSystemStats()
      this.stats = res.data
    }
  }
}
</script>