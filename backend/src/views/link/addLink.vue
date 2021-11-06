<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div class="addLink">
      <el-form ref="form" :model="form" label-width="100px" :rules="rules">
        <el-form-item label="链接名称" prop="title">
          <el-input v-model="form.title" style="width:300px"></el-input>
        </el-form-item>
        <el-form-item label="链接地址" prop="link_url">
          <el-input v-model="form.link_url" style="width:450px"></el-input>
        </el-form-item>
        <el-form-item label="排序">
          <el-input
            type="number"
            v-model="form.sort"
            style="width:80px"
          ></el-input>
        </el-form-item>
        <el-form-item label="LOGO">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/link/icon/edit"
            :show-file-list="false"
            :on-success="handleCoverSuccess"
            :before-upload="beforeCoverUpload"
            :data="upIconParamData"
            multiple="multiple"
          >
            <img
              v-if="form.icon"
              :src="staticUrl + '/' + form.icon"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit('form')"
            >立即创建</el-button
          >
          <el-button @click="resetForm('form')">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import MainTop from "@/components/common/mainTop.vue";
import staticUrl from "@/components/staticUrl.vue";
import {reqInsertLink} from "@/api"
export default {
  name: "addLink",
  components: {
    MainTop,
  },
  data() {
    return {
      staticUrl: staticUrl.static_url,
      managerInfo: {},
      menuTitle: "",
      form: {
        icon: "",
        link_url: "",
        sort: 0,
        title: "",
      },
      rules: {
        title: [{ required: true, message: "请输入链接名称", trigger: "blur" }],
        link_url: [
          { required: true, message: "请输入链接地址", trigger: "blur" },
        ],
      },
      upIconParamData: {
        token: "",
        old_icon_url: "",
        link_id: 0,
      },
    };
  },
  mounted() {
    this.getManagerInfo();
  },
  methods: {
    getManagerInfo() {
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
      this.menuTitle = this.$route.query.title;
    },
    onSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          var token = this.managerInfo.manager_password;
          var id = 0;
          var icon = this.form.icon;
          var link_url = this.form.link_url;
          var sort = Number(this.form.sort);
          var title = this.form.title;
          this.InsertLink(token,id,icon,link_url,sort,title)
        } else {
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    beforeCoverUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      if (this.form.icon != "") {
        this.upIconParamData.old_icon_url = this.form.icon;
      }
      this.upIconParamData.token = this.managerInfo.manager_password;
      return isJPG && isLt2M;
    },
    handleCoverSuccess(res, file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.form.icon = res.data
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    async InsertLink(token,id,icon,link_url,sort,title) {
      const result = await reqInsertLink(token,id,icon,link_url,sort,title) 
      if (result.code === 0) {
        this.$message({
          message: "添加成功",
          type: "success",
        });
        this.$router.push({
          path:"/home/link/list",
          query:{
            title:"链接列表"
          }
        })
      } else {
        this.$message.error("添加失败", result.msg);
      }
    }
  },
};
</script>

<style scoped>
.addLink {
  margin-top: 15px;
}
.avatar-uploader .el-upload {
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  border: 1px dashed #808080;
  border-radius: 6px;
  font-size: 28px;
  color: #8c939d;
  width: 278px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 278px;
  height: 178px;
  display: block;
}
</style>
