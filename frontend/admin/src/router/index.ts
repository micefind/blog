import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import { close, start } from '../utils/nporgress.ts';

const routes: RouteRecordRaw[] = [
    {
        path: '/login',
        meta: { title: '登录' },
        name: 'login',
        component: () => import('@/view/login/Login.vue')
    },
    {
        path: '/',
        meta: { title: '首页' },
        name: 'home',
        component: () => import('@/view/home/Index.vue'),
        redirect: '/overview',
        children: [
            {
                path: '/overview',
                meta: { title: '首页' },
                name: 'overview',
                component: () => import('@/view/overview/Overview.vue')
            },
            {
                path: '/:pathMatch(.*)*',
                meta: { title: '404' },
                name: '404',
                component: () => import('@/view/error/404.vue')
            },
            {
                path: '/profile',
                meta: { title: '个人中心' },
                name: 'profile',
                component: () => import('@/view/profile/Profile.vue')
            },
            {
                path: '/user',
                meta: { title: '用户管理' },
                name: 'user',
                component: () => import('@/view/user/User.vue')
            }
        ]
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

router.beforeEach((to, _, next) => {
    start();
    if (to.meta.title && typeof to.meta.title === 'string') {
        document.title = to.meta.title;
    }
    const token = sessionStorage.getItem('token');
    if (to.name !== 'login' && !token) {
        next('/login');
    } else {
        next();
    }
});

router.afterEach(() => {
    close();
});

export default router;
