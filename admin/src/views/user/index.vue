<script setup>
import { ref } from 'vue'
import { Plus, Search, Refresh } from '@element-plus/icons-vue'
import {
  getUserList,
  getUserDetail,
  createUser,
  updateUser,
  logoutUser,
  recoverUser,
  resetUserPassword,
} from '@/api/user'

const queryInfo = ref({
  page_num: 1,
  page_size: 20,
})
const tableData = ref({})
const queryFormRef = ref()
const formData = ref({})
const formDataRef = ref()
const dialogVisible = ref(false)
const rules = {
  username: [
    { required: true, message: '请输入用户名称', trigger: 'blur' },
    { min: 1, max: 12, message: '最大长度为 12 个字符', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    {
      validator: (_, value, callback) => {
        if (value) {
          const regex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[_.!@#$%^&*()]).{6,30}$/
          if (!regex.test(value)) {
            callback(
              new Error('密码长度在 6-30 位之间，且必须包含大小写字母、数字和特殊字符_.!@#$%^&*()'),
            )
          } else {
            callback()
          }
        } else if (!formData.value.id) {
          callback(new Error('请输入密码'))
        } else {
          callback()
        }
      },
      trigger: 'blur',
    },
  ],
}

// 提交表单
const submitFormData = async (formEl) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (!valid) return
    if (formData.value.id) {
      const res = await updateUser(formData.value)
      if (res.code !== 200) return
      ElMessage.success('更新成功')
    } else {
      const res = await createUser(formData.value)
      if (res.code !== 200) return
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    formEl.resetFields()

    await getTableData()
  })
}

// 显示对话框
const showDialog = async (row) => {
  if (row.id) {
    const res = await getUserDetail({ id: row.id })
    if (res.code !== 200) return
    formData.value = res.data
  } else if (!row.id && formData.value.id) {
    formData.value = {}
  }
  dialogVisible.value = true
}

// 恢复表格行
const recoverRow = async (row) => {
  const res = await recoverUser({ id: row.id })
  if (res.code !== 200) return
  ElMessage.success('恢复成功')
  await getTableData()
}

// 删除表格行
const delRow = async (row) => {
  ElMessageBox.confirm('用户注销后将无法登录系统, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(async () => {
      const res = await logoutUser({ id: row.id })
      if (res.code !== 200) return
      ElMessage.success('注销成功')
      await getTableData()
    })
    .catch(() => {})
}

// 重置密码
const resetPassword = async (row) => {
  ElMessageBox.prompt('请输入新密码', '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    inputPattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[_.!@#$%^&*()]).{6,30}$/,
    inputErrorMessage: '密码长度在 6-30 位之间，且必须包含大小写字母、数字和特殊字符_.!@#$%^&*()',
  })
    .then(async ({ value }) => {
      const res = await resetUserPassword({ id: row.id, password: value })
      if (res.code !== 200) return
      ElMessage.success(res.msg)
      await getTableData()
    })
    .catch(() => {})
}

// 重置表单
const resetQueryForm = (formEl) => {
  if (!formEl) return
  formEl.resetFields()
  getTableData()
}

// 分页：每页条数改变
const handleSizeChange = async (val) => {
  queryInfo.value.page_size = val
  await getTableData()
}

// 分页：页码改变
const handleCurrentChange = async (val) => {
  queryInfo.value.page_num = val
  await getTableData()
}

// 获取表格数据
const getTableData = async () => {
  const res = await getUserList(queryInfo.value)
  tableData.value = res.data
}
getTableData()
</script>

