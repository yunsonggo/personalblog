<template>
  <div>
    <Navbar :userInfo="userInfo"></Navbar>
    <div class="editViews">
      <div class="uploadfile">
        <van-uploader
          class="uploadimg"
          preview-size="100vw"
          :max-size="500 * 1024"
          @oversize="onOversize"
          :before-read="beforeRead"
          :after-read="afterRead"
          multiple
        />
        <PersonalEdit left="头像">
          <img
            v-if="userInfo.icon"
            :src="staticUrl + userInfo.icon"
            alt=""
            slot="right"
          />
          <img v-else src="@/assets/avatar.jpg" alt="" slot="right" />
        </PersonalEdit>
      </div>
      <PersonalEdit left="姓名" @editClick="editHandler">
        <a href="javascript:;" slot="right">{{ userInfo.name }}</a>
      </PersonalEdit>
      <PersonalEdit left="昵称" @editClick="editHandler" rule="^.{0,8}$">
        <a href="javascript:;" slot="right">{{ userInfo.nickname }}</a>
      </PersonalEdit>
      <PersonalEdit
        left="邮箱"
        @editClick="editHandler"
        rule="^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$"
      >
        <a href="javascript:;" slot="right">{{ userInfo.email }}</a>
      </PersonalEdit>
      <PersonalEdit
        left="电话"
        @editClick="editHandler"
        rule="^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$"
      >
        <a href="javascript:;" slot="right">{{ userInfo.phone }}</a>
      </PersonalEdit>
      <PersonalEdit left="性别" @editClick="showGenderSelect" rule="">
        <a href="javascript:;" slot="right">{{ userInfo.gender }}</a>
      </PersonalEdit>
      <PersonalEdit left="签名" @editClick="editHandler" rule=" ^.{0,30}$">
        <a href="javascript:;" slot="right">{{ userInfo.desc }}</a>
      </PersonalEdit>
      <van-dialog
        v-model="show"
        :title="editTitle"
        show-cancel-button
        @confirm="confirmClick"
      >
        <van-field v-model="text" autofocus />
      </van-dialog>
      <van-action-sheet
        v-model="showGender"
        :actions="actionsGender"
        @select="onSelect"
      />
      <div style="margin:10px 0">
        <van-button type="primary" block @touchstart="submitEdit"
          >立即修改</van-button
        >
      </div>
      <div class="editBack" @touchstart="$router.back()">
        返回
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from "@/components/common/Navbar.vue";
import PersonalEdit from "@/components/personalComponents/PersonalEdit.vue";
import WebStaticUrl from "@/components/common/WebStaticUrl";
import { reqEditConsumerInfo } from "@/api";
export default {
  name: "edit",
  components: {
    Navbar,
    PersonalEdit,
    WebStaticUrl,
  },
  data() {
    return {
      userInfo: {},
      uploadParams: {
        token: "",
        oldAvatar: "",
      },
      staticUrl: WebStaticUrl.static_url,
      editTitle: "",
      show: false,
      showGender:false,
      text: "",
      actionsGender:[{ name: '先生' }, { name: '女士' }, { name: '保密' }],
      isEmail: false,
      isPhone: false,
      rule: "",
      isChange: false,
    };
  },
  mounted() {
    this.isEmail = false;
    this.isPhone = false;
    if (this.$route.params.userInfo) {
      this.userInfo = this.$route.params.userInfo;
    } else {
      var userInfoJson = localStorage.getItem("consumerInfo");
      this.userInfo = JSON.parse(userInfoJson);
    }
    if (this.userInfo.password) {
      this.uploadParams.token = this.userInfo.password;
    } else {
      this.uploadParams.token = localStorage.getItem("token");
    }
    if (this.userInfo.icon) {
      this.uploadParams.oldAvatar = this.userInfo.icon;
    }
    if (this.userInfo.email) {
      this.isEmail = true;
    }
    if (this.userInfo.phone) {
      this.isPhone = true;
    }
  },
  methods: {
    showGenderSelect(data,rule) {
      this.$notify({ type: 'warning', message: '需要长按修改性别' })
      this.showGender = true
    },
    onSelect(item) {
      this.showGender = false;
      this.userInfo.gender = item.name
      this.isChange = true
    },
    onOversize(file) {
      this.$toast({
        type: "fail",
        duration: 4000,
        message: "文件大小不能超过 500kb",
      });
    },
    beforeRead(file) {
      if (file.type !== "image/jpeg" && file.type !== "image/png") {
        this.$toast({
          type: "fail",
          duration: 4000,
          message: "只能上传jpg或者png格式的图片",
        });
        return false;
      }
      return true;
    },
    afterRead(file) {
      var formdata = new window.FormData();
      formdata.append("file", file.file);
      formdata.append("token", this.uploadParams.token);
      formdata.append("old_avatar", this.uploadParams.oldAvatar);
      this.uploadAvatar(formdata);
    },
    uploadAvatar(updata) {
      fetch("http://192.168.1.102:8080/api/auth/personal/upload/avatar", {
        method: "POST",
        body: updata,
        credentials: "include"
      })
        .then((res) => res.json())
        .then((data) => {
          if (data.code === 0) {
            this.userInfo.icon = data.data;
            this.uploadParams.oldAvatar = this.userInfo.icon;
            var userInfoJson = JSON.stringify(this.userInfo);
            localStorage.removeItem("consumerInfo");
            localStorage.setItem("consumerInfo", userInfoJson);
            this.$toast({
              type: "success",
              message: "上传成功",
            });
          }
        })
        .catch((err) => {
          this.$toast({
            type: "fail",
            duration: 4000,
            message: err,
          });
        });
    },
    editHandler(data, rule) {
      this.text = "";
      if (
        (data == "邮箱" && this.isEmail == true) ||
        data == "姓名" ||
        (data == "电话" && this.isPhone == true)
      ) {
        this.$toast({
          type: "fail",
          duration: 2000,
          message: data + "不能修改",
        });
        return;
      }
      this.show = !this.show;
      this.editTitle = data;
      this.rule = rule;
    },
    confirmClick() {
      if (this.text == "") {
        this.$toast({
          type: "fail",
          duration: 2000,
          message: this.editTitle + "没有输入内容",
        });
        return;
      }
      if (this.editTitle == "昵称") {
        var res = this.regxHandler(this.text, this.rule);
        if (!res) {
          this.$toast({
            type: "fail",
            duration: 2000,
            message: this.editTitle + "不能超过8个字符",
          });
          return;
        }
        this.userInfo.nickname = this.text;
        this.isChange = true;
      }
      if (this.editTitle == "邮箱") {
        var res = this.regxHandler(this.text, this.rule);
        if (!res) {
          this.$toast({
            type: "fail",
            duration: 2000,
            message: this.editTitle + "格式错误",
          });
          return;
        }
        this.userInfo.email = this.text;
        this.isChange = true;
      }
      if (this.editTitle == "电话") {
        var res = this.regxHandler(this.text, this.rule);
        if (!res) {
          this.$toast({
            type: "fail",
            duration: 2000,
            message: this.editTitle + "格式错误",
          });
          return;
        }
        this.userInfo.phone = this.text;
        this.isChange = true;
      }
      if (this.editTitle == "性别") {
        var res = this.regxHandler(this.text, this.rule);
        if (!res) {
          this.$toast({
            type: "fail",
            duration: 2000,
            message: this.editTitle + "不能超过2个字符",
          });
          return;
        }
        this.userInfo.gender = this.text;
        this.isChange = true;
      }
      if (this.editTitle == "签名") {
        var res = this.regxHandler(this.text, this.rule);
        if (!res) {
          this.$toast({
            type: "fail",
            duration: 2000,
            message: this.editTitle + "不能超过30个字符",
          });
          return;
        }
        this.userInfo.desc = this.text;
        this.isChange = true;
      }
    },
    regxHandler(content, rule) {
      const regRule = new RegExp(rule);
      if (regRule.test(content)) {
        return true;
      } else {
        return false;
      }
    },
    submitEdit() {
      if (this.isChange) {
        this.$dialog
          .confirm({
            title: "提示",
            message: "邮箱或电话只能修改一次,不可再次修改",
          })
          .then(() => {
            // on confirm
            this.editConsumerInfo();
          })
          .catch(() => {
            // on cancel
            return;
          });
      } else {
        this.$toast({
          type: "fail",
          duration: 2000,
          message: "没有修改内容",
        });
        return;
      }
    },
    async editConsumerInfo() {
      var is_email;
      var is_phone;
      if (this.isEmail) {
        is_email = "true";
      } else {
        is_email = "false";
      }
      if (this.isPhone) {
        is_phone = "true";
      } else {
        is_phone = "false";
      }
      const result = await reqEditConsumerInfo(
        this.uploadParams.token,
        this.userInfo,
        is_email,
        is_phone
      );
      if (result.code === 0) {
        var userInfoJson = JSON.stringify(this.userInfo);
        localStorage.removeItem("consumerInfo");
        localStorage.setItem("consumerInfo", userInfoJson);
        this.$toast({
          type: "success",
          message: "修改成功",
        });
        this.$router.push('/personal/center')
      } else {
        this.$toast({
          type: "fail",
          message: "修改失败",
        });
      }
    },
  },
};
</script>

<style scoped>
.editViews {
  margin: 10px 0;
}
.editViews a {
  color: #333;
}
.editViews img {
  height: 46px;
  width: 46px;
  border-radius: 50%;
}
.uploadfile {
  position: relative;
  overflow: hidden;
}
.uploadimg {
  position: absolute;
  opacity: 0;
}
.editBack {
  background-color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #999;
  font-size: 3.333vw;
  padding: 12px 0;
}
</style>
