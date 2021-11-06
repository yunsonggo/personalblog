<template>
  <div>
    <LoginTop :middleTop="middleTopText">
      <div
        slot="right"
        style="font-size:3.333vw;color:#07c160"
        @touchstart="changeSigner"
      >
        {{ rightTop }}
      </div>
    </LoginTop>
    <LoginText
      style="margin:4.167vw 0;"
      v-if="!changeSignerBtn"
      label="邮箱:"
      placeholder="请输入邮箱:"
      type="email"
      rule="^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$"
      required
      @inputChange="(res) => (email = res)"
      errorInfo="请输入正确的邮箱地址"
      :clearErrorInfo="clearErrorInfo"
    />
    <LoginText
      style="margin:4.167vw 0;"
      v-else
      label="手机:"
      placeholder="请输入手机号码:"
      type="tel"
      rule="^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$"
      required
      @inputChange="(res) => (phone = res)"
      errorInfo="请输入正确的手机号码"
      :clearErrorInfo="clearErrorInfo"
    />
    <LoginText
      v-if="!changeSignerBtn"
      label="密码:"
      placeholder="请输入密码:"
      type="password"
      rule="^[a-zA-Z]([-_a-zA-Z0-9]{7,19})+$"
      required
      @inputChange="(res) => (password = res)"
      errorInfo="密码为包含字母的7-19位(字母,数字,_,-)"
      :clearErrorInfo="clearErrorInfo"
    />
    <LoginCode
      v-else
      label="验证码"
      placeholder="请输入验证码"
      type="text"
      rule="^\d{4}$"
      required
      :requestInfo="requestInfo"
      @inputChange="(res) => (code = res)"
      errorInfo="验证码为四位数字"
      :codeMode="phoneCode"
      :clearErrorInfo="clearErrorInfo"
      @getBizId="getBizId"
    ></LoginCode>
    <LoginCaptcha
      v-if="!changeSignerBtn"
      label="验证码:"
      required
      placeholder="请输入验证码"
      @inputChange="(res) => (captcha = res)"
      @getCaptchaId="getCaptchaId"
    ></LoginCaptcha>
    <LoginBtn
      style="margin:4.167vw 0;"
      btnName="立即登录"
      block
      size="large"
      @registerSubmit="submit"
    ></LoginBtn>
    <LoginBottom toPath="register" toName="新用户注册"></LoginBottom>
  </div>
</template>

<script>
import LoginTop from "@/components/common/LoginTop.vue";
import LoginText from "@/components/common/LoginText.vue";
import LoginCode from "@/components/common/LoginCode.vue";
import LoginCaptcha from "@/components/common/LoginCaptcha.vue";
import LoginBtn from "@/components/common/LoginBtn.vue";
import LoginBottom from "@/components/common/LoginBottom.vue";
import { reqLoginByEmail, reqLoginByPhone } from "@/api";
export default {
  name: "login",
  components: {
    LoginTop,
    LoginText,
    LoginCode,
    LoginCaptcha,
    LoginBtn,
    LoginBottom,
  },
  watch: {
    phone() {
      this.requestInfo = this.phone;
    },
  },
  data() {
    return {
      middleTopText: "用户登录",
      changeSignerBtn: false,
      rightTop: "手机登录",
      requestUrl: "email",
      clearErrorInfo: false,
      phoneCode: true,
      bizId: "",
      phone: "none",
      email: "none",
      password: "",
      code: "",
      captcha: "",
      captchaId: "",
      requestInfo: "",
    };
  },
  methods: {
    changeSigner() {
      this.changeSignerBtn = !this.changeSignerBtn;
      this.email = "";
      this.phone = "";
      this.password = "";
      this.code = "";
      if (this.changeSignerBtn) {
        this.rightTop = "邮箱登录";
        this.requestUrl = "phone";
      } else {
        this.rightTop = "手机登录";
        this.requestUrl = "email";
      }
      this.clearErrorInfo = this.changeSignerBtn;
    },
    getBizId(data) {
      this.bizId = data;
    },
    getCaptchaId(data) {
      this.captchaId = data;
    },
    submit() {
      // 邮箱登录
      if (this.requestUrl == "email") {
        if (this.captcha == "") {
          this.$toast({
            type: "fail",
            duration: 4000,
            message: "请输入验证码",
          });
          return;
        }
        if (this.email == "" || this.email == "none") {
          this.$toast({
            type: "fail",
            duration: 4000,
            message: "请输入邮箱地址",
          });
          return;
        }
        if (this.password == "") {
          this.$toast({ type: "fail", duration: 4000, message: "请输入密码" });
          return;
        }
        this.loginByEmail();
      } else {
        if (this.phone == "" || this.phone == "none") {
          this.$toast({
            type: "fail ",
            duration: 4000,
            message: "请输入手机号码",
          });
          return;
        }
        if (this.code == "") {
          this.$toast({
            type: "fail ",
            duration: 4000,
            message: "请输入验证码",
          });
          return;
        }
        this.loginByPhone();
      }
    },
    async loginByEmail() {
      const result = await reqLoginByEmail(
        this.email,
        this.password,
        this.captcha,
        this.captchaId
      );
      console.log(result);
      if (result.code == 1) {
        this.$toast({ type: "fail", duration: 4000, message: result.msg });
        return;
      } else {
        localStorage.setItem("token", result.data.password);
        var consumerInfo = JSON.stringify(result.data);
        localStorage.setItem("consumerInfo", consumerInfo);
        this.$toast({ type: "success", message: "登录成功!" });
        setTimeout(() => {
          this.$router.push("/personal/center");
        }, 1000);
      }
    },
    async loginByPhone() {
      const result = await reqLoginByPhone(this.phone, this.code, this.bizId);
      console.log(result);
      if (result.code == 1) {
        this.$toast({ type: "fail", duration: 4000, message: result.msg });
        return;
      } else {
        localStorage.setItem("token", result.data.password);
        var consumerInfo = JSON.stringify(result.data);
        localStorage.setItem("consumerInfo", consumerInfo);
        this.$toast({ type: "success", message: "登录成功!" });
        setTimeout(() => {
          this.$router.push("/personal/center");
        }, 1000);
      }
    },
  },
};
</script>

<style scoped></style>
