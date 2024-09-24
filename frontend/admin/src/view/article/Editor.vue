
<script setup lang="ts">
import { ref } from 'vue';
import Vditor from '../../components/editor/Vditor.vue';  // 组件的路径
import avatar from "../../assets/image/defaultAvatar.png"
import { Close, FullScreen,Plus } from "@element-plus/icons-vue"
import type { FormInstance } from "element-plus"
import request from '../../utils/request';
import { useRouter,useRoute } from 'vue-router';


const router = useRouter()
const route = useRoute()
const avatarUrl = localStorage.getItem("avatar") || avatar
const dialogVisible = ref<boolean>(false)
const fullscreen = ref(false)
const articleFormRef = ref<FormInstance | null>(null)
// const markdownContent = ref<string>('');  // 初始的 Markdown 内容
interface ArticleInfo {
  id?: number,
  title: string,
  intro: string,
  keywords: string,
  cover_image: string,
  content: string,
  status: string,
}
const articleInfo = ref<ArticleInfo>({
  title: '',
  intro: '',
  keywords: '',
  cover_image: '',
  content: '',
  status: '',
})

const rules = {
  title: [
    { required: true, message: "请输入标题", trigger: "blur" },
    { min: 1, max: 20, message: "长度在 1 到 50 个字符", trigger: "blur" },
  ],
  intro: [
    { required: true, message: "请输入简介", trigger: "blur" },
    { min: 1, max: 100, message: "长度在 1 到 100 个字符", trigger: "blur" },
  ]
}
// 定义上传文件类型
interface FileItem {
  url?: string;
  raw?: File;
}

const fileList = ref<FileItem[]>([]);
  

const getArticleDetails = async () => {
  fileList.value = []
  if (!route.query.id) return
  const { data: res } = await request.post("/article/details",{ id: parseInt(route.query.id as string)})
  if (res.status !== 200) return
  articleInfo.value = res.data
  articleInfo.value.cover_image&&fileList.value.push({url: articleInfo.value.cover_image})
 }
 getArticleDetails()

const btnClick = (num:string) => {
  dialogVisible.value = true
  articleInfo.value.status=num
}

const avatarHandleCommand = async (command: string) => {
  switch (command) {
    case "goback":
    try {
    await ElMessageBox.confirm("现在返回内容不会被保存，是否继续", "提示", {
      confirmButtonText: "确认",
      cancelButtonText: "取消",
      type: "warning",
    })
    router.go(-1)
  } catch {
    ElMessage.info("已取消")
  } 
      break
  }
}

// 上传图片
const uploadImg = async (file: File): Promise<string | null> => {
  const formData = new FormData();
  formData.append("file", file);
  const { data: res } = await request.post("/upload/image", formData);
  return res.status === 200 ? res.data.url : null;
};


const submitData = (formEl: FormInstance | null) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (!valid) return
    if (fileList.value.length > 0 && fileList.value[0].raw) {
        const coverImage = await uploadImg(fileList.value[0].raw);
        articleInfo.value.cover_image = coverImage || "";
      }
      if (fileList.value.length ===0) articleInfo.value.cover_image = ''
    const url = articleInfo.value.id? "/article/edit" : "/article/add"
    const { data: res } = await request.post(url, articleInfo.value)
    if (res.status === 200) {
      dialogVisible.value = false
      ElMessage.success(articleInfo.value.id? "更新成功" : "添加成功")
      router.push("/article")
    }
  })
};
</script>

<template>
  <div>
    <div class="header">
      <div class="input">
        <input placeholder="请输入文章标题" v-model="articleInfo.title"></input>
      </div>
      <div class="btn">
        <el-button size="mini" @click="btnClick('1')" type="primary" plain>保存</el-button>
        <el-button size="mini" @click="btnClick('2')" type=primary>发布</el-button>
      </div>
      <el-dropdown
        style="margin-right: 2%;"
        @command="avatarHandleCommand"
      >
        <el-avatar :size="40" fit="cover" :src="avatarUrl"/>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="goback">
              返回上一页
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
    <div class="content">
      <Vditor v-model="articleInfo.content" />
    </div>
  </div>
  <!-- 对话框 -->
  <el-dialog
    :fullscreen="fullscreen"
    v-model="dialogVisible"
    width="50vw"
    :draggable="true"
    :show-close="false"
    style="padding: 0">
    <template #header="{ close, titleId, titleClass }" style="padding: 0">
      <div
        style="
          height: 54px;
          display: flex;
          align-items: center;
          padding: 0 15px;
          justify-content: space-between;
        ">
        <h4
          :id="titleId"
          :class="titleClass"
          style="font-weight: 500; font-size: 16px">
          {{ articleInfo.id? "编辑文章信息" : "添加文章" }}
        </h4>
        <div>
          <el-icon
            style="cursor: pointer; margin: 0 10px"
            @click="fullscreen = !fullscreen">
            <FullScreen />
          </el-icon>
          <el-icon @click="close" style="cursor: pointer">
            <Close />
          </el-icon>
        </div>
      </div>
    </template>
    <div
      style="
        border-top: 1px solid var(--el-border-color);
        border-bottom: 1px solid var(--el-border-color);
        padding: 20px;
      ">
      <el-form
        ref="articleFormRef"
        label-width="auto"
        :model="articleInfo"
        :rules="rules">
        <el-form-item prop="title" label="标题">
          <el-input placeholder="请输入标题" v-model="articleInfo.title">
          </el-input>
        </el-form-item>
        <el-form-item prop="intro" label="简介">
          <el-input placeholder="请输入简介" v-model="articleInfo.intro">
          </el-input>
        </el-form-item>
        <el-form-item prop="keywords" label="关键字">
          <el-input placeholder="请输入关键字" v-model="articleInfo.keywords">
          </el-input>
        </el-form-item>
        <el-form-item prop="logo" label="封面">
          <el-upload 
    v-model:file-list="fileList"
    limit="1"
    list-type="picture-card"
    :auto-upload="false"
  >
    <el-icon><Plus /></el-icon>
  </el-upload>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="articleInfo.status">
            <el-radio value="1">暂存</el-radio>
            <el-radio value="2">发布</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
    </div>
    <template #footer>
      <div style="padding: 10px 20px 20px">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitData(articleFormRef)">
          确定
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<style scoped lang="scss">
.header {
  height: 50px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid var(--el-border-color);

  .input {
    flex: 1;
    height: 100%;
    width: 50%;
    padding: 0 50px 0 2%;
    input {
      border: none;
      height: 49px;
      width: 100%;
      font-size: 24px;
      &:focus {
        outline: none;
        border: none;
      }
    }
  }
  .btn {
    padding: 0 20px 0 10px;
  }
}
.content {
  height: calc(100vh - 50px);
}
</style>
  