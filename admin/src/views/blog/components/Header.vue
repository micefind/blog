<script setup>
import { Edit, Search, Moon, Sunny, MoreFilled } from '@element-plus/icons-vue'
import { useDark, useToggle } from '@vueuse/core'
import { computed, ref } from 'vue'

const showInput = ref(false)

const isDark = useDark()
const toggleDark = useToggle(isDark)
const icon = computed(() => (isDark.value ? Moon : Sunny))
</script>

<template>
  <div class="header">
    <div class="logo">
      <router-link to="/">
        <img src="/public/logo.png" alt="logo" />
      </router-link>
      <h1>鼠觅</h1>
    </div>
    <div class="menu">
      <el-menu>
        <!-- 搜索按钮 -->
        <el-menu-item>
          <el-input
            v-if="showInput"
            style="width: 200px"
            placeholder="Pick a date"
            :suffix-icon="Search"
          />
          <el-icon size="14" @click="showInput = !showInput" v-if="!showInput"><Search /></el-icon>
        </el-menu-item>
        <!-- 筛选按钮 -->
        <el-menu-item>
          <el-icon size="14"><MoreFilled /></el-icon>
        </el-menu-item>
        <!-- 明暗切换 -->
        <el-menu-item @click="toggleDark()">
          <el-icon size="14">
            <component :is="icon"></component>
          </el-icon>
        </el-menu-item>
        <!-- 前往后台 -->
        <el-menu-item @click="$router.push('/article/list')">
          <el-icon size="14"><Edit /></el-icon>
        </el-menu-item>
      </el-menu>
    </div>
  </div>
</template>

<style scoped lang="scss">
.header {
  height: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .logo {
    height: 100%;
    display: flex;
    gap: 10px;
    align-items: center;
    padding: 0 10px;
    img {
      width: 40px;
    }
    h1 {
      font-size: 20px;
    }
  }
  .menu {
    height: 100%;
    padding: 0 10px;
    .el-menu {
      height: 100%;
      display: flex;
      border: none;
      .el-menu-item {
        height: 100%;
        padding: 0 10px;
        cursor: pointer;
        .el-icon {
          margin: 0;
        }
        &:hover {
          color: var(--el-menu-active-color);
          background-color: transparent;
        }
      }
    }
  }
}
</style>
