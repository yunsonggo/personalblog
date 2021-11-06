<template>
  <div>
    <div class="loginBottom">
      <div></div>
      <div>
        <van-button plain type="primary" size="small" :to="toPath">{{
          toName
        }}</van-button>
      </div>
      <div></div>
    </div>
    <div class="bottomMsg">
        <p>姓名,邮箱,电话注册以后不能修改</p>
        <p>如需修改请联系管理员或者注册新账号</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "LoginBottom",
  props: {
    toPath: {
      type: String,
      default: "",
    },
    toName: {
      type: String,
      default: "",
    },
  },
};
</script>

<style scoped>
.loginBottom {
  display: flex;
}
.loginBottom div {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}
.bottomMsg {
  display: flex;
  margin: 20px 0;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.bottomMsg p {
  padding: 3px;
  flex: 1;
  font-size: 3.333vw;
  color: #646566;
}
</style>
