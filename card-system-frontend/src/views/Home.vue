<template>
  <div class="min-h-screen bg-gray-50 flex flex-col items-center justify-center">
    <div class="text-center mb-8">
      <h1 class="text-[clamp(2rem,5vw,4rem)] font-bold text-gray-900 mb-4">卡密销售系统</h1>
      <p class="text-[clamp(1rem,2vw,1.5rem)] text-gray-600 max-w-2xl mx-auto">
        简单高效的数字产品销售平台，帮助商家快速创建和管理卡密销售页面
      </p>
    </div>
    
    <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 mb-8">
      <router-link 
        to="/login" 
        class="px-8 py-3 bg-primary hover:bg-primary/90 text-white font-medium rounded-lg transition-all duration-300 shadow-lg hover:shadow-xl transform hover:-translate-y-1"
      >
        商家登录
      </router-link>
      <router-link 
        to="/register" 
        class="px-8 py-3 bg-white hover:bg-gray-50 text-primary border border-primary font-medium rounded-lg transition-all duration-300 shadow-lg hover:shadow-xl transform hover:-translate-y-1"
      >
        注册账号
      </router-link>
    </div>
    
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-5xl w-full">
      <div class="bg-white rounded-xl shadow-lg p-6 transition-all duration-300 hover:shadow-xl">
        <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mb-4">
          <i class="fa-solid fa-paint-brush text-primary text-xl"></i>
        </div>
        <h3 class="text-xl font-bold text-gray-900 mb-2">拖拽式页面设计</h3>
        <p class="text-gray-600">无需代码，通过简单拖拽即可创建专业的卡密销售页面</p>
      </div>
      
      <div class="bg-white rounded-xl shadow-lg p-6 transition-all duration-300 hover:shadow-xl">
        <div class="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mb-4">
          <i class="fa-solid fa-credit-card text-success text-xl"></i>
        </div>
        <h3 class="text-xl font-bold text-gray-900 mb-2">智能卡密管理</h3>
        <p class="text-gray-600">批量生成、导入导出卡密，自动核销，轻松管理卡密生命周期</p>
      </div>
      
      <div class="bg-white rounded-xl shadow-lg p-6 transition-all duration-300 hover:shadow-xl">
        <div class="w-12 h-12 bg-orange-100 rounded-full flex items-center justify-center mb-4">
          <i class="fa-solid fa-chart-line text-warning text-xl"></i>
        </div>
        <h3 class="text-xl font-bold text-gray-900 mb-2">实时数据分析</h3>
        <p class="text-gray-600">详细的销售数据统计和分析，帮助您了解业务表现和用户行为</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HomeView'
}
</script>