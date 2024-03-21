<template>
    <div>
        <div v-if="this.CardType=='AskEncrypt'" class="radius-card-ask-encrypt">
            <span>{{this.Content}}</span>
            <!-- <div class="flex-grow" /> -->
            <div class="card-button-div">
                <el-button type="primary" @click="HandleAgree(key,User,App)">Ok</el-button>
            </div>
        </div>
        <div v-if="this.CardType=='AskDecrypt'" class="radius-card-ask-decrypt">
            <span>{{this.Content}}</span>
            <!-- <div class="flex-grow" /> -->
            <div class="card-button-div">
                <el-button type="primary" @click="HandleAgree(key,User,App)">Yes</el-button>
                <el-button  @click="HandleDisAgree(key,User,App)">No</el-button>
            </div>
        </div>
    </div>
</template>

<script>

export default({
    props:{
        Content:{
            type:String,
            default:'',
        },
        CardType:{
            type:String,
        },
        key:{
            type:Number,
            default:0,
        },
        User:{
            type:String,
            default:'',
        },
        App:{
            type:String,
            default:'',
        },
        HandleAgree:{
            type:Function,
        },
        HandleDisAgree:{
            type:Function,
        }
    },
    data(){

    },
    methods:{

    }
})
</script>

<style scoped>
.radius-card-ask-encrypt{
    border-radius: round;
    box-shadow: dark;
    background: gray;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 1000px;
    height: 60px;
    margin-top: 10px;
    border-radius: 10px;
}

.radius-card-ask-decrypt{
    border-radius: round;
    box-shadow: dark;
    background: yellow;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 1000px;
    height: 60px;
    margin-top: 10px;
    border-radius: 10px;
}

.card-button-div{
    margin-left: 20px;
    display: flex;
    justify-content: center;
    justify-items: center;
    align-content: center;
}
</style>
