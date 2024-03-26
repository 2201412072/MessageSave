<template>
    <div>
        <div class="top-div">
            <!-- 添加密码&待用密码 -->
            <div class="add-password">
                <el-button type="primary" @click="AddWindowVisible=true">Add password</el-button>
                <el-button type="primary" @click="SearchPasswordUsing">Search password using</el-button>
            </div>
            <!-- 查询 -->
            <el-form class="form" :model="myform_password" :inline="true">
                <el-form-item class="el-form-item" label="Application">
                    <el-input v-model="myform_password.app" placeholder="Application"/>
                </el-form-item>
                <el-form-item label="Connect User">
                    <el-input v-model="myform_password.connect_user" placeholder="Connect User"/>
                </el-form-item>
                <el-form-item>
                    <el-icon @click="PasswordSearch"><Search /></el-icon>
                </el-form-item>
            </el-form>
        </div>
        <!-- 数据表 -->
        <div class="password-table">
            <el-table :data="tableData" style="width:100%">
                <el-table-column label='Index' width="180">
                    <template v-slot="scope"> {{scope.$index+1}}</template>
                </el-table-column>
                <el-table-column label='Application' width="180">
                    <template v-slot="scope"> {{scope.row.app}}</template>
                </el-table-column>  
                <el-table-column label='Password' width="180">
                    <template v-slot="scope"> {{GetValueByKey(scope.row.app)}}</template>
                </el-table-column>
                <el-table-column label='Connect User' width="600">
                    <template v-slot="scope"> 
                        <div class="users-div">
                            <div class="users-show-div" v-for="connect_user in scope.row.connect_user" :key="connect_user">
                                <i class="status-circle" style="background: red" 
                                    v-if="!stage_map.has(scope.row.app+'-'+scope.row.connect_user)"></i>
                                <i class="status-circle" style="background: yellow" 
                                    v-if="stage_map.has(scope.row.app+'-'+scope.row.connect_user) && stage_map.get(scope.row.app+'-'+scope.row.connect_user)=='hasn\'t completed'"></i>
                                <i class="status-circle" style="background: green" 
                                    v-if="stage_map.has(scope.row.app+'-'+scope.row.connect_user) && stage_map.get(scope.row.app+'-'+scope.row.connect_user)=='has complete'"></i>
                                <span class="user-span" @click="ToggleButton(scope.row,connect_user)">{{connect_user}}</span>
                                <el-button-group class="user-button-group">
                                    <el-button class="user-button" type="primary" v-show="scope.row.ButtonVisible[connect_user]" @click="UnlockPassword(connect_user,scope.row.app)" >
                                        <el-icon><Unlock /></el-icon>
                                    </el-button>
                                    <el-button class="user-button" type="primary" v-show="scope.row.ButtonVisible[connect_user]" @click="DeletePassword(scope.$index,connect_user,scope.row.app)" >
                                        <el-icon><Delete /></el-icon>
                                    </el-button>
                                </el-button-group>
                            </div>
                            <!-- :icon="Lock" :icon="Unlock" :icon="Delete" -->
                            <el-button class="user-button" type="primary" @click="LockPassword(scope.row.app)" >
                                <el-icon><Lock /></el-icon>
                            </el-button>
                            <el-icon class="addpasswd-icon" @click="OpenAddWinowWithAPP(scope.row.app)"><CirclePlus/></el-icon>
                        </div>
                    </template>
                </el-table-column>  
                    
                <!-- <el-table-column label="Operate">
                    <template v-slot="scope">
                        <el-button type="danger" @click="handleDelete_password(scope.$index,scope.row.connect_user,scope.row.app)">delete</el-button>
                    </template>
                </el-table-column> -->
            </el-table>
        </div>
    </div>

    <!-- 添加密码界面 -->
    <AddPasswordUI v-model:AddWindowVisible="AddWindowVisible" :Application="AddUIApp" page="ClientHomePage/PasswordManage" :myfunction="PasswordSearchAll"></AddPasswordUI>
