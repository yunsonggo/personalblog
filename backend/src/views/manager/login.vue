<template>
  <div class="login">
    <div class="loginForm">
      <h3>后台登录</h3>
      <el-form
        :model="ruleForm"
        status-icon
        :rules="rules"
        ref="ruleForm"
        label-width="80px"
        class="demo-ruleForm"
      >
        <el-form-item label="账号" prop="name" required>
          <el-input type="text" v-model="ruleForm.name"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="pass" required>
          <el-input
            type="password"
            v-model="ruleForm.pass"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="验证码" prop="captcha" required style="">
          <el-input
            v-model="ruleForm.captcha"
            type="text"
            style="width:60%;vertical-align: middle;margin-right:26px"
          ></el-input>
          <img
            style="border:1px solid #d7d7d7;vertical-align: middle;display:table-cell；"
            :src="imgCaptcha"
            @click="handleNewCaptcha"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="submitForm('ruleForm')"
            style="width:150px"
            >提交</el-button
          >
          <el-button @click="resetForm('ruleForm')">重置</el-button>
        </el-form-item>
      </el-form>
      <div style="text-align:center">
        <el-button type="text" @click="$router.push('/manager/register')">用户注册</el-button>
        <el-button type="text" style="color:#808080;">找回密码</el-button>
      </div>
    </div>
  </div>
</template>

<script>
import { reqCaptcha,reqLogin } from "@/api";
export default {
  name: "login",
  data() {
    return {
      ruleForm: {
        pass: "",
        name: "",
        captcha: "",
        captchaId: "",
      },
      rules: {
        pass: [{ required: true, message: "请输入密码", trigger: "blur" }],
        name: [{ required: true, message: "请输入账号", trigger: "blur" }],
        captcha: [
          { required: true, message: "请输入验证码", trigger: "blur" },
          { min: 4, max: 4, message: "请输入四位验证码", trigger: "blur" },
        ],
      },
      imgCaptcha: "",
    };
  },
  mounted() {
    this.getCaptcha();
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.login()
        } else {
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    // 图形验证码
    async getCaptcha() {
      const result = await reqCaptcha();
      if (result.code === 0) {
        this.imgCaptcha = result.data.B64s;
        this.ruleForm.captchaId = result.data.Id;
      }
    },
    handleNewCaptcha() {
      this.getCaptcha();
    },
    async login() {
      const result = await reqLogin(this.ruleForm.name,this.ruleForm.pass,this.ruleForm.captchaId,this.ruleForm.captcha)
      if (result.code === 0) {
        var adminToken = JSON.stringify(result.data)
        localStorage.setItem("adminToken",adminToken)
        this.$toast({
          type:'success',
          message:"登录成功"
        })
        this.$router.push('/')
      } else {
        this.$toast({
          type:'fail',
          message:result.msg
        })
      }
    }
  },
};
</script>

<style coped>
.login {
  display: flex;
  justify-content: center;
  position: relative;
}
.loginForm {
  position: absolute;
  top: 200px;
  background-color: white;
  width: 450px;
  height: 350px;
  padding: 25px 35px;
  border-radius: 25px;
}
.loginForm h3 {
  text-align: center;
  margin-bottom: 25px;
}
</style>
