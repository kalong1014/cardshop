<template>
  <div ref="chart" style="height: 400px;"></div>
</template>

<script>
import echarts from 'echarts'

export default {
  mounted() {
    this.renderChart()
  },
  methods: {
    async renderChart() {
      const res = await this.$api.getDailyOrderTrend()
      const chart = echarts.init(this.$refs.chart)
      chart.setOption({
        xAxis: { type: 'category', data: res.dates },
        yAxis: { type: 'value' },
        series: [{
          data: res.orders,
          type: 'line',
          smooth: true
        }]
      })
    }
  }
}
</script>