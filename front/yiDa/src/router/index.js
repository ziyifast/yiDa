import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Login from '../pages/Login'
import Register from '../pages/Register'
import Home from '../pages/Home'
import MyActivity from '../pages/MyActivity'
import CreateActivity from '../pages/CreateActivity'
import { Notification } from 'element-ui'
Vue.prototype.$notify = Notification;

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/index',
      name: 'Index',
      component: Index
    },
    {
      path: '/register',
      name: 'Register',
      component: Register
    },
    {
      path: '/myActivity',
      name: 'MyActivity',
      component: MyActivity
    },
    {
      path: '/createActivity',
      name: 'CreateActivity',
      component: CreateActivity
    }
  ]
})
