<template>
  <div>
    <Navbar :userInfo="userInfo"></Navbar>
    <Banner></Banner>
    <PersonalInfo :userInfo="userInfo"></PersonalInfo>
  </div>
</template>

<script>
import Navbar from "@/components/common/Navbar.vue";
import Banner from "@/components/common/Banner.vue";
import PersonalInfo from "@/components/personalComponents/PersonalInfo.vue";
export default {
  name: "center",
  components: {
    Navbar,
    Banner,
    PersonalInfo,
  },
  mounted() {
    this.getUserInfo()
  },
  data() {
    return {
      userInfo: {},
    };
  },
  methods: {
    getUserInfo() {
       var consumerInfoJson = localStorage.getItem("consumerInfo")
       if (consumerInfoJson) {
         this.userInfo = JSON.parse(consumerInfoJson)
       }
    }
  },
};
</script>

<style scoped></style>
