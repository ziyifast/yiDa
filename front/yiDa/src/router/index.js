import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Login from '../pages/Login'

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
  ]
})
