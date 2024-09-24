<script setup lang="ts">
import { useElementPlusTheme } from "use-element-plus-theme"
import { useAppStore } from './store/modules/app.ts'
import { useDark, useToggle } from "@vueuse/core"

// 获取当前时间
const currentHour = new Date().getHours()

// 判断是否是晚上
const isNight = currentHour >= 18 || currentHour < 6

// 初始化暗黑模式
const isDark = useDark()
const toggleDark = useToggle(isDark)

// 如果是晚上启用暗黑模式，否则关闭暗黑模式
if (isNight) {
  toggleDark(true)
} else {
  toggleDark(false)
}

// 设置全局主题色
const appStore = useAppStore()
const { changeTheme } = useElementPlusTheme(localStorage.getItem('theme') ? localStorage.getItem('theme') || '' : appStore.theme)
changeTheme()
</script>

<template>
  <router-view />
</template>

<style scoped></style>
