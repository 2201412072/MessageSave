<template>
    <div>
        <el-container> 
             <el-header class="header">
                <TopBar />
            </el-header>
            <el-main class="main">
              <span class="span-tip" style="margin-bottom: 20px; font-size:24px;">Login</span>
              <div id="input-div" class="container">
                <div class="input-container">
                  <span class="span-tip">username</span>
                  <el-input class="el-input" type="text" v-model="username" placeholder="input the username."></el-input>
                  <span class="span-tip">password</span>
                  <el-input class="el-input" type="text" v-model="password" placeholder="input the password."></el-input>
                </div>
                <el-button class="button" @click="login">Login</el-button>
              </div>
            </el-main> 
        </el-container>
    </div>
</template>

<script>
import TopBar from './TopBar.vue';
import axios from 'axios';
import Router from '@/router.js'
import { mapMutations } from 'vuex';

export default ({
  name: 'Login',
  components: {
    TopBar,
  },
  data() {
    return {
      username:"",
      password:"",
      whatpage:"",
    };
  },
  methods:{
    ...mapMutations(['setMyParameter']),
    jumppage(){
      //this.$store.commit('setMyParameter', 'example');
      this.setMyParameter(this.username)
      if(this.whatpage=="client"){
        Router.push({name:"ClientHomePage"})
      }
      else if(this.whatpage=="server"){
        Router.push({name:"ServerHomePage"})
      }
    },
    login(){
      console.log("login");
        axios.post("http://localhost:8090/Login",{
            "username":this.username,
            'password':this.password,
        })
        .then(response => {
                console.log('login over.',response);
                this.password = "";
                if(response.data.result==1){
                    //登录成功
                    console.log('login succeed.',response);
                    this.whatpage=response.data.whatpage
                    this.jumppage()
                }
                else{
                    alert("用户或密码错误")
                }
            },
        )
        .catch(response => {
            console.log("error",response);
            alert("login failed.");
        }) 
      
    }
  }
});
</script>

<style scoped>

</style>
