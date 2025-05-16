<template>
  <div class="flex h-screen overflow-hidden bg-gray-50">
    <!-- 左侧组件面板 -->
    <div class="w-64 bg-white border-r border-gray-200 shadow-sm flex flex-col">
      <div class="p-4 border-b border-gray-200">
        <h3 class="font-semibold text-gray-900">组件库</h3>
      </div>
      <div class="flex-1 overflow-y-auto p-4">
        <div class="space-y-3">
          <!-- 文本组件 -->
          <div 
            class="p-3 bg-gray-50 rounded-lg border border-gray-200 cursor-move hover:border-primary transition-colors"
            draggable="true"
            @dragstart="onDragStart('text')"
          >
            <div class="flex items-center">
              <div class="w-8 h-8 bg-blue-100 rounded flex items-center justify-center mr-3">
                <i class="fa-solid fa-font text-primary"></i>
              </div>
              <span class="font-medium text-gray-700">文本</span>
            </div>
          </div>
          
          <!-- 图片组件 -->
          <div 
            class="p-3 bg-gray-50 rounded-lg border border-gray-200 cursor-move hover:border-primary transition-colors"
            draggable="true"
            @dragstart="onDragStart('image')"
          >
            <div class="flex items-center">
              <div class="w-8 h-8 bg-green-100 rounded flex items-center justify-center mr-3">
                <i class="fa-solid fa-image text-success"></i>
              </div>
              <span class="font-medium text-gray-700">图片</span>
            </div>
          </div>
          
          <!-- 按钮组件 -->
          <div 
            class="p-3 bg-gray-50 rounded-lg border border-gray-200 cursor-move hover:border-primary transition-colors"
            draggable="true"
            @dragstart="onDragStart('button')"
          >
            <div class="flex items-center">
              <div class="w-8 h-8 bg-orange-100 rounded flex items-center justify-center mr-3">
                <i class="fa-solid fa-square text-warning"></i>
              </div>
              <span class="font-medium text-gray-700">按钮</span>
            </div>
          </div>
          
          <!-- 产品卡片组件 -->
          <div 
            class="p-3 bg-gray-50 rounded-lg border border-gray-200 cursor-move hover:border-primary transition-colors"
            draggable="true"
            @dragstart="onDragStart('product')"
          >
            <div class="flex items-center">
              <div class="w-8 h-8 bg-purple-100 rounded flex items-center justify-center mr-3">
                <i class="fa-solid fa-shopping-bag text-purple-600"></i>
              </div>
              <span class="font-medium text-gray-700">产品卡片</span>
            </div>
          </div>
          
          <!-- 分隔线组件 -->
          <div 
            class="p-3 bg-gray-50 rounded-lg border border-gray-200 cursor-move hover:border-primary transition-colors"
            draggable="true"
            @dragstart="onDragStart('divider')"
          >
            <div class="flex items-center">
              <div class="w-8 h-8 bg-gray-200 rounded flex items-center justify-center mr-3">
                <i class="fa-solid fa-minus text-gray-600"></i>
              </div>
              <span class="font-medium text-gray-700">分隔线</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 中间画布 -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <div class="p-4 border-b border-gray-200 bg-white flex justify-between items-center">
        <div>
          <input 
            type="text" 
            v-model="pageTitle" 
            placeholder="输入页面标题" 
            class="border border-gray-300 rounded-md px-3 py-1.5 text-sm focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
          >
        </div>
        <div class="flex space-x-2">
          <button class="text-gray-500 hover:text-primary transition-colors">
            <i class="fa-solid fa-eye mr-1"></i> 预览
          </button>
          <button class="text-gray-500 hover:text-primary transition-colors">
            <i class="fa-solid fa-mobile-alt mr-1"></i> 移动视图
          </button>
          <button @click="savePage" class="bg-primary hover:bg-primary/90 text-white px-3 py-1.5 rounded-md text-sm flex items-center">
            <i class="fa-solid fa-save mr-1"></i> 保存
          </button>
        </div>
      </div>
      
      <div class="flex-1 overflow-y-auto p-6 bg-gray-100" id="canvas">
        <div 
          class="max-w-4xl mx-auto min-h-[600px] bg-white shadow-md p-6"
          @dragover="onDragOver"
          @drop="onDrop"
        >
          <!-- 画布区域 -->
          <div v-if="!pageElements.length" class="flex flex-col items-center justify-center py-20 border-2 border-dashed border-gray-300 rounded-lg">
            <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mb-4">
              <i class="fa-solid fa-plus text-gray-400 text-2xl"></i>
            </div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">从左侧拖入组件开始设计</h3>
            <p class="text-gray-500 text-sm">您可以拖放各种组件到这里来构建您的页面</p>
          </div>
          
          <!-- 页面元素 -->
          <div 
            v-for="(element, index) in pageElements" 
            :key="element.id"
            :class="selectedElementId === element.id ? 'border-2 border-primary' : 'border-2 border-transparent'"
            class="mb-4 p-3 rounded-lg transition-all duration-200 hover:border-gray-300 relative group"
            @click="selectElement(element.id)"
          >
            <!-- 元素操作栏 -->
            <div class="absolute top-0 right-0 opacity-0 group-hover:opacity-100 transition-opacity bg-white/90 px-2 py-1 rounded-bl-lg shadow-sm">
              <button class="text-gray-500 hover:text-primary transition-colors mr-1" @click.stop="moveElementUp(index)">
                <i class="fa-solid fa-arrow-up"></i>
              </button>
              <button class="text-gray-500 hover:text-primary transition-colors mr-1" @click.stop="moveElementDown(index)">
                <i class="fa-solid fa-arrow-down"></i>
              </button>
              <button class="text-gray-500 hover:text-danger transition-colors" @click.stop="deleteElement(element.id)">
                <i class="fa-solid fa-trash"></i>
              </button>
            </div>
            
            <!-- 文本组件 -->
            <div v-if="element.type === 'text'" class="p-4">
              <div 
                class="w-full border border-gray-200 rounded p-4 focus:outline-none focus:ring-2 focus:ring-primary/30"
                contenteditable="true"
                @input="updateElementContent(element.id, $event.target.innerHTML)"
              >
                {{ element.content || '输入文本内容...' }}
              </div>
            </div>
            
            <!-- 图片组件 -->
            <div v-if="element.type === 'image'" class="p-4">
              <div class="relative">
                <img 
                  :src="element.content || 'https://picsum.photos/800/400'" 
                  alt="图片" 
                  class="w-full h-auto rounded"
                >
                <button 
                  class="absolute inset-0 w-full h-full bg-black/20 opacity-0 hover:opacity-100 transition-opacity flex items-center justify-center text-white"
                  @click.stop="uploadImage(element.id)"
                >
                  <i class="fa-solid fa-camera"></i>
                </button>
              </div>
              <input 
                type="text" 
                v-model="element.content" 
                placeholder="输入图片URL" 
                class="w-full mt-2 border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
              >
            </div>
            
            <!-- 按钮组件 -->
            <div v-if="element.type === 'button'" class="p-4 flex justify-center">
              <button 
                :class="element.style || 'bg-primary hover:bg-primary/90 text-white px-6 py-2 rounded-md transition-colors'"
                @click.stop=""
              >
                {{ element.content || '按钮文本' }}
              </button>
            </div>
            
            <!-- 产品卡片组件 -->
            <div v-if="element.type === 'product'" class="p-4">
              <div class="border border-gray-200 rounded-lg overflow-hidden">
                <img 
                  :src="element.image || 'https://picsum.photos/600/300'" 
                  alt="产品图片" 
                  class="w-full h-48 object-cover"
                >
                <div class="p-4">
                  <h3 class="font-semibold text-gray-900 text-lg mb-1">{{ element.title || '产品标题' }}</h3>
                  <p class="text-gray-500 text-sm mb-3 line-clamp-2">{{ element.description || '产品描述...' }}</p>
                  <div class="flex justify-between items-center">
                    <span class="text-primary font-bold text-lg">¥{{ element.price || '0.00' }}</span>
                    <button class="bg-primary hover:bg-primary/90 text-white px-4 py-1.5 rounded text-sm transition-colors">
                      购买
                    </button>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- 分隔线组件 -->
            <div v-if="element.type === 'divider'" class="p-4">
              <div class="border-t border-gray-200"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 右侧属性面板 -->
    <div v-if="selectedElement" class="w-72 bg-white border-l border-gray-200 shadow-sm flex flex-col">
      <div class="p-4 border-b border-gray-200">
        <h3 class="font-semibold text-gray-900">属性设置</h3>
      </div>
      <div class="flex-1 overflow-y-auto p-4">
        <div class="space-y-4">
          <!-- 通用设置 -->
          <div>
            <h4 class="text-sm font-medium text-gray-700 mb-2">基本设置</h4>
            <div class="space-y-3">
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">元素类型</label>
                <span class="text-sm text-gray-700">{{ selectedElement.type }}</span>
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">元素ID</label>
                <span class="text-sm text-gray-700">{{ selectedElement.id }}</span>
              </div>
            </div>
          </div>
          
          <!-- 文本设置 -->
          <div v-if="selectedElement.type === 'text'">
            <h4 class="text-sm font-medium text-gray-700 mb-2">文本设置</h4>
            <div class="space-y-3">
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">文本内容</label>
                <textarea 
                  v-model="selectedElement.content" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                  rows="3"
                ></textarea>
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">字体大小</label>
                <input 
                  type="number" 
                  v-model="selectedElement.fontSize" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">字体颜色</label>
                <input 
                  type="color" 
                  v-model="selectedElement.color" 
                  class="w-full h-10 border border-gray-200 rounded"
                >
              </div>
            </div>
          </div>
          
          <!-- 图片设置 -->
          <div v-if="selectedElement.type === 'image'">
            <h4 class="text-sm font-medium text-gray-700 mb-2">图片设置</h4>
            <div class="space-y-3">
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">图片URL</label>
                <input 
                  type="text" 
                  v-model="selectedElement.content" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">图片说明</label>
                <input 
                  type="text" 
                  v-model="selectedElement.alt" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">对齐方式</label>
                <select 
                  v-model="selectedElement.align" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
                  <option value="left">左对齐</option>
                  <option value="center">居中对齐</option>
                  <option value="right">右对齐</option>
                </select>
              </div>
            </div>
          </div>
          
          <!-- 按钮设置 -->
          <div v-if="selectedElement.type === 'button'">
            <h4 class="text-sm font-medium text-gray-700 mb-2">按钮设置</h4>
            <div class="space-y-3">
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">按钮文本</label>
                <input 
                  type="text" 
                  v-model="selectedElement.content" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">按钮样式</label>
                <select 
                  v-model="selectedElement.style" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
                  <option value="bg-primary hover:bg-primary/90 text-white px-6 py-2 rounded-md transition-colors">主色调</option>
                  <option value="bg-secondary hover:bg-secondary/90 text-white px-6 py-2 rounded-md transition-colors">辅助色</option>
                  <option value="bg-white border border-gray-300 hover:bg-gray-50 text-gray-700 px-6 py-2 rounded-md transition-colors">边框按钮</option>
                </select>
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">按钮大小</label>
                <select 
                  v-model="selectedElement.size" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
                  <option value="px-4 py-1.5 text-sm">小</option>
                  <option value="px-6 py-2">中</option>
                  <option value="px-8 py-3 text-lg">大</option>
                </select>
              </div>
            </div>
          </div>
          
          <!-- 产品卡片设置 -->
          <div v-if="selectedElement.type === 'product'">
            <h4 class="text-sm font-medium text-gray-700 mb-2">产品卡片设置</h4>
            <div class="space-y-3">
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">产品标题</label>
                <input 
                  type="text" 
                  v-model="selectedElement.title" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">产品描述</label>
                <textarea 
                  v-model="selectedElement.description" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                  rows="3"
                ></textarea>
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">产品价格</label>
                <input 
                  type="number" 
                  v-model="selectedElement.price" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
              </div>
              <div>
                <label class="block text-xs font-medium text-gray-500 mb-1">产品图片</label>
                <input 
                  type="text" 
                  v-model="selectedElement.image" 
                  class="w-full border border-gray-200 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary/30"
                >
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="p-4 border-t border-gray-200">
        <button 
          @click="deleteSelectedElement" 
          class="w-full bg-danger hover:bg-danger/90 text-white py-2 rounded-md transition-colors flex items-center justify-center"
        >
          <i class="fa-solid fa-trash mr-2"></i> 删除元素
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import http from '@/utils/http'

