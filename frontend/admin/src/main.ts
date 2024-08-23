import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from "./router";
import { createPinia } from 'pinia' 

import './style/element/element.css'

import 'element-plus/theme-chalk/dark/css-vars.css'
import './style/dark/css-vars.css'

createApp(App).use(router).use(createPinia()).mount('#app')

