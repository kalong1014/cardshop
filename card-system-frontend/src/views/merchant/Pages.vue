<template>
  <div class="min-h-screen bg-gray-50">
    <!-- 顶部导航 -->
    <div class="bg-white shadow-sm">
      <!-- 导航内容省略... -->
    </div>

    <!-- 主要内容 -->
    <main>
      <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div class="bg-white overflow-hidden shadow-sm sm:rounded-lg">
          <div class="p-6 bg-white border-b border-gray-200">
            <div class="flex justify-between items-center mb-6">
              <h2 class="text-2xl font-bold text-gray-900">页面管理</h2>
              <button 
                @click="createNewPage" 
                class="bg-primary hover:bg-primary/90 text-white font-medium py-2 px-4 rounded-lg transition-all duration-300 flex items-center"
              >
                <i class="fa-solid fa-plus mr-2"></i> 新建页面
              </button>
            </div>
            
            <!-- 页面列表 -->
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              <div 
                v-for="page in pages" 
                :key="page.id" 
                class="bg-white rounded-lg shadow-md overflow-hidden transition-all duration-300 hover:shadow-xl cursor-pointer"
                @click="editPage(page.id)"
              >
                <div class="h-48 bg-gray-100 relative">
                  <img 
                    :src="page.thumbnail || 'https://picsum.photos/800/400?random=' + page.id" 
                    alt="页面预览" 
                    class="w-full h-full object-cover"
                  >
                  <div class="absolute top-2 right-2 bg-white/80 px-2 py-1 rounded text-xs font-medium">
                    {{ page.status }}
                  </div>
                </div>
                <div class="p-4">
                  <h3 class="font-semibold text-gray-900 text-lg mb-1">{{ page.name }}</h3>
                  <p class="text-gray-500 text-sm mb-3">{{ page.description || '未设置描述' }}</p>
                  <div class="flex justify-between items-center">
                    <span class="text-xs text-gray-400">{{ formatDate(page.updated_at) }}</span>
                    <div class="flex space-x-2">
                      <button class="text-gray-500 hover:text-primary transition-colors">
                        <i class="fa-solid fa-eye"></i>
                      </button>
                      <button class="text-gray-500 hover:text-primary transition-colors">
                        <i class="fa-solid fa-edit"></i>
                      </button>
                      <button class="text-gray-500 hover:text-danger transition-colors">
                        <i class="fa-solid fa-trash"></i>
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- 空状态 -->
            <div v-if="pages.length === 0" class="flex flex-col items-center justify-center py-12">
              <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mb-4">
                <i class="fa-solid fa-file-alt text-gray-400 text-2xl"></i>
              </div>
              <h3 class="text-lg font-medium text-gray-900 mb-2">还没有创建任何页面</h3>
              <p class="text-gray-500 mb-4">点击上方"新建页面"按钮开始设计您的销售页面</p>
              <button 
                @click="createNewPage" 
                class="bg-primary hover:bg-primary/90 text-white font-medium py-2 px-4 rounded-lg transition-all duration-300"
              >
                新建页面
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import http from '@/utils/http'

export default {
  setup() {
    const pages = ref([])
    const loading = ref(true)
    const error = ref('')
    
    // 获取页面列表
    const fetchPages = async () => {
      try {
        const response = await http.get('/merchant/pages')
        pages.value = response.data
      } catch (err) {
        error.value = '获取页面列表失败，请重试'
        console.error(err)
      } finally {
        loading.value = false
      }
    }
    
    // 创建新页面
    const createNewPage = () => {
      window.location.href = '/merchant/pages/builder'
    }
    
    // 编辑页面
    const editPage = (pageId) => {
      window.location.href = `/merchant/pages/builder?pageId=${pageId}`
    }
    
    // 格式化日期
    const formatDate = (dateStr) => {
      if (!dateStr) return ''
      const date = new Date(dateStr)
      return date.toLocaleDateString()
    }
    
    onMounted(() => {
      fetchPages()
    })
    
    return {
      pages,
      loading,
      error,
      createNewPage,
      editPage,
      formatDate
    }
  }
}
</script>