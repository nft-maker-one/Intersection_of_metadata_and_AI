// * +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// * Copyright 2023 The Geek-AI Authors. All rights reserved.
// * Use of this source code is governed by a Apache-2.0 license
// * that can be found in the LICENSE file.
// * @Author yangjian102621@163.com
// * +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

import {createRouter, createWebHistory} from "vue-router";

const routes = [
    {
        name: 'Index',
        path: '/',
        meta: {title: process.env.VUE_APP_TITLE},
        component: () => import('@/views/Index.vue'),
    },
    {
        name: 'home',
        path: '/home',
        redirect: '/chat',
        meta: {title: '首页'},
        component: () => import('@/views/Home.vue'),
        children: [
            {
                name: 'chat',
                path: '/chat',
                meta: {title: '创作中心'},
                component: () => import('@/views/ChatPlus.vue'),
            },
            {
                name:'map',
                path:'/map',
                component:() => import('@/views/Map.vue')
            },{
                name:'market',
                path:'/market',
                meta:{title:'购买 Api 算力'},
                component:() => import('@/views/Market.vue')
            },
        ]
    },
    {
        name:"payment",
        path:"/payment",
        component:() => import('@/views/Payment.vue')
    },
    {
        name:'market',
        path:'/market',
        meta:{title:'购买 Api 算力'},
        component:() => import('@/views/Market.vue')
    },
    {
        name:'map',
        path:'/map',
        meta:{title: '算力提供者分布'},
        component:() => import('@/views/Map.vue'),
    },
    {
        name: 'chat-export',
        path: '/chat/export',
        meta: {title: '导出会话记录'},
        component: () => import('@/views/ChatExport.vue'),
    },
    {
        name: 'login',
        path: '/login',
        meta: {title: '用户登录'},
        component: () => import('@/views/Login.vue'),
    },
    {
        name: 'register',
        path: '/register',

        meta: {title: '用户注册'},
        component: () => import('@/views/Register.vue'),
    },
    {
        path: '/admin/login',
        name: 'admin-login',
        meta: {title: 'ChatPuls 控制台登录'},
        component: () => import('@/views/admin/Login.vue'),
    },
    {
        name: 'admin',
        path: '/admin',
        redirect: '/admin/dashboard',
        component: () => import("@/views/admin/Home.vue"),
        meta: {title: 'ChatPuls 管理后台'},
        children: [
            {
                path: '/admin/dashboard',
                name: 'admin-dashboard',
                meta: {title: '仪表盘'},
                component: () => import('@/views/admin/Dashboard.vue'),
            },
            {
                path: '/admin/system',
                name: 'admin-system',
                meta: {title: '系统设置'},
                component: () => import('@/views/admin/SysConfig.vue'),
            },
            {
                path: '/admin/user',
                name: 'admin-user',
                meta: {title: '用户管理'},
                component: () => import('@/views/admin/Users.vue'),
            },
            {
                path: '/admin/role',
                name: 'admin-role',
                meta: {title: '角色管理'},
                component: () => import('@/views/admin/Roles.vue'),
            },
            {
                path: '/admin/apikey',
                name: 'admin-apikey',
                meta: {title: 'API-KEY 管理'},
                component: () => import('@/views/admin/ApiKey.vue'),
            },
            {
                path: '/admin/chat/model',
                name: 'admin-chat-model',
                meta: {title: '语言模型'},
                component: () => import('@/views/admin/ChatModel.vue'),
            },
            {
                path: '/admin/product',
                name: 'admin-product',
                meta: {title: '充值产品'},
                component: () => import('@/views/admin/Product.vue'),
            },
            {
                path: '/admin/order',
                name: 'admin-order',
                meta: {title: '充值订单'},
                component: () => import('@/views/admin/Order.vue'),
            },
            {
                path: '/admin/reward',
                name: 'admin-reward',
                meta: {title: '众筹管理'},
                component: () => import('@/views/admin/Reward.vue'),
            },
            {
                path: '/admin/loginLog',
                name: 'admin-loginLog',
                meta: {title: '登录日志'},
                component: () => import('@/views/admin/LoginLog.vue'),
            },
            {
                path: '/admin/functions',
                name: 'admin-functions',
                meta: {title: '函数管理'},
                component: () => import('@/views/admin/Functions.vue'),
            },
            {
                path: '/admin/chats',
                name: 'admin-chats',
                meta: {title: '对话管理'},
                component: () => import('@/views/admin/ChatList.vue'),
            },
            {
                path: '/admin/powerLog',
                name: 'admin-power-log',
                meta: {title: '算力日志'},
                component: () => import('@/views/admin/PowerLog.vue'),
            },
            {
                path: '/admin/manger',
                name: 'admin-manger',
                meta: {title: '管理员'},
                component: () => import('@/views/admin/Manager.vue'),
            }
        ]
    },


    {
        name: 'mobile',
        path: '/mobile',
        meta: {title: 'Geek-AI v4.0'},
        component: () => import('@/views/mobile/Home.vue'),
        redirect: '/mobile/index',
        children: [
            {
                path: '/mobile/index',
                name: 'mobile-index',
                component: () => import('@/views/mobile/Index.vue'),
            },
            {
                path: '/mobile/chat',
                name: 'mobile-chat',
                component: () => import('@/views/mobile/ChatList.vue'),
            },
            {
                path: '/mobile/image',
                name: 'mobile-image',
                component: () => import('@/views/mobile/Image.vue'),
            },
            {
                path: '/mobile/profile',
                name: 'mobile-profile',
                component: () => import('@/views/mobile/Profile.vue'),
            },
            {
                path: '/mobile/imgWall',
                name: 'mobile-img-wall',
                component: () => import('@/views/mobile/pages/ImgWall.vue'),
            },
            {
                path: '/mobile/chat/session',
                name: 'mobile-chat-session',
                component: () => import('@/views/mobile/ChatSession.vue'),
            },
            {
                path: '/mobile/chat/export',
                name: 'mobile-chat-export',
                component: () => import('@/views/mobile/ChatExport.vue'),
            },
        ]
    },

    {
        name: 'test',
        path: '/test',
        meta: {title: '测试页面'},
        component: () => import('@/views/Test.vue'),
    },
    {
        name: 'NotFound',
        path: '/:all(.*)',
        meta: {title: '页面没有找到'},
        component: () => import('@/views/404.vue'),
    },
]

// console.log(MY_VARIABLE)
const router = createRouter({
    mode:'history',
    history: createWebHistory(),
    routes: routes,
})

let prevRoute = null
// dynamic change the title when router change
router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = `${to.meta.title} | ${process.env.VUE_APP_TITLE}`
    }
    prevRoute = from
    next()
})

export {router, prevRoute};