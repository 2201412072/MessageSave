<template>
    <div>
        <!-- 添加密码&待用密码 -->
        <div class="add-password">
            <el-button type="primary" @click="AddWindowVisible=true">Add password</el-button>
            <el-button type="primary" @click="SearchPasswordUsing">Search password using</el-button>
        </div>
        <!-- 查询 -->
        <div class="password-search">
            <el-form :model="myform_password" :inline="true">
                <el-form-item label="Application">
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
                    <el-table-column label='Connect User' width="600">
                        <template v-slot="scope"> 
                            <div class="users-div">
                                <div class="users-show-div" v-for="connect_user in scope.row.connect_user" :key="connect_user">
                                    <div class="status-circle" :class="stage_options[scope.row.stage-1]"></div>
                                    <span class="user-span" @click="ToggleButton(scope.row,connect_user)">{{connect_user}}</span>
                                    <el-button-group class="user-button-group">
                                        <el-button class="user-button" type="primary" v-show="scope.row.ButtonVisible[connect_user]"  @click="LockPassword(scope.row.connect_user,scope.row.app)" >
                                            <el-icon><Lock /></el-icon>
                                        </el-button>
                                        <el-button class="user-button" type="primary" v-show="scope.row.ButtonVisible[connect_user]" @click="UnlockPassword(scope.row.connect_user,scope.row.app)" >
                                            <el-icon><Unlock /></el-icon>
                                        </el-button>
                                        <el-button class="user-button" type="primary" v-show="scope.row.ButtonVisible[connect_user]" @click="DeletePassword(scope.$index,scope.row.connect_user,scope.row.app)" >
                                            <el-icon><Delete /></el-icon>
                                        </el-button>
                                    </el-button-group>
                                </div>
                                <!-- :icon="Lock" :icon="Unlock" :icon="Delete" -->
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
    
        <!-- <h2> App-Password Manage </h2>
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
        <!-- <div></div>
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
        </div> -->
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
            myform_result:{
                app:'',
                connect_user:'',
                selectedOption:'',
            },
            AddWindowVisible: false,
            AddUIApp: '',
            // stage_options:[
            //     {label:"has completed",value:1},
            //     {label:"hasn't completed",value:2},
            //     {label:"all",value:3},
            // ],
            stage_options:[
                {label:"unlocked",class:'status-unlocked'},
                {label:"unlockingd",class:'status-unlocking'},
                {label:"locked",class:'status-locked'},
            ],
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
                stage:3,
            }));
            console.log("tableData",this.tableData);
        },
        SearchPasswordUsing(){

        },
        ToggleButton(row,connect_user){
            // console.log("togglbutton tableData",this.tableData);
            console.log("togglbutton row",row);
            row.ButtonVisible[connect_user] = !row.ButtonVisible[connect_user];
            // this.$set(row.ButtonVisible,connect_user,!row.ButtonVisible[connect_user]);
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
        },
        PasswordSearchAll(){
            axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Password-Search",{
                "Application":"",
                "Username":"",
            })
            .then(response=>{
                this.tableData=response.data;
                this.ChangeTableData();
            },)
            .catch(error=>{
                console.log("error",error);
                alert("请求失败");
            })
        },
        LockPassword(connect_user,app){
            // 将某密码复位，即重新回到加密状态
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
            },)
            .catch(response => {
                console.log("error",response);
                alert("请求失败");
            });
        },
        DeletePassword(index,connect_user,app){
            console.log('delete user:',connect_user,' app:',app)
            this.tableData.splice(index, 1);
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
        // PasswordReset(){
        //     this.myform_password.connect_user = '';
        //     this.myform_password.app = '';
        // },
        // handleDelete_password(index,temp1,temp2){
        //     console.log('delete user:',temp1,' app:',temp2)
        //     this.tableData.splice(index, 1);
        //     axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Password-Delete",{
        //         "connect_user": temp1,
        //         "app":temp2,
        //     })
        //     .then(response => {
        //             console.log('delete.');
        //         },
        //     )
        //     .catch(response => {
        //         console.log("error",response);
        //         alert("请求失败");
        //     })
        // },

        // ResultSearchAll(){
        //     axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Result-Search",{
        //         "Username": "",
        //         "Application":"",
        //         "stage":"",
        //     })
        //     .then(response=>{
        //         this.tableData_result=response.data;
        //         console.log('get result data.',this.tableData_result);
        //     },)
        //     .catch(error=>{
        //         console.log("error",error);
        //         alert("请求失败");
        //     })
        // },

        // ResultSearch(){
        //     var myform = this.myform_result
        //     if(myform.app==""&&myform.connect_user=="")
        //     {
        //         self.PasswordSearchAll()
        //         return 
        //     }
        //     axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Result-Search",{
        //         "Username": myform.connect_user,
        //         "Application":myform.app,
        //         "stage":myform.selectedOption,
        //     })
        //     .then(response => {
        //             console.log('get result data.',response.data);
        //             this.tableData_result = response.data;
        //         },
        //     )
        //     .catch(response => {
        //         console.log("error",response);
        //         alert("请求失败");
        //     })
        // },

        // ResultReset(){
        //     this.myform_result.Username = '';
        //     this.myform_result.app = '';
        //     this.myform_result.selectedOption = 1;
        // },
        // handleDelete_result(index,temp1,temp2,temp_stage){
        //     this.tableData_result.splice(index, 1);
        //     axios.post("http://localhost:8090/ClientHomePage/PasswordManage/Result-Delete",{
        //         "Username": temp1,
        //         "Application":temp2,
        //         "stage":temp_stage,
        //     })
        //     .then(response => {
        //             console.log('delete.');
        //         },
        //     )
        //     .catch(response => {
        //         console.log("error",response);
        //         alert("请求失败");
        //     })
        // },
    }
}
</script>

<style scoped>  
.status-circle {  
  width: 20px;  
  height: 20px;  
  border-radius: 50%;  
  display: inline-block;  
  margin: 0 5px;  
  vertical-align: middle;  
}  


.status-unlocked {  
  background-color: green; /* 状态1的颜色 */  
}  
  
.status-unlocking {  
  background-color: yellow; /* 状态2的颜色 */  
}  
  
.status-locked {  
  background-color: red; /* 状态3颜色 */  
}  

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