<script setup lang="ts">
import { useTagsViewStore } from "../../store/modules/tagsView.ts";
import { DArrowLeft, DArrowRight, Close } from '@element-plus/icons-vue';
import { ref } from "vue";
import { useRouter } from 'vue-router';

interface TagItem {
    name: string;
    path: string;
}

const router = useRouter();
const tagsStore = useTagsViewStore();
const scroll = ref<HTMLElement | null>(null); // 定义 ref 类型为 HTMLElement 或 null
const scrollLeft = ref<number>(0);

// 删除标签
const delTag = (i: number) => {
    if (i === tagsStore.visitedViews.length - 1 && i !== 0) {
        gotoPath(tagsStore.visitedViews[tagsStore.visitedViews.length - 2]);
        tagsStore.deleteVisitedView(i);
    }
    if (i < tagsStore.visitedViews.length - 1) {
        gotoPath(tagsStore.visitedViews[i + 1]);
        tagsStore.deleteVisitedView(i);
    }
};

// 跳转页面并保存选中状态
const gotoPath = (item: TagItem) => {
    tagsStore.setActivePath(item);
    router.push(item.path);
};

// 获取滚动条位置
const handleScroll = (e: Event) => {
    const target = e.target as HTMLElement;
    scrollLeft.value = target.scrollLeft;
};

// 点击按钮实现滚动条左右滚动
const scrollTo = (direction: 'left' | 'right') => {
    if (scroll.value) {
        const scrollAmount = direction === 'left' ? scrollLeft.value - 200 : scrollLeft.value + 200;
        scroll.value.scrollTo({ left: scrollAmount, behavior: 'smooth' });
    }
};

// 初始化 tag
tagsStore.setVisitedView();
// 初始化 activePath
tagsStore.setActivePath(null);
</script>

<template>
    <div class="tags">
        <span class="to-left-right" @click="scrollTo('left')">
            <el-icon>
                <d-arrow-left />
            </el-icon>
        </span>
        <div class="tag-box">
            <el-scrollbar @scroll.native="handleScroll" ref="scroll"
                :wrap-style="{ display: 'flex', alignItems: 'center' }">
                <div class="btn-box">
                    <el-button :type="tagsStore.activePath === item.path ? 'primary' : ''"
                        v-for="(item, i) in tagsStore.visitedViews" @click="gotoPath(item)">
                        <span>{{ item.name }}</span>
                        <span class="icon">
                            <el-icon @click="delTag(i)">
                                <Close />
                            </el-icon>
                        </span>
                    </el-button>
                </div>
            </el-scrollbar>
        </div>
        <span class="to-left-right" @click="scrollTo('right')">
            <el-icon>
                <d-arrow-right />
            </el-icon>
        </span>
    </div>
</template>

<style scoped lang="scss">
.tags {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.tag-box {
    flex: 1;
    border-right: 1px solid var(--el-border-color);
    border-left: 1px solid var(--el-border-color);

    .btn-box {
        display: flex;
        align-items: center;
        height: 35px;
        width: 0;
    }
}

.el-button {
    margin: 0 2px;
    border-radius: 2px;
    padding-right: 20px;
    position: relative;

    &:nth-child(1) {
        margin-left: 4px;
    }

    &:hover {
        .icon {
            display: block;
        }
    }

    span {
        font-size: 12px;
    }

    .icon {
        position: absolute;
        right: 5px;
        top: 9px;
        display: flex;
        align-items: center;
        padding-left: 5px;
        display: none;
    }
}

.to-left-right {
    display: flex;
    flex-shrink: 0;
    align-items: center;
    justify-content: center;
    width: 35px;
    height: 35px;
    cursor: pointer;
}
</style>
