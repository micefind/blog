<script setup lang="ts">
import { User } from '@element-plus/icons-vue'
import { ref } from 'vue'
import defaultAvatar from '../../assets/image/defaultAvatar.png'

const activeName = ref('first')

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
                    <li>
                        <span>
                            <el-icon>
                                <User />
                            </el-icon>
                            <span>用户名</span>
                        </span>
                        <span>admin</span>
                    </li>

                    <li>1</li>
                    <li>1</li>
                    <li>1</li>
                    <li>1</li>

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
                    <el-tab-pane label="基本资料" name="first">User</el-tab-pane>
                    <el-tab-pane label="修改密码" name="second">Config</el-tab-pane>
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