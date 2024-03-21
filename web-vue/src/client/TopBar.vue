<template>
    <div class="bar">
        <img class="icon" src="@/assets/images/logo.svg" alt="MessageSave logo"/>
        <div class="flex-container">
            <span>MessageSave</span>
        </div>

        <div class="flex-grow" />

        <el-dropdown>
            <span class="el-dropdown-link">
                {{username}}
                <el-icon class="icon"><arrow-down /></el-icon>
            </span>
            <template #dropdown>
                <el-dropdown-menu>
                    <el-dropdown-item>config</el-dropdown-item>
                    <el-dropdown-item>logout</el-dropdown-item>
                </el-dropdown-menu>
            </template>
        </el-dropdown>

        <el-icon class="icon"><Bell /></el-icon>
    </div>
</template>

<script>

import { mapState } from 'vuex';
import { ArrowDown,Bell } from '@element-plus/icons-vue';

export default ({
    name: "TopBar",
    computed: {
        ...mapState(['username'])
    },
    components:{
        ArrowDown,Bell
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
.bar{
    display: flex;
    align-items: center;
}

.icon{
    width:50px;
    height: 50px;
}

.flex-container {  
  display: flex;  
  justify-content: center; /* 水平居中 */  
  align-items: center; /* 垂直居中 */  
  /* height: 100px; 你需要的任何高度   */
}

.flex-grow{
    flex-grow: 1;
}
</style>
