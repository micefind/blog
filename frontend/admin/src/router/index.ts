import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router"
import { close, start } from "../utils/nporgress.ts"

const routes: RouteRecordRaw[] = [
  {
    path: "/login",
    meta: { title: "登录" },
    name: "login",
    component: () => import("@/view/login/Login.vue"),
  },
  {
    path: "/",
    meta: { title: "首页" },
    name: "home",
    component: () => import("@/view/home/Index.vue"),
    redirect: "/article",
    children: [
      {
        path: "/:pathMatch(.*)*",
        meta: { title: "404" },
        name: "404",
        component: () => import("@/view/error/404.vue"),
      },
      {
        path: "/article",
        meta: { title: "文章管理" },
        name: "article",
        component: () => import("@/view/article/Article.vue"),
      },
      {
        path: "/project",
        meta: { title: "项目管理" },
        name: "project",
        component: () => import("@/view/project/Project.vue"),
      },
      {
        path: "/user",
        meta: { title: "用户管理" },
        name: "user",
        component: () => import("@/view/user/User.vue"),
      },
      {
        path: "/profile",
        meta: { title: "个人中心" },
        name: "profile",
        component: () => import("@/view/profile/Profile.vue"),
      },
    ],
  },
  {
    path: "/article/editor",
    meta: { title: "文章编辑" },
    name: "editor",
    component: () => import("@/view/article/Editor.vue"),
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

router.beforeEach((to, _, next) => {
  start()
  if (to.meta.title && typeof to.meta.title === "string") {
    document.title = to.meta.title
  }
  const token = localStorage.getItem("token")
  if (to.name !== "login" && !token) {
    next("/login")
  } else {
    next()
  }
})

router.afterEach(() => {
  close()
})

export default router
