export default {
    path:"/manager",
    name:"manager",
    component:() => import('@/views/manager/manager.vue'),
    children:[
        {
            path:'/manager/login',
            name:'login',
            component:() => import('@/views/manager/login.vue')
        },
        {
            path:'/manager/register',
            name:'register',
            component:() => import('@/views/manager/register.vue')
        },
        {
            path:'/manager',
            redirect:'/manager/login'
        }
    ]
}