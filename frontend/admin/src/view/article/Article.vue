<script setup lang="ts">
import { ref } from "vue"
import { Search, Plus, Refresh,Picture } from "@element-plus/icons-vue"
import type { FormInstance } from "element-plus"
import request from "../../utils/request"
import { useRouter } from "vue-router"

const router = useRouter()

// 定义类型
interface QueryInfo {
  keyword: string
  status:string
  pageNum: number
  pageSize: number
}

interface ProjectInfo {
  id: number
  title: string
  intro: string
  cover_image: string
  keywords: string
  views: number
  creator_id: number
  create_time: string
  status: string,
  creator: string
}

interface TableData {
  list: ProjectInfo[]
  total: number
}

const queryFormRef = ref<FormInstance | null>(null)
const tableData = ref<TableData>({ list: [], total: 0 })

const queryInfo = ref<QueryInfo>({
  keyword: "",
  status:'',
  pageNum: 1,
  pageSize: 20,
})

const deleteItem=async (id: number) => {
  try {
    await ElMessageBox.confirm("删除后不可恢复，是否继续", "提示", {
      confirmButtonText: "确认",
      cancelButtonText: "取消",
      type: "warning",
    });
    const { data: res } = await request.post("/article/delete",{id})
  if (res.status === 200) {
    ElMessage.success("删除成功");
    await getTableData()
  }
  } catch {
    ElMessage.info("已取消");
  }
}

const gotoEditor = (id?: number) => {
  router.push({
    path: "/article/editor",
    query: { id },
  })
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
  const { data: res } = await request.post("/article/list", queryInfo.value)
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
      <el-form-item label="关键词" prop="keyword">
        <el-input
        @keyup.enter="getTableData"
          placeholder="请输入关键词"
          v-model="queryInfo.keyword"
          clearable></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select placeholder="请选择状态" v-model="queryInfo.status" clearable>
          <el-option label="提交" value="1" />
          <el-option label="发布" value="2" />
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
        <el-button type="primary" plain @click="gotoEditor()">
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
      <el-table-column type="index" label="封面" align="center" width="120">
            <template #default="scope">
                <el-image style="width: 80px;height: 45px;" :src="scope.row.cover_image" fit="cover" v-if="scope.row.cover_image" :preview-src-list="[scope.row.cover_image]" :preview-teleported="true"/>
                <el-image v-else>
        <template #error>
          <div>
            <el-icon style="height: 50px;width: 50px;"><Picture style="width: 2em;height: 2em"/></el-icon>
          </div>
        </template>
      </el-image>
            </template>
        </el-table-column>
      <el-table-column prop="title" label="标题" align="center" width="240"/>
      <el-table-column prop="intro" label="简介" align="center" show-overflow-tooltip/>
      <el-table-column prop="keywords" label="关键词" align="center" show-overflow-tooltip/>
      <el-table-column prop="views" label="阅读量" align="center" />
      <el-table-column label="状态" align="center">
        <template #default="scope">
          <el-tag
            :type="
              scope.row.status === '0'
                ? 'wanning'
                : scope.row.status === '1'
                ? 'primary'
                : 'success'
            "
            >{{
              scope.row.status === "0"
                ? "草稿"
                : scope.row.status === "1"
                ? "暂存"
                : "发布"
            }}</el-tag
          >
        </template>
      </el-table-column>
      <el-table-column prop="create_time" label="创建时间" align="center" width="180"/>
      <el-table-column prop="creator" label="创建人" align="center" />

      <el-table-column label="操作" align="center" fixed="right" width="150">
        <template #default="scope">
          <el-link
            style="margin-right: 10px"
            type="primary"
            @click="gotoEditor(scope.row.id)"
            :underline="false"
            >编辑</el-link
          >
          <el-link
            type="danger"
            @click="deleteItem(scope.row.id)"
            :underline="false"
            >删除</el-link
          >
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
