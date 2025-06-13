<script setup>
import { ref, onMounted } from 'vue'
import Vditor from '@/components/Vditor.vue'
import { getArticleDetail, createArticle, updateArticle } from '@/api/article'
import { useRoute, useRouter } from 'vue-router'
import { getArticleCateList } from '@/api/article/cate'
import { getArticleTagList } from '@/api/article/tag'
import { Plus } from '@element-plus/icons-vue'
import { uploadImage } from '@/api/upload'
import config from '@/config/config'

const route = useRoute()
const router = useRouter()
const article = ref({
  title: '',
  intro: '',
  content: '',
  cover_img: '',
  status: null,
  cate_id: null,
  tag_id: null,
})

const cateList = ref([])
const tagList = ref([])
const drawerVisible = ref(false)
const formRef = ref(null)
const fileList = ref([])
const rules = {
  cate_id: [{ required: true, message: '请选择分类', trigger: 'blur' }],
  tag_id: [{ required: true, message: '请选择标签', trigger: 'blur' }],
  intro: [
    { required: true, message: '请输入摘要', trigger: 'blur' },
    {
      min: 1,
      max: 200,
      message: '摘要长度不能超过200个字符',
      trigger: 'blur',
    },
  ],
}

// 图片超出限制
const uploadHandleExceed = () => {
  ElMessage.warning('只能上传一张图片')
}

//  显示抽屉
const showDrawer = (status) => {
  article.value.status = status
  drawerVisible.value = true
}

// 上传图片
const uploadCoverImage = async (file) => {
  let formData = new FormData()
  formData.append('file', file)
  const res = await uploadImage(formData)
  if (res.code !== 200) return
  return res
}

// 提交文章
const submitArticle = async (formEl) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (!valid) return
    if (fileList.value.length && fileList.value[0].raw) {
      const res = await uploadCoverImage(fileList.value[0].raw)
      if (res.code !== 200) return
      article.value.cover_img = res.data.file_path
    } else {
      article.value.cover_img = ''
    }
    if (article.value.id) {
      const res = await updateArticle(article.value)
      if (res.code !== 200) return
      ElMessage.success('更新成功')
    } else {
      const res = await createArticle(article.value)
      if (res.code !== 200) return
      ElMessage.success('发布成功')
    }
    router.go(-1)
  })
}

// 获取文章分类
const getArticleCate = async () => {
  const res = await getArticleCateList({})
  cateList.value = res.data.list
}

// 获取文章标签
const getArticleTag = async () => {
  const res = await getArticleTagList({})
  tagList.value = res.data.list
}

// 获取文章详情
const getArticle = async (id) => {
  const res = await getArticleDetail({ id })
  if (res.code !== 200) return
  article.value = res.data
  fileList.value = []
  if (article.value.cover_img !== '') {
    fileList.value.push({
      name: 'cover_img',
      url: config.file_base_url + article.value.cover_img,
    })
  }
}

onMounted(async () => {
  getArticleCate()
  getArticleTag()
  if (route.query.id) {
    await getArticle(parseInt(route.query.id))
  } else {
    article.value = {
      title: '',
      intro: '',
      content: '',
      cover_img: '',
      status: null,
      cate_id: null,
      tag_id: null,
    }
  }
})
</script>

<template>
  <div class="editor-page">
    <div class="header">
      <div class="title">
        <input type="text" v-model="article.title" placeholder="请输入文章标题" />
      </div>
      <div class="btn">
        <el-button @click="showDrawer(0)">存草稿</el-button>
        <el-button type="primary" @click="showDrawer(1)">发布</el-button>
        <el-button type="primary" @click="router.go(-1)">返回</el-button>
      </div>
    </div>
    <div class="main">
      <Vditor v-model="article.content" />
    </div>
  </div>
  <el-drawer
    v-model="drawerVisible"
    :title="article.status === 0 ? '保存草稿' : '发布文章'"
    size="400"
  >
    <el-form :model="article" ref="formRef" label-width="auto" :rules="rules">
      <el-form-item label="分类" prop="cate_id">
        <el-select v-model="article.cate_id" placeholder="请选择" clearable>
          <el-option
            v-for="item in cateList"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="标签" prop="tag_id">
        <el-select v-model="article.tag_id" placeholder="请选择" clearable filterable>
          <el-option
            v-for="item in tagList"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="封面" prop="cover_img">
        <el-upload
          v-model:file-list="fileList"
          list-type="picture-card"
          :auto-upload="false"
          :on-exceed="uploadHandleExceed"
          accept=".jpg,.jpeg,.png,.gif"
          :limit="1"
        >
          <el-icon><Plus /></el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="摘要" prop="intro">
        <el-input
          v-model="article.intro"
          :autosize="{ minRows: 4, maxRows: 8 }"
          type="textarea"
          maxlength="200"
          show-word-limit
          placeholder="请输入"
        >
        </el-input>
      </el-form-item>
      <div class="btn-group">
        <el-button @click="drawerVisible = false">取消</el-button>
        <el-button type="primary" @click="submitArticle(formRef)">提交</el-button>
      </div>
    </el-form>
  </el-drawer>
</template>

<style scoped lang="scss">
.editor-page {
  height: 100%;
  .header {
    height: 50px;
    display: flex;
    align-items: center;

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
.btn-group {
  display: flex;
  justify-content: flex-end;
  padding: 20px 0;
}
</style>
