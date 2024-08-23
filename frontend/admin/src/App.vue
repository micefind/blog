<script setup lang="ts">
import { useElementPlusTheme } from "use-element-plus-theme"
import { useAppStore } from './store/modules/app.ts'
import { useDark } from "@vueuse/core"

useDark()
const appStore = useAppStore()

// 设置全局主题色
const { changeTheme } = useElementPlusTheme(localStorage.getItem('theme') ? localStorage.getItem('theme') || '' : appStore.theme)
changeTheme()
</script>

<template>
  <router-view />
</template>

<style scoped></style>
