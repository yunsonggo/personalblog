<template>
  <div>
    <NavBar :userInfo="userInfo"></NavBar>
    <Tabs></Tabs>
  </div>
</template>

<script>
import NavBar from "@/components/common/Navbar.vue";
import Tabs from "@/components/common/Tabs.vue";
export default {
  name: "homePage",
  components: {
    NavBar,
    Tabs,
  },
  mounted() {
    this.getUserInfo();
  },
  data() {
    return {
      userInfo: {},
    };
  },
  methods: {
    getUserInfo() {
      var userInfoJson = localStorage.getItem("consumerInfo");
      if (userInfoJson) {
        this.userInfo = JSON.parse(userInfoJson);
      }
    },
  },
};
</script>

<style scoped>

</style>
