<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-bold text-gray-900">
          创建新账户
        </h2>
      </div>
      <form class="mt-8 space-y-6" @submit.prevent="handleRegister">
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
          <div class="mb-4">
            <label for="email" class="sr-only">邮箱</label>
            <input
              type="email"
              id="email"
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-primary focus:border-primary focus:z-10 sm:text-sm"
              placeholder="邮箱"
              v-model="form.email"
              required
            />
          </div>
          <div class="mb-4">
            <label for="password" class="sr-only">密码</label>
            <input
              type="password"
              id="password"
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-primary focus:border-primary focus:z-10 sm:text-sm"
              placeholder="密码 (至少6个字符)"
              v-model="form.password"
              required
            />
          </div>
          <div>
            <label for="confirmPassword" class="sr-only">确认密码</label>
            <input
              type="password"
              id="confirmPassword"
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-primary focus:border-primary focus:z-10 sm:text-sm"
              placeholder="确认密码"
              v-model="form.confirmPassword"
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
              <i class="fa-solid fa-user-plus text-white"></i>
            </span>
            注册
          </button>
        </div>
        <div class="text-center">
          <router-link to="/login" class="font-medium text-primary hover:text-primary/80">
            已有账户？登录
          </router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import http from '@/utils/http'

export default {
  setup() {
    const form = ref({
      username: '',
      email: '',
      password: '',
      confirmPassword: ''
    })
    const loading = ref(false)
    const error = ref('')

    const handleRegister = async () => {
      loading.value = true
      error.value = ''

      // 验证密码匹配
      if (form.value.password !== form.value.confirmPassword) {
        error.value = '密码不匹配'
        loading.value = false
        return
      }

      try {
        const response = await http.post('/register', {
          username: form.value.username,
          email: form.value.email,
          password: form.value.password
        })

        // 注册成功，自动登录
        window.$message.success('注册成功，请登录')
        window.location.href = '/login'
      } catch (err) {
        if (err.response && err.response.data && err.response.data.error) {
          error.value = err.response.data.error
        } else {
          error.value = '注册失败，请重试'
        }
        console.error(err)
      } finally {
        loading.value = false
      }
    }

    return {
      form,
      loading,
      error,
      handleRegister
    }
  }
}
</script>