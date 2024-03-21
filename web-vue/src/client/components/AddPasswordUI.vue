<template>
    <div class="search">
        <el-dialog :model-value="AddWindowVisible" :before-close="Cancel" title="Add password" width="500">
            <!-- <h3> Add Password </h3> -->
            <el-form :model="myform" :inline="true">
                <el-form-item label="Application" >
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
                    <el-button @click="Cancel">Cancel</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>



<script>
// import { MessageBox } from 'element-ui';
import { ref, watch } from 'vue'

const axios = require('axios');

export default{
    // props:['page','myfunction'],
    props:{
        'page':{
            type:String,
            default:"wrong"
        },
        'myfunction':{
            type:Function,
            required:false
        },
        AddWindowVisible:{
            type:Boolean,
        },
        Application:{
            type:String,
            required:false
        }
    },
    components:{
        //MessageBox
    },
    data(){
        return{
            myform:{
                app:'',//this.Application
                connect_user: '',
                password:'',
                confirm_password:'',
            },
            is_show_password:false,
            is_show_confirm_password:false,
            // visible:false,
        }
    },
    methods:{
        Cancel(){
            // 将对话框数据清空
            this.Reset();
            // 关闭对话框，即将对话框可见性赋为false，传给父组件
            this.$emit('update:AddWindowVisible',false);     
        },
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
                    console.log('add password for '+myform.app,response.data);
                    let inputValue="为 "+myform.app+"添加密码成功";
                    // MessageBox.alert(inputValue, '提示', {
                    //     confirmButtonText: '确定',
                    //     type: 'info'
                    // });
                    alert(inputValue);
                    this.myfunction();
                    this.Reset();
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
            this.Cancel();
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