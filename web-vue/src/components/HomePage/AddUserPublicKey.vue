<template>
    <div id="AddUserPublicKey" title="Add the user public key." class="container">
        <span class="span-tip">Add User Public Key</span>
        <div class="input-container">
            <span class="span-tip">User</span>
            <el-input class="input" type="text" v-model="User" placeholder="please input the user."></el-input>
            <span class="span-tip">Public Key</span>
            <el-input class="input" type="text" v-model="Public_key" placeholder="please input the public key."></el-input>
            <el-button @click="add">Add</el-button>
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
  margin-top: 20px;
}

.input-container {
  display: flex;
}
</style>
