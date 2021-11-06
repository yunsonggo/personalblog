<template>
  <div style="background-color:white">
    <Navbar :userInfo="userInfo"></Navbar>
    <div class="articleTopView">
      <div v-if="articleInfo.video">
        <video :src="staticUrl + articleInfo.video" controls></video>
      </div>
      <div v-else-if="articleInfo.cover">
        <img :src="staticUrl + articleInfo.cover" alt="" />
      </div>
      <div v-else>
        <img src="@/assets/default_top_view.jpg" alt="" />
      </div>
    </div>
    <div class="articleInfo">
      <van-collapse v-model="activeNames" @change="changeInfo">
        <van-collapse-item name="1">
          <template #title>
            <div class="articleTitle">
              <div v-if="articleInfo.is_original == 1" class="titleIcon1">
                <span>原创</span>
              </div>
              <div v-if="articleInfo.link_url" class="titleIcon2">
                <span>引用</span>
              </div>
              <div>
                <span style="  font-weight: bolder; font-size:3.88vw">{{
                  articleInfo.title
                }}</span>
              </div>
            </div>
          </template>
          <div class="articleDesc">
            <div v-if="articleInfo.link_url">
              <span>原文:</span>
              <span
                ><a :href="articleInfo.link_url" target="_blank">{{
                  articleInfo.link_url
                }}</a></span
              >
            </div>
            <div v-if="articleInfo.keyword">
              <span>关键字:</span>
              <span>{{ articleInfo.keyword }}</span>
            </div>
            <div v-if="articleInfo.desc">
              <span>描述:</span>
              <span>{{ articleInfo.desc }}</span>
            </div>
            <div v-if="articleInfo.author">
              <span>作者:</span>
              <span>{{ articleInfo.author }}</span>
            </div>
            <div v-else-if="articleInfo.author_id === userInfo.id">
              <span>作者:</span>
              <span>{{ userInfo.name }}</span>
            </div>
            <div v-else>
              <span>作者:</span>
              <span>暂未设置</span>
            </div>
            <div v-if="articleInfo.update_time">
              <span>时间:</span>
              <span>{{ articleInfo.update_time }}</span>
            </div>
            <div v-else>
              <span>时间:</span>
              <span>{{ articleInfo.create_time }}</span>
            </div>
          </div>
        </van-collapse-item>
      </van-collapse>
      <div class="articleCount">
        <div>
          <van-icon name="browsing-history-o" />
          <span>阅读:{{ articleInfo.read_num }}</span>
        </div>
        <div>
          <van-icon name="good-job-o" />
          <span @click="goodJob(articleInfo.Id, articleInfo.star)"
            >赞:{{ articleInfo.star }}</span
          >
        </div>
        <div>
          <van-icon name="closed-eye" />
          <span @click="badJob(articleInfo.Id, articleInfo.tread)"
            >踩:{{ articleInfo.tread }}</span
          >
        </div>
        <div @click="changeShowShareBtn">
          <van-icon name="share-o" />
          <span>分享</span>
        </div>
        <div @touchstart="toCommentTop">
          <van-icon name="chat-o" />
          <span>评论:{{ commentNumTotal }}</span>
        </div>
      </div>
      <div class="articleContent">
        {{ articleInfo.content }}
      </div>
      <div v-if="articleInfo.banner" class="articleBanner">
        <img :src="staticUrl + articleInfo.banner" alt="" />
      </div>
    </div>
    <BottomArticle
      :articleId="articleInfo.Id"
      @changeArticleId="changeArticleId"
    ></BottomArticle>
    <ArticleComment
      :userInfo="userInfo"
      :commentNum="articleInfo.comment_num"
      @submitComment="submitComment"
      :commentData="commentList"
      @submitCommentChild="submitCommentChild"
      @commentNum="commentNum"
    ></ArticleComment>
    <Share
      :showShareBtn="showShareBtn"
      @closeShare="handleCloseShare"
      :articleUrlId="articleInfo.Id"
      @qrcode="createQRCode"
      :articleInfo="articleInfo"
    ></Share>
    <van-dialog 
    v-model="showQRCode" 
    :title="shareTitle"
    show-cancel-button
    confirmButtonText="保存"
    @confirm="saveQRCode"
    @open="checkBrowser"
    >
    <div class="qrcode">
      <div ref="qrCodeUrl" class="qrcodeImg"></div>
    </div>
    </van-dialog>
  </div>
