<template>
    <div class="top-div">
        <!-- 查询 -->
        <el-form class="form" :model="myform" :inline="true">
            <el-form-item label="User">
                <el-input v-model="myform.user" placeholder="User"/>
            </el-form-item>
            <el-form-item>
                <el-icon @click="Search(myform.user)"><Search /></el-icon>
            </el-form-item>
        </el-form>
    </div>
    <div class="table">
        <el-table :data="tableData" style="width:100%">
            <el-table-column label="Index" width="180">
                <template v-slot="scope"> {{scope.$index+1}}</template>
            </el-table-column>
            <el-table-column label="User" width="180">
                <template v-slot="scope"> {{scope.row.connect_user}}</template>
            </el-table-column>
            <el-table-column label="PublicKey" width="180">
                <template v-slot="scope"> 
                    <span class="text-show">{{scope.row.public_key}}</span>
                </template>
            </el-table-column>
            <el-table-column label="Operate" width="180">
                <template v-slot="scope"> 
                    <el-button type="primary" @click="Copy(scope.row.public_key)">Copy</el-button>
                    <el-button type="primary" @click="Delete(scope.row.index,scope.row.connect_user)">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import axios from "axios";
import {Search} from '@element-plus/icons-vue';

export default({
    data(){
        return {
            tableData:[],
            myform:{
                'user':'',
            },
        };
    },
    components:{
        Search,
    },
    mounted(){
        this.Search('');
    },
    methods:{
        Search(user){
            // 查找指定人员公钥
            // if(user==''){
            //     params = {};
            // }
            // else{
            //     params = {connect_user:user};
            // }
            let params = {connect_user:user};
            axios.post("http://localhost:8090/ClientHomePage/Config/PublicKey-Search",params)
            .then(response => {
                this.tableData = response.data;
                console.log("searched "+ user+"'s publickey successfully.");
                this.myform.user="";
            })
            .catch(error =>{
                console.log("search failed, ",error);
                alert("请求失败");
            });
        },
        Copy(public_key){
            // 将公钥复制到剪切板上

        },
        Delete(index,user){
            // 删除公钥
            axios.post("http://localhost:8090/ClientHomePage/Config/PublicKey-Delete",{
                connect_user:user,
            })
            .then(repsonse=>{
                console.log("delete "+user+" successfully.");
                this.tableData.splice(index, 1);
            })
            .catch(error=>{
                console.log("delete failed, ",error);
                alert("delete failed.");
            });
        }
    }
})
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

.table{
    display: block;
}

.text-show{
    display: block;
    white-space: nowrap;
    max-width: 150px;
    overflow: hidden;
    text-overflow: ellipsis;
}
</style>