<script setup>
import { useAppStore } from '@/store/app'
import routes from '@/config/routes'
import { Expand, Fold } from '@element-plus/icons-vue'

const appStore = useAppStore()
</script>

<template>
  <div class="aside-page">
    <el-scrollbar class="menu-scrollbar">
      <el-menu
        :collapse="appStore.sidebarCollapsed"
        :router="true"
        :default-active="appStore.activePath"
        :collapse-transition="true"
        :unique-opened="true"
      >
        <template v-for="item in routes" :key="item.path">
          <el-sub-menu v-if="item.children" :index="item.path">
            <template #title>
              <el-icon size="14">
                <component :is="item.meta.icon" />
              </el-icon>
              <span>{{ item.meta.title }}</span>
            </template>
            <template v-for="child in item.children" :key="child.path">
              <el-sub-menu :index="child.path" v-if="child.children">
                <template #title>
                  <span>{{ child.meta.title }}</span>
                </template>
                <el-menu-item
                  v-for="grandchild in child.children"
                  :key="grandchild.path"
                  :index="grandchild.path"
                  >{{ grandchild.meta.title }}</el-menu-item
                >
              </el-sub-menu>
              <el-menu-item :index="child.path" v-else>{{ child.meta.title }}</el-menu-item>
            </template>
          </el-sub-menu>
          <el-menu-item :index="item.path" v-else>
            <el-icon size="14">
              <component :is="item.meta.icon" />
            </el-icon>
            <span>{{ item.meta.title }}</span>
          </el-menu-item>
        </template>
      </el-menu>
    </el-scrollbar>

    <div class="aside-bottom">
      <el-tooltip
        effect="light"
        :content="appStore.sidebarCollapsed ? '点击展开' : '点击折叠'"
        placement="right"
      >
        <div class="collapse">
          <el-icon size="18" @click="appStore.toggleSidebar()">
            <component :is="appStore.sidebarCollapsed ? Expand : Fold" />
          </el-icon>
        </div>
      </el-tooltip>
    </div>
  </div>
</template>

<style scoped lang="scss">
.aside-page {
  height: 100%;
  .menu-scrollbar {
    height: calc(100% - 50px);
    .el-menu {
      border: none;
      &:not(.el-menu--collapse) {
        width: 200px; // 未折叠时的菜单宽度
      }
    }
  }
  .aside-bottom {
    height: 50px;
    display: flex;
    align-items: center;
    border-top: 1px solid var(--el-border-color);
    padding: 0 10px;
  }
}
</style>
