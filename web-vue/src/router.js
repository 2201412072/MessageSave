import { createRouter,createWebHistory,createWebHashHistory } from 'vue-router'
import HomePage from '@/start/HomePage.vue';
import Register from '@/start/Register.vue';
import Login from '@/start/Login.vue';
import ClientHomePage from '@/client/ClientHomePage.vue';
import ClientPasswordManage from '@/client/ClientPasswordManage.vue';
import ClientAnnouncementManage from '@/client/ClientAnnouncementManage.vue';
import ClientConfig from '@/client/ClientConfig.vue';

import ServerHomePage from '@/server/ServerHomePage.vue';
import ServerPublicKeyManage from '@/server/PublicKeyManage.vue';
import ServerPasswordManage from '@/server/PasswordManage.vue';
import ServerMessageManage from '@/server/MessageManage.vue';


export default createRouter({
    history: createWebHistory(),
    routes:[
        {
            path:'/',
            name:'HomePage',
            component: HomePage,
            meta:{
                requiresAuth: false
            }
        },
        {
            path:'/Register',
            name:'Register',
            component: Register,
            meta:{
                requiresAuth: false
            }
        },
        {
            path:'/Login',
            name:'Login',
            component: Login,
            meta:{
                requiresAuth: false
            }
        },
        {
            path:'/ClientHomePage',
            name:'ClientHomePage',
            component: ClientHomePage,
            meta:{
                requiresAuth: false
            },
            children:[
                {
                    path:'PasswordManage',
                    name:'ClientPasswordManage',
                    component:ClientPasswordManage,
                },
                {
                    path:'AnnouncementManage',
                    name:'ClientAnnouncementManage',
                    component:ClientAnnouncementManage,
                },
                {
                    path:'Config',
                    name:'ClientConfig',
                    component:ClientConfig,
                },
            ]
        },
        // {
        //     path:'/ClientHomePage/PasswordManage',
        //     name:'ClientPasswordManage',
        //     component: ClientPasswordManage,
        //     meta:{
        //         requiresAuth: false
        //     }
        // },
        // {
        //     path:'/ClientHomePage/AnnouncementManage',
        //     name:'ClientAnnouncementManage',
        //     component: ClientAnnouncementManage,
        //     meta:{
        //         requiresAuth: false
        //     }
        // },
        // {
        //     path:'/ClientHomePage/Config',
        //     name:'ClientConfig',
        //     component: ClientConfig,
        //     meta:{
        //         requiresAuth: false
        //     }
        // },
        {
            path:'/ServerHomePage',
            name:'ServerHomePage',
            component: ServerHomePage,
            meta:{
                requiresAuth: false
            },
            children:[
                {
                    path:'PublicKeyManage',
                    name:'ServerPublicKeyManage',
                    component:ServerPublicKeyManage,
                },
                {
                    path:'PasswordManage',
                    name:'ServerPasswordManage',
                    component:ServerPasswordManage,
                },
                {
                    path:'MessageManage',
                    name:'ServerMessageManage',
                    component:ServerMessageManage,
                },
            ]
        },
    ],
})