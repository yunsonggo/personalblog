<template>
  <div class="homeIndex">
    <HeaderPage :managerInfo="managerInfo"></HeaderPage>
    <div class="homeBody">
      <Tabbar :menusInfo="menusInfo" @breadParam="getBreadParam"></Tabbar>
      <div class="homeContent">
        <div class="bread">
          <Bread :breadParams="breadParams"></Bread>
        </div>
        <div>
          <router-view></router-view>
        </div>
      </div>
    </div>
    <div class="homeFooter">
      <span>后台管理系统</span>
    </div>
  </div>
</template>

<script>
import HeaderPage from "@/components/common/headerPage.vue";
import Tabbar from "@/components/common/tabbar.vue";
import Bread from "@/components/common/bread.vue";
import { reqMenuList ,reqWebsiteInfo} from "@/api";
export default {
  name: "home",
  components: {
    HeaderPage,
    Tabbar,
    Bread,
  },
  data() {
    return {
      managerInfo: {},
      menusInfo: [],
      breadParams:[],
    };
  },
  mounted() {
    this.getManagerInfo();
    this.getMenuList();
    this.getWebsite();
  },
  methods: {
    getManagerInfo() {
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
    },
    async getMenuList() {
      const result = await reqMenuList(this.managerInfo.manager_password);
      if (result.code === 0) {
        this.menusInfo = result.data;
      } else {
        localStorage.removeItem("adminToken")
        this.$message.error("请重新登录")
        this.$router.push("/manager/login")
      }
    },
    async getWebsite() {
      const result = await reqWebsiteInfo();
      if (result.code === 0) {
        localStorage.setItem("staticUrl",result.data.website_static_url) 
      }
    },
    getBreadParam(params) {
      this.breadParams = params
    }
  },
};
</script>

<style>
.homeIndex {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  background-color: white;
}
.homeBody {
  width: 90%;
  margin-top: 5px;
  display: flex;
  height: auto;
  min-height: 800px;
}
.homeContent {
  width: 100%;
  margin-left: 15px;
}
.bread {
  display: flex;
  height: 45px;
  align-items: center;
}
.homeFooter {
  height: 100px;
  background-color: #4285e4;
  width: 100%;
  display: flex;
  align-items: center;
  color: white;
  justify-content: center;
}

</style>
