<script setup>
import { defineModel } from 'vue'
import { ref, watch } from 'vue'
import 'vditor/dist/index.css'
import Vditor from 'vditor'
import config from '@/config/config'
import { useDark } from '@vueuse/core'

const isDark = useDark()
const content = defineModel()
const previewRef = ref(null)
const outlineRef = ref(null)

const renderPreview = () => {
  Vditor.preview(previewRef.value, content.value || '', {
    theme: {
      current: isDark.value ? 'Dark' : 'Ant Design',
    },
    speech: {
      enable: true,
    },
    hljs: {
      enable: true,
      style: isDark.value ? 'github-dark' : 'github',
      lineNumber: false,
    },
    markdown: {
      toc: true,
      mark: true,
      footnotes: true,
      linkBase: config.file_base_url,
    },

    after() {
      // 自动渲染目录
      Vditor.outlineRender(previewRef.value, outlineRef.value)
      if (outlineRef.value.innerText.trim() !== '') {
        outlineRef.value.style.display = 'block'
      }
    },
  })
}

watch(
  () => content.value,
  () => {
    renderPreview()
  },
)
</script>

<template>
  <div class="vditor-preview-page">
    <div ref="outlineRef" class="outline"></div>
    <div ref="previewRef" class="preview"></div>
  </div>
</template>

<style scoped lang="scss">
.vditor-preview-page {
  position: relative;
  .preview {
    border: none;
    padding: 0 50px 100px 300px;
  }

  .outline {
    position: fixed;
    top: 50px;
    left: 0;
    width: 250px;
    height: calc(100% - 50px);
    padding: 20px 0;
    overflow: auto;
    border-right: 1px solid var(--el-border-color);
    --toolbar-icon-hover-color: var(--el-color-primary);
    --textarea-text-color: var(--el-text-color-primary);
  }
}
@media screen and (max-width: 768px) {
  .vditor-preview-page {
    .preview {
      padding: 0 10px 50px;
    }
    .outline {
      width: 0;
      display: none;
    }
  }
}
</style>
