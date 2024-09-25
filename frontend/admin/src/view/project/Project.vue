<script setup lang="ts">
import { ref } from "vue";
import { Search, Plus, Refresh, Close, FullScreen, Picture } from "@element-plus/icons-vue";
import type { FormInstance } from "element-plus";
import request from "../../utils/request";
import { ElMessageBox, ElMessage } from "element-plus";

// 定义查询信息类型
interface QueryInfo {
  project_name: string;
  pageNum: number;
  pageSize: number;
}

// 定义项目信息类型
interface ProjectInfo {
  id?: number;
  project_name: string;
  description: string;
  logo: string;
  url: string;
}

// 定义表格数据类型
interface TableData {
  list: ProjectInfo[];
  total: number;
}

// 定义上传文件类型
interface FileItem {
  url?: string;
  raw?: File;
}

// 表单引用
const queryFormRef = ref<FormInstance | null>(null);
const projectFormRef = ref<FormInstance | null>(null);

// 弹窗控制
const dialogVisible = ref(false);
const fullscreen = ref(false);

// 表格数据
const tableData = ref<TableData>({ list: [], total: 0 });

// 查询信息
const queryInfo = ref<QueryInfo>({
  project_name: "",
  pageNum: 1,
  pageSize: 20,
});

// 项目信息
const projectInfo = ref<ProjectInfo>({
  project_name: "",
  description: "",
  logo: "",
  url: "",
});

// 表单验证规则
const rules = {
  project_name: [{ required: true, message: "请输入项目名称", trigger: "blur" }],
};

// 上传文件列表
const fileList = ref<FileItem[]>([]);

// 删除项目
const deleteItem = async (id: number) => {
  try {
    await ElMessageBox.confirm("删除后不可恢复，是否继续", "提示", {
      confirmButtonText: "确认",
      cancelButtonText: "取消",
      type: "warning",
    });
    const { data: res } = await request.post("/project/delete", { id });
    if (res.status === 200) {
      ElMessage.success("删除成功");
      await getTableData();
    }
  } catch {
  }
};

// 上传图片
const uploadImg = async (file: File): Promise<string | null> => {
  const formData = new FormData();
  formData.append("file", file);
  const { data: res } = await request.post("/upload/image", formData);
  return res.status === 200 ? res.data.url : null;
};

// 提交表单
const submitForm = async (formEl: FormInstance | null) => {
  if (!formEl) return;
  await formEl.validate(async (valid) => {
    if (valid) {
      if (fileList.value.length > 0 && fileList.value[0].raw) {
        const logoUrl = await uploadImg(fileList.value[0].raw);
        projectInfo.value.logo = logoUrl || "";
      }
      if (fileList.value.length ===0) projectInfo.value.logo = ''
      const url = projectInfo.value.id ? "/project/edit" : "/project/add";
      const { data: res } = await request.post(url, projectInfo.value);
      if (res.status === 200) {
        dialogVisible.value = false;
        ElMessage.success(projectInfo.value.id ? "更新成功" : "添加成功");
        await getTableData();
        resetUserInfo();
      }
    }
  });
};

// 获取项目详情
const getDetails = async (id: number) => {
  fileList.value = [];
  const { data: res } = await request.post("/project/details", { id });
  if (res.status === 200) {
    projectInfo.value = res.data;
    if (res.data.logo) {
      fileList.value.push({ url: res.data.logo });
    }
    dialogVisible.value = true;
  }
};

// 重置项目信息
const resetUserInfo = () => {
  projectInfo.value = {
    project_name: "",
    description: "",
    logo: "",
    url: "",
  };
};

// 显示新增项目弹窗
const showAddDialog = () => {
  if (projectInfo.value.id)
    resetUserInfo();
  fileList.value = [];
  dialogVisible.value = true;
};

// 重置查询表单
const resetQueryForm = async (formEl: FormInstance | null) => {
  formEl?.resetFields();
  await getTableData();
};

// 分页处理
const handleSizeChange = async (val: number) => {
  queryInfo.value.pageSize = val;
  await getTableData();
};

const handleCurrentChange = async (val: number) => {
  queryInfo.value.pageNum = val;
  await getTableData();
};

// 获取表格数据
const getTableData = async () => {
  const { data: res } = await request.post("/project/list", queryInfo.value);
  if (res.status === 200) {
    tableData.value = res.data;
  }
};

// 初始化数据
getTableData();
</script>


<template>
  <el-card shadow="never" class="header-box">
    <el-form ref="queryFormRef" :model="queryInfo">
      <el-form-item label="项目名称" prop="project_name">
        <el-input
        @keyup.enter="getTableData"
          placeholder="请输入项目名称"
          v-model="queryInfo.project_name"
          clearable></el-input>
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
        <el-table-column type="index" label="标志" align="center" width="150">
            <template #default="scope">
                <el-image style="width: 50px; height: 50px" :src="scope.row.logo" fit="cover" v-if="scope.row.logo" :preview-src-list="[scope.row.logo]" :preview-teleported="true"/>
                <el-image v-else>
        <template #error>
          <div>
            <el-icon style="height: 50px;width: 50px;"><Picture style="width: 2em;height: 2em"/></el-icon>
          </div>
        </template>
      </el-image>
            </template>
        </el-table-column>
      <el-table-column prop="project_name" label="项目名称" align="center" />
      <el-table-column  label="地址" align="center" >
        <template #default="scope">
            <el-link :href="scope.row.url" target="_blank" :underline="false">{{scope.row.url}}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="description" label="描述" align="center" />
      <el-table-column label="操作" align="center" fixed="right" width="150">
        <template #default="scope">
          <el-link
            style="margin-right: 10px"
            type="primary"
            @click="getDetails(scope.row.id)"
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
         {{ projectInfo.id ? '编辑项目信息' : '新增项目' }}
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
        ref="projectFormRef"
        label-width="auto"
        :model="projectInfo"
        :rules="rules">
        <el-form-item prop="project_name" label="项目名称">
          <el-input placeholder="请输入项目" v-model="projectInfo.project_name">
          </el-input>
        </el-form-item>
        <el-form-item prop="url" label="地址">
          <el-input placeholder="请输入地址" v-model="projectInfo.url">
          </el-input>
        </el-form-item>
        <el-form-item prop="description" label="描述">
          <el-input placeholder="请输入描述" v-model="projectInfo.description" :autosize="{ minRows: 4, maxRows: 8 }"
          type="textarea" maxlength="100" show-word-limit>
          </el-input>
        </el-form-item>
        <el-form-item prop="logo" label="标志">
          <el-upload 
    v-model:file-list="fileList"
    limit="1"
    list-type="picture-card"
    :auto-upload="false"
  >
    <el-icon><Plus /></el-icon>
  </el-upload>
        </el-form-item>
      </el-form>
    </div>
    <template #footer>
      <div style="padding: 10px 20px 20px">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm(projectFormRef)">
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
