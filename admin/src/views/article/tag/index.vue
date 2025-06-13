<script setup>
import { ref } from 'vue'
import { Plus, Search, Refresh } from '@element-plus/icons-vue'
import {
  getArticleTagList,
  getArticleTagDetail,
  createArticleTag,
  updateArticleTag,
  deleteArticleTag,
} from '@/api/article/tag'

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
  name: [
    { required: true, message: '请输入标签名称', trigger: 'blur' },
    { min: 1, max: 12, message: '最大长度为 12 个字符', trigger: 'blur' },
  ],
}

// 提交表单
const submitFormData = async (formEl) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (!valid) return
    if (formData.value.id) {
      const res = await updateArticleTag(formData.value)
      if (res.code !== 200) return
      ElMessage.success('更新成功')
    } else {
      const res = await createArticleTag(formData.value)
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
    const res = await getArticleTagDetail({ id: row.id })
    if (res.code !== 200) return
    formData.value = res.data
  } else if (!row.id && formData.value.id) {
    formData.value = {}
  }
  dialogVisible.value = true
}

// 删除表格行
const delRow = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该项数据, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(async () => {
      const res = await deleteArticleTag({ id: row.id })
      if (res.code !== 200) return
      ElMessage.success('删除成功')
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
  const res = await getArticleTagList(queryInfo.value)
  tableData.value = res.data
}
getTableData()
</script>

<template>
  <div class="tag-page">
    <el-card shadow="never" class="query">
      <el-form :model="queryInfo" ref="queryFormRef" @submit.prevent>
        <el-form-item label="标签名称" prop="name">
          <el-input
            v-model="queryInfo.name"
            clearable
            @keyup.enter="getTableData"
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item>
          <el-button :icon="Search" @click="getTableData">搜索</el-button>
          <el-button :icon="Refresh" @click="resetQueryForm(queryFormRef)">重置</el-button>
          <el-button :icon="Plus" type="primary" plain @click="showDialog">新建标签</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card shadow="never" class="table">
      <el-table :data="tableData.list">
        <el-table-column prop="name" label="标签名称" align="left" />
        <el-table-column prop="creator" label="创建人" align="left" />
        <el-table-column prop="create_time" label="创建时间" align="left" />
        <el-table-column prop="updater" label="更新人" align="left" />
        <el-table-column prop="update_time" label="更新时间" align="left" />
        <el-table-column label="操作" align="left" width="150" fixed="right">
          <template #default="{ row }">
            <div class="operate">
              <el-link :underline="false" type="primary" @click="showDialog(row)">编辑</el-link>
              <el-link :underline="false" type="danger" @click="delRow(row)">删除</el-link>
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
  <el-dialog :title="formData.id ? '编辑标签' : '新建标签'" v-model="dialogVisible" draggable>
    <el-form ref="formDataRef" :rules="rules" :model="formData" label-width="auto" @submit.prevent>
      <el-form-item label="标签名称" prop="name">
        <el-input v-model="formData.name" clearable placeholder="请输入"></el-input>
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
.tag-page {
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
