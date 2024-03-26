<template>
    <div>
        <div class="top-div">
            <!-- 查询 -->
            <el-form class="form" :model="myform" :inline="true">
                <el-form-item class="el-form-item" label="SrcUser" >
                    <el-input v-model="myform.src_user" placeholder="SrcUser"/>
                </el-form-item>
                <el-form-item label="Application">
                    <el-input v-model="myform.app" placeholder="Application"/>
                </el-form-item>
                <el-form-item label="Operate">
                    <el-input v-model="myform.operate" placeholder="Operate"/>
                </el-form-item>
                <el-form-item>
                    <el-icon @click="MessageSearch(myform.src_user,myform.app,myform.operate)"><Search /></el-icon>
                </el-form-item>
            </el-form>  
        </div>
        <!-- 消息卡片 -->
        <template v-for="(row,i) in this.tableData">
            <MessageCard  :key="i" v-if="row.operate=='EncryptAnnocement2Client'" CardType="AskEncrypt" :text="row"
                :Content="'User '+row.connect_user+' send you an encrypt invitation for application '+row.app"
                :User="row.connect_user" :App="row.app"
                :HandleAgree="handleAgreeAnnoncement"/>   
            <MessageCard  :key="i" v-if="row.operate=='DecryptRequest2Client'" CardType="AskDecrypt" :text="row"
                :Content="'User '+row.connect_user+' send you a decrypt request for application '+row.app"
                :User="row.connect_user" :App="row.app"
                :HandleAgree="handleAgreeRequest" :HandleDisAgree="handleDisagreeRequest"/> 
        </template>
    </div>
</template>

<script>
import {Search} from '@element-plus/icons-vue';
import MessageCard from './components/MessageCard.vue';

const axios = require('axios');

export default{
    data(){
        return{
            // tableData:[],
            // tableData_request:[],
            // myform_add:{
            //     app:'',
            //     connect_user: '',
            // },
            // myform_request:{
            //     app:'',
            //     connect_user:'',
            // },
            tableData:[],
            myform:{
                src_user:"",
                app:"",
                operate:"",
            }
        };
    },
    components: {
        MessageCard,
        Search,
    },
    mounted(){
        // this.AddSearchAll(),
        // this.RequestSearchAll()
        this.MessageSearchAll()
    },
    methods:{
        MessageSearchAll(){
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Message-Search",{
                "connect_user":"",
                "app":"",
                "operate":"",
            })
            .then(response=>{
                this.tableData=response.data;
                // src_user  app  operate
                console.log('get password data.',this.tableData);
            })
            .catch(error=>{
                console.log("error",error);
                alert("请求失败");
            })
        },
        MessageSearch(src_user,app,operate){
            // 查找指定的消息
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Message-Search",{
                "connect_user":src_user,
                "app":app,
                "operate":operate,
            })
            .then(response=>{
                this.tableData = response.data;
                console.log("Search over, ",this.tableData);
                this.myform.src_user="";
                this.myform.app="";
                this.myform.operate="";
            })
            .catch(error=>{
                console.log("search failed, ",error);
                alert("Search fialed.");
            });
        },
        handleAgreeAnnoncement(index,src_user,key_word){
            this.tableData.splice(index, 1);
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Add-Delete",{
                "SrcUser":src_user,
                "KeyWord":key_word,
            })
            .then(response => {
                    console.log('delete.');
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },
        handleAgreeRequest(index,src_user,key_word){
            this.tableData.splice(index, 1);
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Request-Agree",{
                "SrcUser":src_user,
                "KeyWord":key_word,
            })
            .then(response => {
                    console.log('delete.');
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },
        handleDisagreeRequest(index,src_user,key_word){
            this.tableData.splice(index, 1);
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Request-Disagree",{
                "SrcUser":src_user,
                "KeyWord":key_word,
            })
            .then(response => {
                    console.log('delete.');
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },
    }
}
</script>

<style scoped>
.top-div{
    background: white;
    height: 50px;
    display: flex;
    margin-top: 5px;
    margin-bottom: 5px;
}

.form{
    display: flex;
    margin-left:5px;
}

.el-form-item{
    display: flex;
    align-items: center;
    align-content: center;
    margin-top: 0px;
    margin-bottom: 0px;
}
</style>