import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    sidebarCollapsed: localStorage.getItem('sidebarCollapsed') === 'true',
    loading: false,
    activePath: localStorage.getItem('activePath') || '/',
  }),
  actions: {
    toggleSidebar() {
      this.sidebarCollapsed = !this.sidebarCollapsed
      localStorage.setItem('sidebarCollapsed', this.sidebarCollapsed)
    },
    setLoading(status) {
      this.loading = status
    },
    setActivePath(path) {
      this.activePath = path
      localStorage.setItem('activePath', path)
    },
  },
})
