<template>
    <div>
        <el-container> 
             <el-header class="header">
                <TopBar />
            </el-header>
            <el-main class="main">
              <span class="span-tip" style="margin-bottom: 20px; font-size:24px;">Register</span>
              <div id="input-div" class="container">
                <div class="input-container">
                  <span class="span-tip">username</span>
                  <el-input class="el-input" type="text" v-model="username" placeholder="input the username."></el-input>
                  <span class="span-tip">password</span>
                  <el-input class="el-input" type="text" v-model="password" placeholder="input the password."></el-input>
                  <span class="span-tip">confirm password</span>
                  <el-input class="el-input" type="text" v-model="confirm_password" placeholder="input the password again."></el-input>
                </div>
                <el-button class="button" @click="register">Register</el-button>
              </div>
            </el-main> 
        </el-container>
    </div>
</template>

<script>
import TopBar from './TopBar.vue';
import axios from 'axios';
import { mapMutations } from 'vuex';

export default ({
  name: 'Register',
  components: {
    TopBar,
  },
  data() {
    return {
      username:"",
      password:"",
      confirm_password:"",
      whatpage:"",
    };
  },
  methods:{
    ...mapMutations(['setMyParameter']),
    jumppage(){
      this.setMyParameter(this.username)
      if(this.whatpage=="client"){
        Router.push({name:"ClientHomePage"})
      }
      else if(this.whatpage=="server"){
        Router.push({name:"ServerHomePage"})
      }
    },
    register(){
      console.log("register");
      if(this.password!=this.confirm_password){
        alert("请确保两次输入密码相同")
        this.confirm_password=""
      }
      else{
        axios.post("http://localhost:8090/Register",{
          "username":this.username,
          'password':this.password,
        })
        .then(response => {
                console.log('register over.',response);
                this.confirm_password = "";
                this.password = "";
                if(response.data.result==1){
                    //注册成功
                    console.log('register succeed.',response);
                    this.whatpage=response.data.whatpage
                    this.jumppage()
                }
                else{
                    alert("该用户已被注册")
                }
            },
        )
        .catch(response => {
            console.log("error",response);
            alert("register failed.");
        })
      }  
      
    }
  }
});
</script>

<style scoped>

</style>