</template>

<script>
import Navbar from "@/components/common/Navbar.vue";
import {
  reqArticleById,
  reqGetArticleComment,
  reqAddArticleComment,
  reqAddArticleCommentChild,
  reqGoodOrBadArticle,
} from "@/api";
import WebStaticUrl from "@/components/common/WebStaticUrl.vue";
import BottomArticle from "@/views/article/bottomArticle.vue";
import ArticleComment from "@/components/article/articleComment.vue";
import Share from "@/components/article/share.vue";
import QRCode from "qrcodejs2";
import html2canvas from "html2canvas"
export default {
  name: "contentArticle",
  components: {
    Navbar,
    BottomArticle,
    ArticleComment,
    Share,
    QRCode,
  },
  data() {
    return {
      userInfo: {},
      articleInfo: {},
      articleId: 0,
      staticUrl: WebStaticUrl.static_url,
      activeNames: [],
      menus: [],
      categories: [],
      commentCount: 8,
      commentStart: 0,
      commentList: [],
      commentNumTotal: 0,
      showShareBtn: false,
      showQRCode: false,
      qrcode:{},
      shareTitle:"文章分享二维码"
    };
  },
  beforeRouteEnter(to, from, next) {
    if (to.params.articleInfo) {
      next((vm) => {
        vm.articleInfo = to.params.articleInfo;
        vm.articleId = to.params.id;
        vm.getArticleComment(to.params.id, vm.commentCount, vm.commentStart);
        vm.getUserInfo();
      });
    } else if (!to.params.articleInfo) {
      if (to.params.id > 0) {
        next((vm) => {
          vm.articleId = to.params.id;
          vm.getUserInfo();
        });
      } else {
        next({
          path: from.path,
        });
      }
    } else {
      next({
        path: from.path,
      });
    }
  },
  mounted() {
    this.getArticleInfo();
  },
  methods: {
    getUserInfo() {
      var userInfoJson = localStorage.getItem("consumerInfo");
      if (userInfoJson) {
        this.userInfo = JSON.parse(userInfoJson);
      }
    },
    getArticleInfo() {
      if (this.articleInfo.Id > 0) {
        this.getArticleComment(
          this.articleInfo.Id,
          this.commentCount,
          this.commentStart
        );
        return;
      } else if (
        this.articleInfo.Id > 0 == false &&
        this.articleId > 0 == true
      ) {
        this.getArticleById();
      } else {
        this.$router.back();
      }
    },
    async getArticleById() {
      console.log("从后台获取文章了....");
      var id = Number(this.articleId);
      const result = await reqArticleById(id);
      if (result.code === 0 && result.data.Id > 0) {
        this.articleInfo = result.data;
        console.log(this.articleInfo);
        this.getArticleComment(
          this.articleInfo.Id,
          this.commentCount,
          this.commentStart
        );
      } else {
        this.articleId -= 1;
        this.getArticleById();
        this.$toast({
          type: "fail",
          duration: 3000,
          message: "没有更多文章了",
        });
      }
    },
    changeInfo(data) {
      console.log(this.articleInfo);
      if (data.length > 0) {
        this.getMenusInfo();
      }
    },
    getMenusInfo() {
      if (this.menus.length == 0) {
        var menuArray = this.$store.getters.menu;
        if (menuArray.length > 0) {
          this.menus = menuArray;
        }
      }
      if (this.categories.length == 0) {
        var category = this.$store.getters.category;
        if (category.length > 0) {
          this.categories = category;
        }
      }
    },
    changeArticleId(changeMessage) {
      if (changeMessage == "pageBack") {
        this.$router.push("/");
      } else if (changeMessage == "pageUp") {
        this.articleId = Number(this.articleId);
        this.articleId -= 1;
        if (this.articleId >= 1) {
          this.articleInfo = {};
          this.getArticleInfo();
        } else {
          this.articleId += 1;
        }
      } else if (changeMessage == "pageDown") {
        this.articleId = Number(this.articleId);
        this.articleId += 1;
        this.articleInfo = {};
        this.getArticleInfo();
      } else {
        this.$toast({
          type: "fail",
          duration: 2000,
          message: "非正常浏览",
        });
        this.$router.push("/");
      }
    },
    toCommentTop() {
      document.getElementById("commentTop").scrollIntoView({
        behavior: "smooth",
        block: "center",
      });
    },
    submitComment(value) {
      this.addCommentOne(value);
    },
    // 添加一级评论
    async addCommentOne(value) {
      const result = await reqAddArticleComment(
        this.articleInfo.Id,
        this.userInfo.id,
        value,
        this.userInfo.nickname,
        this.userInfo.icon
      );
      if (result.code === 0) {
        this.getArticleComment(
          this.articleInfo.Id,
          this.commentCount,
          this.commentStart
        );
      }
    },
    submitCommentChild(value, comment_id) {
      console.log("value=", value);
      console.log("comment_id =", comment_id);
      this.addCommentChild(comment_id, value);
    },
    // 添加二级评论
    async addCommentChild(comment_id, content) {
      const result = await reqAddArticleCommentChild(
        comment_id,
        this.userInfo.id,
        content,
        this.userInfo.nickname,
        this.userInfo.icon
      );
      if (result.code === 0) {
        this.getArticleComment(
          this.articleInfo.Id,
          this.commentCount,
          this.commentStart
        );
      }
    },
    // 获取文章评论
    async getArticleComment(article_id, count, start) {
      const result = await reqGetArticleComment(article_id, count, start);
      if (result.code === 0) {
        this.commentList = result.data;
        console.log(this.commentList);
      }
    },
    commentNum(value) {
      console.log(value);
      this.commentNumTotal = value;
    },
    goodJob(articleId, star) {
      this.editGoodOrBadJob(articleId, star, 0, "good");
    },
    badJob(articleId, tread) {
      this.editGoodOrBadJob(articleId, 0, tread, "bad");
    },
    async editGoodOrBadJob(articleId, star, tread, jobTag) {
      const result = await reqGoodOrBadArticle(articleId, star, tread, jobTag);
      if (result.code === 0) {
        if (jobTag == "good") {
          this.articleInfo.star += 1;
        } else if (jobTag == "bad") {
          this.articleInfo.tread += 1;
        }
      } else {
        this.$toast({
          type: "fail",
          message: result.msg,
        });
      }
    },
    changeShowShareBtn() {
      this.showShareBtn = true;
    },
    handleCloseShare() {
      this.showShareBtn = false;
    },
    createQRCode(shareUrl) {
      this.showQRCode = true;
        setTimeout(() => {
        this.$refs.qrCodeUrl.innerHTML = ""
        this.qrcode = new QRCode(this.$refs.qrCodeUrl, {
          text: shareUrl,
          width: 100,
          height: 100,
          colorDark: "#000000",
          colorLight: "#ffffff",
          correctLevel: QRCode.CorrectLevel.H,
        });
      }, 1000);
    },
    getBrowser() {
      var userAgent = navigator.userAgent;
      var isOpera = userAgent.indexOf("Opera") > -1;
      if (userAgent.indexOf("FireFox") > -1) {
        return "FireFox"
      };
      if (userAgent.indexOf("Chrome") > -1) {
        return "Chrome"
      };
      if (userAgent.indexOf("Opera") > -1) {
        return "Opera"
      };
      if (userAgent.indexOf("Safari") > -1) {
        return "Safari"
      };
      if (userAgent.indexOf("compatible") > -1 && userAgent.indexOf("MSIE") > -1 && !isOpera) {
        return "IE"
      };
      if (userAgent.indexOf("Trident") > -1) {
        return "Edge"
      };
    },
    checkBrowser() {
       var browserTag = this.getBrowser()
      console.log(browserTag);
      if (browserTag == "IE" && browserTag == "Edge") {
        this.IEDownloadQRCode()
      }
    },
    saveQRCode() {
      this.CFDownloadQRCode()     
    },
    // Chrome firefox 下载二维码
    CFDownloadQRCode() {
      html2canvas(this.$refs.qrCodeUrl)
      .then(canvas => {
        var imgData = canvas.toDataURL("image/png");
        var link = document.createElement('a')
        link.href = imgData
        link.setAttribute('download','article_'+this.articleInfo.Id+'.png')
        link.style.display = 'none'
        document.body.appendChild(link)
        link.click()
        link.removeAttribute('download')
        document.body.removeChild(link)
      })
    },
    // IE 下载二维码
    IEDownloadQRCode() {
      document.write("暂不支持IE浏览器")
    },
    shareWeiBo(url) {
      console.log(url);
    }
  },
};
</script>

