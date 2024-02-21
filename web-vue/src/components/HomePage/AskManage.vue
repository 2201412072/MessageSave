<template>
    <div id="AskManage">
        <span class="span-tip">Search</span>
        <div id="input-div" class="input-container">
          <span class="span-tip">Key Word</span>
          <el-input class="el-input" type="text" v-model="key_word" placeholder="please input the key word." style="width: 250px"></el-input>
          <span class="span-tip">Connect User</span>
          <el-input class="el-input" type="text" v-model="connect_user" placeholder="please input the connect user." style="width: 250px"></el-input>
          <el-button class="button" @click="search">Search</el-button>
        </div>
        <div id="origin-show" class="text-container">
          <div class="container-tips">
            <span class="span-tip">Original Password</span>
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
            <span class="span-tip">Encrypted Password</span>
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

        <div id="other-passwd-decrypt" class="text-container">
          <div class="container-tips">
            <span class="span-tip">Other User's Encrypted Password</span>
            <el-input class="input" type="text" v-model="other_password" plcaehodler="please input the password other user sended to you."></el-input>
            <el-button class="button" @click="decrypt_other">Decrypt</el-button> 
          </div>
          <div class="scrollable-container">
            <el-scrollbar class="scrollbar-wrapper">
              <div class="scrollable-content">
              {{ otehr_decrypted_password }}
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
      other_password: '',
      otehr_decrypted_password: '',
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
    },
    decrypt_other(){
      axios.post("http://localhost:8090/DecryptOtherPassword",{
        "encrypted_password":this.other_password,
      })
      .then(reponse =>{
        console.log('decrypted successfully.');
        this.other_decrypted_password = reponse['passwd'];
        this.other_password = '';
      })
      .catch(response => {
          console.log("error",response);
          alert("search failed.");
      })
    },
  },
});
</script>

<style scoped>
.input-container{
    display: flex;
    align-items: center;
    width: 500px;
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
  height: 100px; /* 设置容器的高度 */
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
