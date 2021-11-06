<template>
  <div>
    <van-button 
    :type="type" 
    :block="block" 
    :size="size" 
    :text="btnName"
    @touchstart="registerClick"
    >
    </van-button>
  </div>
</template>

<script>
export default {
  name: "LoginBtn",
  props: {
    btnName: {
      type: String,
      default: "按钮",
    },
    block: {
      type: Boolean,
      default: false,
    },
    type: {
      type: String,
      default: "primary",
    },
    size: {
      type: String,
      default: "small",
    },
  },
  methods:{
    registerClick() {
      this.$emit('registerSubmit')
    },
  }
};
</script>

<style scoped></style>
