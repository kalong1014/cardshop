// src/stores/auth.js
import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: localStorage.getItem('token') || null,
    role: localStorage.getItem('role') || null
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
    isAdmin: (state) => state.role === 'admin'
  },
  actions: {
    async login(username, password) {
      try {
        const response = await axios.post('http://localhost:8080/api/login', {
          username,
          password
        })
        
        this.user = response.data.user
        this.token = response.data.token
        this.role = response.data.user.role
        
        // 存储到localStorage
        localStorage.setItem('token', response.data.token)
        localStorage.setItem('role', response.data.user.role)
        
        // 设置axios的Authorization头
        axios.defaults.headers.common['Authorization'] = `Bearer ${response.data.token}`
        
        return true
      } catch (error) {
        console.error('Login failed:', error)
        return false
      }
    },
    logout() {
      this.user = null
      this.token = null
      this.role = null
      
      localStorage.removeItem('token')
      localStorage.removeItem('role')
      
      // 清除axios的Authorization头
      delete axios.defaults.headers.common['Authorization']
    },
    async fetchUser() {
      try {
        const response = await axios.get('http://localhost:8080/api/user/me')
        this.user = response.data
        return true
      } catch (error) {
        console.error('Fetch user failed:', error)
        this.logout()
        return false
      }
    }
  }
})