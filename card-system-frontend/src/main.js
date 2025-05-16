// 修改后的内容（包含所有配置）
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import ElementPlugin from './plugins/element'

// 导入Tailwind CSS
import './assets/main.css'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(ElementPlugin)

app.mount('#app')