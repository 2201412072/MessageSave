<template>
    <div id="AskManage" title="Search the password app wanted.">
        <span>Search</span>
        <div id="input-div" class="input-container">
          <span>Key Word</span>
          <el-input class="input" type="text" v-model="key_word" placeholder="please input the key word." ></el-input>
          <span>Connect User</span>
          <el-input class="input" type="text" v-model="connect_user" placeholder="please input the connect user." ></el-input>
          <el-button class="button" @click="search">Search</el-button>
        </div>
        <div id="origin-show" class="text-container">
          <div class="container-tips">
            <span>Original Password</span>
            <el-button class="button" @click="copy">Copy</el-button> 
          </div>
          <div class="scrollable-container">
            <el-scrollbar class="scrollbar-wrapper">
              <div id="origin_passwd_div" class="scrollable-content">
              {{ origin_password }}
              </div>
            </el-scrollbar>
          </div>
        </div>

        <div id="passwd-show" class="text-container">
          <div class="container-tips">
            <span>Encrypted Password</span>
            <el-input class="input" type="text" v-model="encrypted_password" plcaehodler="please input the encrypted password."></el-input>
            <el-button class="button" @click="decrypt">Decrypt</el-button> 
          </div>
          <div class="scrollable-container">
            <el-scrollbar class="scrollbar-wrapper">
              <div class="scrollable-content">
              {{ decrypted_password }}
              </div>
            </el-scrollbar>
          </div>
        </div>
    </div>
</template>

<script>
const axios = require('axios');

export default ({
  name: 'AskManage',
  components: {
  },
  data() {
    return {
      key_word: '',
      connect_user: '',
      origin_password: '',
      encrypted_password: '',
      decrypted_password: '',
    };
  },
  methods: {
    search() {
      console.log(this.key_word);
      axios.post("http://localhost:8090/SearchPassword",{
        "key_word":this.key_word,
        'user':this.connect_user,
      })
        .then(response => {
                console.log('searched over.',response);
                this.textContent = response['passwd'];
                this.key_word = '';
                this.connect_user = '';
            },
        )
        .catch(response => {
            console.log("error",response);
            alert("search failed.");
        })
    },
    copy(){
      let text = document.getElementById('origin_passwd_div').innerText
      let InputElement = document.createElement("input")
      InputElement.value = text;
      document.body.appendChild(InputElement);
      InputElement.select(); // 选中文本
      document.execCommand("copy"); // 执行浏览器复制命令
      InputElement.remove(); // 移除该临时输入框
      console.log("copy over.", text);

      this.origin_password = '';
    },
    decrypt(){
      axios.post("http://localhost:8090/DecryptPassword",{
        "encrypted_password":this.encrypted_password,
      })
      .then(reponse =>{
        console.log('decrypted successfully.');
        this.decrypted_password = reponse['passwd'];
        this.encrypted_password = '';
      })
      .catch(response => {
          console.log("error",response);
          alert("search failed.");
      })
    }
  },
});
</script>

<style scoped>
.input-container{
  display: flex;
  align-items: center;
  width: 100%;
}

.text-container{
  /* display: flex; */
  align-items: center;
  width: 100%;
}

.container-tips{
  display: flex;
  align-items: center;
  width: 100%;
}

.scrollable-container {
  flex: 1;
  height: 200px; /* 设置容器的高度 */
  overflow: auto; /* 添加滚动条 */
  border: 1px solid;
}

.scrollbar-wrapper{
  height: 100%;
  overflow-x: hidden !important;
}

.scrollable-content {
  padding: 10px; /* 可根据需求设置内容的内边距 */
  color: #8a8a8a;
}

.button{
  margin-left: auto;
}
</style>
