import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Login from '../pages/Login'
import Register from '../pages/Register'
import Home from '../pages/Home'
import MyActivity from '../pages/MyActivity'
import {Notification} from 'element-ui'
Vue.prototype.$notify = Notification;

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
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
      path: '/home',
      name: 'Home',
      component: Home
    },
    {
      path: '/myActivity',
      name: 'MyActivity',
      component: MyActivity
    }
  ]
})
