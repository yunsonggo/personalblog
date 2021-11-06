<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div class="addArticleBody">
      <el-form
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"
        label-width="100px"
        class="demo-ruleForm"
      >
        <el-form-item label="文章标题" prop="title">
          <el-input
            v-model="ruleForm.title"
            placeholder="请输入文章标题, 至少三个字符"
          ></el-input>
        </el-form-item>
        <el-form-item label="文章描述" prop="desc">
          <el-input
            v-model="ruleForm.desc"
            placeholder="请输入文章描述"
          ></el-input>
        </el-form-item>
        <el-form-item label="关键词" prop="keyword">
          <el-input
            v-model="ruleForm.keyword"
            placeholder="请输入关键词,以逗号隔开"
          ></el-input>
        </el-form-item>
        <el-form-item label="文章类型" prop="is_original">
          <el-radio-group v-model="ruleForm.is_original">
            <el-radio label="0" border size="medium">原创</el-radio>
            <el-radio label="1" border size="medium">引用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="ruleForm.is_original == '1'" label="作者或出处">
          <el-input
            v-model="ruleForm.author"
            placeholder="请输入作者或者出处"
          ></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.is_original == '1'" label="引用URL">
          <el-input
            v-model="ruleForm.link_url"
            placeholder="请输入引用地址"
          ></el-input>
        </el-form-item>
        <el-form-item label="是否激活" prop="stateBtn">
          <el-switch v-model="stateBtn" @change="changeState"></el-switch>
        </el-form-item>
        <el-form-item label="一级菜单" prop="menu_id">
          <el-select
            v-model="ruleForm.menu_id"
            placeholder="请选择一级菜单"
            @change="selectMenu"
          >
            <el-option
              v-for="(menu, index) in menuList"
              :key="index"
              :label="menu.title"
              :value="menu.Id"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="showCategory" label="二级菜单" prop="category_id">
          <el-select
            v-model="ruleForm.category_id"
            placeholder="请选择二级级菜单"
          >
            <el-option
              v-for="(cate, index) in categoryList"
              :key="index"
              :label="cate.title"
              :value="cate.Id"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="展示位置" prop="type">
          <el-checkbox-group v-model="type" @change="selectPosition">
            <el-checkbox
              v-for="position in positionArr"
              :key="position.label"
              :label="position.value"
              >{{ position.label }}</el-checkbox
            >
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="头部图片" prop="banner">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/article/upload/image/banner"
            :show-file-list="false"
            :on-success="handleBannerSuccess"
            :before-upload="beforeBannerUpload"
            :data="upBannerParamData"
            multiple="multiple"
          >
            <img
              v-if="ruleForm.banner"
              :src="staticUrl + '/' + ruleForm.banner"
              class="banner"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="封面图片" prop="cover">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/article/upload/image/cover"
            :show-file-list="false"
            :on-success="handleCoverSuccess"
            :before-upload="beforeCoverUpload"
            :data="upCoverParamData"
            multiple="multiple"
          >
            <img
              v-if="ruleForm.cover"
              :src="staticUrl + '/' + ruleForm.cover"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="视频" prop="video">
          <span style="color:#808080;font-size:12px;">视频优先于封面图片</span>
          <!-- 视频上传 -->
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/article/upload/video"
            :data="upVideoParamData"
            :on-progress="uploadVideoProcess"
            :on-success="handleVideoSuccess"
            :before-upload="beforeUploadVideo"
            :show-file-list="false"
          >
            <video
              v-if="ruleForm.video != '' && !videoFlag"
              :src="staticUrl + '/' + ruleForm.video"
              class="video-avatar"
              controls
            >
              您的浏览器不支持视频播放
            </video>
            <i
              v-else-if="ruleForm.video == '' && !videoFlag"
              class="el-icon-plus avatar-uploader-icon"
            ></i>
            <el-progress
              v-if="videoFlag == true"
              type="circle"
              v-bind:percentage="videoUploadPercent"
              style="margin-top:7px;"
            ></el-progress>
          </el-upload>
          <!-- 视频上传结束 -->
        </el-form-item>
        <el-form-item label="文章内容">
          <WangEditor @onEditor="onEditorChange"></WangEditor>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')"
            >立即创建</el-button
          >
          <el-button @click="resetForm('ruleForm')">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import MainTop from "@/components/common/mainTop.vue";
