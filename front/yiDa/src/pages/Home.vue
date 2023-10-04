<template>
  <div>
    <div align="right">
      <el-dropdown trigger="click" @command="handleCommand">
        <span>
          当前登录用户：{{username}}
        </span>
        <el-dropdown-menu slot="dropdown" >
          <el-dropdown-item command="toMyActivity">个人中心</el-dropdown-item>
          <el-dropdown-item command="logout">退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
    <div>
      <h1>this is home page</h1>
      <!-- <div @click="handleClick()">testJson</div> -->
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
            <el-button type="primary" size="mini" @click="addActivity(scope.row)">报名</el-button>
            <el-button type="danger" size="mini" @click="exitActivity(scope.row)">退出活动</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>
<script>
import { updateActitityDetail } from '../api/index'
import { getActivityList } from '../api/index'
export default {
  name: 'home',
  data() {
    return {
      tableData: [],
    }
  },
  computed:{
    username(){
      return localStorage.getItem('username')
    }
  },
  mounted() {
    this.getData()
  },
  methods: {
    // 获取数据
    getData() {
      // 调用接口
      getActivityList().then(res => {
        this.handleActivityDetail(res.content)
      })
    },
    addActivity(row) {
      //根据当前登录的用户信息获取用户id，然后将其与活动关联
      let uId = localStorage.getItem('userId');
      row.update_time = ''
      let activityDetail = row
      activityDetail.user_list = activityDetail.user_list + ',' + uId
      console.log("detail user_list=", activityDetail.user_list);
      updateActitityDetail(activityDetail).then(res=>{
        console.log("res=", res);
        if (res.code == '1') {
          this.$notify({ type: 'success', message: '报名成功' })
          this.getData()
        }
      })
    },
    exitActivity(row){
      //根据当前登录的用户信息获取用户id，然后将其与活动解绑
      let uId = localStorage.getItem('userId');
      let activityDetail = row
      let arr = activityDetail.user_list.split(",")
      let newUserList=[]
      for (let i = 0; i < arr.length; i++){
        if (arr[i] != uId) {
          newUserList.push(arr[i])
        }
      }
      console.log("newUserList=", newUserList);
      activityDetail.user_list = newUserList.join(",")
      console.log("detail user_list=", activityDetail.user_list);
      updateActitityDetail(activityDetail).then(res=>{
        console.log("res=", res);
        if (res.code == '1') {
          this.$notify({ type: 'success', message: '退出活动成功' })
          this.getData()
        }
      })
    },
    handleActivityDetail(data) {
      //TODO 将报名人员转换为username
      this.tableData = data
    },
    handleCommand(command){
      if (command == "toMyActivity"){
        this.$router.push({ path: '/myActivity' })
      } else if (command == "logout"){
        //清空缓存，退出登录
        localStorage.removeItem('userId');
        this.$router.push({ path: '/' })
      }
    }
  }
}
</script>

<style scoped>
.el-row {
  margin-bottom: 20px;

  &:last-child {
    margin-bottom: 0;
  }
}

.el-col {
  border-radius: 4px;
}

.bg-purple-dark {
  background: #99a9bf;
}

.bg-purple {
  background: #d3dce6;
}

.bg-purple-light {
  background: #e5e9f2;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}

.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}
</style>