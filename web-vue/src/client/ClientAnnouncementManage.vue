<template>
    <div>
        <el-container> 
             <el-header class="header">
                <TopBar />
            </el-header>
            <el-main class="main">
                <!-- 查询 -->
                <h2> Annonuncement Manage </h2>
                <h3> Add-Password Annonuncement </h3>
                    <div class="app-user-search">
                        <el-form :model="myform_add" :inline="true">
                            <el-form-item label="Application">
                                <el-input v-model="myform_add.app" placeholder="Application"/>
                            </el-form-item>
                            <el-form-item label="Connect User">
                                <el-input v-model="myform_add.connect_user" placeholder="Connect User"/>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" @click="AddSearch">Search</el-button>
                                <el-button type="primary" @click="AddSearchAll">SearchAll</el-button>
                                <el-button @click="AddReset">Reset</el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                    <!-- 数据表 -->
                    <div class="app-user-table">
                        <el-table :data="tableData" style="width:100%">
                            <el-table-column label='Index' width="180">
                                <template v-slot="scope"> {{scope.$index}}</template>
                            </el-table-column>
                            <el-table-column label='Connect User' width="180">
                                <template v-slot="scope"> {{scope.row.connect_user}}</template>
                            </el-table-column>  
                            <el-table-column label='Application' width="180">
                                <template v-slot="scope"> {{scope.row.app}}</template>
                            </el-table-column>    
                            <el-table-column label="tip">
                                <span>进行了加密</span>
                            </el-table-column>
                            <el-table-column label="Operate">
                                <template v-slot="scope">
                                    <el-button type="danger" @click="handleAnnounce(scope.$index,scope.row.connect_user,scope.row.app)">receive</el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                <h3> Request table </h3>
                    <div class="request-search">
                        <el-form :model="myform_request" :inline="true">
                            <el-form-item label="Application">
                                <el-input v-model="myform_request.app" placeholder="Application"/>
                            </el-form-item>
                            <el-form-item label="Connect User">
                                <el-input v-model="myform_request.cosnnect_user" placeholder="Connect User"/>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" @click="RequestSearch">Search</el-button>
                                <el-button type="primary" @click="RequestSearchAll">SearchAll</el-button>
                                <el-button @click="RequestReset">Reset</el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                    <!-- 数据表 -->
                    <div class="request-table">
                        <el-table :data="tableData" style="width:100%">
                            <el-table-column label='Index' width="180">
                                <template v-slot="scope"> {{scope.$index}}</template>
                            </el-table-column>
                            <el-table-column label='Connect User' width="180">
                                <template v-slot="scope"> {{scope.row.connect_user}}</template>
                            </el-table-column>  
                            <el-table-column label='Application' width="180">
                                <template v-slot="scope"> {{scope.row.app}}</template>
                            </el-table-column>
                            <el-table-column label="tip">
                                <span>发送了解密请求</span>
                            </el-table-column>    
                            <el-table-column label="agree">
                                <template v-slot="scope">
                                    <el-button type="danger" @click="handleAgree(scope.$index,scope.row.connect_user,scope.row.app)">agree</el-button>
                                </template>
                            </el-table-column>
                            <el-table-column label="disagree">
                                <template v-slot="scope">
                                    <el-button type="danger" @click="handleDisagree(scope.$index,scope.row.connect_user,scope.row.app)">disagree</el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
            </el-main> 
        </el-container>
    </div>
</template>

<script>
import TopBar from './TopBar.vue';
const axios = require('axios');

export default{
    data(){
        return{
            tableData:[],
            tableData_request:[],
            myform_add:{
                app:'',
                connect_user: '',
            },
            myform_request:{
                app:'',
                connect_user:'',
            },
        };
    },
    components: {
    TopBar,
    },
    mounted(){
        this.AddSearchAll(),
        this.RequestSearchAll()
    },
    methods:{
        AddSearchAll(){
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Add-Search",{
                "SrcUser":"",
                "Keyword":"",
            })
            .then(response=>{
                this.tableData=response.data;
                console.log('get password data.',this.tableData);
            },)
            .catch(error=>{
                console.log("error",error);
                alert("请求失败");
            })
        },

        AddSearch(){
            var myform = this.myform_add
            if(myform.app==""&&myform.connect_user=="")
            {
                self.AddSearchAll()
                return 
            }
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Add-Search",{
                "SrcUser": myform.connect_user,
                "Keyword":myform.app,
            })
            .then(response => {
                    console.log('get password data.',response.data);
                    this.tableData = response.data;
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },

        AddReset(){
            this.myform_add.connect_user = '';
            this.myform_add.app = '';
        },
        handleAnnounce(index,temp1,temp2){
            this.tableData.splice(index, 1);
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Add-Delete",{
                "SrcUser": temp1,
                "KeyWord":temp2,
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

        RequestSearchAll(){
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Request-Search",{
                "SrcUser":"",
                "KeyWord":"",
            })
            .then(response=>{
                this.tableData_request=response.data;
                console.log('get result data.',this.tableData_request);
            },)
            .catch(error=>{
                console.log("error",error);
                alert("请求失败");
            })
        },

        RequestSearch(){
            var myform = this.myform_request
            if(myform.app==""&&myform.connect_user=="")
            {
                self.PasswordSearchAll()
                return 
            }
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Request-Search",{
                "SrcUser": myform.connect_user,
                "KeyWord":myform.app,
            })
            .then(response => {
                    console.log('get result data.',response.data);
                    this.tableData_request = response.data;
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },

        RequestReset(){
            this.myform_request.connect_user = '';
            this.myform_request.app = '';
        },
        handleAgree(index,temp1,temp2){
            this.tableData_request.splice(index, 1);
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Request-Agree",{
                "SrcUser": temp1,
                "KeyWord":temp2,
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
        handleDisagree(){
            this.tableData_request.splice(index, 1);
            axios.post("http://localhost:8090/ClientHomePage/AnnouncementManage/Request-Disagree",{
                "SrcUser": temp1,
                "KeyWord":temp2,
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