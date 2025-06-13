<script setup>
import { ref, onMounted } from 'vue'
import VditorPreview from '@/components/VditorPreview.vue'
import { getArticleDetail } from '@/api/article'
import { useRoute, useRouter } from 'vue-router'
import { Back } from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()

const article = ref({})
// 获取文章详情
const getArticle = async (id) => {
  const res = await getArticleDetail({ id })
  if (res.code !== 200) return
  article.value = res.data
}

// 获取文章详情
onMounted(async () => {
  route.query.id && (await getArticle(parseInt(route.query.id)))
})
</script>

<template>
  <div class="article-page">
    <div class="header">
      <div class="title">
        <input type="text" v-model="article.title" placeholder="请输入文章标题" />
      </div>
      <div class="btn">
        <el-button plain @click="router.go(-1)" :icon="Back" size="small"></el-button>
      </div>
    </div>
    <div class="main">
      <VditorPreview v-model="article.content" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.article-page {
  height: 100%;
  position: relative;
  padding-top: 50px;

  .header {
    height: 50px;
    width: 100%;
    display: flex;
    align-items: center;
    border-bottom: 1px solid var(--el-border-color);
    position: fixed;
    top: 0;
    z-index: 999;
    background-color: var(--page-bg-color);

    .title {
      height: 100%;
      padding: 0 20px;
      flex: 1;
      input {
        border: none;
        height: 100%;
        width: 100%;
        font-size: 24px;
        background-color: transparent;
        &:focus {
          outline: none;
        }
      }
    }
    .btn {
      padding: 0 20px;
    }
  }
  .main {
    height: calc(100% - 50px);
  }
}
</style>
