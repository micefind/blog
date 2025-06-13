import './styles/main.scss'
import 'element-plus/theme-chalk/dark/css-vars.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import { router, initRouter } from './router'
import setupDirectives from './directive'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'

const initApp = () => {
  const app = createApp(App)
  // 注册指令
  setupDirectives(app)
  // 初始化路由
  initRouter()
  app.use(createPinia())
  app.use(router)
  app.use(ElementPlus, {
    locale: zhCn,
  })

  app.mount('#app')
}

initApp()
