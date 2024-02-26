<template>
    <div class="search">
        <h3> Password request </h3>
        <el-form :model="myform" :inline="true">
            <el-form-item label="Application">
                <el-input v-model="myform.app" placeholder="Application"/>
            </el-form-item>
            <el-form-item label="Connect User">
                <el-input v-model="myform.connect_user" placeholder="Connect User"/>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="Search">Search</el-button>
                <el-button @click="Reset">Reset</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>

const axios = require('axios');
export default{
    props:['page','myfunction'],
    data(){
        return{
            myform:{
                app:'',
                connect_user: '',
            }
        };
    },
    methods:{
        Search(){
            var myform = this.myform
            axios.post("http://localhost:8090/"+this.page+"/SearchPassword",{
                "connect_user": myform.connect_user,
                "app":myform.app,
            })
            .then(response => {
                    console.log('ask password ',response.data);
                    inputValue="已经向 "+myform.connect_user+" 请求"+myform.app+"密码"
                    // MessageBox.alert(inputValue, '提示', {
                    //     confirmButtonText: '确定',
                    //     type: 'info'
                    // });
                    alert(inputValue)
                    myfunction()
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },

        Reset(){
            this.myform.connect_user = '';
            this.myform.app='';
        },
    }
}
</script>