<template>
  <div class="min-h-screen bg-gray-50">
    <!-- 顶部工具栏 -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-10">
      <div class="flex items-center justify-between px-6 py-3">
        <div class="flex items-center">
          <button @click="goBack" class="mr-4 text-gray-600 hover:text-gray-900">
            <i class="fa-solid fa-arrow-left"></i>
          </button>
          <h2 class="text-xl font-bold text-gray-900">页面构建器</h2>
        </div>
        
        <div class="flex items-center space-x-3">
          <button class="px-4 py-2 text-sm text-gray-700 border border-gray-300 rounded-md hover:bg-gray-50">
            <i class="fa-solid fa-eye mr-1"></i> 预览
          </button>
          <button class="px-4 py-2 text-sm text-gray-700 border border-gray-300 rounded-md hover:bg-gray-50">
            <i class="fa-solid fa-paint-brush mr-1"></i> 设计
          </button>
          <button 
            class="px-4 py-2 text-sm font-medium text-white bg-primary rounded-md hover:bg-primary/90 transition-colors"
            @click="savePage"
          >
            <i class="fa-solid fa-save mr-1"></i> 保存
          </button>
        </div>
      </div>
    </div>
    
    <div class="flex h-[calc(100vh-64px)]">
      <!-- 左侧组件面板 -->
      <div class="w-64 bg-white border-r border-gray-200 overflow-y-auto">
        <div class="p-4">
          <h3 class="text-sm font-medium text-gray-900 mb-3">页面组件</h3>
          <div class="space-y-2">
            <div 
              class="p-3 bg-gray-50 rounded-md border border-gray-200 cursor-move hover:bg-gray-100 transition-colors"
              draggable="true"
              @dragstart="dragStart('heading')"
            >
              <i class="fa-solid fa-heading mr-2 text-gray-600"></i> 标题
            </div>
            <div 
              class="p-3 bg-gray-50 rounded-md border border-gray-200 cursor-move hover:bg-gray-100 transition-colors"
              draggable="true"
              @dragstart="dragStart('paragraph')"
            >
              <i class="fa-solid fa-paragraph mr-2 text-gray-600"></i> 段落
            </div>
            <div 
              class="p-3 bg-gray-50 rounded-md border border-gray-200 cursor-move hover:bg-gray-100 transition-colors"
              draggable="true"
              @dragstart="dragStart('image')"
            >
              <i class="fa-solid fa-image mr-2 text-gray-600"></i> 图片
            </div>
            <div 
              class="p-3 bg-gray-50 rounded-md border border-gray-200 cursor-move hover:bg-gray-100 transition-colors"
              draggable="true"
              @dragstart="dragStart('button')"
            >
              <i class="fa-solid fa-square mr-2 text-gray-600"></i> 按钮
            </div>
            <div 
              class="p-3 bg-gray-50 rounded-md border border-gray-200 cursor-move hover:bg-gray-100 transition-colors"
              draggable="true"
              @dragstart="dragStart('card')"
            >
              <i class="fa-solid fa-credit-card mr-2 text-gray-600"></i> 卡片
            </div>
            <div 
              class="p-3 bg-gray-50 rounded-md border border-gray-200 cursor-move hover:bg-gray-100 transition-colors"
              draggable="true"
              @dragstart="dragStart('list')"
            >
              <i class="fa-solid fa-list mr-2 text-gray-600"></i> 列表
            </div>
            <div 
              class="p-3 bg-gray-50 rounded-md border border-gray-200 cursor-move hover:bg-gray-100 transition-colors"
              draggable="true"
              @dragstart="dragStart('form')"
            >
              <i class="fa-solid fa-file-invoice mr-2 text-gray-600"></i> 表单
            </div>
          </div>
        </div>
      </div>
      
      <!-- 中间编辑区域 -->
      <div class="flex-1 overflow-y-auto p-6 bg-gray-100">
        <div id="dropzone" class="max-w-4xl mx-auto min-h-[800px] bg-white shadow-md p-8"
             @dragover="dragOver"
             @drop="drop">
          <div v-if="!pageContent.length" class="h-64 flex items-center justify-center border-2 border-dashed border-gray-300 rounded-lg">
            <div class="text-center">
              <i class="fa-solid fa-arrow-down text-gray-400 text-3xl mb-3"></i>
              <p class="text-gray-500">从左侧拖入组件到这里开始设计您的页面</p>
            </div>
          </div>
          
          <div 
            v-for="(element, index) in pageContent" 
            :key="index"
            :class="{ 'border-2 border-primary bg-blue-50': selectedElementIndex === index }"
            class="mb-4 p-3 border border-gray-200 rounded-md relative group"
            @click="selectElement(index)"
          >
            <div v-if="element.type === 'heading'" class="py-2">
              <h2 class="text-2xl font-bold text-gray-900">{{ element.content || '这是一个标题' }}</h2>
            </div>
            
            <div v-if="element.type === 'paragraph'" class="py-2">
              <p class="text-gray-700 leading-relaxed">{{ element.content || '这是一段文本内容。您可以在此处添加产品描述、介绍或其他信息。' }}</p>
            </div>
            
            <div v-if="element.type === 'image'" class="py-2">
              <img 
                :src="element.src || 'https://picsum.photos/800/400'" 
                alt="页面图片" 
                class="w-full h-auto rounded-lg"
              />
              <p class="text-sm text-gray-500 mt-1 text-center">{{ element.alt || '图片描述' }}</p>
            </div>
            
            <div v-if="element.type === 'button'" class="py-2 text-center">
              <button class="px-6 py-3 bg-primary text-white rounded-md hover:bg-primary/90 transition-colors">
                {{ element.text || '点击按钮' }}
              </button>
            </div>
            
            <div v-if="element.type === 'card'" class="py-2">
              <div class="bg-gray-50 rounded-lg p-4 shadow-sm">
                <h3 class="text-lg font-semibold text-gray-900 mb-2">{{ element.title || '卡片标题' }}</h3>
                <p class="text-gray-600 mb-3">{{ element.content || '卡片内容描述' }}</p>
                <button class="text-primary hover:text-primary/80 text-sm">
                  {{ element.buttonText || '了解更多' }} <i class="fa-solid fa-arrow-right ml-1"></i>
                </button>
              </div>
            </div>
            
            <div v-if="element.type === 'list'" class="py-2">
              <ul class="list-disc pl-5 text-gray-700">
                <li v-for="(item, i) in element.items || ['列表项 1', '列表项 2', '列表项 3']" :key="i">{{ item }}</li>
              </ul>
            </div>
            
            <div v-if="element.type === 'form'" class="py-2">
              <div class="bg-gray-50 rounded-lg p-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-3">{{ element.title || '联系我们' }}</h3>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">姓名</label>
                    <input type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent">
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">邮箱</label>
                    <input type="email" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent">
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">留言</label>
                    <textarea rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"></textarea>
                  </div>
                  <button class="w-full py-2 bg-primary text-white rounded-md hover:bg-primary/90 transition-colors">
                    {{ element.buttonText || '提交' }}
                  </button>
                </div>
              </div>
            </div>
            
            <!-- 编辑控件 -->
            <div class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity">
              <button class="text-gray-500 hover:text-gray-700 mr-2" @click.stop="editElement(index)">
                <i class="fa-solid fa-pencil"></i>
              </button>
              <button class="text-gray-500 hover:text-red-600" @click.stop="deleteElement(index)">
                <i class="fa-solid fa-trash"></i>
              </button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 右侧属性面板 -->
      <div class="w-72 bg-white border-l border-gray-200 overflow-y-auto">
        <div class="p-4">
          <h3 class="text-sm font-medium text-gray-900 mb-3">属性设置</h3>
          
          <div v-if="selectedElementIndex !== -1" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">元素类型</label>
              <p class="text-sm text-gray-500 bg-gray-50 p-2 rounded-md">
                {{ elementTypeText(pageContent[selectedElementIndex].type) }}
              </p>
            </div>
            
            <div v-if="pageContent[selectedElementIndex].type === 'heading'">
              <label class="block text-sm font-medium text-gray-700 mb-1">标题内容</label>
              <input 
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
                v-model="pageContent[selectedElementIndex].content"
              >
            </div>
            
            <div v-if="pageContent[selectedElementIndex].type === 'paragraph'">
              <label class="block text-sm font-medium text-gray-700 mb-1">段落内容</label>
              <textarea 
                rows="3" 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
                v-model="pageContent[selectedElementIndex].content"
              ></textarea>
            </div>
            
            <div v-if="pageContent[selectedElementIndex].type === 'image'">
              <label class="block text-sm font-medium text-gray-700 mb-1">图片URL</label>
              <input 
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
                v-model="pageContent[selectedElementIndex].src"
              >
              <label class="block text-sm font-medium text-gray-700 mb-1 mt-3">图片描述</label>
              <input 
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
                v-model="pageContent[selectedElementIndex].alt"
              >
            </div>
            
            <div v-if="pageContent[selectedElementIndex].type === 'button'">
              <label class="block text-sm font-medium text-gray-700 mb-1">按钮文本</label>
              <input 
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
                v-model="pageContent[selectedElementIndex].text"
              >
              <label class="block text-sm font-medium text-gray-700 mb-1 mt-3">按钮颜色</label>
              <div class="grid grid-cols-4 gap-2 mt-2">
                <button 
                  class="h-8 rounded-md bg-primary"
                  @click="pageContent[selectedElementIndex].color = 'primary'"
                  :class="{ 'ring-2 ring-offset-2 ring-black': pageContent[selectedElementIndex].color === 'primary' }"
                ></button>
                <button 
                  class="h-8 rounded-md bg-secondary"
                  @click="pageContent[selectedElementIndex].color = 'secondary'"
                  :class="{ 'ring-2 ring-offset-2 ring-black': pageContent[selectedElementIndex].color === 'secondary' }"
                ></button>
                <button 
                  class="h-8 rounded-md bg-success"
                  @click="pageContent[selectedElementIndex].color = 'success'"
                  :class="{ 'ring-2 ring-offset-2 ring-black': pageContent[selectedElementIndex].color === 'success' }"
                ></button>
                <button 
                  class="h-8 rounded-md bg-danger"
                  @click="pageContent[selectedElementIndex].color = 'danger'"
                  :class="{ 'ring-2 ring-offset-2 ring-black': pageContent[selectedElementIndex].color === 'danger' }"
                ></button>
              </div>
            </div>
            
            <div class="pt-4 border-t border-gray-200">
              <button 
                class="w-full py-2 bg-primary text-white rounded-md hover:bg-primary/90 transition-colors"
                @click="savePage"
              >
                保存设置
              </button>
            </div>
          </div>
          
          <div v-else class="text-center py-8 text-gray-500">
            <i class="fa-solid fa-info-circle mb-2 text-gray-400"></i>
            <p>选择一个元素以编辑其属性</p>
          </div>
        </div>
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
    const pageContent = ref([])
    const selectedElementIndex = ref(-1)
    const draggedElementType = ref('')
    const pageId = ref(null)
    const pageTitle = ref('')
    
    // 页面加载时获取页面数据
    onMounted(() => {
      const query = new URLSearchParams(window.location.search)
      pageId.value = query.get('id')
      
      if (pageId.value) {
        // 编辑已存在的页面
        fetchPageData()
      } else {
        // 新建页面
        pageTitle.value = '新建页面'
      }
    })
    
    // 获取页面数据
    const fetchPageData = async () => {
      try {
        const response = await http.get(`/merchant/pages/${pageId.value}`)
        pageContent.value = response.data.content || []
        pageTitle.value = response.data.name || '未命名页面'
      } catch (err) {
        ElMessage.error('获取页面数据失败')
        console.error(err)
      }
    }
    
    // 拖放相关方法
    const dragStart = (type) => {
      draggedElementType.value = type
    }
    
    const dragOver = (e) => {
      e.preventDefault()
    }
    
    const drop = (e) => {
      e.preventDefault()
      
      // 添加新元素
      const newElement = {
        type: draggedElementType.value,
        id: Date.now().toString(),
        content: '',
        // 根据元素类型设置默认属性
        ...(draggedElementType.value === 'button' && { text: '点击按钮', color: 'primary' }),
        ...(draggedElementType.value === 'image' && { src: 'https://picsum.photos/800/400', alt: '图片描述' }),
        ...(draggedElementType.value === 'card' && { title: '卡片标题', content: '卡片内容描述', buttonText: '了解更多' }),
        ...(draggedElementType.value === 'list' && { items: ['列表项 1', '列表项 2', '列表项 3'] }),
        ...(draggedElementType.value === 'form' && { title: '联系我们', buttonText: '提交' }),
      }
      
      pageContent.value.push(newElement)
      selectedElementIndex.value = pageContent.value.length - 1
    }
    
    // 元素操作方法
    const selectElement = (index) => {
      selectedElementIndex.value = index
    }
    
    const deleteElement = (index) => {
      if (confirm('确定要删除这个元素吗？')) {
        pageContent.value.splice(index, 1)
        if (selectedElementIndex.value === index) {
          selectedElementIndex.value = -1
        } else if (selectedElementIndex.value > index) {
          selectedElementIndex.value--
        }
      }
    }
    
    const editElement = (index) => {
      // 编辑元素逻辑，可以打开模态框等
      selectedElementIndex.value = index
    }
    
    // 辅助方法
    const elementTypeText = (type) => {
      const types = {
        'heading': '标题',
        'paragraph': '段落',
        'image': '图片',
        'button': '按钮',
        'card': '卡片',
        'list': '列表',
        'form': '表单'
      }
      return types[type] || type
    }
    
    // 保存页面
    const savePage = async () => {
      try {
        const pageData = {
          name: pageTitle.value,
          content: pageContent.value,
          status: 'draft'
        }
        
        let response
        
        if (pageId.value) {
          // 更新现有页面
          response = await http.put(`/merchant/pages/${pageId.value}`, pageData)
        } else {
          // 创建新页面
          response = await http.post('/merchant/pages', pageData)
          pageId.value = response.data.id
        }
        
        ElMessage.success('页面已保存')
      } catch (err) {
        ElMessage.error('保存页面失败')
        console.error(err)
      }
    }
    
    // 返回
    const goBack = () => {
      window.history.back()
    }
    
    return {
      pageContent,
      selectedElementIndex,
      draggedElementType,
      pageId,
      pageTitle,
      dragStart,
      dragOver,
      drop,
      selectElement,
      deleteElement,
      editElement,
      elementTypeText,
      savePage,
      goBack
    }
  }
}
</script>