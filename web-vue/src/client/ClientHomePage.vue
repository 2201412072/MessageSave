<template>
  <div class="common-layout">
    <el-container class="container">
      <el-header class="header">
        <TopBar />
      </el-header>
      <el-container>
        <el-aside class="aside">
          <el-scrollbar>
            <el-menu class="menu" router :default-active="$route.path">
              <MenuTree :menuList="menuList"/>
            </el-menu>
          </el-scrollbar>
        </el-aside>
        <el-main class="main">
          <!-- 路由占位符 -->
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>


<script>
import TopBar from './TopBar.vue';
import MenuTree from './MenuTree.vue';
import AddPasswordUI from './components/AddPasswordUI.vue';
import SearchPasswordUI from './components/SearchPasswordUI.vue';
export default ({
  name: 'ClientHomePage',
  data(){
    return {
      menuList:[
        {
          id:1,
          name:'PasswordManage',
          url:"/ClientHomePage/PasswordManage",
        },
        {
          id:2,
          name:'MessageManage',
          url:"/ClientHomePage/AnnouncementManage",
        },
        {
          id:3,
          name:'Config',
          url:"/ClientHomePage/Config",
        },
        {
          id:4,
          name:'PublicKeyManage',
          url:"/ClientHomePage/PublicKeyManage",
        },
      ]
    }
  },
  components: {
    TopBar,MenuTree,AddPasswordUI,SearchPasswordUI,
  },
  methods:{
    tempfunction(){}
  }
});
</script>

<style scoped>


.container{
  height: 100vh;
}

.aside{
  background: #797979;
  width:200px;
}

.header{
  background: #a1e9d2;
  height:50px;
  display: block;
  align-content: center;
}

</style>
