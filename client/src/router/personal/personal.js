export default {
    path:'/personal',
    name:'personal',
    component:() => import('@/views/personal/personal.vue'),
    children:[
        {
            path:'/personal/login',
            name:'login',
            component:() => import('@/views/personal/login.vue')
        },
        {
            path:'/personal/register',
            name:'register',
            component:() => import('@/views/personal/register.vue')
        },
        {
            path:'/personal/center',
            name:'center',
            component:() => import('@/views/personal/center.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/personal/edit',
            name:'personalEdit',
            component:() => import('@/views/personal/edit.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/personal',
            redirect: 'login'
        }
    ]
}