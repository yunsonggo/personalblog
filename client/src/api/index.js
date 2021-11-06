import ajax from './ajax'
const BASE_URL = "http://192.168.1.102:8080/api"

// export const reqCities = () => ajax(BASE_URL + "/consumer/cities");
// export const reqEmailCode = (consumer_email) => ajax(BASE_URL + '/consumer/email/captcha',{consumer_email},'POST')
// 获取邮箱验证码
export const reqCodeByEmail = (consumer_email) => ajax(BASE_URL + "/captcha/email",{consumer_email},"POST")
// 获取手机验证码
export const reqCodeByPhone = (consumer_phone,tpl_id,key) =>ajax(BASE_URL + "/captcha/phone",{consumer_phone,tpl_id,key},"POST")
// 邮箱注册
export const reqRegisterByEmail = (name,email,password,checkpass,code) => ajax(BASE_URL + "/register/email",{name,email,password,checkpass,code},"POST")
// 手机注册
export const reqRegisterByPhone = (name,bizid,phone,code) => ajax(BASE_URL + "/register/phone",{name,bizid,phone,code},"POST")
// 获取验证码
export const reqCaptcha = () => ajax(BASE_URL + "/captcha/img")
// 邮箱密码登录
export const reqLoginByEmail = (email,password,captcha,captcha_id) => ajax(BASE_URL + "/login/email",{email,password,captcha,captcha_id},"POST")
// 手机登录
export const reqLoginByPhone = (phone,code,bizid) => ajax(BASE_URL + "/login/phone",{phone,code,bizid},'POST')
// 修改个人信息
export const reqEditConsumerInfo = (token,consumer,is_email,is_phone) => ajax(BASE_URL + "/auth/personal/edit",{token,consumer,is_email,is_phone},'POST')
// 获取所有激活菜单
export const reqMenu = () => ajax(BASE_URL + "/menus",{},'POST')
// 获取所有激活类别
export const reqCategory = () => ajax(BASE_URL + "/category",{},'POST')
// 获取首页文章
export const reqIndexArticle = (count,start) => ajax(BASE_URL + "/article/index",{count,start},'POST')
// 获取非首页 其他大类 首页展示文章 
export const reqOtherIndexArticle = (count,start) => ajax(BASE_URL + "/article/other/index",{count,start},'POST')
// 获取指定类别文章
export const reqMenuCategoryArticle = (menu_id,category_id,count,start) => ajax(BASE_URL+"/article/menu/category",{menu_id,category_id,count,start},'POST')
// 获取指定大类无小类文章
export const reqMenuArticle = (menu_id,count,start) => ajax(BASE_URL + "/article/menu",{menu_id,count,start},'POST')
// 根据ID 获取文章
export const reqArticleById = (article_id) => ajax(BASE_URL + "/article/info",{article_id},'POST')
// 根据文章ID 获取评论
export const reqGetArticleComment = (article_id,count,start) => ajax(BASE_URL + "/article/comment/list",{article_id,count,start},'POST')
// 根据文章ID 添加评论
export const reqAddArticleComment = (article_id,author_id,content,author_nickname,author_icon) => ajax(BASE_URL + "/article/comment/insert",{article_id,author_id,content,author_nickname,author_icon},'POST')
// 添加二级评论
export const reqAddArticleCommentChild = (comment_id,author_id,content,author_nickname,author_icon) => ajax(BASE_URL + "/article/comment/insert/child",{comment_id,author_id,content,author_nickname,author_icon},'POST')
// 更新赞或者踩
export const reqGoodOrBadArticle = (article_id,star,tread,job_tag) => ajax(BASE_URL + "/article/good/or",{article_id,star,tread,job_tag},'POST')
// 搜索文章
export const reqSearchArticle = (keyword,count,start) => ajax(BASE_URL + "/article/search/keyword" ,{keyword,count,start},'POST')
// 获取网站信息
export const reqWebsiteInfo = () => ajax(BASE_URL + "/website/info",{},'POST')
// 获取友情链接
export const reqLinksInfo = () => ajax(BASE_URL + "/links",{},'POST')