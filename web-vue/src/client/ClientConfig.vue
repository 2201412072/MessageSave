<template>
   <div class="top-div">
   </div>
</template>

<script>
import TopBar from './TopBar.vue';
const axios = require('axios');

export default{
    data(){
        return{
            public_key:'',
            tableData:[],
            myform:{
                connect_user: '',
            }
        };
    },
    components: {
    TopBar,
    },
    mounted(){
        this.SearchAll()
    },
    methods:{
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
</style>
