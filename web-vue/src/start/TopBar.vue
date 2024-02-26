<template>
    <!-- <el-menu class="menu" :default-active="activeIndex" mode="horizontal" :elipsis="false" @select="handleMenuSelect"> -->
    <el-menu class="menu" router :default-active="$route.path" mode="horizontal" :ellipsis="false">
        <el-menu-item index="/">
            <img style="width:50px" src="@/assets/images/logo.svg" alt="MessageSave logo"/>
        </el-menu-item>
        <div class="flex-grow" />
        <el-menu-item index="/">
            <span>HomePage</span>
        </el-menu-item>
        <el-menu-item index="/Login">
            <span>Login</span>
        </el-menu-item>
        <el-menu-item index="/Register">
            <span>Register</span>
        </el-menu-item>
        <el-menu-item index="/">
            <img style="width:50px" src="@/assets/images/user.svg"/>
        </el-menu-item>
    </el-menu>
</template>

<script>

export default ({
    name: "TopBar",
    data(){
        return {
            // activeIndex:'1',
        };
    },
    methods:{
        handleMenuSelect(index){ // 处理导航栏项点击事件
            if (index==0) {
                this.activeIndex = 1;
            } else{
                this.activeIndex = index;
            }
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
