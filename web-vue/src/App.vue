<template>
  <div class="common-layout">
    <el-container>
      <!-- <el-header class="header"><TopBar /></el-header> -->
      <el-header class="header">
        <TopBar />
      </el-header>
      <el-main class="main"><router-view></router-view></el-main> 
    </el-container>
  </div>
  <!-- <HomePage /> -->
</template>

<script>
import TopBar from './components/TopBar.vue';


export default {
  name: 'App',
  components: {
    TopBar
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
.header{
  background: #a1e9d2;
  height:80px;
}
.main{

}
</style>