<style scoped>
.articleTopView {
  widows: 100%;
}
.articleTopView div video {
  width: 100%;
}
.articleTopView div img {
  width: 100%;
}
.articleTitle {
  display: flex;
  align-items: center;
  overflow: hidden;
}
.titleIcon1 {
  display: flex;
  color: white;
  background-color: rgba(237, 25, 65, 0.8);
  height: 5.556vw;
  align-items: center;
  margin-right: 1.389vw;
  border-radius: 2.778vw;
  width: 12.5vw;
  justify-content: center;
}
.titleIcon2 {
  display: flex;
  color: white;
  background-color: rgba(34, 139, 34, 0.8);
  height: 5.556vw;
  align-items: center;
  margin-right: 1.389vw;
  border-radius: 2.778vw;
  width: 12.5vw;
  justify-content: center;
}
.titleIcon1 span {
  padding: 0 1.389vw;
  font-size: 3.056vw;
}
.titleIcon2 span {
  padding: 0 1.389vw;
  font-size: 3.056vw;
}
.articleDesc {
  display: flex;
  flex-wrap: wrap;
}
.articleDesc div {
  display: flex;
  width: 100%;
  padding: 0.833vw;
  overflow: hidden;
}
.articleDesc div span:nth-child(1) {
  font-weight: bolder;
  text-align: justify;
  text-align-last: justify;
  padding-right: 3.167vw;
  width: 14.278vw;
}
.articleDesc div span a {
  font-style: italic;
  font-family: Verdana, Geneva, Tahoma, sans-serif;
  color: #4169e1;
}
.articleCount {
  background-color: white;
  display: flex;
  height: 45px;
}
.articleCount div {
  display: flex;
  align-items: center;
  color: #808080;
  position: relative;
}
.articleCount div:hover {
  display: flex;
  align-items: center;
  color: #4169e1;
  position: relative;
}
.articleCount div:nth-child(1) {
  margin-left: 4.167vw;
}
.articleCount div:nth-child(5) {
  margin-left: 16.278vw;
}
.articleCount span {
  font-size: 3.333vw;
  margin-right: 2.222vw;
}
.articleContent {
  background-color: white;
  padding: 4.167vw;
}
.articleBanner {
  background-color: white;
  width: 100%;
}
.articleBanner {
  background-color: white;
}
.articleBanner img {
  width: 100%;
  height: 39.444vw;
}
.qrcode {
  display: flex;
  justify-content: center;
  margin:6.944vw 0;
}
.qrcodeImg {
  display: flex;
  height: 100%;
  padding:1.667vw;
  justify-content: center;
  background-color: #fff;
}

</style>
