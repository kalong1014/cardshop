<template>
  <el-upload
    :action="$api.uploadLogo"
    :show-file-list="false"
    :on-success="handleUploadSuccess"
    :before-upload="beforeUpload"
  >
    <el-button type="primary">上传LOGO</el-button>
  </el-upload>
  <img v-if="merchant.logo" :src="merchant.logo" style="max-width: 200px;">
</template>

<script>
export default {
  data() {
    return {
      merchant: {}
    }
  },
  methods: {
    beforeUpload(file) {
      const isJPG = file.type === 'image/jpeg'
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isJPG) {
        this.$message.error('请上传JPG格式图片')
      }
      if (!isLt2M) {
        this.$message.error('图片大小不超过2MB')
      }
      return isJPG && isLt2M
    },
    handleUploadSuccess(res) {
      this.merchant.logo = res.url
      this.saveMerchantInfo() // 保存商户信息时更新LOGO路径
    }
  }
}
</script>