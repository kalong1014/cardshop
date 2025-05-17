<template>
  <div class="container">
    <el-form ref="form" :model="form" rules="rules" label-width="80px">
      <el-form-item label="商户名称" prop="name">
        <el-input v-model="form.name" placeholder="请输入商户名称"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitApply">提交申请</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { applyMerchant } from '@/api/merchant'

export default {
  data() {
    return {
      form: {
        name: ''
      },
      rules: {
        name: [{ required: true, message: '请输入商户名称', trigger: 'blur' }]
      }
    }
  },
  methods: {
    async submitApply() {
      try {
        await applyMerchant(this.form)
        this.$message.success('申请已提交，等待审核')
        this.$router.push('/merchant/dashboard')
      } catch (error) {
        this.$message.error(error.response.data.error)
      }
    }
  }
}
</script>