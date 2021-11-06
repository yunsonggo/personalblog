import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeRouter from './home/home.js'
import personalRouter from './personal/personal.js'
import articleRouter from './article/article.js'

Vue.use(VueRouter)

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this,location).catch(err=>err)
}

const routes = [
 HomeRouter,
 personalRouter,
 articleRouter,
 {
   path:'/*',
   redirect: '/'
 }
]

const router = new VueRouter({
  mode:'history',
  base:process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (!localStorage.getItem("token") && to.meta.isAuth === true) {
    router.push('/personal/login')
    Vue.prototype.$toast.fail("请重新登录")
    return
  }
  next()
})

export default router
