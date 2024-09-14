<script setup lang="ts">
import { User,Iphone,EditPen,Help,CollectionTag,Sunrise } from '@element-plus/icons-vue'
import { ref } from 'vue'
import request from '../../utils/request'
import defaultAvatar from '../../assets/image/defaultAvatar.png'
import type { FormInstance } from 'element-plus'
import { useRouter } from 'vue-router'

const router = useRouter()
const activeName = ref('first')
const userInfo = ref<any>({})
const userInfoArr = ref<any>([])
const passwordFormData = ref<any>({
    password: '',
    newPassword: '',
    newPasswordTwo: '',
})
const passwordFormRef = ref<FormInstance>()
    const userFormRef = ref<FormInstance>()

    const userInfoRules = ref<any>({
    username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 1, max: 20, message: '密码为1-20位字符', trigger: 'blur' }
    ]
    })


const passwordFormRules = ref<any>({
    password: [
        { required: true, message: '密码不能为空', trigger: 'blur' },
        { min: 6, max: 18, message: '密码为6-30位字符', trigger: 'blur' }
    ],
    newPassword: [
        { required: true, message: '密码不能为空', trigger: 'blur' },
        {
      validator: (_: any, value: string, callback: Function) => {
        // 仅在密码有值时校验密码规则
        if (value) {
          const regex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&.])[A-Za-z\d@$!%*?&.]{6,30}$/
          if (!regex.test(value)) {
            callback(new Error("密码长度在 6-30 位之间，且必须包含大小写字母、数字和特殊字符"));
          } else {
            callback();
          }
        } else if (userInfo.value.id === -1) {
          // 新增用户必须输入密码
          callback(new Error("请输入密码"));
        } else {
          // 编辑用户时允许密码为空
          callback();
        }
      },
      trigger: "blur"
    }
    ],
    newPasswordTwo: [
        { required: true, message: '密码不能为空', trigger: 'blur' },
        {
      validator: (_: any, value: string, callback: Function) => {
        // 仅在密码有值时校验密码规则
        if (value) {
          const regex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&.])[A-Za-z\d@$!%*?&.]{6,30}$/
          if (!regex.test(value)) {
            callback(new Error("密码长度在 6-30 位之间，且必须包含大小写字母、数字和特殊字符"));
          } else {
            callback();
          }
        } else if (userInfo.value.id === -1) {
          // 新增用户必须输入密码
          callback(new Error("请输入密码"));
        } else {
          // 编辑用户时允许密码为空
          callback();
        }
      },
      trigger: "blur"
    }
    ]
})

const submitUserInfo = (formEl: any) => {
    formEl.validate(async (valid: any) => {
        if (!valid) return
        const { data: res } = await request.post('/user/edit', userInfo.value)
        if (res.status !== 200) return
        ElMessage.success(`修改成功`)
        getUserInfo()
    })
}
const resetUserInfoForm= () => {
    getUserInfo()
}

// 修改用户密码
const changePassword = (formEl: any) => {
    formEl.validate(async (valid: any) => {
        if (!valid) return
        if (passwordFormData.value.newPassword !== passwordFormData.value.newPasswordTwo) return ElMessage('密码输入不一致')
        if (passwordFormData.value.newPassword === passwordFormData.value.password) return ElMessage('新密码不能与旧密码相同')
        const { data: res } = await request.post('/user/password/change', passwordFormData.value)
        if (res.status !== 200) return
        ElMessage.success(`修改密码成功，请重新登录`)
        router.push({
            path: '/login'
        })
        sessionStorage.clear()
    })
}

// 重置密码表单
const resetPasswordForm = (formEl: any) => {
    formEl.resetFields()
}

// 从token中获取用户id
const getUserID = ()=> {
    const token = localStorage.getItem('token')
    let strings = token?.split(".")||'';//通过split()方法将token转为字符串数组
    //取strings[1]数组中的第二个字符进行解析
    var userinfo = JSON.parse(decodeURIComponent(escape(window.atob(strings[1].replace(/-/g, "+").replace(/_/g, "/")))));
    return userinfo.userID
}

// 设置用户信息组
const setUserInfoArr = (userInfo:any) => {
    userInfoArr.value = [
        {name:'用户名', value: userInfo['username'], icon: User},
        {name:'姓名', value: userInfo['real_name'], icon: EditPen},
        {name:'手机号', value: userInfo['phone_number'], icon: Iphone},
        {name:'邮箱', value: userInfo['email'], icon: Help},
        {name:'角色', value: userInfo['role']==='0'?'管理员':'游客', icon: CollectionTag},
        {name:'状态', value: userInfo['status']==='0'?'正常':userInfo['status']==='1'?'限制':'注销', icon: Sunrise}
    ]
}

// 获取用户信息
const getUserInfo = async () => {
    const userID = getUserID()
    const {data:res} = await request.post('/user/details', {id:userID})
    userInfo.value = res.data
    setUserInfoArr(userInfo.value)
}
getUserInfo()

</script>

<template>
    <div class="container-box">
        <el-card class="pro-file" shadow="never">
            <template #header>
                <div class="card-header">
                    <span>个人信息</span>
                </div>
            </template>
            <div class="card-body">
                <ul>
                    <li>
                        <el-avatar :size="100" :src="defaultAvatar" />
                    </li>
                    <li v-for="item in userInfoArr">
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
                        <el-form ref="userFormRef" label-width="auto" :model="userInfo"
                            :rules="userInfoRules">
                            <el-form-item prop="username" label="用户名">
                                <el-input v-model="userInfo.username">
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="real_name" label="姓名">
                                <el-input v-model="userInfo.real_name">
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="phone_number" label="手机号">
                                <el-input v-model="userInfo.phone_number" >
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="email" label="邮箱">
                                <el-input v-model="userInfo.email" >
                                </el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" @click="submitUserInfo(userFormRef)">
                                    保存
                                </el-button>
                                <el-button type="primary" @click="resetUserInfoForm()">
                                    重置
                                </el-button>
                            </el-form-item>
                        </el-form>
                    </el-tab-pane>
                    <el-tab-pane label="修改密码" name="second">
                        <el-form ref="passwordFormRef" label-width="auto" :model="passwordFormData"
                            :rules="passwordFormRules">
                            <el-form-item prop="password" label="旧密码">
                                <el-input v-model="passwordFormData.password" :show-password="true">
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="newPassword" label="新密码">
                                <el-input v-model="passwordFormData.newPassword" :show-password="true">
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="newPasswordTwo" label="确认密码">
                                <el-input v-model="passwordFormData.newPasswordTwo" :show-password="true">
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
.container-box {
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

.base-file {
    width: calc(65% - 10px);
    padding: 15px 20px 20px;

    &:hover {
        box-shadow: var(--el-box-shadow-light);
    }
}
</style>