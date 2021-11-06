import ajax from './ajax'
const BASE_URL = "http://192.168.1.102:8081/backend"

// 服务器信息
export const reqServerInfo = () => ajax(BASE_URL + "/server/info",{},'POST')
// 获取验证码
export const reqCaptcha = () => ajax(BASE_URL + "/captcha/img")
// 用户注册
export const reqRegisterManager = (name,password,captcha_id,captcha) => ajax(BASE_URL + "/manager/register",{name,password,captcha_id,captcha},'POST')
// 用户登录
export const reqLogin = (name,password,captcha_id,captcha) => ajax(BASE_URL + "/manager/login",{name,password,captcha_id,captcha},'POST')
// 获取网站信息
export const reqWebsiteInfo = () => ajax(BASE_URL + "/website/info",{},'POST')
// 获取后台菜单
export const reqMenuList = (token) => ajax(BASE_URL + "/auth/menu/list",{token},'POST')
// 获取前台菜单
export const reqWebMenuList = (token) => ajax(BASE_URL + "/auth/menu/web/list",{token},'POST') 
// 修改前台菜单
export const reqWebMenuEdit = (token,Id,title,sort,state,parent_id) => ajax(BASE_URL + "/auth/menu/web/edit",{token,Id,title,sort,state,parent_id},'POST')
// 添加前台菜单
export const reqWebMenuAdd = (token,title,sort,state,parent_id) => ajax(BASE_URL + "/auth/menu/web/add",{token,title,sort,state,parent_id},'POST')
// 删除前台菜单 
export const reqWebMenuRemove = (token,id,parent_id) => ajax(BASE_URL + "/auth/menu/web/remove",{token,id,parent_id},'POST')
// 获取所有文章
export const reqArticleAll = (token) => ajax(BASE_URL + "/auth/article/all",{token},'POST')
// 批量删除文章
export const reqRemoveArticleArr = (token,id_str_arr) => ajax(BASE_URL + "/auth/article/remove/arr",{token,id_str_arr},'POST')
// 添加文章
export const reqAddArticle = (token,article_info) => ajax(BASE_URL + "/auth/article/add",{token,article_info},'POST')
// 根据ID 获取文章数据
export const reqArticleById = (token,article_id) => ajax(BASE_URL + "/auth/article/info",{token,article_id},'POST')
// 根据ID 修改文章
export const reqArticleEditById = (token,article_info) => ajax(BASE_URL + "/auth/article/edit",{token,article_info},'POST')
// 获取用户列表
export const reqConsumerList = (token) => ajax(BASE_URL + "/auth/consumer/list",{token},'POST')
// 更新用户状态
export const reqConsumerState = (token,id,state) => ajax(BASE_URL + "/auth/consumer/edit/state",{token,id,state},'POST')
// 获取所有评论数据
export const reqCommentAll = (token) => ajax(BASE_URL + "/auth/comment/list",{token},'POST')
// 删除评论
export const reqCommentRemove = (token,id,comment_id) => ajax(BASE_URL + "/auth/comment/remove",{token,id,comment_id},'POST')
// 获取所有友情链接
export const reqLinkList = (token) => ajax(BASE_URL + "/auth/link/list",{token},'POST')
// 修改链接文本
export const reqLinkText = (token,id,icon,link_url,sort,title) => ajax(BASE_URL + "/auth/link/edit/text",{token,id,icon,link_url,sort,title},'POST')
// 删除链接
export const reqRemoveLink = (token,id) => ajax(BASE_URL + "/auth/link/remove",{token,id},'POST')
// 添加链接
export const reqInsertLink = (token,id,icon,link_url,sort,title) => ajax(BASE_URL + "/auth/link/insert/one",{token,id,icon,link_url,sort,title},'POST')
// 更新网站基本信息
export const reqWebsiteInfoEdit = (token,website_info) => ajax(BASE_URL + "/auth/website/info/edit",{token,website_info},'POST')
// 更新关于我们
export const reqEditWebsiteAbout = (token,about) => ajax(BASE_URL + "/auth/website/about/edit",{token,about},'POST')
// 更新用户协议
export const reqEditWebsiteAgreement = (token,agreement) => ajax(BASE_URL + "/auth/website/agreement/edit",{token,agreement},'POST')
// 更新隐私声明
export const reqEditWebsitePrivacy = (token,privacy) => ajax(BASE_URL + "/auth/website/privacy/edit",{token,privacy},'POST')
