<template>
    <div>
    <div><h1>this is a register vue</h1></div>
    <div class="container">
        <div class="left">
            <RegisterLogo />
        </div>
        <div class="center"></div>
        <div class="right">
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
                <el-form-item label="昵称" prop="username">
                    <el-input v-model="ruleForm.username"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model="ruleForm.password"></el-input>
                </el-form-item>
                <el-form-item label="微信" prop="wechat">
                    <el-input v-model="ruleForm.wechat"></el-input>
                </el-form-item>
                <el-form-item label="电话" prop="phone">
                    <el-input v-model="ruleForm.phone"></el-input>
                </el-form-item>
                <el-form-item label="城市" prop="city">
                    <el-input v-model="ruleForm.city"></el-input>
                </el-form-item>
                <el-form-item label="兴趣点" prop="hobbies">
                    <el-select v-model="ruleForm.hobbies" placeholder="请选择你的爱好">
                        <el-option label="篮球" value="basketball"></el-option>
                        <el-option label="城市漫步" value="city walk"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="生日" required>
                    <el-col :span="11">
                        <el-form-item prop="birthday">
                            <el-date-picker type="date" placeholder="请选择出生日期" v-model="ruleForm.birthday"
                                style="width: 100%;"></el-date-picker>
                        </el-form-item>
                    </el-col>
                </el-form-item>
                <el-form-item label="个人标签" prop="type">
                    <el-checkbox-group v-model="ruleForm.type">
                        <el-checkbox label="文艺青年" username="type"></el-checkbox>
                        <el-checkbox label="rapper" username="type"></el-checkbox>
                        <el-checkbox label="社恐" username="type"></el-checkbox>
                        <el-checkbox label="社牛" username="type"></el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="个人描述" prop="desc">
                    <el-input type="textarea" v-model="ruleForm.desc"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
                    <el-button @click="resetForm('ruleForm')">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</div>
</template>
<script>
import RegisterLogo from '../components/RegisterLogo.vue'
export default {
    username: 'Register',
    components: {
        RegisterLogo
    },
    data() {
        return {
            ruleForm: {
                username: '',
                password: '',
                wechat: '',
                phone: '',
                city: '',
                description: '',
                hobbies: '',
                birthday: '',
                type: [],
                desc: ''
            },
            rules: {
                username: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    { min: 1, max: 15, message: '长度在 1 到 15 个字符', trigger: 'blur' }
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                    {min: 6, max: 15, message: '长度在 6 到 15 个字符', trigger: 'blur'}
                ],
                hobbies: [
                    { required: true, message: '请输入你的爱好', trigger: 'change' }
                ],
                birthday: [
                    { type: 'date', required: true, message: '请选择日期', trigger: 'change' }
                ],
                desc: [
                    { required: false, message: '请填写活动形式', trigger: 'blur' }
                ]
            }
        };
    },
    methods: {
        submitForm(formusername) {
            this.$refs[formusername].validate((valid) => {
                if (valid) {
                    //TODO 向后端发起请求
                    this.$notify({ type: 'success', message: '注册成功' })
                    this.$router.push({ path: '/home' })
                } else {
                    this.$notify({ type: 'error', message: '注册失败' })
                    console.log('error submit!!');
                    return false;
                }
            });
        },
        resetForm(formusername) {
            this.$refs[formusername].resetFields();
        }
    }
}
</script>

<style scoped>
.container {
    display: flex;
    height: 100vh;
}

.left {
    flex-basis: 40%;
    background-color: lightgray;
}

.center{
    flex-basis: 5%;
}

.right {
    flex-basis: 40%;
    background-color: white;
}
</style>