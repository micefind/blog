<script setup lang="ts">
import { ref } from 'vue'
import { Search, Plus, Refresh, Close, FullScreen } from '@element-plus/icons-vue'
import type { FormInstance } from 'element-plus'

// 定义类型
interface QueryInfo {
    username: string
}

interface UserInfo {
    username: string
    password: string
    real_name: string
    mobile: string
    email: string
    avatar: string
    role: string
}

const queryFormRef = ref<FormInstance>()
const userFormRef = ref<FormInstance>()
const dialogVisible = ref<boolean>(false)
const fullscreen = ref<boolean>(false)
const tableData = ref<Array<{
    username: string
    mobile: string
    email: string
    role: string
}>>([])
const queryInfo = ref<QueryInfo>({
    username: ''
})
const userInfo = ref<UserInfo>({
    username: '',
    password: '',
    real_name: '',
    mobile: '',
    email: '',
    avatar: '',
    role: '1'
})

// 重置查询表单
const resetQueryForm = async (formEl: FormInstance | undefined) => {
    formEl ? formEl.resetFields() : ''
}

// 获取列表数据
const getTableData = async () => {
    tableData.value = []
    for (let i = 0; i < 9; i++) {
        tableData.value.push({
            username: 'user' + i,
            mobile: '18288888888',
            email: 'micefind@qq.com',
            role: (i % 2).toString()
        })
    }
}

getTableData()

</script>

<template>
    <el-card shadow="never" class="header-box">
        <el-form ref="queryFormRef" :model="queryInfo">
            <el-form-item label="用户名" prop="username">
                <el-input placeholder="请输入用户名" v-model="queryInfo.username" clearable></el-input>
            </el-form-item>
            <el-form-item label="角色">
                <el-select placeholder="请选择角色">
                    <el-option label="管理员" value="0" />
                    <el-option label="普通用户" value="1" />
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button @click="getTableData">
                    <el-icon>
                        <Search />
                    </el-icon>
                    <span>搜索</span>
                </el-button>
                <el-button @click="resetQueryForm(queryFormRef)">
                    <el-icon>
                        <Refresh />
                    </el-icon>
                    <span>重置</span>
                </el-button>
                <el-button type="primary" plain @click="dialogVisible = true">
                    <el-icon>
                        <Plus />
                    </el-icon>
                    <span>新增</span>
                </el-button>
            </el-form-item>
        </el-form>
    </el-card>
    <el-card shadow="never" class="content-box">
        <el-table :data="tableData">
            <el-table-column prop="username" label="用户名" align="center" />
            <el-table-column prop="mobile" label="手机号" align="center" />
            <el-table-column prop="email" label="邮箱" align="center" />
            <el-table-column label="角色" align="center">
                <template #default="scope">
                    <el-tag type="primary" v-if="scope.row.role === '1'">普通用户</el-tag>
                    <el-tag type="success" v-if="scope.row.role === '0'">管理员</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center" fixed="right">
                <template #default="scope">
                    <el-link type="danger" :underline="false" v-if="scope">重置密码</el-link>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
    <!-- 对话框 -->
    <el-dialog :fullscreen="fullscreen" v-model="dialogVisible" width="50vw" :draggable="true" :show-close="false"
        style="padding: 0;">
        <template #header="{ close, titleId, titleClass }" style="padding: 0;">
            <div style="height: 54px;display: flex;align-items: center;padding: 0 15px;justify-content: space-between;">
                <h4 :id="titleId" :class="titleClass" style="font-weight: 500;font-size: 16px;">新增用户</h4>
                <div>
                    <el-icon style="cursor: pointer;margin: 0 10px;" @click="fullscreen = !fullscreen">
                        <FullScreen />
                    </el-icon>
                    <el-icon @click="close" style="cursor: pointer;">
                        <Close />
                    </el-icon>
                </div>
            </div>
        </template>
        <div
            style="border-top: 1px solid var(--el-border-color);border-bottom: 1px solid var(--el-border-color);padding: 20px">
            <el-form ref="userFormRef" label-width="auto" :model="userInfo">
                <el-form-item prop="username" label="用户名">
                    <el-input placeholder="请输入用户名" v-model="userInfo.username">
                    </el-input>
                </el-form-item>
                <el-form-item prop="password" label="密码">
                    <el-input placeholder="请输入密码" v-model="userInfo.password">
                    </el-input>
                </el-form-item>
                <el-form-item prop="real_name" label="姓名">
                    <el-input placeholder="请输入姓名" v-model="userInfo.real_name">
                    </el-input>
                </el-form-item>
                <el-form-item prop="mobile" label="手机号">
                    <el-input placeholder="请输入手机号" v-model="userInfo.mobile">
                    </el-input>
                </el-form-item>
                <el-form-item prop="email" label="邮箱">
                    <el-input placeholder="请输入邮箱" v-model="userInfo.email">
                    </el-input>
                </el-form-item>
                <el-form-item label="角色">
                    <el-radio-group v-model="userInfo.role">
                        <el-radio value="0">管理员</el-radio>
                        <el-radio value="1">普通用户</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
        </div>
        <template #footer>
            <div style="padding: 10px 20px 20px;">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="dialogVisible = false">
                    确定
                </el-button>
            </div>
        </template>
    </el-dialog>
</template>

<style scoped lang="scss">
.header-box {
    padding: 20px;
    padding-bottom: 5px;
    margin-bottom: 15px;

    .el-form {
        display: flex;
        flex-wrap: wrap;

        .el-form-item {
            margin-right: 40px;

            .el-select {
                width: 240px;
            }

            .el-input {
                width: 240px;

                .el-button {
                    span {
                        padding-left: 5px;
                    }
                }
            }
        }
    }
}

.content-box {
    padding: 20px;
}
</style>
