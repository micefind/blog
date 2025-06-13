<script setup>
import { User, Iphone, EditPen, Help, CollectionTag, Sunrise } from '@element-plus/icons-vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/store/user'
import { getUserDetail, updateUser, changeUserPassword } from '@/api/user'

const userStore = useUserStore()
const router = useRouter()
const dialogVisible = ref(false)
const activeName = ref('first')
const userInfo = ref({})
const userInfoArr = ref([])
const passwordFormData = ref({
  password: '',
  new_password: '',
  confirm_password: '',
})
const passwordFormRef = ref()
const userFormRef = ref()

const userInfoRules = ref({
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 1, max: 20, message: '密码为1-20位字符', trigger: 'blur' },
  ],
})

const passwordFormRules = ref({
  password: [
    { required: true, message: '密码不能为空', trigger: 'blur' },
    { min: 6, max: 18, message: '密码为6-30位字符', trigger: 'blur' },
  ],
  new_password: [
    { required: true, message: '密码不能为空', trigger: 'blur' },
    {
      validator: (_, value, callback) => {
        if (value) {
          const regex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&.])[A-Za-z\d@$!%*?&.]{6,30}$/
          if (!regex.test(value)) {
            callback(new Error('密码长度在 6-30 位之间，且必须包含大小写字母、数字和特殊字符'))
          } else {
            callback()
          }
        }
      },
      trigger: 'blur',
    },
  ],
  confirm_password: [{ required: true, message: '密码不能为空', trigger: 'blur' }],
})

// 提交用户信息
const submitUserInfo = (formEl) => {
  formEl.validate(async (valid) => {
    if (!valid) return
    const res = await updateUser(userInfo.value)
    if (res.code !== 200) return
    ElMessage.success(`更新成功`)
    userStore.setUserInfo(userInfo.value)
    getUserInfo()
  })
}

// 重置用户信息表单
const resetUserInfoForm = () => {
  getUserInfo()
}

// 修改用户密码
const changePassword = (formEl) => {
  formEl.validate(async (valid) => {
    if (!valid) return
    if (passwordFormData.value.new_password !== passwordFormData.value.confirm_password)
      return ElMessage('密码输入不一致')
    if (passwordFormData.value.new_password === passwordFormData.value.password)
      return ElMessage('新密码不能与旧密码相同')
    const res = await changeUserPassword(passwordFormData.value)
    if (res.code !== 200) return
    ElMessage.success(`修改密码成功，请重新登录`)
    sessionStorage.removeItem('token')
    router.push('/login')
  })
}

// 重置密码表单
const resetPasswordForm = (formEl) => {
  formEl.resetFields()
}

// 设置用户信息组
const setUserInfoArr = (userInfo) => {
  userInfoArr.value = [
    { name: '用户名', value: userInfo['username'], icon: User },
    { name: '姓名', value: userInfo['name'], icon: EditPen },
    { name: '手机号', value: userInfo['mobile'], icon: Iphone },
    { name: '邮箱', value: userInfo['email'], icon: Help },
    {
      name: '角色',
      value: userInfo['role'] === '0' ? '管理员' : '普通用户',
      icon: CollectionTag,
    },
    {
      name: '状态',
      value: userInfo['is_deleted'] === 0 ? '正常' : '注销',
      icon: Sunrise,
    },
  ]
}

// 获取用户信息
const getUserInfo = async () => {
  const res = await getUserDetail({ id: userStore.userInfo.id })
  if (res.code !== 200) return
  userInfo.value = res.data
  setUserInfoArr(userInfo.value)
}
getUserInfo()
</script>

<template>
  <div class="profile-page">
    <el-card class="pro-file" shadow="never">
      <template #header>
        <div class="card-header">
          <span>个人信息</span>
        </div>
      </template>
      <div class="card-body">
        <ul>
          <li @click="dialogVisible = true" class="avatar-box">
            <div class="img-box">
              <el-avatar :size="100" :src="userInfo.avatar" fit="cover" v-if="userInfo.avatar" />
              <el-avatar :size="100" fit="cover" v-else>
                {{ userInfo.username ? userInfo.username[0].toUpperCase() : 'A' }}
              </el-avatar>
            </div>
          </li>
          <li v-for="item in userInfoArr" :key="item.name">
            <span>
              <el-icon>
                <component :is="item.icon" />
              </el-icon>
              <span>{{ item.name }}</span>
            </span>
            <span>{{ item.value }}</span>
          </li>
        </ul>
      </div>
    </el-card>
    <el-card class="base-file" shadow="never">
      <template #header>
        <div class="card-header">
          <span>基本信息</span>
        </div>
      </template>
      <div class="card-body">
        <el-tabs v-model="activeName" class="demo-tabs">
          <el-tab-pane label="基本资料" name="first">
            <el-form ref="userFormRef" label-width="auto" :model="userInfo" :rules="userInfoRules">
              <el-form-item prop="username" label="用户名">
                <el-input v-model="userInfo.username"> </el-input>
              </el-form-item>
              <el-form-item prop="name" label="姓名">
                <el-input v-model="userInfo.name"> </el-input>
              </el-form-item>
              <el-form-item prop="mobile" label="手机号">
                <el-input v-model="userInfo.mobile"> </el-input>
              </el-form-item>
              <el-form-item prop="email" label="邮箱">
                <el-input v-model="userInfo.email"> </el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="submitUserInfo(userFormRef)"> 保存 </el-button>
                <el-button type="primary" @click="resetUserInfoForm()"> 重置 </el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="修改密码" name="second">
            <el-form
              ref="passwordFormRef"
              label-width="auto"
              :model="passwordFormData"
              :rules="passwordFormRules"
            >
              <el-form-item prop="password" label="旧密码">
                <el-input v-model="passwordFormData.password" :show-password="true"> </el-input>
              </el-form-item>
              <el-form-item prop="new_password" label="新密码">
                <el-input v-model="passwordFormData.new_password" :show-password="true"> </el-input>
              </el-form-item>
              <el-form-item prop="confirm_password" label="确认密码">
                <el-input v-model="passwordFormData.confirm_password" :show-password="true">
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="changePassword(passwordFormRef)">
                  保存
                </el-button>
                <el-button type="primary" @click="resetPasswordForm(passwordFormRef)">
                  重置
                </el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-card>
  </div>
</template>

<style scoped lang="scss">
.profile-page {
  padding: var(--page-padding);
  display: flex;
  justify-content: space-between;
}

.pro-file {
  width: 35%;
  padding: 15px 20px 20px;

  &:hover {
    box-shadow: var(--el-box-shadow-light);
  }
}

.card-header {
  display: flex;
  justify-content: center;

  span {
    font-size: 16px;
    text-align: center;
  }
}

.card-body {
  padding: 15px;

  ul li {
    padding: 11px 0;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1px solid var(--el-border-color);
    font-size: 13px;

    span {
      font-size: 13px;
    }

    span span {
      padding: 0 0 0 10px;
    }
  }

  ul li:nth-child(1) {
    justify-content: center;
    padding-top: 0;
  }
}

.avatar-box {
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;

  .img-box {
    position: relative;
    width: 100px;
    height: 100px;

    .el-avatar {
      border: 3px solid var(--el-border-color);
    }
  }
}

.base-file {
  width: calc(65% - 10px);
  padding: 15px 20px 20px;

  &:hover {
    box-shadow: var(--el-box-shadow-light);
  }
}
</style>
