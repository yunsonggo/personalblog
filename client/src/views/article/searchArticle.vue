<template>
  <div class="searchStyle" ref="searchBody">
    <LoginTop :middleTop="middleTopText">
      <div slot="right">
        <div class="userInfo">
          <van-popover
            v-model="showPopover"
            trigger="click"
            :actions="actions"
            @select="onSelect"
            @open="checkLogin"
          >
            <template #reference>
              <img
                v-if="userInfo.icon"
                :src="staticUrl + userInfo.icon"
                alt=""
              />
              <img v-else src="@/assets/default_avatar.jpg" alt="" />
            </template>
          </van-popover>
          <p>下载App</p>
        </div>
      </div>
    </LoginTop>
    <div>
      <van-search
        v-model="value"
        placeholder="点击键盘上的搜索/回车按钮后搜索"
        @search="onSearch"
        @clear="clearInput"
      />
    </div>
    <MenuArticle
      :theArticles="articleList"
      :articleType="0"
      :articleStart="start"
      @getMoreArticle="getMore"
      :resLoading="resLoading"
      :resFinished="resFinished"
    ></MenuArticle>
    <div id="push" v-if="showPush">
      <van-empty description="搜索文章..." />
    </div>
  </div>
</template>

<script>
import LoginTop from "@/components/common/LoginTop.vue";
import WebStaticUrl from "@/components/common/WebStaticUrl";
import { reqSearchArticle } from "@/api";
import MenuArticle from "@/views/article/menuArticle.vue";
export default {
  name: "searchArticle",
  components: {
    LoginTop,
    MenuArticle,
  },
  beforeRouteEnter(to, from, next) {
    if (to.params.userInfo) {
      next((vm) => {
        vm.userInfo = to.params.userInfo;
      });
    } else {
      next((vm) => {
        vm.getUserInfo();
      });
    }
  },
  computed: {
    actions() {
      if (this.isLogin) {
        return [
          { text: "个人中心", icon: "manager-o" },
          { text: "退出登录", icon: "close" },
        ];
      } else {
        return [{ text: "登录", icon: "manager-o" }];
      }
    },
  },
  watch: {
    value() {
      if (this.value == "") {
        this.clearInput();
      } 
    },
    articleList() {
      if (this.articleList.length > 0 || this.value !== "") {
        this.showPush = false
      } else {
        this.showPush = true
      }
    }
  },
  data() {
    return {
      middleTopText: "搜索文章",
      staticUrl: WebStaticUrl.static_url,
      showPopover: false,
      isLogin: false,
      userInfo: {},
      value: "",
      count: 6,
      start: 1,
      articleList: [],
      resLoading: false,
      resFinished: false,
      showPush:true,
    };
  },
  methods: {
    getUserInfo() {
      var userInfoJson = localStorage.getItem("consumerInfo");
      if (userInfoJson) {
        this.userInfo = JSON.parse(userInfoJson);
      }
    },
    goHome() {
      if (this.$route.path != "/home/index") {
        this.$router.push("/home/index");
      }
    },
    goCenter() {
      if (this.$route.path != "/personal/center") {
        this.$router.push("/personal/center");
      }
    },
    checkLogin() {
      if (localStorage.getItem("token")) {
        this.isLogin = true;
      } else {
        this.isLogin = false;
      }
    },
    onSelect(action) {
      console.log(action.text);
      if (action.text === "个人中心") {
        this.goCenter();
      } else if (action.text === "登录") {
        this.$router.push("/personal/login");
      } else {
        localStorage.removeItem("token");
        localStorage.removeItem("consumerInfo");
        this.$router.push("/");
        this.$toast({
          type: "success",
          message: "退出成功!",
        });
      }
    },
    clearInput() {
      this.resLoading = false;
      this.resFinished = false;
      this.articleList = [];
      this.start = 1;
    },
    onSearch() {
      if (this.value == "") {
        this.clearInput();
        this.$toast({
          type: "fail",
          message: "没有收入关键词...",
        });
        return;
      } else {
        this.getArticleList("0", this.start);
      }
    },
    async getArticleList(articleType, start) {
      this.value = String(this.value);
      this.value = this.value.trim();
      const result = await reqSearchArticle(this.value, this.count, start);
      if (result.code === 0 && result.data !== null) {
        this.resLoading = false;
        this.resFinished = false;
        this.articleList.push(...result.data);
      } else {
        this.resLoading = false;
        this.resFinished = true;
        this.$toast({
          type: "fail",
          message: "没有搜索到更多内容...",
        });
      }
      if (!this.articleList) {
        this.$toast({
          type: "fail",
          message: "没有搜索到更多内容...",
        });
      }
    },
    getMore(articleType, newStart) {
      this.resLoading = true;
      this.start = newStart;
      setTimeout(() => {
        this.getArticleList(articleType, newStart);
      }, 1000);
    },
  },
};
</script>

<style scoped>
.searchStyle {
  height: 100%;
  height: auto !important;
  height: 100%; /*IE6不识别min-height*/
  margin: 0 auto;
  background-color: white;
}
.userInfo img {
  width: 6.944vw;
  height: 6.944vw;
}
.userInfo p {
  padding: 5px 10px;
  margin: 0 8px;
  background-color: #fb7299;
  color: white;
  font-size: 13px;
  border-radius: 3px;
}
#push {
  height: calc(100vh - 340px) ;
}
</style>
