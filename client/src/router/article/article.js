export default {
    path:'/article',
    name:'article',
    component:() => import('@/views/article/detailsArticle.vue'),
    children:[
        {
            path:'/article/content/:id',
            name:'contentArticle',
            component:() => import('@/views/article/contentArticle.vue')
        },
        {
            path:'/article/search',
            name:"articleSearch",
            component:() => import('@/views/article/searchArticle.vue'),
            meta:{
                keepAlive: true,
            }
        },
        {
            path:'/article',
            redirect: '/home'
        }
    ]
}