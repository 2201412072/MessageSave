<template>
  <div id="AddManage" title="Add the passwd you want to save.">
    <span class="span-tip" style="font-size:24px;">Add</span>
    <div class="container">
      <div class="input-container">
        <span class="span-tip">Key Word</span>
        <el-input class="el-input" type="text" v-model="key_word" placeholder="input the key word."></el-input>
        <span class="span-tip">Password</span>
        <el-input class="input" type="text" v-model="passwd" placeholder="input the passwd."></el-input>
        <span class="span-tip">User</span>
        <el-input class="input" type="text" v-model="user" placeholder="input the user."></el-input>
      </div>
      <el-button class="button" @click="add">Add</el-button>
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
          "Key_word":this.key_word,
          "Passwd":this.passwd,
          "User":this.user,
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
  width: 600 px;
  display: flex;
  align-items: center;
  margin-top: 10px;
  margin-bottom: 20px;
}

.input-container{
    display: flex;
}

.button{
  margin-left: auto;
  width:70px;
}
</style>
