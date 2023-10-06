<template>
    <div>
        <h1>{{ username }} 的 个人中心</h1>
        <el-table :data="tableData" stripe style="width: 100%">
            <el-table-column prop="id" label="id" width="50">
            </el-table-column>
            <el-table-column prop="theme" label="主题" width="180">
            </el-table-column>
            <el-table-column prop="owner" label="发起人" width="180">
            </el-table-column>
            <el-table-column prop="description" label="描述" width="180">
            </el-table-column>
            <el-table-column prop="user_list" label="报名人员" width="180">
            </el-table-column>
            <el-table-column prop="status" label="活动状态" width="180">
            </el-table-column>
            <el-table-column prop="update_time" label="更新·时间" width="180">
            </el-table-column>
            <el-table-column label="操作" width="300" align="center">
                <template slot-scope="scope">
                    <el-button type="danger" size="mini" @click="exitActivity(scope.row)">退出活动</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import { getActivityListByUid } from '../api/index'
import { delUserActivity } from '../api/index'
export default {
    data() {
        return {
            tableData: []
        }
    },
    created() {
        if (localStorage.getItem("userId") == null) {
            this.$router.push('/login')
        }
    },
    mounted() {
        if (localStorage.getItem("userId") != null) {
            this.getData()
        }
    },
    methods: {
        getData() {
            let uid = localStorage.getItem('userId')
            getActivityListByUid(uid).then(res => {
                this.tableData = res.content
            })
        },
        exitActivity(row) {
            //根据当前登录的用户信息获取用户id，然后将其与活动解绑
            let uId = localStorage.getItem("userId")
            let aId = row.id
            delUserActivity(uId, aId).then(res => {
                if (res.code == '1') {
                    this.$notify({ type: 'success', message: '退出活动成功' })
                    this.getData()
                }
            })
        },
    },
    computed: {
        username() {
            return localStorage.getItem('username')
        }
    }
}

</script>

