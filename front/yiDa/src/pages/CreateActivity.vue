<template>
    <div>
        <div>
            <h1>发起活动页面</h1>
        </div>
        <div class="container">
            <div class="left">
                <RegisterLogo />
            </div>
            <div class="center"></div>
            <div class="right">
                <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
                    <el-form-item label="主题" prop="theme">
                        <el-input v-model="ruleForm.theme"></el-input>
                    </el-form-item>
                    <el-form-item label="描述" prop="description">
                        <el-input v-model="ruleForm.description"></el-input>
                    </el-form-item>
                    <el-form-item label="状态" prop="status">
                        <el-input v-model="ruleForm.status"></el-input>
                    </el-form-item>
                    <!-- <el-form-item label="活动性质">
                        <el-checkbox-group v-model="ruleForm.hobbies">
                            <el-checkbox label="篮球" name="type"></el-checkbox>
                            <el-checkbox label="足球" name="type"></el-checkbox>
                            <el-checkbox label="散步" name="type"></el-checkbox>
                            <el-checkbox label="游泳" name="type"></el-checkbox>
                        </el-checkbox-group>
                    </el-form-item> -->
                    <!-- <el-form-item label="兴趣点" prop="hobbies">
                        <el-select v-model="ruleForm.hobbies" placeholder="请选择你的爱好">
                            <el-option label="篮球" value="basketball"></el-option>
                            <el-option label="城市漫步" value="city walk"></el-option>
                        </el-select>
                    </el-form-item> -->
                    <!-- <el-form-item label="个人标签" prop="type">
                        <el-checkbox-group v-model="ruleForm.type">
                            <el-checkbox label="文艺青年" username="type"></el-checkbox>
                            <el-checkbox label="rapper" username="type"></el-checkbox>
                            <el-checkbox label="社恐" username="type"></el-checkbox>
                            <el-checkbox label="社牛" username="type"></el-checkbox>
                        </el-checkbox-group>
                    </el-form-item> -->
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
import {addActivityDetail} from '../api/index'
export default {
    name: 'CreateActivity',
    components: {
        RegisterLogo
    },
    data() {
        return {
            ruleForm: {
                theme: '唱歌',
                owner: 0,
                description: '',
                user_list: '',
                status: 1,
                create_time: '',
                update_time: '',
            },
            rules: {
                theme: [
                    { required: true, message: '请输入主题', trigger: 'blur' },
                    { min: 1, max: 15, message: '长度在 1 到 15 个字符', trigger: 'blur' }
                ],
            }
        };
    },
    methods: {
        submitForm(formusername) {
            this.$refs[formusername].validate((valid) => {
                if (valid) {
                    //TODO 向后端发起请求
                    let data = this.ruleForm
                    this.ruleForm.owner = parseInt(localStorage.getItem('userId'))
                    this.ruleForm.user_list = localStorage.getItem('userId')
                    console.log(data);
                    addActivityDetail(this.ruleForm).then(res => {
                        console.log(res);
                    })
                    this.$notify({ type: 'success', message: '发起成功' })

                } else {
                    this.$notify({ type: 'error', message: '发起失败' })
                    console.log('error submit!!');
                    return false;
                }
            });
        },
        resetForm(formusername) {
            this.$refs[formusername].resetFields();
        },

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

.center {
    flex-basis: 5%;
}

.right {
    flex-basis: 40%;
    background-color: white;
}
</style>