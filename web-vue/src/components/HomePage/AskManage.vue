<template>
    <div id="AskManage">
        <span class="span-tip" style="margin-bottom: 20px; font-size:24px;">Search</span>
        <div id="input-div" class="container">
          <div class="input-container">
            <span class="span-tip">Key Word</span>
            <el-input class="el-input" type="text" v-model="key_word" placeholder="input the key word."></el-input>
            <span class="span-tip">Connect User</span>
            <el-input class="el-input" type="text" v-model="connect_user" placeholder="input the connect user."></el-input>
          </div>
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
            <span class="span-tip" style="margin-right: 10px">Encrypted Password</span>
            <el-input class="input" type="text" v-model="encrypted_password" placeholder="input the encrypted password." style="margin-right: 10px;"></el-input>
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
            <span class="span-tip" style="margin-right: 10px">Other User's Encrypted Password</span>
            <el-input class="input" type="text" v-model="other_password" placeholder="input the password other user sended to you." style="margin-right: 10px"></el-input>
            <el-button class="button" @click="decrypt_other">Decrypt</el-button> 
          </div>
          <div class="scrollable-container">
            <el-scrollbar class="scrollbar-wrapper">
              <div class="scrollable-content">
              {{ other_decrypted_password }}
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
      other_decrypted_password: '',
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
                this.origin_password = response.data.passwd;
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

      // this.origin_password = '';
    },
    decrypt(){
      axios.post("http://localhost:8090/DecryptPassword",{
        "encrypted_password":this.encrypted_password,
      })
      .then(reponse =>{
        console.log('decrypted successfully.');
        this.decrypted_password = reponse.data.passwd;
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
        this.other_decrypted_password = reponse.data.passwd;
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
.container{
  width: 600 px;
  display: flex;
  align-items: center;
  margin-top: 15px;
}

.input-container{
    display: flex;
}

.text-container{
  align-items: center;
  width: 100%;
}

.container-tips{
  display: flex;
  align-items: center;
  width: 100%;
  margin-bottom: 5px;
}

.scrollable-container {
  flex: 1;
  height: 100px; /* 设置容器的高度 */
  overflow: auto; /* 添加滚动条 */
  border: 1px solid;
  margin-bottom: 10px;
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
  width:70px;
}
</style>
