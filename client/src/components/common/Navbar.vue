<template>
  <div class="navbar">
    <div @touchstart="goHome">
      <img :src="websiteInfo.website_static_url + '/' + websiteInfo.website_logo" alt="" />
    </div>
    <div @touchstart="goSearch">
      <div><van-icon name="search" class="iconNav" />请输入搜索内容</div>
    </div>
    <div>
      <van-popover
        v-model="showPopover"
        trigger="click"
        :actions="actions"
        @select="onSelect"
        @open="checkLogin"
      >
        <template #reference>
          <img v-if="userInfo.icon" :src="websiteInfo.website_static_url + userInfo.icon" alt="" />
          <img v-else src="@/assets/default_avatar.jpg" alt="" />
        </template>
      </van-popover>
      <p>下载App</p>
    </div>
  </div>
</template>

<script>
import {reqWebsiteInfo} from "@/api"
export default {
  name: "Navabar",
  props: {
    userInfo: {
      type: Object,
      default: () => ({}),
    },
  },
  data() {
    return {
      showPopover: false,
      isLogin: false,
      websiteInfo:{},
    };
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
  mounted() {
    this.getWebsiteInfo()
  },
  methods: {
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
        this.$router.push("/personal/login")
      }else {
        localStorage.removeItem("token");
        localStorage.removeItem("consumerInfo");
        this.$router.push("/");
        this.$toast({
          type: "success",
          message: "退出成功!",
        });
      }
    },
    goSearch() {
      this.$router.push({
        name:"articleSearch",
        params:{
          userInfo:this.userInfo
        }
        })
    },
    async getWebsiteInfo() {
      const result = await reqWebsiteInfo()
      if (result.code === 0) {
        this.websiteInfo = result.data
        this.$store.dispatch("setWebsite", result.data);
      }
      console.log(this.websiteInfo);
    }
  },
};
</script>

<style scoped>
.navbar {
  height: 45px;
  background-color: white;
  display: flex;
}
.navbar div:nth-child(1) {
  width: 30.556vw;
  display: flex;
  justify-content: center;
  align-items: center;
}
.navbar div:nth-child(1) img {
  width: 100%;
  height: 100%;
}
.navbar div:nth-child(2) {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 1.389vw;
}
.navbar div:nth-child(2) div {
  background-color: #e0e0e0;
  height: 6.667vw;
  border-radius: 3px;
  position: relative;
  width: 100%;
  font-size: 13px;
  color: #aaa;
}
.navbar div:nth-child(2) div .iconNav {
  color: #aaa;
  position: absolute;
  left: 2.778vw;
  top: 50%;
  transform: translate(0, -50%);
}
.navbar div:nth-child(3) {
  display: flex;
  justify-content: center;
  align-items: center;
}
.navbar div:nth-child(3) img {
  width: 6.944vw;
  height: 6.944vw;
}
.navbar div:nth-child(3) p {
  padding: 5px 10px;
  margin: 0 8px;
  background-color: #fb7299;
  color: white;
  font-size: 13px;
  border-radius: 3px;
}
</style>
