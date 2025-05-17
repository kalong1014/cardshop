<template>
  <div class="refund-form">
    <el-form :model="form" rules="rules">
      <el-form-item label="退款原因" prop="reason">
        <el-textarea v-model="form.reason" rows="3" placeholder="请输入原因"></el-textarea>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitRefund">提交申请</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        orderId: this.$route.params.orderId,
        reason: ''
      },
      rules: {
        reason: [{ required: true, message: '请填写退款原因', trigger: 'blur' }]
      }
    }
  },
  methods: {
    async submitRefund() {
      try {
        await this.$api.applyRefund(this.form)
        this.$message.success('退款申请已提交')
        this.$router.go(-1)
      } catch (error) {
        this.$message.error('申请失败：' + error.response.data.error)
      }
    }
  }
}
</script>