import WangEditor from "@/components/common/wangEditor.vue";
import { reqWebMenuList,reqAddArticle } from "@/api";
import staticUrl from "@/components/staticUrl.vue";
export default {
  name: "addArticle",
  components: {
    MainTop,
    WangEditor,
  },
  data() {
    return {
      staticUrl: staticUrl.static_url,
      menuTitle: "",
      managerInfo: {},
      menuList: [],
      showCategory: false,
      categoryList: [],
      positionArr: [
        { label: "首页展示", value: "is_index" },
        { label: "本类推荐", value: "is_top" },
        { label: "文章底部", value: "is_bottom" },
      ],
      type: [],
      stateBtn: true,
      upBannerParamData: {
        token: "",
        old_banner_url: "",
        article_id: 0,
      },
      upCoverParamData: {
        token: "",
        old_cover_url: "",
        article_id: 0,
      },
      upVideoParamData: {
        token: "",
        old_Video_url: "",
        article_id: 0,
      },
      ruleForm: {
        title: "",
        desc: "",
        keyword: "",
        is_original: "",
        link_url: "",
        state: 0,
        menu_id: null,
        category_id: null,
        is_index: null,
        is_top: null,
        is_bottom: null,
        author: "",
        author_id: null,
        content: null, //内容
        cover: "",
        banner: "",
        video: "",
      },
      rules: {
        title: [
          { required: true, message: "请输入文章标题", trigger: "blur" },
          { min: 3, message: "最少三个字符", trigger: "blur" },
        ],
        menu_id: [
          { required: true, message: "请选择一级菜单", trigger: "change" },
        ],
      },
      videoFlag: false,
      //是否显示进度条
      videoUploadPercent: "",
      //进度条的进度，
      isShowUploadVideo: false,
      //显示上传按钮
      videoForm: {
        showVideoPath: "",
      },
    };
  },
  mounted() {
    this.getManagerInfo();
    this.getMenuList();
  },
  methods: {
    getManagerInfo() {
      this.menuTitle = this.$route.query.title;
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
    },
    getMenuList() {
      if (this.$store.getters.webMenu.length) {
        this.menuList = this.$store.getters.webMenu;
      } else {
        this.getMenuData();
      }
    },
    async getMenuData() {
      const result = await reqWebMenuList(this.managerInfo.manager_password);
      if (result.code === 0) {
        result.data.forEach((item, index) => {
          if (item.WebMenuCategories) {
            item.WebMenu.hasChildren = true;
            item.WebMenu.children = [];
            item.WebMenu.children.push(...item.WebMenuCategories);
          } else {
            item.WebMenu.hasChildren = false;
          }
          this.menuList.push(item.WebMenu);
        });
        this.$store.dispatch("setWebmenu", this.menuList);
      }
    },
    selectMenu(value) {
      this.categoryList = [];
      this.menuList.forEach((menu) => {
        if (menu.Id == value) {
          if (menu.hasChildren) {
            this.showCategory = true;
            this.categoryList = menu.children;
          } else {
            this.showCategory = false;
          }
        }
      });
    },
    selectPosition(value) {
      this.type.forEach((item) => {
        if (item == "is_index") {
          this.ruleForm.is_index = 1;
        }
        if (item == "is_top") {
          this.ruleForm.is_top = 1;
        }
        if (item == "is_bottom") {
          this.ruleForm.is_bottom = 1;
        }
      });
    },
    changeState() {
      if (this.stateBtn == true) {
        this.ruleForm.state = 0;
      }
      if (this.stateBtn == false) {
        this.ruleForm.state = 1;
      }
    },
    beforeBannerUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传头像图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      if (this.ruleForm.banner != "") {
        this.upBannerParamData.old_banner_url = this.ruleForm.banner;
      }
      this.upBannerParamData.token = this.managerInfo.manager_password;
      return isJPG && isLt2M;
    },
    handleBannerSuccess(res, file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.ruleForm.banner = res.data;
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    beforeCoverUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传头像图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      if (this.ruleForm.cover != "") {
        this.upCoverParamData.old_cover_url = this.ruleForm.cover;
      }
      this.upCoverParamData.token = this.managerInfo.manager_password;
      return isJPG && isLt2M;
    },
    handleCoverSuccess(res, file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.ruleForm.cover = res.data;
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    /**** 上传视频 */
    //上传前回调
    beforeUploadVideo(file) {
      var fileSize = file.size / 1024 / 1024 < 50;
      if (
        [
          "video/mp4",
          "video/ogg",
          "video/flv",
          "video/avi",
          "video/wmv",
          "video/rmvb",
          "video/mov",
        ].indexOf(file.type) == -1
      ) {
        this.$message.error("请上传正确的视频格式");
        return false;
      }
      if (!fileSize) {
        this.$message.error("视频大小不能超过50MB");
        return false;
      }
      this.isShowUploadVideo = false;
      if (this.ruleForm.video != "") {
        this.upVideoParamData.old_video_url = this.ruleForm.video;
      }
      this.upVideoParamData.token = this.managerInfo.manager_password;
      return true;
    },
    //进度条
    uploadVideoProcess(event, file, fileList) {
      this.videoFlag = true;
      this.videoUploadPercent = file.percentage.toFixed(0) * 1;
    },
    //上传成功回调
    handleVideoSuccess(res, file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.ruleForm.video = res.data;
        this.isShowUploadVideo = true;
        this.videoFlag = false;
        this.videoUploadPercent = 0;
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.submitAddArticle()
        } else {
          this.$message.error("请输入必填项")
          return false;
        }
      });
    },
    // 提交新文章
    async submitAddArticle() {
      var token = this.managerInfo.manager_password
      var orig = Number(this.ruleForm.is_original )
      this.ruleForm.is_original = orig
      if (this.ruleForm.author == "") {
        this.ruleForm.author_id = this.managerInfo.Id
      }
      const result = await reqAddArticle(token,this.ruleForm)
      if (result.code === 0) {
        this.$toast({
          type:'success',
          message:"添加成功"
        })
        this.$router.push({
          path:"/home/article/list",
            query:{
                title:"文章列表",
            }
        })
      } else {
        this.$toast({
          type:'fail',
          message:"添加失败" + result.msg
        })
      }
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    onEditorChange(value) {
      this.ruleForm.content = value
    },
  },
};
</script>

<style scoped>
.addArticleBody {
  margin-top: 25px;
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
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
.banner {
  height: 178px;
  display: block;
}
.video-avatar {
  height: 178px;
  display: block;
}
</style>
