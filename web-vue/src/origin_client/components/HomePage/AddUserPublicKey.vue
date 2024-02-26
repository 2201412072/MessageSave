<template>
    <div id="AddUserPublicKey" title="Add the user public key.">
        <span class="span-tip" style="font-size:24px;">Add User Public Key</span>
        <div class="container">
            <div class="input-container">
                <span class="span-tip">User</span>
                <el-input class="input" type="text" v-model="User" placeholder="input the user."></el-input>
                <span class="span-tip">Public Key</span>
                <el-input class="input" type="text" v-model="Public_key" placeholder="input the public key."></el-input>
            </div>
            <el-button class="button" @click="add">Add</el-button>
        </div>
    </div>
</template>

<script>
const axios = require("axios");

export default ({
    name: 'AddUserPublicKey',
    data(){
        return {
            User:"",
            Public_key:"",
        }
    },
    methods:{
        add(){
            axios.post("http://localhost:8090/AddUserPublicKey",{
                'Username':this.User,
                'Public_key':this.Public_key,
            })
            .then(response => {
                console.log('added user public key successfully.');
                this.User = '';
                this.Public_key = '';
            },
        )
        .catch(error => {
            console.log("error",error);
            alert("add failed.");
        })
        }
    }
});
</script>


<style scoped>
.container{
  width: 100%;
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