<template>
  <div class="cate-page">
    <el-card shadow="never" class="query">
      <el-form :model="queryInfo" ref="queryFormRef">
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="queryInfo.username"
            clearable
            @keyup.enter="getTableData"
            placeholder="请输入"
          ></el-input>
        </el-form-item>
        <el-form-item label="手机号" prop="mobile">
          <el-input
            v-model="queryInfo.mobile"
            clearable
            @keyup.enter="getTableData"
            placeholder="请输入"
          ></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input
            v-model="queryInfo.email"
            clearable
            @keyup.enter="getTableData"
            placeholder="请输入"
          ></el-input>
        </el-form-item>
        <el-form-item label="状态" prop="is_deleted">
          <el-select
            v-model="queryInfo.is_deleted"
            placeholder="请选择"
            @change="getTableData"
            clearable
          >
            <el-option label="正常" :value="0"></el-option>
            <el-option label="已注销" :value="1"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button :icon="Search" @click="getTableData">搜索</el-button>
          <el-button :icon="Refresh" @click="resetQueryForm(queryFormRef)">重置</el-button>
          <el-button :icon="Plus" type="primary" plain @click="showDialog"
            >新建管理员用户</el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
    <el-card shadow="never" class="table">
      <el-table :data="tableData.list">
        <el-table-column prop="username" label="用户名" align="left" />
        <el-table-column prop="name" label="姓名" align="left" />
        <el-table-column prop="mobile" label="手机号" align="left" />
        <el-table-column prop="email" label="邮箱" align="left" />
        <el-table-column label="角色" align="left">
          <template #default="{ row }">
            <el-tag v-if="row.role === 0" type="primary">管理员</el-tag>
            <el-tag v-if="row.role === 1" type="info">普通用户</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="状态" align="left">
          <template #default="{ row }">
            <el-tag v-if="row.is_deleted === 0" type="success">正常</el-tag>
            <el-tag v-if="row.is_deleted === 1" type="danger">已注销</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="left" width="150" fixed="right">
          <template #default="{ row }">
            <div class="operate">
              <el-link :underline="false" type="primary" @click="showDialog(row)">编辑</el-link>
              <el-link
                :underline="false"
                type="danger"
                @click="delRow(row)"
                v-if="row.is_deleted === 0"
                >注销</el-link
              >
              <el-link
                :underline="false"
                type="success"
                @click="recoverRow(row)"
                v-if="row.is_deleted === 1"
                >恢复</el-link
              >
              <el-link :underline="false" type="warning" @click="resetPassword(row)"
                >重置密码</el-link
              >
            </div>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
          v-model:current-page="queryInfo.page_num"
          v-model:page-size="queryInfo.page_size"
          :page-sizes="[10, 20, 50, 100]"
          :background="true"
          layout="total,sizes, prev, pager, next, jumper"
          :total="tableData.total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
  <el-dialog :title="formData.id ? '编辑用户' : '新建用户'" v-model="dialogVisible" draggable>
    <el-form ref="formDataRef" :rules="rules" :model="formData" label-width="auto">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="formData.username" clearable placeholder="请输入"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password" v-if="!formData.id">
        <el-input v-model="formData.password" clearable placeholder="请输入"></el-input>
      </el-form-item>
      <el-form-item label="姓名" prop="name">
        <el-input v-model="formData.name" clearable placeholder="请输入"></el-input>
      </el-form-item>
      <el-form-item label="手机号" prop="mobile">
        <el-input v-model="formData.mobile" clearable placeholder="请输入"></el-input>
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="formData.email" clearable placeholder="请输入"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <div>
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button @click="submitFormData(formDataRef)" type="primary" v-prevent-click
          >确 定</el-button
        >
      </div>
    </template>
  </el-dialog>
</template>

<style scoped lang="scss">
.cate-page {
  padding: var(--page-padding);
  .query {
    margin-bottom: 15px;
    .el-form {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      gap: 20px;
      .el-form-item {
        margin: 0;
        padding: 0;
        .el-input {
          width: 200px;
        }
        .el-select {
          width: 200px;
        }
      }
    }
  }
  .table {
    .el-table {
      margin-bottom: 20px;

      .operate {
        display: flex;
        align-items: center;
        gap: 5px;
      }
    }
    .pagination {
      display: flex;
      justify-content: flex-end;
    }
  }
}
</style>
