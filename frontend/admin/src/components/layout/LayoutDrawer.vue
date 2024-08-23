<script setup lang="ts">
import { ref, watch } from 'vue';
import { useDark, useToggle } from "@vueuse/core";
import { useElementPlusTheme } from "use-element-plus-theme";
import { useAppStore } from '../../store/modules/app.ts';

const appStore = useAppStore();

// 明暗切换
const isDark = useDark();
const toggleDark = useToggle(isDark);

// 预定义的主题颜色数组
const predefineColors: string[] = ['#409eff', '#009688', '#ff5c93', '#0096c7', '#9c27b0', '#8a2c29', '#daa520'];

// 获取默认的主题颜色
const defaultTheme = ref<string>(localStorage.getItem('theme') ?? appStore.theme);

// 使用 Element Plus 主题更改功能
const { changeTheme } = useElementPlusTheme(defaultTheme.value);

// 监听主题颜色的变化，并将其存储在本地缓存中
watch(defaultTheme, (newTheme: string) => {
    localStorage.setItem('theme', newTheme);
});
</script>

<template>
    <el-divider>
        主题
    </el-divider>
    <div class="centent-box">
        <el-switch v-model="isDark" inline-prompt active-color="var(--el-fill-color-dark)" active-text="昼"
            inactive-text="夜" inactive-color="var(--el-color-primary)" @change="toggleDark" />
    </div>
    <el-divider>
        系统主题
    </el-divider>
    <div class="centent-box">
        <el-color-picker :predefine="predefineColors" v-model="defaultTheme" @change="changeTheme" />
    </div>
</template>

<style scoped lang="scss">
.centent-box {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 5px 0;
}
</style>
