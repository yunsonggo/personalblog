<template>
  <div>
    <van-field
      v-model="code"
      center
      clearable
      :label="label"
      :placeholder="placeholder"
      :required="required"
      :error-message="errorMessage"
    >
      <template #button>
        <van-button 
        size="small" 
        type="primary" 
        @touchstart="getCode"
        :disabled="disable"
        >
        {{btnMsg}}
        </van-button>
      </template>
    </van-field>
    <div v-if="showCode" style="margin:10px;font-size:14px;"> {{showCode}} </div>
  </div>
</template>

<script>
import { reqCodeByEmail,reqCodeByPhone } from "@/api";
export default {
  name: "LoginCode",
  props: {
    label: {
      type: String,
      default: "发送",
    },
    placeholder: {
      type: String,
      default: "请输入...",
    },
    type: {
      type: String,
      default: "text",
    },
    requestInfo: {
      type: String,
      default: "",
    },
    required: {
      type: Boolean,
      default: false,
    },
    errorInfo: {
      type: String,
      default: "",
    },
    rule: {
      type: String,
      default: "",
    },
    codeMode: {
      type: Boolean,
      default: true,
    },
    clearErrorInfo: {
      type: Boolean,
      default: false,
    },
  },
  watch: {
    code() {
      this.regxHandler();
    },
    clearErrorInfo() {
      this.code = "",
      this.errorMessage = ""
    },
  },
  data() {
    return {
      code: "",
      errorMessage: "",
      disable:false,
      btnMsg:"获取验证码",
      disableTime:60,
      showCode:"",
    };
  },
  methods: {
    regxHandler() {
      const regRule = new RegExp(this.rule);
      if (regRule.test(this.code)) {
        this.errorMessage = "";
        this.$emit("inputChange", this.code);
      } else {
        this.errorMessage = this.errorInfo;
      }
    },
    getCode() {
      if (!this.codeMode) {
        if (this.disable == false) {
          this.getCodeByEmail();
        } 
      } else {
        if (this.disable == false) {
           this.getCodeByPhone();
        }
      }
    },
    async getCodeByEmail() {
      if (this.requestInfo != "" ) {
        const result = await reqCodeByEmail(this.requestInfo)
        if (result.code === 0) {
          this.$notify({ type: 'success', message: '发送成功!'});
          this.timer()
        } else {
          this.$notify({ type: 'warning', message: '发送验证码失败,请重试'});
        }
      } else {
        this.$notify({ type: 'warning', message: '请输入邮箱地址'});
      }
    },
    async getCodeByPhone() {
      if (this.requestInfo != "") {
        const result = await reqCodeByPhone(this.requestInfo,"id","secret")
        if (result.code === 0) {
          this.showCode = result.data.code
          this.$notify({ type: 'success', message: '发送成功!'});
          this.$emit("getBizId",result.data.bizId)
          this.timer2()
        } else {
          this.$notify({ type: 'warning', message: '发送验证码失败,请重试'});
        }
      } else {
        this.$notify({ type: 'warning', message: '请输入手机号码' });
      }
    },
    timer() {
      this.disable = true
      this.timerSet = setInterval(() => {
        this.disableTime --
        if (this.disableTime === 0) {
          this.disableTime = 60
          this.disable = false
          this.btnMsg = "获取验证码"
          clearInterval(this.timerSet)
        } else {
          this.btnMsg = "倒计时:(" + this.disableTime + ")"
        }
      },1000)
    },
    timer2() {
      this.disable = true
      this.timerSet2 = setInterval(() => {
        this.disableTime --
        if (this.disableTime === 0) {
          this.disableTime = 60
          this.disable = false
          this.btnMsg = "获取验证码"
          clearInterval(this.timerSet2)
        } else {
          this.btnMsg = "倒计时:(" + this.disableTime + ")"
        }
      },1000)
    },
  },
};
</script>

<style scoped></style>
