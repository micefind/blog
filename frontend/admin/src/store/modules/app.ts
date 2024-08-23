import { defineStore } from "pinia";

export const useAppStore = defineStore("app", {
    state: () =>({
        theme: '#409eff', // 系统主题色
        isMobile: false, // 设备是否手机
        isCollapse: true, // 侧边栏折叠展开
    }),
    actions: {
        setIsMobile(type:boolean) {
            this.isMobile = type
        },
        changeCollapse() {
            this.isCollapse = !this.isCollapse
        }

    }
})