<script setup lang="ts">
import {
  FullScreen,
  Setting,
  Fold,
  Expand,
  Avatar,
  SwitchButton,
} from "@element-plus/icons-vue"
import { onMounted, ref } from "vue"
import LayoutDrawer from "./LayoutDrawer.vue"
import { useAppStore } from "../../store/modules/app.ts"
import { useTagsViewStore } from "../../store/modules/tagsView.ts"
import type { DropdownInstance } from "element-plus"
import { useRouter } from "vue-router"
import avatar from "../../assets/image/defaultAvatar.png"
import logo from "../../assets/image/logo.png"

const router = useRouter()
const appStore = useAppStore()
const tagsViewStore = useTagsViewStore()
const drawer = ref(false)
const fullscreen = ref(false)
const dropdown = ref<DropdownInstance>()

const showClick = () => {
  if (!dropdown.value) return
  dropdown.value.handleOpen()
}

const changeCollapse = () => {
  appStore.changeCollapse()
}

const avatarHandleCommand = (command: string) => {
  switch (command) {
    case "profile":
      router.push({ path: "/profile" })
      tagsViewStore.setActivePath({ path: "/profile", name: "个人中心" })
      break
    case "logout":
      router.push({ path: "/login" })
      break
  }
}

// 进入退出全屏操作
const screen = () => {
  const element: HTMLElement & {
    webkitRequestFullScreen?: () => Promise<void>
    mozRequestFullScreen?: () => Promise<void>
    msRequestFullscreen?: () => Promise<void>
    webkitCancelFullScreen?: () => void
    mozCancelFullScreen?: () => void
    msExitFullscreen?: () => void
  } = document.documentElement

  if (fullscreen.value) {
    if (document.exitFullscreen) {
      document.exitFullscreen()
    } else if (element.webkitCancelFullScreen) {
      element.webkitCancelFullScreen()
    } else if (element.mozCancelFullScreen) {
      element.mozCancelFullScreen()
    } else if (element.msExitFullscreen) {
      element.msExitFullscreen()
    }
  } else {
    if (element.requestFullscreen) {
      element.requestFullscreen()
    } else if (element.webkitRequestFullScreen) {
      element.webkitRequestFullScreen()
    } else if (element.mozRequestFullScreen) {
      element.mozRequestFullScreen()
    } else if (element.msRequestFullscreen) {
      element.msRequestFullscreen()
    }
  }

  document.addEventListener("keydown", (e: KeyboardEvent) => {
    if (e.key === "F11") {
      e.preventDefault()
      screen()
    }
  })

  fullscreen.value = !fullscreen.value
}

onMounted(() => {
  window.addEventListener("resize", () => {
    fullscreen.value = window.innerHeight === window.screen.height
  })
})
</script>

<template>
  <el-menu mode="horizontal" :ellipsis="false">
    <el-menu-item class="increase-spacing">
      <img :src="logo" style="height: 40px" />
      <h1>鼠觅奇物</h1>
    </el-menu-item>
    <el-menu-item @click="changeCollapse" v-if="!appStore.isMobile">
      <el-icon v-show="appStore.isCollapse" class="collapse-btn">
        <Fold />
      </el-icon>
      <el-icon v-show="!appStore.isCollapse" class="collapse-btn">
        <Expand />
      </el-icon>
    </el-menu-item>
    <div class="center-box" v-if="!appStore.isMobile">
      <!-- Vite + Vue3 + TypeScript -->
    </div>
    <el-menu-item @click="screen">
      <el-icon>
        <FullScreen />
      </el-icon>
    </el-menu-item>
    <el-menu-item @click="drawer = true">
      <el-icon>
        <Setting />
      </el-icon>
    </el-menu-item>
    <el-menu-item class="increase-spacing" @click="showClick">
      <el-dropdown
        ref="dropdown"
        trigger="contextmenu"
        @command="avatarHandleCommand"
      >
        <el-avatar :size="40" fit="cover" :src="avatar" />
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile">
              <el-icon>
                <Avatar />
              </el-icon>
              个人中心
            </el-dropdown-item>
            <el-divider style="margin: 6px 0" />
            <el-dropdown-item command="logout">
              <el-icon>
                <SwitchButton />
              </el-icon>
              退出系统
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </el-menu-item>
  </el-menu>

  <el-drawer v-model="drawer" title="项目配置" size="350px">
    <el-scrollbar>
      <layout-drawer />
    </el-scrollbar>
  </el-drawer>
</template>

<style scoped lang="scss">
.el-menu {
  height: 100%;
  display: flex;
  justify-content: space-between;
}

.center-box {
  flex: 1;
  height: 100%;
  display: flex;
  align-items: center;
  padding: 0 10px;
}

.el-menu-item {
  padding: 0 5px;
}

.increase-spacing {
  padding: 0 15px;

  h1 {
    font-size: 18px;
    padding: 0 10px;
    height: 49px;
    line-height: 49px;
  }
}

.el-menu-item:hover {
  background-color: #f6f6f6 !important;
}

.el-menu-item:focus {
  background-color: transparent !important;
}

.collapse-btn {
  svg {
    font-size: 20px;
  }
}
</style>
