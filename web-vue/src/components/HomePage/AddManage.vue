<template>
  <div id="AddManage" title="Add the passwd you want to save." class="container">
    <span class="span-tip">Add</span>
    <div class="input-container">
      <span class="span-tip">Key Word</span>
      <el-input class="el-input" type="text" v-model="key_word" placeholder="please input the key word."></el-input>
      <span class="span-tip">Password</span>
      <el-input class="input" type="text" v-model="passwd" placeholder="please input the passwd."></el-input>
      <span class="span-tip">User</span>
      <el-input class="input" type="text" v-model="user" placeholder="please input the user."></el-input>
      <el-button class="input-button" @click="add">Add</el-button>
    </div>
  </div>
</template>

<script>
const axios = require('axios');

export default ({
  name: 'AddManage',
  components: {
  },
  data(){
    return{
        key_word: '',
        passwd: '',
        user: '',   
    };
  },
  methods: {
    add(){
        console.log("Add:",this.key_word,this.passwd,this.user);
        
          // axios.post("http://localhost:8090/AddPassword")
        axios.post("http://localhost:8090/AddPassword",{
          "key_word":this.key_word,
          "passwd":this.passwd,
          "user":this.user,
        })
        .then(response => {
                console.log('added password data successfully.');
                this.key_word = '';
                this.passwd = '';
                this.user = '';
            },
        )
        .catch(response => {
            console.log("error",response);
            alert("add failed.");
        })
    },
  }
});
</script>

<style scoped>
.container{
  margin-top: 20px;
}

.input-container {
  display: flex;
}
</style>
