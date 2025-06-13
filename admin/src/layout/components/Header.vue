<script setup>
import { computed } from 'vue'
import {
  FullScreen,
  Avatar,
  SwitchButton,
  ArrowDown,
  Reading,
  Sunny,
  Moon,
} from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { useDark, useToggle } from '@vueuse/core'
import { useUserStore } from '@/store/user'

const userStore = useUserStore()
const isDark = useDark()
const toggleDark = useToggle(isDark)
const icon = computed(() => (isDark.value ? Moon : Sunny))

const router = useRouter()

// 进入/退出全屏
const toggleFullscreen = () => {
  if (!document.fullscreenElement) {
    document.documentElement.requestFullscreen()
  } else {
    document.exitFullscreen()
  }
}

// 处理菜单命令
const handleCommand = (command) => {
  // 退出登录
  if (command === 'logout') {
    router.push('/login')
    localStorage.removeItem('token')
  }
  // 个人信息
  if (command === 'profile') {
    router.push('/user/profile')
  }
}
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
        <!-- 前往博客 -->
        <el-menu-item @click="$router.push('/')">
          <el-icon size="14"><Reading /></el-icon>
        </el-menu-item>

        <!-- 明暗切换 -->
        <el-menu-item @click="toggleDark()">
          <el-icon size="14">
            <component :is="icon"></component>
          </el-icon>
        </el-menu-item>

        <!-- 全屏按钮 -->
        <el-menu-item @click="toggleFullscreen">
          <el-icon size="14"><FullScreen /></el-icon>
        </el-menu-item>

        <!-- 用户信息 -->
        <el-menu-item>
          <el-dropdown trigger="click" @command="handleCommand">
            <div class="user-info">
              <el-avatar
                :size="35"
                :src="userStore.userInfo.avatar"
                fit="cover"
                v-if="userStore.userInfo.avatar"
              />
              <el-avatar :size="35" v-else>
                {{
                  userStore.userInfo.username ? userStore.userInfo.username[0].toUpperCase() : 'A'
                }}
              </el-avatar>
              <span class="username">
                {{ userStore.userInfo.username }}
                <el-icon size="14"><ArrowDown /></el-icon>
              </span>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">
                  <el-icon size="14"><Avatar /></el-icon>
                  个人信息
                </el-dropdown-item>
                <el-divider style="margin: 6px 0" />
                <el-dropdown-item command="logout">
                  <el-icon size="14"><SwitchButton /></el-icon>
                  退出登录
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
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
        }
      }
      .el-dropdown {
        .user-info {
          display: flex;
          align-items: center;
          .username {
            padding: 0 0 0 10px;
          }
          &:hover {
            color: var(--el-menu-active-color);
          }
        }
      }
    }
  }
}
</style>
