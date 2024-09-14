<script setup lang="ts">
import { ref } from "vue"
import { Search, Plus, Refresh, Close, FullScreen } from "@element-plus/icons-vue"
import type { FormInstance } from "element-plus"
import request from "../../utils/request"

// 定义类型
interface QueryInfo {
  username: string
  status: string
  pageNum: number
  pageSize: number
}

interface UserInfo {
  id: number
  username: string
  password?: string
  phone_number: string
  email: string
  real_name: string
  avatar: string
  status: string
  role: string
}

interface TableData {
  list: UserInfo[]
  total: number
}

const queryFormRef = ref<FormInstance | null>(null)
const userFormRef = ref<FormInstance | null>(null)
const dialogVisible = ref(false)
const fullscreen = ref(false)
const tableData = ref<TableData>({ list: [], total: 0 })

const queryInfo = ref<QueryInfo>({
  username: "",
  status: "",
  pageNum: 1,
  pageSize: 20,
})

const userInfo = ref<UserInfo>({
  id: -1,
  username: "",
  phone_number: "",
  email: "",
  real_name: "",
  avatar: "",
  role: "0",
  status: "0",
})

const rules = {
  username: [
    { required: true, message: "请输入用户名", trigger: "blur" },
    { min: 1, max: 20, message: "长度在 1 到 20 个字符", trigger: "blur" },
  ],
  password: [
  { required: true, message: "请输入密码", trigger: "blur" },
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
}

// 重置密码
const resetPassword = async (id: number) => {
  try {
    await ElMessageBox.confirm("是否将密码重置为123456", "提示", {
      confirmButtonText: "确认",
      cancelButtonText: "取消",
      type: "warning",
    })
    const { data: res } = await request.post("/user/password/reset", { id })
    if (res.status === 200) ElMessage.success("重置成功")
  } catch {
    ElMessage.info("已取消")
  }
}

// 提交数据
const submitData = async (formEl: FormInstance | null) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (!valid) return
    const url = userInfo.value.id === -1 ? "/user/add" : "/user/edit"
    const { data: res } = await request.post(url, userInfo.value)
    if (res.status === 200) {
      ElMessage.success(userInfo.value.id === -1 ? "添加成功" : "更新成功")
      await getTableData()
      resetUserInfo()
      dialogVisible.value = false
    }
  })
}

// 获取用户详情
const getDetails = async (id: number) => {
  const { data: res } = await request.post("/user/details", { id })
  if (res.status === 200) {
    userInfo.value = res.data
    dialogVisible.value = true
  }
}

// 重置用户信息
const resetUserInfo = () => {
  userInfo.value = {
    id: -1,
    username: "",
    phone_number: "",
    email: "",
    real_name: "",
    avatar: "",
    role: "0",
    status: "0",
  }
}

// 显示新增弹窗
const showAddDialog = () => {
if (userInfo.value.id !== -1) {
  resetUserInfo()
}
  dialogVisible.value = true
}

// 重置查询表单
const resetQueryForm = async (formEl: FormInstance | null) => {
  formEl?.resetFields()
  await getTableData()
}

// 分页处理
const handleSizeChange = async (val: number) => {
  queryInfo.value.pageSize = val
  await getTableData()
}

const handleCurrentChange = async (val: number) => {
  queryInfo.value.pageNum = val
  await getTableData()
}

// 获取列表数据
const getTableData = async () => {
  const { data: res } = await request.post("/user/list", queryInfo.value)
  if (res.status === 200) {
    tableData.value = res.data
  }
}

// 初始化时获取数据
getTableData()
</script>

<template>
  <el-card shadow="never" class="header-box">
    <el-form ref="queryFormRef" :model="queryInfo">
      <el-form-item label="用户名" prop="username">
        <el-input
        @keyup.enter="getTableData"
          placeholder="请输入用户名"
          v-model="queryInfo.username"
          clearable></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select placeholder="请选择状态" v-model="queryInfo.status">
          <el-option label="正常" value="0" />
          <el-option label="限制" value="1" />
          <el-option label="注销" value="2" />
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
        <el-button type="primary" plain @click="showAddDialog">
          <el-icon>
            <Plus />
          </el-icon>
          <span>新增</span>
        </el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <el-card shadow="never" class="content-box">
    <el-table :data="tableData.list">
      <el-table-column prop="username" label="用户名" align="center" />
      <el-table-column prop="real_name" label="姓名" align="center" />
      <el-table-column prop="phone_number" label="手机号" align="center" />
      <el-table-column prop="email" label="邮箱" align="center" />
      <el-table-column label="角色" align="center">
        <template #default="scope">
          <el-tag :type="scope.row.role === '0' ? 'primary' : 'info'">{{
            scope.row.role === "0" ? "管理员" : "游客"
          }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center">
        <template #default="scope">
          <el-tag
            :type="
              scope.row.status === '0'
                ? 'success'
                : scope.row.status === '1'
                ? 'warning'
                : 'danger'
            "
            >{{
              scope.row.status === "0"
                ? "正常"
                : scope.row.status === "1"
                ? "限制"
                : "注销"
            }}</el-tag
          >
        </template>
      </el-table-column>
      <el-table-column
        prop="register_time"
        label="注册时间"
        align="center"
        width="180" />
      <el-table-column label="操作" align="center" fixed="right" width="150">
        <template #default="scope">
          <el-link
            style="margin-right: 10px"
            type="primary"
            :underline="false"
            @click="getDetails(scope.row.id)"
            >编辑</el-link
          >
          <el-link
            type="danger"
            :underline="false"
            @click="resetPassword(scope.row.id)"
            >重置密码</el-link>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination">
      <el-pagination
        v-model:current-page="queryInfo.pageNum"
        v-model:page-size="queryInfo.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :background="true"
        layout="total,sizes, prev, pager, next, jumper"
        :total="tableData.total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange">
      </el-pagination>
    </div>
  </el-card>
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
          {{ userInfo.id === -1 ? "新增用户" : "编辑用户" }}
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
        ref="userFormRef"
        label-width="auto"
        :model="userInfo"
        :rules="rules">
        <el-form-item prop="username" label="用户名">
          <el-input placeholder="请输入用户名" v-model="userInfo.username">
          </el-input>
        </el-form-item>
        <el-form-item prop="password" label="密码" v-if="userInfo.id === -1">
          <el-input placeholder="请输入密码" v-model="userInfo.password">
          </el-input>
        </el-form-item>
        <el-form-item prop="real_name" label="姓名">
          <el-input placeholder="请输入姓名" v-model="userInfo.real_name">
          </el-input>
        </el-form-item>
        <el-form-item prop="phone_number" label="手机号">
          <el-input placeholder="请输入手机号" v-model="userInfo.phone_number">
          </el-input>
        </el-form-item>
        <el-form-item prop="email" label="邮箱">
          <el-input placeholder="请输入邮箱" v-model="userInfo.email">
          </el-input>
        </el-form-item>
        <el-form-item label="角色" prop="role">
          <el-radio-group v-model="userInfo.role">
            <el-radio value="0">管理员</el-radio>
            <el-radio value="1">游客</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="userInfo.status">
            <el-radio value="0">正常</el-radio>
            <el-radio value="1">限制</el-radio>
            <el-radio value="2">注销</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
    </div>
    <template #footer>
      <div style="padding: 10px 20px 20px">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitData(userFormRef)">
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

.pagination {
  padding: 20px 0 10px 0;
  float: right;
}
</style>
