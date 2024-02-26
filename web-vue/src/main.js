import { createApp, render } from 'vue';
import App from './App.vue';
import Router from '@/router.js'
import '@/assets/css/common.css'
import store from '@/store.js'

import ElementPlus from 'element-plus' //全局引入
import 'element-plus/dist/index.css'

Router.beforeEach((to, from, next) => {
    const requiresAuth = to.meta.requiresAuth; // 检查路由是否需要验证
  
    if (requiresAuth && !isLoggedIn()) {
      // 如果需要验证且用户未登录，则不跳转
      next(false);
    } else {
      next();
    }
  });

createApp(App).use(ElementPlus).use(Router).use(store).mount('#app');

// createApp({
//   data() {
//     return {
//     };
//   },
// }).use(ElementPlus).use(Router).mount('#app')

// new Vue({
//   el:'#App',
//   components:{App},
//   template: '<App/>'
// })