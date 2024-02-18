import { createRouter,createWebHistory,createWebHashHistory } from 'vue-router'
import HomePage from '@/components/HomePage.vue';
import PasswordManage from '@/components/PasswordManage.vue';
import MessageManage from '@/components/MessageManage.vue';

export default createRouter({
    history: createWebHistory(),
    routes:[
        {
            path:'/',
            name:'HomePage',
            component: HomePage,
        },
        {
            path:'/PasswordManage',
            name:'PasswordManage',
            component: PasswordManage,
        },
        {
            path:'/MessageManage',
            name:'MessageManage',
            component: MessageManage,
        },
    ],
})