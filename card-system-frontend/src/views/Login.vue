<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-bold text-gray-900">
          登录到您的账户
        </h2>
      </div>
      <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
        <div class="rounded-md shadow-sm -space-y-px">
          <div class="mb-4">
            <label for="username" class="sr-only">用户名</label>
            <input
              type="text"
              id="username"
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-primary focus:border-primary focus:z-10 sm:text-sm"
              placeholder="用户名"
              v-model="form.username"
              required
            />
          </div>
          <div>
            <label for="password" class="sr-only">密码</label>
            <input
              type="password"
              id="password"
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-primary focus:border-primary focus:z-10 sm:text-sm"
              placeholder="密码"
              v-model="form.password"
              required
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-primary hover:bg-primary/90 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary transition-all duration-300"
          >
            <span class="absolute left-0 inset-y-0 flex items-center pl-3">
              <i class="fa-solid fa-sign-in-alt text-white"></i>
            </span>
            登录
          </button>
        </div>
        <div class="text-center">
          <router-link to="/register" class="font-medium text-primary hover:text-primary/80">
            没有账户？注册
          </router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'

export default {
  setup() {
    const authStore = useAuthStore()
    const form = ref({
      username: '',
      password: ''
    })
    const loading = ref(false)
    const error = ref('')

    const handleLogin = async () => {
      loading.value = true
      error.value = ''

      try {
        const success = await authStore.login(form.value.username, form.value.password)
        if (success) {
          // 登录成功，重定向到首页或商户后台
          if (authStore.isAdmin) {
            window.location.href = '/admin'
          } else {
            window.location.href = '/merchant'
          }
        } else {
          error.value = '用户名或密码不正确'
        }
      } catch (err) {
        error.value = '登录失败，请重试'
        console.error(err)
      } finally {
        loading.value = false
      }
    }

    return {
      form,
      loading,
      error,
      handleLogin
    }
  }
}
</script>