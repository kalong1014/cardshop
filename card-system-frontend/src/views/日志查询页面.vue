<template>
  <el-table :data="logs" stripe>
    <el-table-column prop="createdAt" label="时间" :formatter="formatDate"></el-table-column>
    <el-table-column prop="username" label="用户"></el-table-column>
    <el-table-column prop="type" label="类型" :formatter="formatType"></el-table-column>
    <el-table-column prop="action" label="操作"></el-table-column>
  </el-table>
</template>

<script>
export default {
  data() {
    return {
      logs: []
    }
  },
  mounted() {
    this.fetchLogs()
  },
  methods: {
    async fetchLogs() {
      const res = await this.$api.getOperationLogs()
      this.logs = res.data
    }
  },
  filters: {
    formatDate(time) {
      return new Date(time).toLocaleString()
    },
    formatType(type) {
      return type === 'login' ? '登录' : type === 'merchant' ? '商户管理' : '其他'
    }
  }
}
</script>