<script setup lang="ts">
import LayoutHeader from '@/components/layout/LayoutHeader.vue';
import LayoutAside from '@/components/layout/LayoutAside.vue';
import LayoutTab from '@/components/layout/LayoutTab.vue';
import { useAppStore } from "../../store/modules/app.ts"

const appStore = useAppStore()

if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
    // 当前设备是手机
    appStore.setIsMobile(true)
} else {
    // 当前设备是电脑
    appStore.setIsMobile(false)
}
</script>

<template>
    <div class="root-container">
        <el-container>
            <el-header>
                <layout-header></layout-header>
            </el-header>
            <el-container class="view-container">
                <el-aside style="width: auto;" v-if="!appStore.isMobile">
                    <layout-aside></layout-aside>
                </el-aside>
                <el-main>
                    <div class="tab-box">
                        <layout-tab></layout-tab>
                    </div>
                    <el-scrollbar class="scroll-box">
                        <div class="main-box">
                            <router-view />
                        </div>
                    </el-scrollbar>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>

<style scoped lang="scss">
.root-container {
    width: 100%;
    height: 100vh;
    position: fixed;
    top: 0;
    left: 0;
}

.view-container {
    height: calc(100vh - 50px);
}

.el-header {
    height: 50px;
    padding: 0 !important;
}

.el-main {
    padding: 0;
    overflow: hidden;
    display: flex;
    flex-direction: column;
}

.tab-box {
    height: 36px;
    border-bottom: 1px solid var(--el-border-color);
}

.scroll-box {
    flex: 1;
    background-color: #f5f7f9;
}

.main-box {
    padding: 20px;
}
</style>