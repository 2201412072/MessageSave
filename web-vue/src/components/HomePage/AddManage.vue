<template>
  <div id="AddManage" title="Add the passwd you want to save." class="container">
    <span>Add</span>
    <div class="input-container">
      <el-input class="input" type="text" v-model="key_word" placeholder="please input the key word."></el-input>
      <el-input class="input" type="text" v-model="passwd" placeholder="please input the passwd."></el-input>
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
        
        axios.post("http://localhost:8090/AddPassword")
        .then(response => {
                console.log('added password data successfully.');
                alert('added successfully.')
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
