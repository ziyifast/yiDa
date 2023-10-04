//存储用户信息
const user ={
    state: {
      userId: '',   //用户id
      username: '', //用户名
      avator: '',   //用户头像
    },
    getters: {
      userId: state => {
        let userId = state.userId
        if(!userId){//如果userId不为空则返回
          userId = JSON.parse(window.localStorage.getItem('userId'))
        }
        return userId;
      },
      username: state => {
        let username = state.username
        if(!username){
          username = JSON.parse(window.localStorage.getItem('username'))
        }
        return username
      },
      avator: state => {
        let avator = state.avator
        if(!avator){
          avator = JSON.parse(window.localStorage.getItem('avator'));
        }
        return avator;
      }
    },
    mutations: {
      setUserId: (state, userId) => {
        state.userId = userId;
        window.localStorage.setItem('userId', JSON.stringify(userId))
      },
      setUsername: (state, username) => {
        state.username = username
        window.localStorage.setItem('username', JSON.stringify(username))
      },
      setAvator: (state, avator) => {
        state.avator = avator;
        window.localStorage.setItem('avator', JSON.stringify(avator));
      }
  
    }
  }
  export default user
  