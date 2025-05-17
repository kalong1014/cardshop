<template>
  <div class="container">
    <el-button type="primary" @click="openGenerateDialog">生成卡密</el-button>
    
    <el-dialog title="批量生成卡密" :visible.sync="generateDialogVisible">
      <el-form ref="genForm" :model="genForm" rules="genRules">
        <el-form-item label="产品分类" prop="productId">
          <el-input v-model="genForm.productId"></el-input>
        </el-form-item>
        <el-form-item label="生成数量" prop="count">
          <el-input-number v-model="genForm.count" :min="1" :max="10000"></el-input-number>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="generateDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleGenerate">生成</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { generateCards } from '@/api/card'

export default {
  data() {
    return {
      generateDialogVisible: false,
      genForm: {
        productId: 'default',
        count: 10
      },
      genRules: {
        productId: [{ required: true, message: '请输入产品分类', trigger: 'blur' }],
        count: [{ required: true, message: '请选择生成数量', trigger: 'blur' }]
      }
    }
  },
  methods: {
    openGenerateDialog() {
      this.generateDialogVisible = true
    },
    async handleGenerate() {
      try {
        await generateCards(this.genForm)
        this.$message.success('卡密生成中，可在列表查看')
        this.generateDialogVisible = false
      } catch (error) {
        this.$message.error('生成失败：' + error.response.data.error)
      }
    }
  }
}
</script>