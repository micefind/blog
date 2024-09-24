import { defineStore } from "pinia";

interface TagView {
  name: string;
  path: string;
  [key: string]: any; // 你可以根据需要添加其他字段的类型
}

export const useTagsViewStore = defineStore("tags", {
  state: () => ({
    visitedViews: [] as TagView[], // 明确声明 visitedViews 的类型
    activePath: ''
  }),
  actions: {
    // 设置 activePath
    setActivePath(view: TagView | null) {
      this.activePath = sessionStorage.getItem('activePath') 
        ? sessionStorage.getItem('activePath') || '' 
        : this.activePath;
      if (view !== null) {
        this.activePath = view.path;
        sessionStorage.setItem('activePath', view.path);
      }
    },
    // 初始化 tag 列表
    setVisitedView() {
      this.visitedViews = sessionStorage.getItem('visitedViews')
        ? JSON.parse(sessionStorage.getItem('visitedViews') || '[]') as TagView[]
        : [];
    },
    // 新增 tag
    addVisitedView(view: TagView) {
      this.visitedViews.push(view);
      const map = new Map<string, TagView>();
      for (const item of this.visitedViews) {
        if (!map.has(item.path)) {
          map.set(item.path, item);
        }
      }
      this.visitedViews = [...map.values()];
      sessionStorage.setItem('visitedViews', JSON.stringify(this.visitedViews));
    },
    // 删除 tag
    deleteVisitedView(index: number) {
      this.visitedViews.splice(index, 1);
      sessionStorage.setItem('visitedViews', JSON.stringify(this.visitedViews));
    },
  }
});
