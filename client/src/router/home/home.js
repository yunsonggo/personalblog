export default {
    path:'/',
    name:'home',
    component:() => import('@/views/home/home.vue'),
    children:[
        {
            path:'/home/index',
            name:'homePage',
            component:() => import('@/views/home/homeIndex.vue')
        },
        {
            path:'/',
            redirect: '/home/index'
        }
    ]
}