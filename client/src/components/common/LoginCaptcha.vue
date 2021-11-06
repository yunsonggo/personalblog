<template>
  <div>
    <van-field
    v-model="captcha"
    :label="label"
    :required="required"
    :placeholder="placeholder"
    clearable
    center
    >
      <template #button>
        <van-image
        width="100%"
        height="9.722vw"
        :src="imgCaptcha"
        fit="cover"
        @click="handleNewCaptcha"
        ></van-image>
      </template>
    </van-field>
  </div>
</template>

<script>
import {reqCaptcha} from "@/api"
export default {
  name: "LoginCaptcha",
  props:{
      label:{
          type:String,
          default:""
      },
      required:{
          type:Boolean,
          default:false,
      },
      placeholder:{
          type:String,
          default:""
      }
  },
  watch:{
      captcha() {
          this.$emit("inputChange",this.captcha)
      }
  },
  data() {
      return {
          captcha:"",
          imgCaptcha:"",
          captchaId:""
      }
  },
  mounted() {
      this.getCaptcha()
  },
  methods:{
      async getCaptcha() {
          const result = await reqCaptcha()
          if (result.code === 0) {
              this.imgCaptcha = result.data.B64s
              this.captchaId = result.data.Id
              this.$emit("getCaptchaId",this.captchaId)
          }
      },
      handleNewCaptcha() {
          this.getCaptcha()
      }
  }
};
</script>

<style scoped></style>
