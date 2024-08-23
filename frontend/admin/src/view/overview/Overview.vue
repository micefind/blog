<script setup>
import { onMounted, ref } from 'vue';
import request from '@/utils/request';
import { useDark, useToggle } from "@vueuse/core";


const onClick = () => {
    ElMessage({
        message: 'This is a message.',
        grouping: true,
        type: 'success',
    })
}

const isDark = useDark()
const toggleDark = () => useToggle(isDark)


import { useTagsViewStore } from "../../store/modules/tagsView.ts"

const tagsStore = useTagsViewStore()


const msg = ref('')

const getList = async () => {
    const { data: res } = await request.post('api/public/user/login', {
        username: 'admin',
        password: '123456'
    })
}

onMounted(async () => {
    // await getList()
})

</script>

<template>

    <el-switch v-model="isDark" inline-prompt active-color="var(--el-fill-color-dark)"
        inactive-color="var(--el-color-primary)" @change="toggleDark" />
    <div>
        <el-button @click="onClick">Default</el-button>
        <el-button type="primary">Primary</el-button>
        <el-button type="success">Success</el-button>
        <el-button type="info">Info</el-button>
        <el-button type="warning">Warning</el-button>
        <el-button type="danger">Danger</el-button>
        <el-button>{{ msg }}</el-button>
        <el-card shadow="never" style="height: 200px;margin: 20px 0"></el-card>
    </div>
</template>

<style scoped lang="scss">
.box {
    width: 400px;
    height: 400px;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #789;

    .box-item {
        width: 200px;
        height: 200px;
        background-color: #998;
    }
}
</style>