export default {
  setup() {
    // 页面数据
    const pageTitle = ref('未命名页面')
    const pageElements = ref([])
    const selectedElementId = ref(null)
    const selectedElement = computed(() => {
      return pageElements.value.find(el => el.id === selectedElementId.value) || null
    })
    
    // 初始化页面数据
    const initPageData = () => {
      // 检查URL参数，加载已存在的页面
      const urlParams = new URLSearchParams(window.location.search)
      const pageId = urlParams.get('pageId')
      
      if (pageId) {
        loadPageData(pageId)
      } else {
        // 新建页面
        pageTitle.value = '新建页面'
      }
    }
    
    // 加载页面数据
    const loadPageData = async (pageId) => {
      try {
        const response = await http.get(`/merchant/pages/${pageId}`)
        pageTitle.value = response.data.name
        pageElements.value = response.data.elements || []
      } catch (err) {
        window.$message.error('加载页面失败，请重试')
        console.error(err)
      }
    }
    
    // 组件拖拽相关
    const draggedElementType = ref('')
    
    const onDragStart = (type) => {
      draggedElementType.value = type
    }
    
    const onDragOver = (e) => {
      e.preventDefault()
    }
    
    const onDrop = (e) => {
      e.preventDefault()
      
      if (!draggedElementType.value) return
      
      // 创建新元素
      const newElement = {
        id: Date.now().toString(), // 简单的ID生成
        type: draggedElementType.value,
        content: '',
        createdAt: new Date().toISOString(),
        // 根据元素类型设置默认属性
        ...getDefaultElementProperties(draggedElementType.value)
      }
      
      // 添加到页面元素
      pageElements.value.push(newElement)
      
      // 选中新元素
      selectedElementId.value = newElement.id
      
      // 重置拖拽状态
      draggedElementType.value = ''
    }
    
    // 获取元素默认属性
    const getDefaultElementProperties = (type) => {
      switch (type) {
        case 'text':
          return {
            fontSize: 16,
            color: '#333333'
          }
        case 'image':
          return {
            content: 'https://picsum.photos/800/400',
            alt: '图片描述',
            align: 'center'
          }
        case 'button':
          return {
            content: '按钮文本',
            style: 'bg-primary hover:bg-primary/90 text-white px-6 py-2 rounded-md transition-colors',
            size: 'px-6 py-2'
          }
        case 'product':
          return {
            title: '产品标题',
            description: '产品描述...',
            price: 0,
            image: 'https://picsum.photos/600/300'
          }
        default:
          return {}
      }
    }
    
    // 元素操作
    const selectElement = (id) => {
      selectedElementId.value = id
    }
    
    const updateElementContent = (id, content) => {
      const elementIndex = pageElements.value.findIndex(el => el.id === id)
      if (elementIndex !== -1) {
        pageElements.value[elementIndex].content = content
      }
    }
    
    const moveElementUp = (index) => {
      if (index > 0) {
        [pageElements.value[index], pageElements.value[index-1]] = [pageElements.value[index-1], pageElements.value[index]]
      }
    }
    
    const moveElementDown = (index) => {
      if (index < pageElements.value.length - 1) {
        [pageElements.value[index], pageElements.value[index+1]] = [pageElements.value[index+1], pageElements.value[index]]
      }
    }
    
    const deleteElement = (id) => {
      if (confirm('确定要删除这个元素吗？')) {
        pageElements.value = pageElements.value.filter(el => el.id !== id)
        if (selectedElementId.value === id) {
          selectedElementId.value = null
        }
      }
    }
    
    const deleteSelectedElement = () => {
      if (selectedElementId.value) {
        deleteElement(selectedElementId.value)
      }
    }
    
    // 图片上传
    const uploadImage = (elementId) => {
      // 实际项目中应该实现图片上传功能
      // 这里仅作为示例
      window.$message.info('图片上传功能尚未实现')
    }
    
    // 保存页面
    const savePage = async () => {
      try {
        const pageData = {
          name: pageTitle.value,
          elements: pageElements.value
        }
        
        const response = await http.post('/merchant/pages/save', pageData)
        
        window.$message.success('页面保存成功')
        
        // 更新URL中的pageId
        if (response.data.id) {
          const url = new URL(window.location.href)
          url.searchParams.set('pageId', response.data.id)
          window.history.replaceState(null, '', url)
        }
      } catch (err) {
        window.$message.error('页面保存失败，请重试')
        console.error(err)
      }
    }
    
    onMounted(() => {
      initPageData()
    })
    
    return {
      pageTitle,
      pageElements,
      selectedElementId,
      selectedElement,
      onDragStart,
      onDragOver,
      onDrop,
      selectElement,
      updateElementContent,
      moveElementUp,
      moveElementDown,
      deleteElement,
      deleteSelectedElement,
      uploadImage,
      savePage
    }
  }
}
</script>