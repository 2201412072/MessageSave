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
          icon:'Key',
          url:"/ClientHomePage/PasswordManage",
        },
        {
          id:2,
          name:'MessageManage',
          icon:'Service',
          url:"/ClientHomePage/AnnouncementManage",
        },
        {
          id:3,
          name:'PublicKeyManage',
          icon:'Box',
          url:"/ClientHomePage/PublicKeyManage",
        },
        {
          id:4,
          name:'Config',
          icon:'Setting',
          url:"/ClientHomePage/Config",
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
  background: #32a4fb;
  width:300px;
}

.header{
  background: 'white';
  height:50px;
  display: block;
  align-content: center;
}

.main{
  background: rgb(241, 244, 246);
}

.el-main{
  width: 100%;
  height: 100%;
  text-align: center;
  background: #efefef;
  padding: 0px !important;
}

.el-icon{
    font-size:25px;
}
</style>
