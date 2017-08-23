import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/home'
import AllBooks from '@/components/allbooks'
import GetInfo from '@/components/getinfo'
import PostInfo from '@/components/postinfo'

Vue.use(Router)
export default new Router({
  routes: [
    {
      path: '/home',
      name: 'home',
      component: Home
    },
    {
      path: '/book/pullbooksinfo/all',
      name: 'allbooks',
      component: AllBooks
    },
    {
      path: '/book/pullbookinfo/:id',
      name: 'getinfo',
      component: GetInfo
    },
    {
      path: '/book/postbookinfo/',
      name: 'postinfo',
      component: PostInfo
    },
    {
      path: '*',
      redirect: '/home'
    }
  ]
})
