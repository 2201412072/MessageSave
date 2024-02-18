<template>
    <el-menu class="menu" :default-active="activeIndex" mode="horizontal" @select="handleMenuSelect">
        <el-menu-item index="1">
            <router-link to="/">HomePage</router-link>
        </el-menu-item>
        <el-menu-item index="2">
            <router-link to="/PasswordManage">PasswordManage</router-link>
        </el-menu-item>
        <el-menu-item index="3">
            <router-link to="/MessageManage">MessageManage</router-link>
        </el-menu-item>
    </el-menu>
</template>

<script>

export default ({
    name: "TopBar",
    data(){
        return {
            activeIndex:'1',
        };
    },
    methods:{
        handleMenuSelect(index){ // 处理导航栏项点击事件
            this.activeIndex = index;
        }
    }
})
</script>

<style scoped>
.menu{
    background-color: transparent !important ;
}
</style>
