<template>
  <div>
    <LoginTop
      :middleTop="middleTopText"
    >
    <div slot="right" style="font-size:3.333vw;color:#07c160"
    @touchstart="changeRegister"
    >
    {{rightTop}}
    </div>
    </LoginTop>
    <LoginText
      style="margin:4.167vw 0;"
      label="姓名:"
      placeholder="请输入姓名:"
      rule="^[\u4E00-\u9FA5]{2,10}(·[\u4E00-\u9FA5]{2,10}){0,2}$"
      required
      @inputChange="(res) => (name = res)"
      errorInfo="姓名为2-10个中文字符"
      :clearErrorInfo="clearErrorInfo"
    />
    <LoginText
      v-if="!changeRegisterBtn"
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
      v-if="!changeRegisterBtn"
      label="密码:"
      placeholder="请输入密码:"
      type="password"
      rule="^[a-zA-Z]([-_a-zA-Z0-9]{7,19})+$"
      required
      @inputChange="(res) => (password = res)"
      errorInfo="密码为包含字母的7-19位(字母,数字,_,-)"
      :clearErrorInfo="clearErrorInfo"
    />
    <LoginText
      v-if="!changeRegisterBtn"
      label="确认密码:"
      placeholder="请再次输入密码:"
      type="password"
      rule="^[a-zA-Z]([-_a-zA-Z0-9]{7,19})+$"
      required
      @inputChange="(res) => (checkpass = res)"
      errorInfo="密码为包含字母的7-19位(字母,数字,_,-)"
      :clearErrorInfo="clearErrorInfo"
    />
    <LoginCode
      label="验证码"
      placeholder="请输入验证码"
      type="text"
      rule="^\d{4}$"
      required
      :requestInfo="requestInfo"
      @inputChange="(res) => (code = res)"
      errorInfo="验证码为四位数字"
      :codeMode="changeRegisterBtn"
      :clearErrorInfo="clearErrorInfo"
      @bizId="getBizId"
    ></LoginCode>
    <LoginBtn
      style="margin:4.167vw 0;"
      btnName="立即注册"
      block
      size="large"
      @registerSubmit="submit"
    ></LoginBtn>
    <LoginBottom
    toPath="login"
    toName="已有账号登录"
    ></LoginBottom>
  </div>
</template>

<script>
import LoginTop from "@/components/common/LoginTop.vue";
import LoginText from "@/components/common/LoginText.vue";
import LoginCode from "@/components/common/LoginCode.vue";
import LoginBtn from "@/components/common/LoginBtn.vue";
import { reqRegisterByEmail, reqRegisterByPhone } from "@/api";
import LoginBottom from "@/components/common/LoginBottom.vue"
export default {
  name: "register",
  components: {
    LoginTop,
    LoginText,
    LoginCode,
    LoginBtn,
    LoginBottom
  },
  data() {
    return {
      middleTopText: "新用户注册",
      changeRegisterBtn: false,
      rightTop: "手机注册",
      name: "",
      phone: "none",
      email: "none",
      password: "",
      checkpass: "",
      code: "",
      clearErrorInfo: false,
      requestInfo: "",
      requestUrl: "",
      bizId: "",
    };
  },
  watch: {
    phone() {
      this.requestInfo = this.phone;
    },
    email() {
      this.requestInfo = this.email;
    },
  },
  methods: {
    changeRegister() {
      this.changeRegisterBtn = !this.changeRegisterBtn;
      if (this.changeRegisterBtn) {
        this.rightTop = "邮箱注册";
        this.requestUrl = "phone";
      } else {
        this.rightTop = "手机注册";
        this.requestUrl = "email";
      }
      this.clearErrorInfo = this.changeRegisterBtn;
      this.name = "";
      this.phone = "";
      this.email = "";
      this.password = "";
      this.checkpass = "";
      this.code = "";
    },
    submit() {
      console.log(this.requestUrl);
      if (this.requestUrl == "phone") {
        this.submitByPhone();
      } else if (this.requestUrl == "" || this.requestUrl == "email") {
        if (!this.verifyCheckPass()) {
          this.$notify({ type: "warning", message: "两次密码不一致" });
          return;
        } else {
          this.submitByEmail();
        }
      }
    },
    async submitByEmail() {
      const result = await reqRegisterByEmail(
        this.name,
        this.email,
        this.password,
        this.checkpass,
        this.code
      );
      console.log(result);
      if (result.code == 1) {
        this.$toast({type:"fail ",duration:4000,message:result.msg})
      } else {
        localStorage.setItem("token",result.data.password)
        var consumerInfo = JSON.stringify(result.data)
        localStorage.setItem("consumerInfo",consumerInfo)
        this.$toast({type:"success",message:"注册成功!"})
        setTimeout(()=>{
          this.$router.push('/personal/center')
        },1000)
      }
      return;
    },
    async submitByPhone() {
      const result = await reqRegisterByPhone(
        this.name,
        this.bizId,
        this.phone,
        this.code
      );
      console.log(result);
      if (result.code == 1) {
        this.$toast({type:"fail ",duration:4000,message:result.msg})
      } else {
        localStorage.setItem("token",result.data.password)
        var consumerInfo = JSON.stringify(result.data)
        localStorage.setItem("consumerInfo",consumerInfo)
        this.$toast({type:"success",message:"注册成功!"})
        setTimeout(()=>{
          this.$router.push('/personal/center')
        },1000)
      }
      return;
    },
    verifyCheckPass() {
      if (this.password == "" || this.password != this.checkpass) {
        return false;
      } else {
        return true;
      }
    },
    getBizId(data) {
      this.bizId = data;
    },
  },
};
</script>

<style scoped></style>
