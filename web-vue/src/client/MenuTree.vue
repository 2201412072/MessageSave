<template>
    <div>
        <template v-for="item in menuList"> 
            <!-- 有次级菜单的,则展开子选项-->
            <el-sub-menu v-if="item.children && item.children.length>0" :key="item.id" :index="item.url">
                <!-- <template #title>
                    <el-icon v-if="item.icon!=''"><component :is="item.icon" /></el-icon>
                    <span>{{item.name}}</span>
                </template> -->
                <!-- 递归,实现无限级展开 -->
                <el-icon v-if="item.icon!=''"><component :is="item.icon" /></el-icon>
                <span>{{item.name}}</span>
                <MenuTree :menuList="item.children"/>
            </el-sub-menu>
            <!-- 没有次级菜单的 -->
            <el-menu-item v-if="!item.children" :key="item.id" :index="item.url">
                <el-icon v-if="item.icon!=''"><component :is="item.icon" /></el-icon>
                <span>{{item.name}}</span>
            </el-menu-item>
        </template>
    </div>
</template>

<script>
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
    },
    methods:{
        tempfunction(){}
    },
})
</script>

<style scoped>

</style>