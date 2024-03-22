<template>
    <div>
        <el-menu background-color="#32a4fb" text-color="white">
            <template v-for="item in menuList"> 
                <!-- 有次级菜单的,则展开子选项-->
                <el-sub-menu class="menu-item" v-if="item.children && item.children.length>0" :key="item.id" :index="item.url">
                    <!-- <template #title>
                        <el-icon v-if="item.icon!=''"><component :is="item.icon" /></el-icon>
                        <span>{{item.name}}</span>
                    </template> -->
                    <!-- 递归,实现无限级展开 -->
                    <el-icon v-if="item.icon!=''"><component :is="item.icon" /></el-icon>
                    <span class="menu-item-span">{{item.name}}</span>
                    <MenuTree :menuList="item.children"/>
                </el-sub-menu>
                <!-- 没有次级菜单的 -->
                <el-menu-item class="menu-item" v-if="!item.children" :key="item.id" :index="item.url">
                    <el-icon v-if="item.icon!=''"><component :is="item.icon" /></el-icon>
                    <span class="menu-item-span" @click="redirectToUrl(item.url)">{{item.name}}</span>
                </el-menu-item>
            </template>
        </el-menu>
    </div>
</template>

<script>
import { ArrowDown,Key,Service,Setting,Box } from '@element-plus/icons-vue';

export default ({
    name: 'MenuTree',
    props:{
        menuList:{
            type:Array,
            default(){
                return []
            },
        }
    },
    components: {
        Key,Service,Setting,Box,
    },
    methods:{
        tempfunction(){},
        redirectToUrl(url){
            window.location.href = url; 
        },
    },
})
</script>

<style scoped>
.el-menu{
    background-color: var(--el-menu-bg-color, transparent); /* 使用变量，并设置默认值 */ 
}

.menu-item{
    /* margin-bottom: 10px; */
    height: 70px;
}

.menu-item-span{
    font-size: 20px;
}
</style>