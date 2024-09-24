<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, watch, defineEmits, defineProps,nextTick } from 'vue';
import Vditor from 'vditor';
import 'vditor/dist/index.css';
import { netConfig } from "../../config/net.config";

// 判断是否黑夜模式
const isDarkMode = () => document.documentElement.classList.contains('dark');

const props = defineProps({
  modelValue: String,
});

const emit = defineEmits(['update:modelValue']);
const vditor = ref<Vditor | null>(null);

// 将modelValue与vditor内容同步
watch(() => props.modelValue, (newValue) => {
  if (vditor.value?.getValue() !== newValue) {
    vditor.value?.setValue(newValue || '');
  }
});

// 初始化编辑器
const initVditor = (isDark: boolean) => {
  vditor.value = new Vditor('vditor', {
    height: '100%',
    value: props.modelValue || '',
    theme: isDark ? 'dark' : 'classic',
    toolbar: [
      'emoji', 'headings', 'bold', 'italic', 'strike', '|', 'line', 'quote', 'list',
      'ordered-list', 'check', 'outdent', 'indent', 'code', 'inline-code', 'insert-before',
      'insert-after', 'undo', 'redo', 'upload', 'link', 'table', 'record', 'edit-mode',
      'both', 'preview', 'fullscreen', 'outline', 'code-theme', 'content-theme', 'export',
      'devtools', 'br'
    ],
    input: (value: string) => emit('update:modelValue', value),
    preview: {
      theme: { current: isDark ? 'dark' : 'Ant Design' },
    },
    counter: { enable: true },
    cache: { enable: false },
    typewriterMode: true,
    upload: {
      url: netConfig.baseURL + '/upload/image',
      headers: { Authorization: localStorage.getItem('token') || '' },
      accept: 'image/jpeg,image/png,image/gif,image/jpg',
      max: 20 * 1024 * 1024,
      multiple: false,
      fieldName: 'file',
      filename: (name) => name.replace(/[^\w\u4e00-\u9fa5\.)]/g, '').replace(/[\?\\/:|<>\*\[\]\(\)\$%\{\}@~]/g, ''),
      format: (_, responseText) => {
        const res = JSON.parse(responseText);
        vditor.value?.insertValue(`![${res.data.file_name}](${res.data.url})`);
        return JSON.stringify({ code: 0, data: { errFiles: '', succMap: {} } });
      },
    },
    after: () => vditor.value?.focus(),
  });
};

// 确保vditor元素在dom中已经存在后再初始化
onMounted(() => {
  setTimeout(() => {
    nextTick(() => {
      initVditor(isDarkMode());
    });
  }, 100);
});

onBeforeUnmount(() => {
  vditor.value?.destroy();
});
</script>

<template>
  <div id="vditor"></div>
</template>

<style scoped lang="scss">
#vditor {
  height: 100%;
  border: none;
}
</style>
