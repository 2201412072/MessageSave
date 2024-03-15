<template>
    <div>
        <el-container> 
             <el-header class="header">
                <TopBar />
            </el-header>
            <el-main class="main">
                <div id="main-function">
                    <AddPasswordUI page="ClientHomePage" :myfunction="tempfunction"></AddPasswordUI>
                    <SearchPasswordUI page="ClientHomePage" :myfunction="tempfunction"></SearchPasswordUI>
                </div>
            </el-main> 
        </el-container>
    </div>
</template>

<script>
import TopBar from './TopBar.vue';
import AddPasswordUI from './components/AddPasswordUI.vue';
import SearchPasswordUI from './components/SearchPasswordUI.vue';

export default ({
  name: 'ClientHomePage',
  components: {
    TopBar,AddPasswordUI,SearchPasswordUI,
  },
  methods:{
    tempfunction(){}
  }
});
</script>

<style scoped>

</style>
