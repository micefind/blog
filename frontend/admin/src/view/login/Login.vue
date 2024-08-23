<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import { useTagsViewStore } from "../../store/modules/tagsView.ts"
import { Lock, Avatar } from "@element-plus/icons-vue"
import type { FormInstance, FormRules } from "element-plus"
import logo from "../../assets/image/logo.png"
import request from "../../utils/request.ts"

// 使用 Vue 的 ref 和类型定义
const tagsViewStore = useTagsViewStore()
const router = useRouter()

// 明确定义 userInfo 的结构
interface UserInfo {
  username: string
  password: string
}

const userInfo = ref<UserInfo>({
  username: "micefind",
  password: "Ms19991115.",
})

const loginFormRef = ref<FormInstance>()

// 明确定义表单验证规则的类型
const rules = ref<FormRules>({
  username: [
    { required: true, message: "用户名不能为空", trigger: "blur" },
    { min: 2, max: 11, message: "用户名为2-11位字符", trigger: "blur" },
  ],
  password: [{ required: true, message: "密码不能为空", trigger: "blur" }],
})

// 用户登录
const login = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid: boolean) => {
    if (!valid) return
    const { data: res } = await request.post("user/login", userInfo.value)
    if (res.status !== 200) return
    ElMessage.success(`登录成功`)
    sessionStorage.setItem("token", res.data.token)

    router.push({
      path: "/overview",
    })
    tagsViewStore.setActivePath({ path: "/overview", name: "首页" })
  })
}
</script>

<template>
  <el-container class="container-box">
    <el-card class="login-box" :style="{ color: '#900' }" shodaw="never">
      <!--头像标志区域-->
      <div class="logo-box">
        <el-image :src="logo" fit="cover" />
      </div>
      <div class="title">登录 - 鼠觅奇物 - 博客后台</div>
      <!--登录表单区域-->
      <el-form
        ref="loginFormRef"
        class="login-form"
        label-width="auto"
        :rules="rules"
        :model="userInfo"
      >
        <!--用户名-->
        <el-form-item prop="username">
          <el-input
            placeholder="请输入用户名"
            v-model="userInfo.username"
            :prefix-icon="Avatar"
          >
          </el-input>
        </el-form-item>
        <!--密码-->
        <el-form-item prop="password">
          <el-input
            placeholder="请输入密码"
            v-model="userInfo.password"
            :show-password="true"
            :prefix-icon="Lock"
          >
          </el-input>
        </el-form-item>
        <!--登录按钮-->
        <el-form-item>
          <el-button type="primary" @click="login(loginFormRef)">
            登录
          </el-button>
        </el-form-item>
      </el-form>
      <div class="tag">
        <el-link> 忘记密码？ </el-link>
      </div>
      <el-divider> 其他作品 </el-divider>
      <div class="other-box">
        <el-link target="_blank" href="https://gitee.com/micefind/museum">
          线上博物馆
        </el-link>
      </div>
    </el-card>
  </el-container>
</template>

<style scoped lang="scss">
.container-box {
  width: 100%;
  height: 100vh;
  overflow: hidden;
  position: fixed;
  top: 0;
  left: 0;
}

.login-box {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 400px;
  height: 470px;
  padding: 20px;
  background-color: #fff;
  border-radius: 5px;
  box-shadow: 0 1px 5px rgba(55, 122, 255, 0.3);

  .logo-box {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;

    .el-image {
      width: 80px;
      height: 80px;
      border-radius: 5px;
    }
  }

  .title {
    width: 100%;
    padding: 15px 0;
    color: #666;
    font-weight: 700;
    font-size: 18px;
    text-align: center;
  }

  .login-form {
    width: 100%;
    padding: 15px 0 10px 0;

    .el-button {
      width: 100%;
    }
  }

  .tag {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0 0 10px 0;

    .el-link {
      padding: 0 10px;
    }
  }

  .other-box {
    display: flex;
    align-items: center;

    .el-link {
      padding: 5px 5px;
    }
  }
}
</style>
