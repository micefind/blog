<script setup>
import Main from './components/Main.vue'
import Info from './components/Info.vue'
</script>

<template>
  <div class="blog-page">
    <div class="container">
      <Main />
    </div>
    <div class="profile">
      <Info></Info>
    </div>
  </div>
</template>

<style scoped lang="scss">
.blog-page {
  display: flex;
  height: 100%;
  width: 100%;
  position: fixed;
  top: 0;
  left: 0;
  .container {
    width: calc(100% - 400px);
    height: 100%;
    .header {
      height: 50px;
      border-bottom: 1px solid var(--el-border-color);
    }
    .main {
      height: calc(100% - 50px);
      background-color: var(--content-bg-color);
    }
  }
  .profile {
    width: 400px;
    height: 100%;
    border-left: 1px solid var(--el-border-color);
  }
}

@media screen and (max-width: 768px) {
  .blog-page {
    .container {
      width: 100%;
    }
    .profile {
      display: none;
    }
  }
}
</style>
