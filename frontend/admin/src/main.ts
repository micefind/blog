import { createApp } from "vue"
import "./style.css"
import App from "./App.vue"
import router from "./router"
import { createPinia } from "pinia"
import "./style/element/element.css"
import "element-plus/theme-chalk/dark/css-vars.css"
import "./style/dark/css-vars.css"
// 引入element-plus插件与样式
import ElementPlus from "element-plus"
import locale from "element-plus/es/locale/lang/zh-cn"
import "element-plus/dist/index.css"

const app = createApp(App)

app.use(router)
app.use(createPinia())
app.use(ElementPlus, { locale })

app.mount("#app")
