import { createRouter, createWebHashHistory } from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import { useAppStore } from '@/store/app'
import routes from '@/config/routes'

const modules = import.meta.glob('../views/**/*.vue')

//全局进度条的配置
NProgress.configure({
  easing: 'ease', // 动画方式，和css动画规格一样（默认：ease）
  speed: 200, // 递增进度条的速度，单位ms（默认： 200）
  showSpinner: false, // 是否显示加载ico
  trickle: true, //是否自动递增
  trickleSpeed: 200, // 自动递增间隔
  minimum: 0.08, // 初始化时的最小百分比，0-1（默认：0.08）
  parent: 'body', //指定进度条的父容器
})

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL), // 使用 Hash 模式创建路由
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login/index.vue'),
      meta: { title: '登录' },
    },
    {
      path: '/admin',
      name: 'layout',
      component: () => import('@/layout/index.vue'),
      redirect: '/home',
      children: [
        {
          path: '/:pathMatch(.*)',
          name: '404',
          component: () => import('@/views/error/404.vue'),
          meta: { title: '404' },
        },
        {
          path: '/user/profile',
          name: 'userProfile',
          component: () => import('@/views/user/profile/index.vue'),
          meta: { title: '个人中心' },
        },
      ],
    },
    {
      path: '/article/editor',
      name: 'articleEditor',
      component: () => import('@/views/article/editor/index.vue'),
      meta: { title: '文章编辑', keepAlive: false },
    },
    {
      path: '/',
      name: 'blog',
      component: () => import('@/views/blog/index.vue'),
      meta: { title: '鼠觅', keepAlive: true },
    },
    {
      path: '/blog/article',
      name: 'blogArticle',
      component: () => import('@/views/blog/article/index.vue'),
      meta: { title: '文章-鼠觅', keepAlive: false },
    },
  ],
})

// 递归扁平化路由
const flattenRoutes = (routes) => {
  let result = []
  routes.forEach((item) => {
    if (item.children) {
      result.push(...flattenRoutes(item.children))
    } else {
      result.push(item)
    }
  })
  return result
}

// 初始化路由
const initRouter = () => {
  const newRoutes = flattenRoutes(routes)
  newRoutes.forEach((item) => {
    if (!router.hasRoute(item.name)) {
      router.addRoute('layout', {
        ...item,
        component: modules[/* @vite-ignore */ `../${item.component}`],
      })
    }
  })
}

// 路由守卫
router.beforeEach(async (to, from, next) => {
  // 进度条开始
  NProgress.start()
  // 设置页面标题
  to.meta.title ? (document.title = to.meta.title) : null
  // 判断是否已登录且没有 token，未登录时重定向到登录页
  const token = localStorage.getItem('token')
  const whiteList = ['/login', '/', '/blog/article']
  if (!whiteList.includes(to.path) && !token) {
    return next({ name: 'login' })
  }
  // 设置管理后台路由的菜单激活
  const appStore = useAppStore()
  appStore.setActivePath(to.path)

  next()
})

router.afterEach(() => {
  // 进度条结束
  NProgress.done()
})

export { router, initRouter }
