import { HomeFilled, UserFilled, Reading } from '@element-plus/icons-vue'

const routes = [
  {
    path: '/home',
    name: 'home',
    component: 'views/home/index.vue',
    meta: {
      title: '首页',
      icon: HomeFilled,
      keepAlive: false,
    },
  },
  {
    path: '/article',
    name: 'article',
    component: null,
    meta: {
      title: '文章管理',
      icon: Reading,
      keepAlive: false,
    },
    children: [
      {
        path: '/article/list',
        name: 'articleList',
        component: 'views/article/list/index.vue',
        meta: {
          title: '文章列表',
          keepAlive: false,
        },
      },
      {
        path: '/article/cate',
        name: 'articleCate',
        component: 'views/article/cate/index.vue',
        meta: {
          title: '文章分类',
          keepAlive: false,
        },
      },
      {
        path: '/article/tag',
        name: 'articleTag',
        component: 'views/article/tag/index.vue',
        meta: {
          title: '文章标签',
          keepAlive: false,
        },
      },
    ],
  },

  {
    path: '/user',
    name: 'user',
    component: 'views/user/index.vue',
    meta: {
      title: '用户管理',
      icon: UserFilled,
      keepAlive: false,
    },
  },
]

export default routes
