<template>
    <div>
        <el-container> 
             <el-header class="header">
                <TopBar />
            </el-header>
            <el-main class="main">
                <!-- 查询 -->
                <h2> App-Password Manage </h2>
                <h3> App-Password Search </h3>
                    <div class="password-search">
                        <el-form :model="myform_password" :inline="true">
                            <el-form-item label="Application">
                                <el-input v-model="myform_password.app" placeholder="Application"/>
                            </el-form-item>
                            <el-form-item label="Connect User">
                                <el-input v-model="myform_password.connect_user" placeholder="Connect User"/>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" @click="PasswordSearch">Search</el-button>
                                <el-button type="primary" @click="PasswordSearchAll">SearchAll</el-button>
                                <el-button @click="PasswordReset">Reset</el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                    <!-- 数据表 -->
                    <div class="password-table">
                        <el-table :data="tableData" style="width:100%">
                            <el-table-column label='Index' width="180">
                                <template v-slot="scope"> {{scope.$index+1}}</template>
                            </el-table-column>
                            <el-table-column label='Connect User' width="180">
                                <template v-slot="scope"> {{scope.row.connect_user}}</template>
                            </el-table-column>  
                            <el-table-column label='Application' width="180">
                                <template v-slot="scope"> {{scope.row.app}}</template>
                            </el-table-column>    
                            <el-table-column label="Operate">
                                <template v-slot="scope">
                                    <el-button type="danger" @click="handleDelete_password(scope.$index,scope.row.connect_user,scope.row.app)">delete</el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                <AddPasswordUI page="ClientHomePage/PasswordManage" :myfunction="PasswordSearch"></AddPasswordUI>
                <SearchPasswordUI page="ClientHomePage/PasswordManage" :myfunction="ResultSearch"></SearchPasswordUI>
                <!-- <div></div> 下面代码，未完待续-->
                <h3> Password Search Result table </h3>
                    <div class="result-search">
                        <el-form :model="myform_result" :inline="true">
                            <el-form-item label="Application">
                                <el-input v-model="myform_result.app" placeholder="Application"/>
                            </el-form-item>
                            <el-form-item label="Connect User">
                                <el-input v-model="myform_result.cosnnect_user" placeholder="Connect User"/>
                            </el-form-item>
                            <el-form-item label="Stage"> 
                                <el-select v-model="myform_result.selectedOption" placeholder="select">
                                    <el-option
                                        v-for="option in stage_options"
                                        :key="option.value"
                                        :label="option.label"
                                        :value="option.value"
                                        :selected="option.value==1"
                                    ></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" @click="ResultSearch">Search</el-button>
                                <el-button type="primary" @click="ResultSearchAll">SearchAll</el-button>
                                <el-button @click="ResultReset">Reset</el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                    <!-- 数据表 -->
                    <div class="password-table">
                        <el-table :data="tableData_result" style="width:100%">
                            <el-table-column label='Index' width="180">
                                <template v-slot="scope"> {{scope.$index+1}}</template>
                            </el-table-column>
                            <el-table-column label='Connect User' width="180">
                                <template v-slot="scope"> {{scope.row.connect_user}}</template>
                            </el-table-column>  
                            <el-table-column label='Application' width="180">
                                <template v-slot="scope"> {{scope.row.app}}</template>
                            </el-table-column>    
                            <el-table-column label='Stage' width="180">
                                <template v-slot="scope"> {{scope.row.stage}}</template>
                            </el-table-column> 
                            <el-table-column label='Password' width="180">
                                <template v-slot="scope"> {{scope.row.password}}</template>
                            </el-table-column> 
                            <el-table-column label="Operate">
                                <template v-slot="scope">
                                    <el-button type="danger" @click="handleDelete_result(scope.$index,scope.row.connect_user,scope.row.app,scope.row.stage)">delete</el-button>
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
import AddPasswordUI from './components/AddPasswordUI.vue';
import SearchPasswordUI from './components/SearchPasswordUI.vue';
const axios = require('axios');

export default{
    data(){
        return{
            tableData:[],
            tableData_result:[],
            myform_password:{
                app:'',
                connect_user: '',
            },
            myform_result:{
                app:'',
                connect_user:'',
                selectedOption:'',
            },
            stage_options:[
                {label:"has completed",value:1},
                {label:"hasn't completed",value:2},
                {label:"all",value:3},
            ],
        };
    },
    components: {
    TopBar,AddPasswordUI,SearchPasswordUI
    },
    mounted(){
        this.PasswordSearchAll(),
        this.ResultSearchAll()
    },
    methods:{
        PasswordSearchAll(){
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Password-Search",{
                "Application":"",
                "Username":"",
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

        PasswordSearch(){
            var myform = this.myform_password
            if(myform.app==""&&myform.connect_user=="")
            {
                self.PasswordSearchAll()
                return 
            }
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Password-Search",{
                "Username": myform.connect_user,
                "Application":myform.app,
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

        PasswordReset(){
            this.myform_password.connect_user = '';
            this.myform_password.app = '';
        },
        handleDelete_password(index,temp1,temp2){
            console.log('delete user:',temp1,' app:',temp2)
            this.tableData.splice(index, 1);
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Password-Delete",{
                "connect_user": temp1,
                "app":temp2,
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

        ResultSearchAll(){
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Result-Search",{
                "Username": "",
                "Application":"",
                "stage":"",
            })
            .then(response=>{
                this.tableData_result=response.data;
                console.log('get result data.',this.tableData_result);
            },)
            .catch(error=>{
                console.log("error",error);
                alert("请求失败");
            })
        },

        ResultSearch(){
            var myform = this.myform_result
            if(myform.app==""&&myform.connect_user=="")
            {
                self.PasswordSearchAll()
                return 
            }
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Result-Search",{
                "Username": myform.connect_user,
                "Application":myform.app,
                "stage":myform.selectedOption,
            })
            .then(response => {
                    console.log('get result data.',response.data);
                    this.tableData_result = response.data;
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },

        ResultReset(){
            this.myform_result.Username = '';
            this.myform_result.app = '';
            this.myform_result.selectedOption = 1;
        },
        handleDelete_result(index,temp1,temp2,temp_stage){
            this.tableData_result.splice(index, 1);
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Result-Delete",{
                "Username": temp1,
                "Application":temp2,
                "stage":temp_stage,
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