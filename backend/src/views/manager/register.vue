<template>
  <div class="login">
    <div class="loginForm">
      <h3>用户注册</h3>
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
        <el-form-item label="确认密码" prop="checkPass" required>
          <el-input
            type="password"
            v-model="ruleForm.checkPass"
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
        <el-button type="text" @click="$router.push('/manager/login')"
          >用户登录</el-button
        >
        <el-button type="text" style="color:#808080;">找回密码</el-button>
      </div>
    </div>
  </div>
</template>

<script>
import { reqCaptcha ,reqRegisterManager} from "@/api";
export default {
  name: "register",
  data() {
    var validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      } else {
        if (this.ruleForm.checkPass !== "") {
          this.$refs.ruleForm.validateField("checkPass");
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请再次输入密码"));
      } else if (value !== this.ruleForm.pass) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        pass: "",
        name: "",
        captcha: "",
        captchaId: "",
        checkPass: "",
      },
      rules: {
        pass: [{ validator: validatePass, trigger: "blur" }],
        checkPass: [{ validator: validatePass2, trigger: "blur" }],
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
          this.registerManager()
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
    // 用户注册
    async registerManager() {
        const result = await reqRegisterManager(this.ruleForm.name,this.ruleForm.pass,this.ruleForm.captchaId,this.ruleForm.captcha)
        if (result.code === 0 && result.data === "注册成功" && result.msg==="OK") {
            this.$toast({
                type:"success",
                message:result.data,
            })
            
            this.$router.push('/manager/login')
            
        } else {
            this.handleNewCaptcha()
            this.$toast({
                type:"fail",
                message:result.msg
            })
        }
    }
  },
};
</script>

<style scoped>
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
  height: 450px;
  padding: 25px 35px;
  border-radius: 25px;
}
.loginForm h3 {
  text-align: center;
  margin-bottom: 25px;
}
</style>
