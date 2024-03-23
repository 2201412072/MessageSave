<template>
    <div>
        <el-container> 
            <el-main class="main">
                <!-- 查询 -->
                <div class="search">
                    <el-form :model="myform" :inline="true">
                        <el-form-item label="Source User">
                            <el-input v-model="myform.src_user" placeholder="Source User"/>
                        </el-form-item>
                        <el-form-item label="Destination User">
                            <el-input v-model="myform.dst_user" placeholder="Destination User"/>
                        </el-form-item>
                        <el-form-item label="Application">
                            <el-input v-model="myform.keyword" placeholder="Application"/>
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
                        <el-table-column label='Application' width="180">
                            <template v-slot="scope"> {{scope.row.keyword}}</template>
                        </el-table-column>  
                        <el-table-column label='Source User' width="180">
                            <template v-slot="scope"> {{scope.row.src_user}}</template>
                        </el-table-column>  
                        <el-table-column label='Destination User' width="180">
                            <template v-slot="scope"> {{scope.row.dst_user}}</template>
                        </el-table-column> 
                        <el-table-column label='Type' width="180">
                            <template v-slot="scope"> {{scope.row.operator}}</template>
                        </el-table-column>     
                        <el-table-column label="Operate">
                            <template v-slot="scope">
                                <el-button type="danger" @click="handleDelete(scope.$index,scope.row.src_user,scope.row.dst_user,scope.row.keyword,scope.row.operator)">delete</el-button>
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
                src_user: '',
                dst_user:'',
                keyword:'',
                operator:'',
            }
        };
    },
    mounted(){
        this.SearchAll()
    },
    methods:{
        SearchAll(){
            axios.post("http://localhost:8090/ServerHomePage/MessageManage/Message-SearchAll",{
                "src_user": "",
                "dst_user": "",
                "keyword":"",
                "operator":"",
            })
            .then(response=>{
                this.tableData=response.data;
                console.log('get message data.',this.tableData);
            },)
            .catch(error=>{
                console.log("error",error);
                alert("请求失败");
            })
        },

        Search(){
            var myform = this.myform
            if(myform.src_user=="" && myform.dst_user=="")
            {
                this.SearchAll()
            }else{
                axios.post("http://localhost:8090/ServerHomePage/MessageManage/Message-Search",{
                    "src_user": myform.src_user,
                    "dst_user": myform.dst_user,
                    "keyword":myform.keyword,
                    "operator":myform.operator,
                })
                .then(response => {
                        console.log('get message data.',response.data);
                        this.tableData = response.data;
                    },
                )
                .catch(response => {
                    console.log("error",response);
                    alert("请求失败");
                })
            }
        },

        Reset(){
            this.myform.src_user = '';
            this.myform.dst_user = '';
        },
        handleDelete(index,temp1,temp2,temp3,temp4){
            this.tableData.splice(index, 1);
            axios.post("http://localhost:8090/ServerHomePage/MessageManage/Message-Delete",{
                "src_user": temp1,
                "dst_user":temp2,
                "keyword":temp3,
                "operator":temp4,
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