</template>

<script>
import TopBar from './TopBar.vue';
import AddPasswordUI from './components/AddPasswordUI.vue';
import SearchPasswordUI from './components/SearchPasswordUI.vue';
import { Search,CirclePlus,Lock,Unlock,Delete,Edit } from '@element-plus/icons-vue';

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
            password_map:new Map(),
            stage_map:new Map(), // key:app+user value:stage (只记录解密的和正在解密的)
            myform_result:{
                app:'',
                connect_user:'',
                selectedOption:'',
            },
            AddWindowVisible: false,
            AddUIApp: '',
            stage_options:[
                {label:"has completed",value:1},
                {label:"hasn't completed",value:2},
                {label:"all",value:3},
            ],
            // stage_options:[
            //     {label:"unlocked",class:'status-unlocked'},
            //     {label:"unlockingd",class:'status-unlocking'},
            //     {label:"locked",class:'status-locked'},
            // ],
        };
    },
    components: {
        TopBar,AddPasswordUI,SearchPasswordUI, 
        Search,CirclePlus,Lock,Unlock,Delete,Edit,
    },
    created(){ // 对表格数据进行修饰，添加显示隐藏变量
        this.ChangeTableData();
    },
    mounted(){
        this.PasswordSearchAll()
        // this.ResultSearchAll()
    },
    methods:{
        InitButtonVisible(row){
            return row.connect_user.reduce((ButtonVisible,user)=>{
                ButtonVisible[user] = true; //false;
                return ButtonVisible;
            },{});
        },
        ChangeTableData(){
            this.tableData = this.tableData.map(row =>({
                ...row, // 复制原有的属性
                ButtonVisible : this.InitButtonVisible(row),
                // stage:3,
            }));
            console.log("tableData",this.tableData);
        },
        SearchPasswordUsing(){
            // 查找当前正在请求或者已经解锁的密码
            var new_table = [];
            for(let i=0;i<this.tableData.length;i++){
                var app=this.tableData[i].app;
                var connect_user=this.tableData[i].connect_user;
                var new_users = [];
                for(let j=0;j<connect_user.length;j++){
                    var value = connect_user[j];
                    var stage = this.stage_map.get(app+"-"+value);
                    if(stage=="has complete" || stage=="hasn't completed"){
                        new_users.push(value);
                    }
                }
                if(new_users.length>0){
                    new_table.push({"app":app,connect_user:new_users});
                }
            }
            this.tableData = new_table;
            this.ChangeTableData();
        },
        ToggleButton(row,connect_user){
            // console.log("togglbutton tableData",this.tableData);
            console.log("togglbutton row",row);
            row.ButtonVisible[connect_user] = !row.ButtonVisible[connect_user];
            // this.$set(row.ButtonVisible,connect_user,!row.ButtonVisible[connect_user]);
        },
        InitPasswordMap(){
            //初始化密码map，如果不在该map里，说明该应用没有收到解密的消息，或者压根没请求
            this.password_map.clear()
            this.stage_map.clear()
            // console.log(this.tableData_result)
            for(let i=0;i<this.tableData_result.length;i++){
                var app=this.tableData_result[i].app;
                var user=this.tableData_result[i].connect_user;
                var stage = this.tableData_result[i].stage;
                this.stage_map.set(app+'-'+user,stage);
                if(stage=="has complete")
                {
                    this.password_map.set(app, this.tableData_result[i].password);
                }
            }
            // console.log(this.password_map);
            console.log('stage map:',this.stage_map);
        },
        GetValueByKey(key){
                if (this.password_map.has(key)) {
                    return this.password_map.get(key);
                } else {
                    return "******";
                }
        },
        PasswordSearch(){
            var myform = this.myform_password
            if(myform.app==""&&myform.connect_user=="")
            {
                this.PasswordSearchAll()
                return 
            }
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Password-Search",{
                "connect_user": myform.connect_user,
                "app":myform.app,
            })
            .then(response => {
                    console.log('password search, APP:',myform.app,' Username:',myform.connect_user);
                    console.log('get password data.',response.data);
                    this.tableData = response.data;
                    this.ChangeTableData();
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })

            //获取密码
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Result-Search",{
                "connect_user": myform.connect_user,
                "app":myform.app,
            })
            .then(response => {
                    this.tableData_result = response.data;
                    this.InitPasswordMap();
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },
        PasswordSearchAll(){
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Password-Search",{
                "app":"",
                "connect_user":"",
            })
            .then(response=>{
                this.tableData=response.data;
                this.ChangeTableData();
            },)
            .catch(error=>{
                console.log("error",error);
                alert("请求失败");
            })

             //获取密码
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Result-Search",{
                "connect_user": "",
                "app":"",
            })
            .then(response => {
                    this.tableData_result = response.data;
                    this.InitPasswordMap();
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },
        LockPassword(app){
            // 将某密码复位，即重新回到加密状态
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Result-Delete",{
                "user_dst": "",
                "application":app,
            })
            .then(response => {
                console.log('delete result ans ',response.data);
            },)
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            });
            this.password_map.delete(app)
        },
        UnlockPassword(connect_user,app){
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/SearchPassword",{
                "connect_user": connect_user,
                "app":app,
            })
            .then(response => {
                console.log('ask password ',response.data);
                let inputValue="已经向 "+connect_user+" 请求"+app+"密码";
                alert(inputValue);
                this.stage_map.set(app+'-'+connect_user,"hasn't completed");
            },)
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            });
        },
        DeletePassword(index,connect_user,app){
            console.log('delete user:',connect_user,' app:',app)
            console.log("tabledata[index]:",this.tableData[index])
            if(this.tableData[index].connect_user.length==1)
                this.tableData.splice(index, 1);
            else{
                //只能删tabledata list里面的一个
                let listindex = this.tableData[index].connect_user.indexOf(connect_user);
                if (listindex !== -1) {
                this.tableData[index].connect_user.splice(listindex, 1);
                }
            }
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Password-Delete",{
                "connect_user":connect_user,
                "app":app,
            })
            .then(response => {
                    console.log('delete over.');
                },
            )
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            })
        },
        OpenAddWinowWithAPP(app){
            this.AddUIApp = app;
            this.AddWindowVisible = true;
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

.add-password{
    margin-left: 10px;
    margin-right: 10px;
    display: flex;
    align-items: center;
}

.password-search{
    display: flex;
    align-items: center;
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

.status-circle {  
  width:10px;
  height:10px;
  border-radius: 50%;
  display: block;
  margin-left: 10px;
  margin-right: 5px;
}  

/* .status-unlocked {  
  width:10px;
  height:10px;
  border-radius: 50%;
  display: block;
  margin-left: 10px;  
  background-color: green; 
}  
  
.status-unlocking {  
  width:10px;
  height:10px;
  border-radius: 50%;
  display: block;
  margin-left: 10px; 
  background-color: yellow;  
}  
  
.status-locked {  
  width:10px;
  height:10px;
  border-radius: 50%;
  display: block;
  margin-left: 10px; 
  background-color: red; 
}   */

.users-div{
    display: flex;
    /* justify-content: center; */
    align-content: center;
}

.users-show-div{
    display: flex;
    justify-content: center;
    align-content: center;
    align-items: center;
    margin-right: 10px;
}

.user-span{
    cursor: pointer;
    /* margin-right: 10px; */
    /* font-size: 20px; */
    align-content: center;
}

.user-button-group{
    display: flex;
    margin-left:10px;
}

.user-button{
}

.addpasswd-icon{   
    color: red;
    margin: 10 10 px;
    font-size:30px; 
}

</style>