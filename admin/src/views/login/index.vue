<script setup>
import { useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'
import { login } from '@/api/user'
import CryptoJS from 'crypto-js'
import { useUserStore } from '@/store/user'

const userStore = useUserStore()
const SECRET_KEY = 'b4f91f1f5db4cc03aabdf18b015fda3b4d74658a7f6b12e0f4a110b9a5680fc4'
const router = useRouter()
const year = ref(new Date().getFullYear())
const loginFormData = ref({})
const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' },
  ],
}
const loginFormRef = ref(null)
const isRemember = ref(localStorage.getItem('isRemember') === 'true')

// 加密函数
function encryptPassword(password) {
  return CryptoJS.AES.encrypt(password, SECRET_KEY).toString()
}

// 解密函数
function decryptPassword(cipherText) {
  const bytes = CryptoJS.AES.decrypt(cipherText, SECRET_KEY)
  return bytes.toString(CryptoJS.enc.Utf8)
}

// 初始化时判断是否记住密码
onMounted(() => {
  if (localStorage.getItem('isRemember') === 'true' && localStorage.getItem('loginFormData')) {
    loginFormData.value = JSON.parse(localStorage.getItem('loginFormData')) || {}
    loginFormData.value.password = decryptPassword(loginFormData.value.password)
  }
})

// 登录
const onLogin = async (formEl) => {
  formEl.validate(async (valid) => {
    if (!valid) return
    const res = await login(loginFormData.value)
    if (res.code !== 200) return
    localStorage.setItem('token', res.data.token)
    isRemember.value
      ? localStorage.setItem('isRemember', true)
      : localStorage.setItem('isRemember', false)
    isRemember.value
      ? localStorage.setItem(
          'loginFormData',
          JSON.stringify({
            username: loginFormData.value.username,
            password: encryptPassword(loginFormData.value.password),
          }),
        )
      : localStorage.removeItem('loginFormData')
    userStore.setUserInfo(res.data.userInfo)
    window.history.state.back ? router.go(-1) : router.push('/home')
  })
}

const forgetPassword = () => {
  ElMessage.info('请联系管理员')
}
</script>

<template>
  <div class="login-page">
    <div class="header">
      <div class="logo">
        <router-link to="/">
          <img src="/public/logo.png" alt="logo" />
        </router-link>
        <h1>鼠觅</h1>
      </div>
    </div>
    <el-card class="login-box" shodaw="never" shadow="never">
      <div class="logo-box">
        <img src="/public/logo.png" alt="" />
      </div>
      <div class="title">登录 - 鼠觅</div>
      <el-form
        ref="loginFormRef"
        class="login-form"
        label-width="auto"
        :rules="rules"
        :model="loginFormData"
      >
        <el-form-item prop="username">
          <el-input
            @keyup.enter="onLogin(loginFormRef)"
            placeholder="请输入用户名"
            v-model="loginFormData.username"
            :prefix-icon="Avatar"
          >
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            @keyup.enter="onLogin(loginFormRef)"
            placeholder="请输入密码"
            v-model="loginFormData.password"
            :show-password="true"
            :prefix-icon="Lock"
          >
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-checkbox v-model="isRemember" label="记住密码" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onLogin(loginFormRef)" v-prevent-click>
            登录
          </el-button>
        </el-form-item>
      </el-form>
      <div class="tag">
        <el-link @click="forgetPassword"> 已有账号，忘记密码？ </el-link>
      </div>
    </el-card>
    <div class="footer">
      <el-link :underline="false"> © {{ year }} Micefind.com </el-link>
      <el-link href="https://beian.miit.gov.cn/" target="_blank" :underline="false">
        渝ICP备2022007989号
      </el-link>
    </div>
  </div>
</template>

<style scoped lang="scss">
.login-page {
  height: 100%;
  position: relative;

  .header {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 50px;
    display: flex;
    align-items: center;
    border-bottom: 1px solid var(--el-border-color);
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
  }

  .login-box {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 40%;
    padding: 35px 20px;
    border-radius: 5px;
    @media (max-width: 768px) {
      width: 100%;
      border: none;
      background-color: transparent;
    }

    .logo-box {
      width: 100%;
      display: flex;
      align-items: center;
      justify-content: center;

      img {
        width: 80px;
        height: 80px;
        border-radius: 5px;
      }
    }

    .title {
      width: 100%;
      padding: 30px 0;
      font-weight: 700;
      font-size: 18px;
      text-align: center;
    }

    .login-form {
      width: 100%;
      padding: 20px 0 20px 0;

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

  .footer {
    height: 50px;
    display: flex;
    align-items: center;
    justify-content: center;
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    .el-link {
      padding: 0 5px;
    }
  }
}
</style>
