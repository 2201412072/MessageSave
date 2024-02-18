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
        console.log(key_word,passwd,user);
    }
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
