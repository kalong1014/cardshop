<template>
  <div class="min-h-screen bg-gray-50">
    <!-- 顶部导航栏 -->
    <div class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <router-link to="/merchant" class="flex-shrink-0 flex items-center">
              <i class="fa-solid fa-store text-primary text-2xl mr-2"></i>
              <span class="font-bold text-xl">商户中心</span>
            </router-link>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <router-link
                to="/merchant"
                class="border-primary text-gray-900 border-b-2 px-1 pt-1 inline-flex items-center text-sm font-medium"
              >
                <i class="fa-solid fa-tachometer-alt mr-1"></i> 概览
              </router-link>
              <router-link
                to="/merchant/pages"
                class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 border-b-2 px-1 pt-1 inline-flex items-center text-sm font-medium transition-colors"
              >
                <i class="fa-solid fa-file-alt mr-1"></i> 页面管理
              </router-link>
              <router-link
                to="/merchant/cards"
                class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 border-b-2 px-1 pt-1 inline-flex items-center text-sm font-medium transition-colors"
              >
                <i class="fa-solid fa-credit-card mr-1"></i> 卡密管理
              </router-link>
              <router-link
                to="/merchant/orders"
                class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 border-b-2 px-1 pt-1 inline-flex items-center text-sm font-medium transition-colors"
              >
                <i class="fa-solid fa-shopping-cart mr-1"></i> 订单管理
              </router-link>
            </div>
          </div>
          <div class="flex items-center">
            <div class="ml-3 relative">
              <div>
                <button type="button" class="flex text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary" id="user-menu-button">
                  <span class="sr-only">打开用户菜单</span>
                  <img class="h-8 w-8 rounded-full" src="https://picsum.photos/200/200" alt="用户头像">
                </button>
              </div>
              <div class="hidden origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none" id="user-menu" role="menu" aria-orientation="vertical" aria-labelledby="user-menu-button" tabindex="-1">
                <router-link to="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" role="menuitem" tabindex="-1" id="user-menu-item-0">
                  <i class="fa-solid fa-user mr-1"></i> 个人资料
                </router-link>
                <router-link to="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" role="menuitem" tabindex="-1" id="user-menu-item-1">
                  <i class="fa-solid fa-cog mr-1"></i> 设置
                </router-link>
                <button @click="logout" class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" role="menuitem" tabindex="-1" id="user-menu-item-2">
                  <i class="fa-solid fa-sign-out-alt mr-1"></i> 退出登录
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 主要内容 -->
    <main>
      <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div class="bg-white overflow-hidden shadow-sm sm:rounded-lg">
          <div class="p-6 bg-white border-b border-gray-200">
            <h2 class="text-2xl font-bold text-gray-900 mb-6">商户概览</h2>
            
            <!-- 数据卡片 -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
              <div class="bg-white rounded-lg shadow-md p-6 transition-all duration-300 hover:shadow-lg">
                <div class="flex items-center">
                  <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mr-4">
                    <i class="fa-solid fa-file-alt text-primary text-xl"></i>
                  </div>
                  <div>
                    <p class="text-sm font-medium text-gray-500">页面总数</p>
                    <p class="text-2xl font-bold text-gray-900">12</p>
                  </div>
                </div>
              </div>
              
              <div class="bg-white rounded-lg shadow-md p-6 transition-all duration-300 hover:shadow-lg">
                <div class="flex items-center">
                  <div class="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mr-4">
                    <i class="fa-solid fa-credit-card text-success text-xl"></i>
                  </div>
                  <div>
                    <p class="text-sm font-medium text-gray-500">可用卡密</p>
                    <p class="text-2xl font-bold text-gray-900">1,254</p>
                  </div>
                </div>
              </div>
              
              <div class="bg-white rounded-lg shadow-md p-6 transition-all duration-300 hover:shadow-lg">
                <div class="flex items-center">
                  <div class="w-12 h-12 bg-orange-100 rounded-full flex items-center justify-center mr-4">
                    <i class="fa-solid fa-shopping-cart text-warning text-xl"></i>
                  </div>
                  <div>
                    <p class="text-sm font-medium text-gray-500">今日订单</p>
                    <p class="text-2xl font-bold text-gray-900">32</p>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- 快速操作 -->
            <div class="mb-8">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">快速操作</h3>
              <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
                <button class="bg-primary hover:bg-primary/90 text-white font-medium py-2 px-4 rounded-lg transition-all duration-300 flex items-center justify-center">
                  <i class="fa-solid fa-plus mr-2"></i> 新建页面
                </button>
                <button class="bg-secondary hover:bg-secondary/90 text-white font-medium py-2 px-4 rounded-lg transition-all duration-300 flex items-center justify-center">
                  <i class="fa-solid fa-plus mr-2"></i> 生成卡密
                </button>
                <button class="bg-info hover:bg-info/90 text-white font-medium py-2 px-4 rounded-lg transition-all duration-300 flex items-center justify-center">
                  <i class="fa-solid fa-paint-brush mr-2"></i> 设计模板
                </button>
                <button class="bg-success hover:bg-success/90 text-white font-medium py-2 px-4 rounded-lg transition-all duration-300 flex items-center justify-center">
                  <i class="fa-solid fa-credit-card mr-2"></i> 查看收入
                </button>
              </div>
            </div>
            
            <!-- 最近订单 -->
            <div>
              <h3 class="text-lg font-semibold text-gray-900 mb-4">最近订单</h3>
              <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">订单ID</th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">商品</th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">金额</th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">支付方式</th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">状态</th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">时间</th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">#ORD-12345</td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">游戏点卡</td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">¥99.00</td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">微信支付</td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">已支付</span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">2025-05-16 10:23</td>
                    </tr>
                    <tr>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">#ORD-12344</td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">软件激活码</td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">¥199.00</td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">支付宝</td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">已支付</span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">2025-05-16 09:47</td>
                    </tr>
                    <tr>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">#ORD-12343</td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">会员月卡</td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">¥39.00</td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">微信支付</td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">已支付</span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">2025-05-16 08:15</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import http from '@/utils/http'

export default {
  setup() {
    const authStore = useAuthStore()
    const user = ref(null)
    const stats = ref({
      pageCount: 0,
      cardCount: 0,
      todayOrders: 0
    })
    const recentOrders = ref([])
    const loading = ref(true)
    const error = ref('')

    const logout = () => {
      authStore.logout()
      window.location.href = '/login'
    }

    const fetchDashboardData = async () => {
      try {
        // 获取用户信息
        const userResponse = await http.get('/user/me')
        user.value = userResponse.data
        
        // 获取统计数据
        const statsResponse = await http.get('/merchant/dashboard/stats')
        stats.value = statsResponse.data
        
        // 获取最近订单
        const ordersResponse = await http.get('/merchant/orders/recent')
        recentOrders.value = ordersResponse.data
      } catch (err) {
        error.value = '获取数据失败，请重试'
        console.error(err)
      } finally {
        loading.value = false
      }
    }

    onMounted(() => {
      fetchDashboardData()
    })

    return {
      user,
      stats,
      recentOrders,
      loading,
      error,
      logout
    }
  }
}
</script>