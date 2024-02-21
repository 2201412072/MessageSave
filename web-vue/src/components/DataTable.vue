<template>
    <div>
        <!-- 查询 -->
        <div class="search">
            <el-form :model="myform" :inline="true">
                <el-form-item label="Key Word">
                    <el-input v-model="myform.key_word" placeholder="Key Word"/>
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
        <!-- 数据表 -->
        <div class="table">
            <el-table :data="tableData" style="width:100%">
                <el-table-column label='Index' width="180">
                    <template v-slot="scope"> {{scope.$index}}</template>
                </el-table-column>
                <el-table-column label='Key Word' width="200">
                    <template v-slot="scope"> {{scope.row.key_word}}</template>
                </el-table-column>
                <el-table-column label='Connect User' width="180">
                    <template v-slot="scope"> {{scope.row.connect_user}}</template>
                </el-table-column>
                
                <!-- <el-table-column label='Encrypted Password' width="300">
                    <template v-slot="scope">
                        <el-input v-model="scope.row.encrypted_passwd" :disabled=true show-password></el-input>
                    </template>
                </el-table-column>
                <el-table-column label='Decrypted Password' width="300">
                    <template v-slot="scope">
                        <el-input v-show="scope.row.got" v-bind="scope.row.decrypted_passwd"  :disabled=true show-password></el-input>
                    </template>
                </el-table-column>
                <el-popover
                            :show="scope.row.got & scope.row.display"
                            trigger="click"
                            placement="top">
                        <p v-show="show">{{ scope.row.decrypted_passwd }}</p>
                        <i v-show="show" slot="suffix" class="el-input__icon el-icon-search"></i>
                        <el-input v-show="!show" placeholder="password" type="password"></el-input>
                        </el-popover> -->
                        
                <el-table-column label="Operate">
                    <template v-slot="scope">
                        <el-button type="danger" @click="handleDelete(scope.$index, scope.row)">delete</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>

<script>
const axios = require('axios');

export default{
    props:{
        page:String,
    },
    data(){
        return{
            tableData:[],
            myform:{
                key_word: '',
                connect_user: '',
            }
        };
    },
    mounted(){
        this.fetchdata()
    },
    methods:{
        fetchdata(){
            axios.get("http://localhost:8090/"+this.page)
            .then(response => {
                    this.tableData = response.data;
                    console.log('get password data.', this.tableData);
                },
            )
            .catch(error => {
                console.log("error",error);
                alert("请求失败");
            })
        },

        Search(){
            var myform = this.myform
            axios.post("http://localhost:8090/"+this.page+"/Search",{
                "key_word": myform.key_word,
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
            this.myform.key_word = '';
            this.myform.connect_user = '';
        },
    }
}
</script>
