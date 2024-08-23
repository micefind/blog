<script setup lang="ts">
import { Avatar, UserFilled, House } from '@element-plus/icons-vue';
import { useTagsViewStore } from "../../store/modules/tagsView.ts";
import { useAppStore } from "../../store/modules/app.ts";

// Define the type for a menu item
interface MenuItem {
    name: string;
    path: string;
    icon: typeof Avatar | typeof UserFilled | typeof House | string;
    children: MenuItem[];
}

const tagsStore = useTagsViewStore();
const appStore = useAppStore();

const menuList: MenuItem[] = [
    {
        name: '首页',
        path: '/overview',
        icon: House,
        children: []
    },
    {
        name: '用户管理',
        path: '',
        icon: UserFilled,
        children: [
            {
                name: '用户列表',
                path: '/user',
                icon: '',
                children: []
            }
        ]
    },
    {
        name: '个人中心',
        path: '/profile',
        icon: Avatar,
        children: []
    }
];

// Define the type for the item parameter based on MenuItem
const saveStatus = (item: MenuItem) => {
    if (item.path !== tagsStore.activePath) {
        tagsStore.addVisitedView(item);
    }
    tagsStore.setActivePath(item);
};

// Initialize activePath
tagsStore.setActivePath(null);
</script>

<template>
    <el-scrollbar>
        <el-menu :default-active="tagsStore.activePath" :collapse="!appStore.isCollapse" :router="true"
            class="aside-menu" :collapse-transition="true" :unique-opened="true">
            <template v-for="item in menuList">
                <el-sub-menu v-if="item.path === ''" :index="item.name">
                    <template #title>
                        <el-icon>
                            <component :is="item.icon" />
                        </el-icon>
                        <span>{{ item.name }}</span>
                    </template>
                    <el-menu-item :index="item2.path" v-for="item2 in item.children" @click="saveStatus(item2)">
                        <el-icon>
                            <component :is="item2.icon" />
                        </el-icon>
                        <template #title>
                            <span>{{ item2.name }}</span>
                        </template>
                    </el-menu-item>
                </el-sub-menu>
                <el-menu-item :index="item.path" v-else @click="saveStatus(item)">
                    <el-icon>
                        <component :is="item.icon" />
                    </el-icon>
                    <template #title>
                        <span>{{ item.name }}</span>
                    </template></el-menu-item>
            </template>
        </el-menu>
    </el-scrollbar>
</template>

<style scoped>
.aside-menu {
    height: calc(100vh - 50px);
    user-select: none;
}

.aside-menu:not(.el-menu--collapse) {
    width: 200px;
}
</style>
