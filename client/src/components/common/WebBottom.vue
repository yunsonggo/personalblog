<template>
  <div class="webBottom">
    <div class="contentBody">
      <div class="webContent">
        <p>{{ websiteInfo.website_notice }}</p>
      </div>
      <div class="webContent">
        <p>{{ websiteInfo.website_keyword }}</p>
      </div>
      <div class="webContent">
        <p>{{ websiteInfo.website_desc }}</p>
      </div>
      <div class="webContent">
        <span>{{ websiteInfo.website_licence_one }}</span>
        <span>
          <img src="@/assets/wangjing.png" alt="" />
          <a href="http://www.beian.gov.cn/portal/index" target="_blank">{{
            websiteInfo.website_licence_two
          }}</a>
        </span>
      </div>
      <div class="webContent">
        <span>网站制作:{{ websiteInfo.website_nickname }}</span>
        <span
          >{{ websiteInfo.website_copy }}{{ websiteInfo.website_title }}</span
        >
      </div>
      <div class="webContent">
        <span>email:{{ websiteInfo.website_email }}</span>
        <span>QQ:{{ websiteInfo.website_qq }}</span>
      </div>
      <div class="webContent">
        <span>地址:{{ websiteInfo.website_address }}</span>
        <span>版本:{{ websiteInfo.website_version }}</span>
      </div>
    </div>
    <div class="links">
      <div class="linkItem" v-for="(link, index) in LinksInfo" :key="index">
        <a :href="link.link_url" target="_blank">
          <img :src="websiteInfo.website_static_url + '/' + link.icon" alt="" />
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import { reqWebsiteInfo, reqLinksInfo } from "@/api";
export default {
  name: "webBottom",
  data() {
    return {
      websiteInfo: {},
      LinksInfo: [],
    };
  },
  mounted() {
    this.getWebsite();
    this.getLinksInfo();
  },
  methods: {
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
    async getLinksInfo() {
      const result = await reqLinksInfo();
      if (result.code === 0) {
        this.LinksInfo = result.data;
        console.log(this.LinksInfo);
      } else {
        this.$toast({
          type: "fail",
          message: "获取友情链接数据失败",
        });
      }
    },
  },
};
</script>

<style scoped>
.webBottom {
  margin: 3.167vw 0;
  background-color: #f4f4f4;
  height: 77.778vw;
  clear: both;
}
.contentBody {
    padding-bottom: 2.778vw;
}
.webContent {
  background-color: #f4f4f4;
  color: #808080;
  font-size: 3.333vw;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}
.webContent p {
  padding: 0 4.889vw;
  text-align: center;
  line-height: 6.944vw;
  font-weight: bold;
}
.webContent span {
  width: 40%;
  display: flex;
  text-align: center;
  margin: 0 1.389vw;
  padding-top: 1.389vw;
}
.webContent span a {
  color: #808080;
}
.links {
  background-color: #f4f4f4;
  padding: 10px 10px;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
.linkItem {
  display: flex;
  margin: 0 0.833vw;
  flex-direction: column;
  width: 30%;
}
.linkItem img {
  width: 27.778vw;
  height: 12.5vw;
  margin-bottom: 3px;
}

</style>
