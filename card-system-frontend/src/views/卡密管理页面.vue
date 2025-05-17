<template>
  <div>
    <el-table :data="cards" stripe>
      <el-table-column prop="id" label="卡密ID"></el-table-column>
      <el-table-column prop="productId" label="产品分类"></el-table-column>
      <el-table-column :formatter="formatStatus" label="状态"></el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button @click="invalidateCard(scope.row.id)">失效</el-button>
        </template>
      </el-table-column>
    </el-table>
    
    <el-pagination
      :total="total"
      :current-page="query.page"
      :page-size="query.pageSize"
      @current-change="handlePageChange"
    ></el-pagination>
  </div>
</template>

<script>
export default {
  data() {
    return {
      query: {
        page: 1,
        pageSize: 20,
        status: ''
      },
      cards: [],
      total: 0
    }
  },
  mounted() {
    this.fetchCards()
  },
  methods: {
    async fetchCards() {
      const res = await this.$api.getCards(this.query)
      this.cards = res.data.cards
      this.total = res.data.total
    },
    async invalidateCard(id) {
      if (confirm('确认要失效此卡密吗？')) {
        await this.$api.invalidateCard(id)
        this.fetchCards()
      }
    }
  },
  filters: {
    formatStatus(status) {
      return status === 'unused' ? '未使用' : status === 'used' ? '已使用' : '已过期'
    }
  }
}
</script>