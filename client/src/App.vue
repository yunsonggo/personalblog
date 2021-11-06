<template>
  <keep-alive>
    <router-view />
  </keep-alive>
</template>

<script>
export default {
  mounted() {
    if (!this._isMobile()) {
      this.$dialog.alert({
        message: "尚未开发PC端，请按F12后按CTRL+SHIFT+M",
      });
    } 
  },
  methods: {
    _isMobile() {
      let flag = navigator.userAgent.match(
        /(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i
      );
      return flag;
    },
  },
};
</script>

<style>
html,
body {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  background-color: #f4f4f4;
  height: 100%;
}
p {
  margin: 0;
}
</style>
