export default {
    path:'/',
    name:"home",
    meta:{
        isAuth: true,
    },
    component:() => import('@/views/home/home.vue'),
    children:[
        {
            path:'/home/index',
            name:'homePage',
            component:() => import('@/views/home/homeIndex.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/menu/list',
            name:'menuList',
            component:() => import('@/views/menu/menuList.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/menu/add',
            name:'addMenu',
            component:() => import('@/views/menu/addMenu.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/article/list',
            name:'articleList',
            component:() => import('@/views/article/articleList.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/article/add',
            name:'addArticle',
            component:() => import('@/views/article/addArticle.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/article/edit',
            name:'editArticle',
            component:() => import('@/views/article/editArticle.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/user',
            name:'user',
            component:() => import('@/views/user/user.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/comment',
            name:'comment',
            component:() => import('@/views/comment/comment.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/link/list',
            name:'linkList',
            component:() => import('@/views/link/link.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/link/add',
            name:'addLink',
            component:() => import('@/views/link/addLink.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/website',
            name:'website',
            component:() => import('@/views/website/website.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/',
            redirect: '/home/index'
        }
    ]
}