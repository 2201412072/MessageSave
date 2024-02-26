<template>
    <div>
        <el-container> 
             <el-header class="header">
                <TopBar />
            </el-header>
            <el-main class="main">
                <router-view></router-view>
            </el-main> 
        </el-container>
    </div>
</template>

<script>
import TopBar from './TopBar.vue';

export default ({
  name: 'ServerHomePage',
  components: {
    TopBar,
  },
});
</script>

<style scoped>

</style>
