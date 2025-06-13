// 引入 ESLint 的扁平配置方法
import { defineConfig, globalIgnores } from 'eslint/config'

// 引入浏览器全局变量定义（如 window、document）
import globals from 'globals'

// 引入官方 JS 推荐配置（eslint:recommended）
import js from '@eslint/js'

// 引入 Vue 的 ESLint 插件
import pluginVue from 'eslint-plugin-vue'

// 跳过格式化相关规则（用于配合 Prettier）
import skipFormatting from '@vue/eslint-config-prettier/skip-formatting'

// 导出扁平配置数组
export default defineConfig([
  {
    // 指定要 lint 的文件类型（JS/JSX/Vue 等）
    name: 'app/files-to-lint',
    files: ['**/*.{js,mjs,jsx,vue}'],
  },

  // 忽略某些目录（如构建产物、覆盖率报告）
  globalIgnores(['**/dist/**', '**/dist-ssr/**', '**/coverage/**']),

  {
    // 设置语言环境中的全局变量（如 window、document）
    languageOptions: {
      globals: {
        ...globals.browser,
        ElMessage: 'readonly',
        ElMessageBox: 'readonly',
        ElLoading: 'readonly',
      },
    },
  },

  // 启用 ESLint 官方 JS 推荐规则
  js.configs.recommended,

  // 启用 Vue 3 的基础规则配置（flat 模式）
  ...pluginVue.configs['flat/essential'],

  // 关闭与 Prettier 冲突的格式化规则
  skipFormatting,

  // 自定义 Vue 规则
  {
    name: 'custom/vue-rules',
    files: ['**/*.vue'],
    rules: {
      // 允许组件名为 index.vue（单词组件名）
      'vue/multi-word-component-names': 'off', // 也可以设置为 'warn'
    },
  },
])
