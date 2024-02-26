<template>
    <div>
        <el-container> 
             <el-header class="header">
                <TopBar />
            </el-header>
            <el-main class="main">
                <img style="width:50px" src="@/assets/images/user.svg"/>
            </el-main> 
        </el-container>
    </div>
</template>

<script>
import TopBar from './TopBar.vue';

export default ({
  name: 'HomePage',
  components: {
    TopBar,
  },
});
</script>

<style scoped>

</style>
