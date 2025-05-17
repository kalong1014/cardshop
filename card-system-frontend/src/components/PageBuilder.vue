<template>
  <div class="page-builder">
    <div class="sidebar">
      <el-card v-for="component in components" :key="component.id" class="component-item" draggable>
        {{ component.name }}
      </el-card>
    </div>
    
    <div class="canvas" v-draggable:parent="true">
      <div v-for="(item, index) in layout" :key="index" 
           :style="{ left: `${item.x}px`, top: `${item.y}px` }"
           v-draggable="{ group: 'components', animation: 300 }">
        {{ item.content }}
      </div>
    </div>
  </div>
</template>

<script>
import draggable from 'vuedraggable'

export default {
  components: { draggable },
  data() {
    return {
      components: [
        { id: 1, name: '标题组件' },
        { id: 2, name: '图片组件' },
        { id: 3, name: '按钮组件' }
      ],
      layout: []
    }
  },
  methods: {
    saveLayout() {
      const data = JSON.stringify(this.layout.map(item => ({
        type: item.type,
        x: item.x,
        y: item.y,
        content: item.content
      })))
      // 调用后端接口保存模板
      this.$api.saveTemplate({ name: '首页', data })
    }
  }
}
</script>