<template>
    <div>
        <el-container> 
             <el-header class="header">
                <TopBar />
            </el-header>
            <el-main class="main">
                <div id="main-function">
                    <AddPasswordUI :page="ClientHomePage"></AddPasswordUI>
                    <SearchPasswordUI :page="ClientHomePage"></SearchPasswordUI>
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
});
</script>

<style scoped>

</style>
