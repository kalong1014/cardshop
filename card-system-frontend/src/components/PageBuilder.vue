<template>
  <div class="page-builder">
    <div class="drag-container" v-draggable="draggableOptions" @start="onDragStart" @end="onDragEnd">
      <!-- 页面元素 -->
      <div 
        v-for="(element, index) in pageElements" 
        :key="element.id || index" 
        class="element-card"
        :class="{'dragging': element.isDragging}"
      >
        <div class="element-header">
          <span>{{ element.name || `Element ${index + 1}` }}</span>
          <div class="actions">
            <button @click="editElement(index)">编辑</button>
            <button @click="removeElement(index)">删除</button>
          </div>
        </div>
        <div class="element-content" v-html="element.content"></div>
      </div>
    </div>
    
    <!-- 添加元素按钮 -->
    <div class="add-element-btn">
      <button @click="addElement">添加元素</button>
    </div>
    
    <!-- 编辑元素模态框 -->
    <div v-if="showEditModal" class="modal-overlay">
      <div class="modal-content">
        <h3>编辑元素</h3>
        <div class="form-group">
          <label>元素名称:</label>
          <input v-model="editingElement.name" type="text" />
        </div>
        <div class="form-group">
          <label>元素内容:</label>
          <textarea v-model="editingElement.content" rows="5"></textarea>
        </div>
        <div class="form-actions">
          <button @click="saveElement">保存</button>
          <button @click="cancelEdit">取消</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import draggable from 'vue-draggable-next';

export default {
  components: {
    draggable,
  },
  data() {
    return {
      pageElements: [
        { id: 1, name: '标题', content: '<h1>欢迎使用页面构建器</h1>' },
        { id: 2, name: '段落', content: '<p>这是一个可拖拽的页面构建器示例</p>' },
        { id: 3, name: '图片', content: '<img src="https://picsum.photos/800/400" alt="示例图片" />' }
      ],
      draggableOptions: {
        animation: 200,
        handle: '.element-header',
        ghostClass: 'ghost-element'
      },
      showEditModal: false,
      editingElement: {},
      editingIndex: -1
    };
  },
  methods: {
    onDragStart(evt) {
      // 记录拖拽开始
      console.log('Drag started', evt);
      this.pageElements[evt.newIndex].isDragging = true;
    },
    onDragEnd(evt) {
      // 记录拖拽结束
      console.log('Drag ended', evt);
      // 重置拖拽状态
      this.pageElements.forEach(el => el.isDragging = false);
      // 保存页面布局
      this.savePageLayout();
    },
    addElement() {
      // 添加新元素
      const newElement = {
        id: Date.now(),
        name: `新元素 ${this.pageElements.length + 1}`,
        content: '<p>这是一个新元素</p>'
      };
      this.pageElements.push(newElement);
      // 自动打开编辑
      this.editElement(this.pageElements.length - 1);
    },
    editElement(index) {
      // 编辑元素
      this.editingIndex = index;
      this.editingElement = JSON.parse(JSON.stringify(this.pageElements[index]));
      this.showEditModal = true;
    },
    removeElement(index) {
      // 删除元素
      if (confirm('确定要删除这个元素吗？')) {
        this.pageElements.splice(index, 1);
        this.savePageLayout();
      }
    },
    saveElement() {
      // 保存编辑的元素
      this.pageElements[this.editingIndex] = JSON.parse(JSON.stringify(this.editingElement));
      this.showEditModal = false;
      this.savePageLayout();
    },
    cancelEdit() {
      // 取消编辑
      this.showEditModal = false;
    },
    savePageLayout() {
      // 保存页面布局到本地存储
      localStorage.setItem('pageLayout', JSON.stringify(this.pageElements));
      console.log('页面布局已保存');
    }
  },
  mounted() {
    // 从本地存储加载页面布局
    const savedLayout = localStorage.getItem('pageLayout');
    if (savedLayout) {
      this.pageElements = JSON.parse(savedLayout);
    }
  }
};
</script>

<style scoped>
.page-builder {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.drag-container {
  min-height: 300px;
  border: 1px dashed #ccc;
  padding: 10px;
  border-radius: 4px;
}

.element-card {
  background-color: #fff;
  border: 1px solid #ddd;
  border-radius: 4px;
  margin-bottom: 10px;
  transition: all 0.2s ease;
}

.element-card.dragging {
  opacity: 0.7;
  transform: scale(1.02);
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}

.element-header {
  padding: 10px;
  background-color: #f5f5f5;
  border-bottom: 1px solid #ddd;
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: move;
}

.element-content {
  padding: 10px;
}

.actions button {
  margin-left: 5px;
  padding: 4px 8px;
  background-color: #f0f0f0;
  border: 1px solid #ddd;
  border-radius: 3px;
  cursor: pointer;
}

.add-element-btn {
  margin-top: 20px;
  text-align: center;
}

.add-element-btn button {
  padding: 8px 16px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0,0,0,0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 4px;
  width: 400px;
  max-width: 90%;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
}

.form-group input, .form-group textarea {
  width: 100%;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.form-actions button {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.form-actions button:first-child {
  background-color: #4CAF50;
  color: white;
}

.form-actions button:last-child {
  background-color: #f44336;
  color: white;
}

.ghost-element {
  background-color: #e0e0e0;
}
</style>