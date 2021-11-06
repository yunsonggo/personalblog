<template>
  <div>
    <div class="loginTop">
      <div @touchstart="goHome">
        <img
          :src="websiteInfo.website_static_url + '/' + websiteInfo.website_logo"
          alt=""
        />
      </div>
      <div>{{ middleTop }}</div>
      <div>
        <slot name="right"></slot>
      </div>
    </div>
  </div>
</template>

<script>
import { reqWebsiteInfo } from "@/api";
export default {
  name: "LoginTop",
  props: {
    middleTop: {
      type: String,
      default: "新用户注册",
    },
  },
  data() {
    return {
      websiteInfo: {},
    };
  },
  // beforeRouteEnter(to, from, next) {
  //   next((vm) => {
  //     vm.getWebsite();
  //   });
  // },
  mounted() {
    this.getWebsite()
  },
  methods: {
    goHome() {
      this.$router.push("/home/index");
    },
    getWebsite() {
      if (this.$store.getters.website.Id) {
        this.websiteInfo = this.$store.getters.website;
      } else {
        this.getWebsiteInfo();
      }
    },
    async getWebsiteInfo() {
      const result = await reqWebsiteInfo();
      if (result.code === 0) {
        this.websiteInfo = result.data;
        this.$store.dispatch("setWebsite", result.data);
      }
    },
  },
};
</script>

<style scoped>
.loginTop {
  height: 11.111vw;
  background-color: white;
  display: flex;
}
.loginTop div {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 3.896vw;
}
.loginTop div:nth-child(1) {
  width: 30.556vw;
  display: flex;
  justify-content: center;
  align-items: center;
}
.loginTop img {
  width: 95%;
  height: 95%;
}
</style>
