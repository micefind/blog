<script setup>
import { onMounted, ref } from 'vue'
import { Plus, Search, Refresh } from '@element-plus/icons-vue'
import { getArticleList, deleteArticle, recoverArticle } from '@/api/article'
import { getArticleCateList } from '@/api/article/cate'
import { getArticleTagList } from '@/api/article/tag'
import config from '@/config/config'

const queryInfo = ref({
  page_num: 1,
  page_size: 20,
})
const tableData = ref({})
const queryFormRef = ref()
const cateList = ref([])
const tagList = ref([])

// 获取分类列表
const getCateList = async () => {
  const res = await getArticleCateList({})
  cateList.value = res.data.list
}

// 获取标签列表
const getTagList = async () => {
  const res = await getArticleTagList({})
  tagList.value = res.data.list
}

// 恢复表格行
const recoverRow = async (row) => {
  const res = await recoverArticle({ id: row.id })
  if (res.code !== 200) return
  ElMessage.success('恢复成功')
  await getTableData()
}

// 删除表格行
const delRow = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该项数据, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(async () => {
      const res = await deleteArticle({ id: row.id })
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
  const res = await getArticleList(queryInfo.value)
  tableData.value = res.data
}
onMounted(async () => {
  await getCateList()
  await getTagList()
  await getTableData()
})
</script>

<template>
  <div class="article-page">
    <el-card shadow="never" class="query">
      <el-form :model="queryInfo" ref="queryFormRef">
        <el-form-item label="关键词" prop="keyword">
          <el-input
            v-model="queryInfo.keyword"
            clearable
            @keyup.enter="getTableData"
            placeholder="请输入"
          ></el-input>
        </el-form-item>
        <el-form-item label="分类" prop="cate_id">
          <el-select
            v-model="queryInfo.cate_id"
            placeholder="请选择"
            @change="getTableData"
            clearable
          >
            <el-option
              :label="item.name"
              :value="item.id"
              v-for="item in cateList"
              :key="item.id"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="标签" prop="tag_id">
          <el-select
            filterable
            v-model="queryInfo.tag_id"
            placeholder="请选择"
            @change="getTableData"
            clearable
          >
            <el-option
              :label="item.name"
              :value="item.id"
              v-for="item in tagList"
              :key="item.id"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="发布状态" prop="status">
          <el-select
            v-model="queryInfo.status"
            placeholder="请选择"
            @change="getTableData"
            clearable
          >
            <el-option label="草稿" :value="0"></el-option>
            <el-option label="已发布" :value="1"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="文章状态" prop="is_deleted">
          <el-select
            v-model="queryInfo.is_deleted"
            placeholder="请选择"
            @change="getTableData"
            clearable
          >
            <el-option label="正常" :value="0"></el-option>
            <el-option label="已删除" :value="1"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button :icon="Search" @click="getTableData">搜索</el-button>
          <el-button :icon="Refresh" @click="resetQueryForm(queryFormRef)">重置</el-button>
          <el-button :icon="Plus" type="primary" plain @click="$router.push('/article/editor')"
            >新建文章</el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
    <el-card shadow="never" class="table">
      <el-table :data="tableData.list">
        <el-table-column label="封面" align="left" width="100">
          <template #default="{ row }">
            <el-image
              v-if="row.cover_img"
              fit="cover"
              style="width: 75px; height: 48px"
              :src="config.file_base_url + row.cover_img"
              :preview-teleported="true"
              :preview-src-list="[config.file_base_url + row.cover_img]"
            ></el-image>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="文章标题" align="left" min-width="300" />
        <el-table-column prop="cate_name" label="分类" align="left" min-width="150" />
        <el-table-column label="标签" align="left" min-width="150">
          <template #default="{ row }">
            <el-tag type="primary">{{ row.tag_name }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="views" label="浏览量" align="left" min-width="100" />
        <el-table-column label="发布状态" align="left" min-width="100">
          <template #default="{ row }">
            <el-tag v-if="row.status === 0" type="warning">草稿</el-tag>
            <el-tag v-if="row.status === 1" type="success">已发布</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="文章状态" align="left" min-width="100">
          <template #default="{ row }">
            <el-tag v-if="row.is_deleted === 0" type="success">正常</el-tag>
            <el-tag v-if="row.is_deleted === 1" type="danger">已删除</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="creator" label="创建人" align="left" min-width="150" />
        <el-table-column prop="create_time" label="创建时间" align="left" min-width="180" />
        <el-table-column prop="updater" label="更新人" align="left" min-width="150" />
        <el-table-column prop="update_time" label="更新时间" align="left" min-width="180" />
        <el-table-column label="操作" align="left" width="150" fixed="right">
          <template #default="{ row }">
            <div class="operate">
              <el-link
                :underline="false"
                type="primary"
                @click="$router.push({ path: '/article/editor', query: { id: row.id } })"
                >编辑</el-link
              >
              <el-link
                :underline="false"
                type="danger"
                @click="delRow(row)"
                v-if="row.is_deleted === 0"
                >删除</el-link
              >
              <el-link
                :underline="false"
                type="success"
                @click="recoverRow(row)"
                v-if="row.is_deleted === 1"
                >恢复</el-link
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
</template>

<style scoped lang="scss">
.article-page {
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
