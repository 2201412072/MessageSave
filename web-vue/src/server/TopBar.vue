<template>
    <!-- <el-menu class="menu" :default-active="activeIndex" mode="horizontal" :elipsis="false" @select="handleMenuSelect"> -->
    <el-menu class="menu" router :default-active="$route.path" mode="horizontal" :ellipsis="false">
        <el-menu-item index="/ServerHomePage">
            <img style="width:50px" src="@/assets/images/logo.svg" alt="MessageSave logo"/>
        </el-menu-item>
        <div class="flex-grow" />
        <el-menu-item index="/ServerHomePage">
            <span>HomePage</span>
        </el-menu-item>
        <el-menu-item index="/ServerHomePage/PublicKeyManage">
            <span>PublicKeyManage</span>
        </el-menu-item>
        <el-menu-item index="/ServerHomePage/PasswordManage">
            <span>UserPasswordManage</span>
        </el-menu-item>
         <el-menu-item index="/ServerHomePage/MessageManage">
            <span>MessageManage</span>
        </el-menu-item>
        <el-menu-item @click="NothandleMenuSelect">
            {{ username}} 
        </el-menu-item>
    </el-menu>
</template>

<script>

import { mapState } from 'vuex';


export default ({
    name: "TopBar",
    computed: {
        ...mapState(['username'])
    },
    data(){
        return {
            // activeIndex:'1',
        };
    },
    methods:{
        
        NothandleMenuSelect(){ // 不处理导航栏的点击事件
        }
    }
})
</script>

<style scoped>
.menu{
    background-color: transparent !important ;
    display: flex;
    align-items: center;
}

.flex-grow{
    flex-grow: 1;
}
</style>
