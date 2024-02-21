import { createApp } from 'vue';
import App from './App.vue';
import Router from '@/router/router.js'

import '@/assets/css/common.css'

import ElementPlus from 'element-plus' //全局引入
import 'element-plus/dist/index.css'


createApp(App).use(ElementPlus).use(Router).mount('#app');

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