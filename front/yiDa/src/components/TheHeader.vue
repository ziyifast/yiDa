<template>
    <div>
        <ul>
            <li v-for="(item, index) in navMsg" :key="index" @click="gonavMsg(item.path)">{{item.name}}</li>
        </ul>
    </div>
</template>

<script>
import {navMsg} from '../assets/data/header'
export default{
    name:'TheHeader',
    data(){
        return {
            navMsg:[]
        }
    },
    created(){
        this.navMsg = navMsg
    },
    methods:{
        gonavMsg(path){
            this.$router.push(path)
        }
    }
}

</script>