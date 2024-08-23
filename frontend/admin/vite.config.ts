import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
import path from "path"
import AutoImport from "unplugin-auto-import/vite"
import Components from "unplugin-vue-components/vite"
import { ElementPlusResolver } from "unplugin-vue-components/resolvers"

// https://vitejs.dev/config/
export default defineConfig({
  base: "./", // 文件根目录为相对路径
  plugins: [
    vue(),
    // 按需自动引入element-plus
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
  resolve: {
    // Vite路径别名配置
    alias: {
      "@": path.resolve("./src"),
    },
  },
  // 项目启动后自动打开浏览器
  server: {
    host: "127.0.0.1", // 指定服务器应该监听哪个 IP 地址
    port: 9500, // 指定开发服务器端口
    strictPort: false, // 设为 true 时若端口已被占用则会直接退出，而不是尝试下一个可用端口
    open: true, // 开发服务器启动时，自动在浏览器中打开应用程序
  },
})
