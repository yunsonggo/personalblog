import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeRouter from './home/homeRouter.js'
import ManagerRouter from './manager/managerRouter.js'

Vue.use(VueRouter)

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this,location).catch(err=>err)
}

const routes = [
  HomeRouter,
  ManagerRouter,
  {
    path:'/*',
    redirect: '/'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (!localStorage.getItem("adminToken") && to.meta.isAuth === true) {
    router.push('/manager/login')
    Vue.prototype.$toast.fail("请重新登录")
    return
  }
  next()
})

export default router
