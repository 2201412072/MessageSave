<template>
    <div>
        <el-container> 
            <el-main class="main">
                <!-- 查询 -->
                <div class="search">
                    <el-form :model="myform" :inline="true">
                        <el-form-item label="Connect User">
                            <el-input v-model="myform.connect_user" placeholder="Connect User"/>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" @click="Search">Search</el-button>
                            <el-button type="primary" @click="SearchAll">SearchAll</el-button>
                            <el-button @click="Reset">Reset</el-button>
                        </el-form-item>
                    </el-form>
                </div>
                <!-- 数据表 -->
                <div class="table">
                    <el-table :data="tableData" style="width:100%">
                        <el-table-column label='Index' width="180">
                            <template v-slot="scope"> {{scope.$index}}</template>
                        </el-table-column>
                        <el-table-column label='Connect User' width="180">
                            <template v-slot="scope"> {{scope.row.connect_user}}</template>
                        </el-table-column>  
                        <el-table-column label='Public Key' width="180">
                            <template v-slot="scope"> {{scope.row.password}}</template>
                        </el-table-column>    
                        <el-table-column label="Operate">
                            <template v-slot="scope">
                                <el-button type="danger" @click="handleDelete(scope.$index,scope.row.connect_user)">delete</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </div>
            </el-main> 
        </el-container>
    </div>
</template>

<script>
const axios = require('axios');

export default{
    data(){
        return{
            tableData:[],
            myform:{
                connect_user: '',
            }
        };
    },
    mounted(){
        this.SearchAll()
    },
    methods:{
        SearchAll(){
            axios.get("http://localhost:8090/ServerHomePage/PasswordManage/Password-SearchAll")
            .then(response=>{
                this.tableData=response.data;
                console.log('get password data.',this.tableData);
            },)
            .catch(error=>{
                console.log("error",error);
                alert("请求失败");
            })
        },

        Search(){
            var myform = this.myform
            axios.post("http://localhost:8090/ServerHomePage/PasswordManage/Password-Search",{
                "connect_user": myform.connect_user,
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

        Reset(){
            this.myform.connect_user = '';
        },
        handleDelete(index,temp){
            this.tableData.splice(index, 1);
            axios.post("http://localhost:8090/ServerHomePage/PasswordManage/Password-Delete",{
                "connect_user": temp,
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
