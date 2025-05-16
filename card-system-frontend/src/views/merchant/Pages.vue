<template>
  <div class="p-6">
    <h2 class="text-xl font-bold mb-4">页面管理</h2>
    <div class="flex justify-between mb-4">
      <div class="text-sm text-gray-600">管理您的所有页面</div>
      <button 
        class="bg-primary text-white px-4 py-2 rounded-md hover:bg-primary/90 transition-colors flex items-center"
        @click="createNewPage"
      >
        <i class="fa-solid fa-plus mr-2"></i> 新建页面
      </button>
    </div>
    
    <!-- 页面列表 -->
    <div class="bg-white rounded-lg shadow-md overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">页面名称</th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">URL</th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">状态</th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">创建时间</th>
              <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">操作</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="page in pages" :key="page.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="h-8 w-8 bg-gray-200 rounded-md flex items-center justify-center mr-3">
                    <i class="fa-solid fa-file-alt text-gray-500"></i>
                  </div>
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ page.name }}</div>
                    <div class="text-sm text-gray-500">{{ page.type }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ page.url }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span 
                  class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                  :class="{
                    'bg-green-100 text-green-800': page.status === 'published',
                    'bg-yellow-100 text-yellow-800': page.status === 'draft',
                    'bg-red-100 text-red-800': page.status === 'archived'
                  }"
                >
                  {{ page.status === 'published' ? '已发布' : page.status === 'draft' ? '草稿' : '已归档' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ page.createdAt }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <router-link 
                  :to="{ name: 'PageBuilder', query: { id: page.id } }"
                  class="text-primary hover:text-primary/80 mr-4"
                >
                  编辑
                </router-link>
                <a href="#" class="text-gray-600 hover:text-gray-900 mr-4">预览</a>
                <button 
                  class="text-red-600 hover:text-red-900"
                  @click="deletePage(page.id)"
                >
                  删除
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import http from '@/utils/http'
import { ElMessage } from 'element-plus'

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
        error.value = '获取页面列表失败'
        console.error(err)
        ElMessage.error(error.value)
      } finally {
        loading.value = false
      }
    }
    
    // 创建新页面
    const createNewPage = () => {
      // 跳转到页面构建器，不带ID表示新建
      window.location.href = '/merchant/pages/builder'
    }
    
    // 删除页面
    const deletePage = (pageId) => {
      if (confirm('确定要删除这个页面吗？')) {
        http.delete(`/merchant/pages/${pageId}`)
          .then(() => {
            ElMessage.success('页面已删除')
            fetchPages()
          })
          .catch(err => {
            ElMessage.error('删除页面失败')
            console.error(err)
          })
      }
    }
    
    onMounted(() => {
      fetchPages()
    })
    
    return {
      pages,
      loading,
      error,
      createNewPage,
      deletePage
    }
  }
}
</script>