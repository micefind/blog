<script setup>
import { defineModel } from 'vue'
import { debounce } from 'lodash' // 引入防抖函数
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import 'vditor/dist/index.css'
import Vditor from 'vditor'
import config from '@/config/config'
import { useDark } from '@vueuse/core'

const isDark = useDark()
const content = defineModel()
const vditorRef = ref(null)
const vditor = ref(null)
const isInitEditor = ref(false)
const isFirstUpdate = ref(true)

// 监听 content 改变
watch(
  () => content.value,
  (newValue) => {
    if (isFirstUpdate.value && isInitEditor.value) {
      isInitEditor.value && vditor.value.setValue(newValue)
      isFirstUpdate.value = false
    }
  },
)

onMounted(() => {
  vditor.value = new Vditor(vditorRef.value, {
    toolbar: [
      'emoji',
      'headings',
      'bold',
      'italic',
      'strike',
      '|',
      'line',
      'quote',
      'list',
      'ordered-list',
      'check',
      'outdent',
      'indent',
      'code',
      'inline-code',
      'upload',
      'link',
      'table',
      '|',
      'edit-mode',
      'both',
      'preview',
      'fullscreen',
      'outline',
      'code-theme',
      'export',
      '|',
      'undo',
      'redo',
    ],
    value: content.value,
    height: '100%',
    theme: isDark.value ? 'dark' : 'classic',
    preview: {
      theme: { current: isDark.value ? 'Dark' : 'Ant Design' },
      markdown: {
        linkBase: config.file_base_url,
      },
      hljs: {
        style: isDark.value ? 'github-dark' : 'github',
        enable: true,
        lineNumber: false,
      },
    },
    mode: 'ir', // 即时渲染模式
    typewriterMode: true, // 启用打字机模式
    counter: { enable: true, type: 'text' }, // 启用计数器功能
    cache: { enable: false }, // 禁用缓存功能
    // 启用大纲并设置为默认显示
    outline: {
      enable: true, // 启用大纲功能
      position: 'left', // 大纲位置：left 或 right
    },
    // 添加内容变化监听
    input: debounce(async (value) => {
      content.value = value
    }, 2000),
    upload: {
      url: config.api_base_url + '/upload/image',
      headers: { Authorization: localStorage.getItem('token') || '' },
      accept: 'image/jpeg,image/png,image/gif,image/jpg',
      max: 20 * 1024 * 1024,
      multiple: false,
      fieldName: 'file',
      filename: (name) =>
        name
          .replace(/[^\w\u4e00-\u9fa5.)]/g, '') // 保留字母、数字、下划线、中文和点、括号
          .replace(/[?\\/:|<>*[\]()$%{}@~]/g, ''), // 移除不合法符号
      format: (_, responseText) => {
        const res = JSON.parse(responseText)
        vditor.value?.insertValue(`![${res.data.file_name}](${res.data.file_path})`)
        return JSON.stringify({ code: 0, data: { errFiles: '', succMap: {} } })
      },
    },
    // Vditor初始化完成后的回调
    after: () => {
      if (content.value) {
        vditor.value.setValue(content.value)
        isFirstUpdate.value = false
      } else {
        vditor.value?.focus()
      }
      isInitEditor.value = true
    },
  })
})

onBeforeUnmount(() => {
  if (vditor.value) {
    vditor.value.destroy()
    vditor.value = null
  }
})
</script>

<template>
  <div class="vditor-page">
    <div ref="vditorRef" class="vditor"></div>
  </div>
</template>

<style scoped lang="scss">
.vditor-page {
  height: 100%;
  :deep(.vditor) {
    height: 100%;
    border: none;
    /*定义滚动条高宽及背景 高宽分别对应横竖滚动条的尺寸*/
    ::-webkit-scrollbar {
      width: 7px;
      height: 7px;
    }
    /*定义滑块 内阴影*/
    ::-webkit-scrollbar-thumb {
      background: var(--el-border-color);
      border-radius: 4px;
    }
    .vditor-toolbar {
      border-top: 1px solid var(--border-color);
    }
  }
}
</style>
