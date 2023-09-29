<template>
    <div >
        <div >
            <span>账号登录</span>
        </div>
        <div class="container">
            this is a login page
        </div>
        <el-form :model="loginForm" ref="loginForm" label-width="100px" class="demo-ruleForm">
            <el-form-item label="用户名" prop="username" >
                <el-input v-model="loginForm.username" placeholder="用户名" label-width="50px"></el-input>
            </el-form-item>
        <el-form-item prop="password" label="密码">
          <el-input type="password" v-model="loginForm.password" placeholder="密码"  @keyup.enter.native="handleLoginIn"></el-input>
        </el-form-item>
        <div class="login-btn">
            <el-button @click="goSignUp">注册</el-button>
            <el-button type="primary" @click="handleLoginIn">登录</el-button>
        </div>
        </el-form>
        <div>
    </div>
    </div>
</template>

<script>
import {loginIn} from '../api/index'
import {getData} from '../api/index'
import {mixin} from '../../mixins'
export default {
    mixins: [mixin],
    data() {
        return {
            loginForm: {
                username: '',
                password: ''
            },
        }
    },
    methods: {
        handleLoginIn() {
            let _this = this;
            console.log(this.loginForm.username)
            console.log(this.loginForm.password)
            loginIn(this.loginForm.username, this.loginForm.password)
            .then(res => {
                if (res.code == "1") {
                    console.log(res.data)
                    _this.notify('登录成功', 'success')
                    this.$router.push({path: '/home'})
                }else {
                    _this.notify('用户名或密码错误', 'error')
                    console.log(res.msg)
                }
            })
            .catch(err => {
                console.log(err)
            })
            
        },
        goSignUp(){

        }
    }
}
</script>