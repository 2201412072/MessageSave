<template>
    <div class="search">
        <h3> Add Password </h3>
        <el-form :model="myform" :inline="true">
            <el-form-item label="Application">
                <el-input v-model="myform.app" placeholder="Application"/>
            </el-form-item>
            <el-form-item label="Connect User">
                <el-input v-model="myform.connect_user" placeholder="Connect User"/>
            </el-form-item>
            <el-form-item label="Password">
                <el-input v-model="myform.password" type="password" placeholder="Password"/>
            </el-form-item>
            <el-form-item label="Confirm Password">
                <el-input v-model="myform.confirm_password" type="password" placeholder="Confirm Password"/>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="Add">Add</el-button>
                <el-button @click="Reset">Reset</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>



<script>
// import { MessageBox } from 'element-ui';

const axios = require('axios');
export default{
    // props:['page','myfunction'],
    props:{
        'page':
        //page:
        {
            type:String,
            default:"wrong"
        },
        'myfunction':{
            type:Function,
            required:false
        }
    },
    components:{
        //MessageBox
    },
    data(){
        return{
            myform:{
                app:'',
                connect_user: '',
                password:'',
                confirm_password:'',
            },
            is_show_password:false,
            is_show_confirm_password:false,
        };
    },
    methods:{
        Add(){
            var myform = this.myform
            if(myform.password!=myform.confirm_password)
            {
                // MessageBox.alert("请确定密码输入时正确", '提示', {
                //         confirmButtonText: '确定',
                //         type: 'error'
                //     });
                alert("请确定密码输入时正确")
                    return
            }
            axios.post("http://localhost:8090/"+this.page+"/AddPassword",{
                "connect_user": myform.connect_user,
                "app":myform.app,
                "password":myform.password,
            })
            .then(response => {
                    console.log('add password ',response.data);
                    inputValue="为 "+myform.app+"添加密码成功"
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
            this.myform.password='';
            this.myform.confirm_password='';
        },
    }
}
</script>