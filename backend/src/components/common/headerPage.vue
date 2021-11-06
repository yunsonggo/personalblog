<template>
  <div class="header">
    <div>
      <img src="@/assets/logo.jpg" alt="" />
    </div>
    <div>
      <h3>后台管理系统</h3>
    </div>
    <div>
      <van-popover
        v-model="showPopover"
        trigger="click"
        :actions="actions"
        @select="onSelect"
      >
        <template #reference>
          <van-icon name="user-o"></van-icon>
          <span>
            {{ managerInfo.manager_name }}
          </span>
        </template>
      </van-popover>
    </div>
  </div>
</template>

<script>
export default {
  name: "headerPage",
  props: {
    managerInfo: {
      type: Object,
      default: () => ({}),
    },
  },
  data() {
      return {
        showPopover:false,
        actions:[{text:'个人中心'},{text:'退出登录'}]
      }
  },
  methods:{
      onSelect(value){
        if (value.text === "个人中心") {
          console.log('1');
        } else if (value.text === "退出登录") {
          localStorage.removeItem("adminToken")
          this.$message.info("退出成功")
          this.$router.push("/manager/login")
        }
      }
  }
};
</script>

<style scoped>
.header {
  display: flex;
  background-color: #4285e4;
  width: 100%;
  height: 65px;
}
.header div:nth-child(1) {
  display: flex;
  width: 20%;
  justify-content: center;
}
.header div:nth-child(2) {
  display: flex;
  width: 60%;
  justify-content: center;
  align-items: center;
  color: white;
}
.header div:nth-child(3) {
  display: flex;
  width: 20%;
  justify-content: center;
  align-items: center;
  color: white;
}
.header div:nth-child(3) span {
  margin-left: 5px;
}
</style>
