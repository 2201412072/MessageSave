<template>
  <div class="common-layout">
    <!-- <el-container> -->
      <!-- <el-header class="header"><TopBar /></el-header> -->
      <!-- <el-header class="header">
        <TopBar />
      </el-header>
      <el-main class="main"><router-view></router-view></el-main>  -->
      <router-view></router-view>
    <!-- </el-container> -->
  </div>
  <!-- <HomePage /> -->
</template>

<script>



export default {
  name: 'App',
  components: {
  },
  data() {
    return {
      message: 'MessageSave',
    };
  },
  methods:{
  },
};
</script>

<style scoped>
/* .header{
  background: #a1e9d2;
  height:80px;
} */

.main{

}
</style>