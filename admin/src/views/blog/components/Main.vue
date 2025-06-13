<script setup>
import {
  Edit,
  Search,
  Moon,
  Sunny,
  CircleClose,
  Loading,
  Clock,
  View,
} from '@element-plus/icons-vue'
import { useDark, useToggle } from '@vueuse/core'
import { computed, ref, onMounted, nextTick } from 'vue'
import { getArticleList } from '@/api/article'
import { debounce } from 'lodash' // 引入防抖函数
import config from '@/config/config'

const articleList = ref([])
const searchInputRef = ref(null)
const total = ref(0)
const queryInfo = ref({
  page_num: 1,
  page_size: 10,
})
const keyword = ref('')
const loading = ref(false)
const is_error = ref(false)
const showInput = ref(false)

// 是否还有更多
const hasMore = computed(() => {
  return (queryInfo.value.page_num - 1) * queryInfo.value.page_size <= total.value
})

// 设置输入框显示状态
const setShowInput = async (status) => {
  if (status) {
    showInput.value = status
    await nextTick() // 等待DOM更新
    if (searchInputRef.value) {
      // 再次检查引用是否存在
      searchInputRef.value.focus()
    }
  } else {
    if (!keyword.value) {
      showInput.value = status
    }
  }
}

// 搜索文章
const searchArticle = debounce(async () => {
  if (!keyword.value && !queryInfo.value.keyword) return
  articleList.value = []
  queryInfo.value = {
    page_num: 1,
    page_size: 10,
    keyword: keyword.value,
  }
  await getArticleData()
}, 500)

// 加载更多
const load = () => {
  if (!hasMore.value || is_error.value) return
  getArticleData()
}

// 获取文章列表
const getArticleData = async () => {
  // 禁止重复请求
  if (loading.value) return
  // 重置数据
  loading.value = true
  is_error.value = false
  // 请求数据
  try {
    const res = await getArticleList(queryInfo.value)
    if (res.code !== 200) return
    total.value = res.data.total
    articleList.value = [...articleList.value, ...res.data.list]
    queryInfo.value.page_num++
    loading.value = false
    is_error.value = false
  } catch {
    is_error.value = true
  } finally {
    loading.value = false
  }
}
onMounted(async () => {
  await getArticleData()
})

const isDark = useDark()
const toggleDark = useToggle(isDark)
const icon = computed(() => (isDark.value ? Moon : Sunny))
</script>

<template>
  <div class="main-page">
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
          <el-menu-item v-if="showInput">
            <el-input
              clearable
              ref="searchInputRef"
              @blur="setShowInput(false)"
              @input="searchArticle"
              @keydown.enter="searchArticle"
              v-model="keyword"
              placeholder="请输入关键词搜索"
              :prefix-icon="Search"
            />
          </el-menu-item>
          <el-menu-item @click="setShowInput(true)" v-if="!showInput">
            <el-icon size="14"><Search /></el-icon>
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
    <el-backtop target=".content-box .el-scrollbar__wrap" />
    <div class="content-box">
      <el-scrollbar>
        <div class="article-box">
          <div
            class="article-list"
            v-infinite-scroll="load"
            infinite-scroll-distance="50"
            v-if="articleList.length"
          >
            <div
              class="article-item"
              v-for="item in articleList"
              :key="item.id"
              @click="$router.push({ path: '/blog/article', query: { id: item.id } })"
            >
              <div class="content">
                <div class="title">{{ item.title }}</div>
                <div class="intro">{{ item.intro }}</div>
                <div class="others">
                  <div class="tags">
                    <el-text type="info">
                      <el-icon><Clock /></el-icon>{{ item.create_time?.split(' ')[0] }}
                    </el-text>
                    <el-text type="info">
                      <el-icon><View /></el-icon>{{ item.views }}
                    </el-text>
                  </div>
                  <div class="tags">
                    <el-tag type="info" size="small">{{ item.cate_name }}</el-tag>
                    <el-tag type="info" size="small">{{ item.tag_name }}</el-tag>
                  </div>
                </div>
              </div>
              <img
                :src="config.file_base_url + item.cover_img"
                alt=""
                loading="lazy"
                class="cover"
                v-if="item.cover_img"
              />
            </div>
          </div>
          <div class="loading">
            <div v-if="loading" class="loading-text">
              <el-icon size="14px"><Loading /></el-icon>加载中
            </div>
            <div v-if="is_error" class="loading-text">
              <el-icon size="14px" color="#f56c6c"><CircleClose /></el-icon>加载失败
              <el-link type="primary" :underline="false" @click="getArticleData">重新加载</el-link>
            </div>
            <div class="loading-text" v-if="!hasMore">没有更多了...</div>
          </div>
        </div>
      </el-scrollbar>
    </div>
  </div>
</template>

<style scoped lang="scss">
.main-page {
  height: 100%;
  .header {
    height: 50px;
    border-bottom: 1px solid var(--el-border-color);
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
        @media screen and (max-width: 768px) {
          & {
            display: none;
          }
        }
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
          .el-input {
            width: 240px;
            :deep(svg) {
              font-size: 14px;
            }
          }
        }
      }
    }
  }
  .content-box {
    height: calc(100% - 50px);
    background-color: var(--content-bg-color);
    .article-box {
      padding: 20px;
      @media screen and (max-width: 768px) {
        & {
          padding: 0;
        }
      }

      .article-list {
        background-color: var(--el-bg-color-overlay);
        padding: 10px 0;
        .article-item {
          height: 127px;
          padding: 15px;
          color: var(--el-text-color-primary);
          border-bottom: 1px solid var(--el-border-color-lighter);
          cursor: pointer;
          display: flex;
          gap: 20px;
          &:last-child {
            border-bottom: none;
          }
          &:hover {
            background-color: var(--hover-bg-color);
          }
          .content {
            flex: 1;
            overflow: hidden;

            .title {
              font-size: 18px;
              font-weight: 700;
              text-overflow: ellipsis;
              white-space: nowrap;
              overflow: hidden;
            }
            .intro {
              padding: 15px 0;
              overflow: hidden;
              text-overflow: ellipsis;
              white-space: nowrap;
            }
            .others {
              display: flex;
              justify-content: space-between;
              align-items: center;
              overflow: hidden;
              white-space: nowrap;
              .tags {
                display: flex;
                gap: 10px;
                :deep(.el-tag__content) {
                  font-size: 12px;
                }
                .el-icon {
                  margin-right: 5px;
                }
                .el-text {
                  font-size: 13px;
                }
              }
            }
          }
          .cover {
            width: 150px;
            height: 100%;
            border: 1px solid var(--el-border-color-lighter);
            border-radius: 3px;
            object-fit: cover;
          }
        }
      }
      .loading {
        display: flex;
        align-items: center;
        justify-content: center;
        margin-top: 20px;
        .loading-text {
          display: flex;
          align-items: center;
          gap: 10px;
          color: var(--el-text-color-secondary);
        }
      }
    }
  }
}
</style>
