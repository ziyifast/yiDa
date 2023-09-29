<template>
    <div class="main">
        <!-- <div class="demo-image__placeholder">
            <div class="block">
                <el-image :src="src" >
                    <div slot="placeholder" class="image-slot">
                        加载中<span class="dot">...</span>
                    </div>
                </el-image>
            </div>
        </div> -->
        <div class="login-container">
        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="70%" class="demo-ruleForm">
            <el-form-item>
                <div class="signUp-head"> yi-Da </div>
            </el-form-item>
            <el-form-item label="用户名" prop="username" label-width="20%">
                <el-input v-model.number="ruleForm.username" label-width="auto"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="pass" label-width="20%">
                <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label-width="20%">
                <el-button type="primary" @click="submitForm('ruleForm')">登录</el-button>
                <el-button @click="signUp">注册</el-button>
            </el-form-item>
        </el-form>
    </div>
    </div>
</template>
<script>
export default {
    components: {
    },
    data() {
        var validatePass = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请输入密码'));
            } else {
                if (this.ruleForm.checkPass !== '') {
                    this.$refs.ruleForm.validateField('checkPass');
                }
                callback();
            }
        };
        var validateUsername = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请输入用户名'));
            } else {
                if (this.ruleForm.checkPass !== '') {
                    this.$refs.ruleForm.validateField('checkPass');
                }
                callback();
            }
        }
        return {
            src: 'https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg',
            ruleForm: {
                pass: '',
                username: ''
            },
            rules: {
                pass: [
                    { validator: validatePass, trigger: 'blur' }
                ],
                username: [
                    { validator: validateUsername, trigger: 'blur' }
                ]
            }
        };
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    alert('submit!');
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        },
        signUp() {

        }
    }
}
</script>

<style scoped> 
.login-container {
    border-radius: 15px;
    background-clip: padding-box;
    margin: 0 auto;
    width: 350px;
    padding: 35px 35px 15px 35px;
    background: #fff;
    /* background-image: url('../assets/lo.jpg'); */
    border: 1px solid #eaeaea;
    box-shadow: 0 0 25px #cac6c6;
}
.title{
    margin: 0px auto 40px auto;
    text-align: center;
    color: red;
    font-size: larger;
    align-items: center;
}

.signUp-head {
    margin-left: -200px;
    margin-bottom: 10px;
    font-size: 20px;
    font-weight: 600;
  }

.main{
    background-image: url('../assets/logo.png');
    background-size: 100% 100%;
    height: 100%;
}
</style>