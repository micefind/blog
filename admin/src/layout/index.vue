<script setup>
import { RouterView } from 'vue-router'
import Header from './components/Header.vue'
import Aside from './components/Aside.vue'
</script>

<template>
  <el-container class="layout">
    <el-header class="header">
      <Header />
    </el-header>
    <el-container>
      <el-aside class="aside">
        <Aside />
      </el-aside>
      <el-backtop target=".main .el-scrollbar__wrap" />
      <el-main class="main">
        <el-scrollbar>
          <RouterView />
        </el-scrollbar>
      </el-main>
    </el-container>
  </el-container>
</template>

<style scoped lang="scss">
.layout {
  height: 100%;
  width: 100%;
  position: fixed;
  top: 0;
  left: 0;

  .header {
    height: 50px;
    padding: 0;
    border-bottom: 1px solid var(--el-border-color);
  }
  .aside {
    border-right: 1px solid var(--el-border-color);
    width: auto;
    height: calc(100vh - 50px);
  }
  .main {
    padding: 0;
    height: calc(100vh - 50px);
    background-color: var(--content-bg-color);
  }
}
</style>
