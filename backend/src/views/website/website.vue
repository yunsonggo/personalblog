<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div class="websiteBtn">
      <el-button type="primary" plain @click="editFormBtn"
        >编辑网站基本信息</el-button
      >
      <el-button type="primary" plain @click="showEditAbout = true"
        >编辑关于我们</el-button
      >
      <el-button type="primary" plain @click="showEditAgreement = true"
        >编辑用户协议</el-button
      >
      <el-button type="primary" plain @click="showEditPrivacy = true"
        >编辑隐私声明</el-button
      >
    </div>
    <el-descriptions :column="1" border>
      <el-descriptions-item
        label="网站名"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_title }}
      </el-descriptions-item>
      <el-descriptions-item
        label="网站域名"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_url }}
      </el-descriptions-item>
      <el-descriptions-item
        label="作者"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_manager }}
      </el-descriptions-item>
      <el-descriptions-item
        label="昵称"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_nickname }}
      </el-descriptions-item>
      <el-descriptions-item
        label="简介"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_desc }}
      </el-descriptions-item>
      <el-descriptions-item
        label="关键词"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        <el-tag size="small" v-for="(word, index) in keyword" :key="index">
          {{ word }}
        </el-tag>
      </el-descriptions-item>
      <el-descriptions-item
        label="通知"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_notice }}
      </el-descriptions-item>
      <el-descriptions-item
        label="图片前缀"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_static_url }}
      </el-descriptions-item>
      <el-descriptions-item
        label="电话"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_phone }}
      </el-descriptions-item>
      <el-descriptions-item
        label="QQ"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_qq }}
      </el-descriptions-item>
      <el-descriptions-item
        label="Email"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_email }}
      </el-descriptions-item>
      <el-descriptions-item
        label="地址"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_address }}
      </el-descriptions-item>
      <el-descriptions-item
        label="备案号"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_licence_one }}
      </el-descriptions-item>
      <el-descriptions-item
        label="网安备案"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_licence_two }}
      </el-descriptions-item>
      <el-descriptions-item
        label-class-name="my-label"
        content-class-name="my-content"
        label="许可证"
      >
        {{ webSiteInfo.website_licence_three }}
      </el-descriptions-item>
      <el-descriptions-item
        label="其他证书"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_licence_Four }}
      </el-descriptions-item>
      <el-descriptions-item
        label="版权信息"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_copy }}
      </el-descriptions-item>
      <el-descriptions-item
        label="版本号"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_version }}
      </el-descriptions-item>
      <el-descriptions-item
        label="默认头像"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        <img
          :src="
            webSiteInfo.website_static_url + '/' + webSiteInfo.website_avatar
          "
          alt=""
          style="width:78px;height:78px"
        />
      </el-descriptions-item>
      <el-descriptions-item
        label="LOGO"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        <img
          :src="webSiteInfo.website_static_url + '/' + webSiteInfo.website_logo"
          alt=""
          style="width:178px;height:78px"
        />
      </el-descriptions-item>
      <el-descriptions-item
        label="Banner"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        <img
          :src="
            webSiteInfo.website_static_url + '/' + webSiteInfo.website_banner
          "
          alt=""
          style="width:178px;height:78px"
        />
      </el-descriptions-item>
      <el-descriptions-item
        label="微博"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        <img
          :src="webSiteInfo.website_static_url + '/' + webSiteInfo.website_webo"
          alt=""
          style="width:78px;height:78px"
        />
      </el-descriptions-item>
      <el-descriptions-item
        label="微信"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        <img
          :src="
            webSiteInfo.website_static_url + '/' + webSiteInfo.website_wechar
          "
          alt=""
          style="width:78px;height:78px"
        />
      </el-descriptions-item>
      <el-descriptions-item
        label="APK"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_apk }}
      </el-descriptions-item>
      <el-descriptions-item
        label="IOS"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        {{ webSiteInfo.website_ios }}
      </el-descriptions-item>
      <el-descriptions-item
        label="文本信息"
        label-class-name="my-label"
        content-class-name="my-content"
      >
        <el-button type="text" @click="showAbout = true">关于我们</el-button>
        <el-button type="text" @click="showAgreement = true"
          >用户协议</el-button
        >
        <el-button type="text" @click="showPrivacy = true">隐私声明</el-button>
      </el-descriptions-item>
    </el-descriptions>
    <el-drawer title="关于我们" :visible.sync="showAbout" :with-header="false">
      <div class="textContent">
        {{ webSiteInfo.website_about }}
      </div>
    </el-drawer>
    <el-drawer
      title="用户协议"
      :visible.sync="showAgreement"
      :with-header="false"
    >
      <div class="textContent">
        {{ webSiteInfo.website_agreement }}
      </div>
    </el-drawer>
    <el-drawer
      title="隐私声明"
      :visible.sync="showPrivacy"
      :with-header="false"
    >
      <div class="textContent">
        {{ webSiteInfo.website_privacy }}
      </div>
    </el-drawer>
    <el-dialog title="编辑网站基本信息" :visible.sync="showEdit">
      <el-form :model="form">
        <el-form-item label="网站标题" :label-width="formLabelWidth">
          <el-input v-model="form.website_title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="网站域名" :label-width="formLabelWidth">
          <el-input v-model="form.website_url" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="作者" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_manager"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="昵称" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_nickname"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="简介" :label-width="formLabelWidth">
          <el-input v-model="form.website_desc" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="关键词" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_keyword"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="通知" :label-width="formLabelWidth">
          <el-input v-model="form.website_notice" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="图片前缀" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_static_url"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="电话" :label-width="formLabelWidth">
          <el-input v-model="form.website_phone" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="QQ" :label-width="formLabelWidth">
          <el-input v-model="form.website_qq" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="Email" :label-width="formLabelWidth">
          <el-input v-model="form.website_email" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="地址" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_address"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="备案号" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_licence_one"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="网安备案" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_licence_two"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="许可证" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_licence_three"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="其他证书" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_licence_Four"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="版权信息" :label-width="formLabelWidth">
          <el-input v-model="form.website_copy" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="版本号" :label-width="formLabelWidth">
          <el-input
            v-model="form.website_version"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="默认头像" :label-width="formLabelWidth">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/website/avatar/edit"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
            :data="upAvatarData"
            multiple="multiple"
          >
            <img
              v-if="form.website_avatar"
              :src="form.website_static_url + '/' + form.website_avatar"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="LOGO" :label-width="formLabelWidth">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/website/logo/edit"
            :show-file-list="false"
            :on-success="handleLogoSuccess"
            :before-upload="beforeLogoUpload"
            :data="upLogoData"
            multiple="multiple"
          >
            <img
              v-if="form.website_logo"
              :src="form.website_static_url + '/' + form.website_logo"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="Banner" :label-width="formLabelWidth">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/website/banner/edit"
            :show-file-list="false"
            :on-success="handleBannerSuccess"
            :before-upload="beforeBannerUpload"
            :data="upBannerData"
            multiple="multiple"
          >
            <img
              v-if="form.website_banner"
              :src="form.website_static_url + '/' + form.website_banner"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="微博" :label-width="formLabelWidth">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/website/webo/edit"
            :show-file-list="false"
            :on-success="handleWeboSuccess"
            :before-upload="beforeWeboUpload"
            :data="upWeboData"
            multiple="multiple"
          >
            <img
              v-if="form.website_webo"
              :src="form.website_static_url + '/' + form.website_webo"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="微信" :label-width="formLabelWidth">
          <el-upload
            class="avatar-uploader"
            action="http://192.168.1.102:8081/backend/auth/website/wechar/edit"
            :show-file-list="false"
            :on-success="handleWecharSuccess"
            :before-upload="beforeWecharUpload"
            :data="upWecharData"
            multiple="multiple"
          >
            <img
              v-if="form.website_wechar"
              :src="form.website_static_url + '/' + form.website_wechar"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="APK" :label-width="formLabelWidth">
          <el-upload
            class="upload-demo"
            action="http://192.168.1.102:8081/backend/auth/website/apk/edit"
            :on-success="handleApkSuccess"
            :before-upload="beforeApkUpload"
            :data="apk"
            multiple="multiple"
            :limit="1"
          >
          <el-button size="small" type="primary">点击上传apk</el-button>
          </el-upload>
          <span>{{form.website_apk}}</span>
        </el-form-item>
        <el-form-item label="IOS" :label-width="formLabelWidth">
          <span>APPStore下载链接</span>
          <el-input
            v-model="form.website_ios"
            autocomplete="off"
          ></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="showEdit = false">取 消</el-button>
        <el-button type="primary" @click="submitEditInfo">确 定</el-button>
      </div>
    </el-dialog>
    <el-dialog title="编辑关于我们" :visible.sync="showEditAbout">
      <WangEditor
        @onEditor="onEditorAboutChange"
        :value="aboutForm.about"
      ></WangEditor>
       <el-button type="primary" @click="submitEditAbout">确 定</el-button>
       <el-button @click="showEditAbout = false">取 消</el-button>
    </el-dialog>
    <el-dialog title="编辑用户协议" :visible.sync="showEditAgreement">
      <WangEditor
        @onEditor="onEditorAgreementChange"
        :value="agreementForm.agreement"
      ></WangEditor>
       <el-button type="primary" @click="submitEditAgreement">确 定</el-button>
       <el-button @click="showEditAgreement = false">取 消</el-button>
    </el-dialog>
    <el-dialog title="编辑隐私声明" :visible.sync="showEditPrivacy">
      <WangEditor
        @onEditor="onEditorPrivacyChange"
        :value="privacyForm.privacy"
      ></WangEditor>
       <el-button type="primary" @click="submitEditPrivacy">确 定</el-button>
       <el-button @click="showEditPrivacy = false">取 消</el-button>
    </el-dialog>
  </div>
</template>

<script>
import MainTop from "@/components/common/mainTop.vue";
import staticUrl from "@/components/staticUrl.vue";
import { reqWebsiteInfo, reqWebsiteInfoEdit,reqEditWebsiteAbout,reqEditWebsiteAgreement,reqEditWebsitePrivacy } from "@/api";
import WangEditor from "@/components/common/wangEditor.vue";
export default {
  name: "website",
  components: {
    MainTop,
    WangEditor,
  },
  data() {
    return {
      staticUrl: staticUrl.static_url,
      managerInfo: {},
      menuTitle: "",
      webSiteInfo: {},
      keyword: [],
      showAbout: false,
      showAgreement: false,
      showPrivacy: false,
      showEdit: false,
      showEditAbout: false,
      showEditAgreement: false,
      showEditPrivacy: false,
      formLabelWidth: "120px",
      form: {
        // Id: 0,
        // website_address: "",
        // website_apk: "",
        // website_avatar: "",
        // website_banner: "",
        // website_copy: "",
        // website_desc: "",
        // website_email: "",
        // website_ios: "",
        // website_keyword: "",
        // website_licence_Four: "",
        // website_licence_one: "",
        // website_licence_three: "",
        // website_licence_two: "",
        // website_logo: "",
        // website_manager: "",
        // website_nickname: "",
        // website_notice: "",
        // website_phone: "",
        // website_qq: "",
        // website_static_url: "",
        // website_title: "",
        // website_url: "",
        // website_version: "",
        // website_webo: "",
        // website_wechar: "",
      },
      upAvatarData: {
        token: "",
        old_avatar_url: "",
        id: 1,
      },
      upLogoData: {
        token: "",
        old_logo_url: "",
        id: 1,
      },
      upBannerData: {
        token: "",
        old_banner_url: "",
        id: 1,
      },
      upWeboData: {
        token: "",
        old_webo_url: "",
        id: 1,
      },
      upWecharData: {
        token: "",
        old_wechar_url: "",
        id: 1,
      },
      aboutForm: {
        about: "",
      },
      agreementForm: {
        agreement: "",
      },
      privacyForm: {
        privacy: "",
      },
      apk:{
        token: "",
        old_apk_url: "",
        id: 1,
      },
    };
  },
  mounted() {
    this.getManagerInfo();
    this.getWebsite();
  },
  methods: {
    getManagerInfo() {
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
      this.menuTitle = this.$route.query.title;
    },
    async getWebsite() {
      const result = await reqWebsiteInfo();
      if (result.code === 0) {
        this.webSiteInfo = result.data;
        this.keyword = this.webSiteInfo.website_keyword.split(",");
        this.form = this.webSiteInfo;
        this.aboutForm.about = this.webSiteInfo.website_about
        this.agreementForm.agreement = this.webSiteInfo.website_agreement
        this.privacyForm.privacy = this.webSiteInfo.website_privacy
        localStorage.removeItem("staticUrl");
        localStorage.setItem("staticUrl", this.webSiteInfo.staticUrl);
      }
    },
    editFormBtn() {
      this.showEdit = true;
    },
    submitEditInfo() {
      this.editWebsiteInfo();
    },
    // 更新基本信息
    async editWebsiteInfo() {
      const result = await reqWebsiteInfoEdit(
        this.managerInfo.manager_password,
        this.form
      );
      if (result.code === 0) {
        this.$message.success(result.msg);
      } else {
        this.$message.error(result.msg);
      }
      this.showEdit = false;
    },
    submitEditAbout() {
      this.editWebsiteAbout()
    },
    // 更新关于我们
    async editWebsiteAbout() {
      const result = await reqEditWebsiteAbout(this.managerInfo.manager_password,this.aboutForm.about)
      if (result.code === 0) {
        this.$message.success(result.msg)
        this.getWebsite()
      } else {
        this.$message.error(result.msg)
      }
      this.showEditAbout = false
    },
    submitEditAgreement() {
      this.editWebsiteAgreement()
    },
    // 更新用户协议
    async editWebsiteAgreement() {
      const result = await reqEditWebsiteAgreement(this.managerInfo.manager_password,this.agreementForm.agreement)
      if (result.code === 0) {
        this.$message.success(result.msg)
        this.getWebsite()
      } else {
        this.$message.error(result.msg)
      }
      this.showEditAgreement = false
    },
    submitEditPrivacy() {
      this.editWebsitePrivacy()
    },
    // 更新隐私
    async editWebsitePrivacy() {
      const result = await reqEditWebsitePrivacy(this.managerInfo.manager_password,this.privacyForm.privacy)
      if (result.code === 0) {
        this.$message.success(result.msg)
        this.getWebsite()
      } else {
        this.$message.error(result.msg)
      }
      this.showEditPrivacy = false
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      if (this.form.website_avatar != "") {
        this.upAvatarData.old_avatar_url = this.form.website_avatar;
      }
      this.upAvatarData.token = this.managerInfo.manager_password;
      return isJPG && isLt2M;
    },
    handleAvatarSuccess(res, file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.form.website_avatar = res.data;
        this.getWebsite();
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    beforeLogoUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      if (this.form.website_logo != "") {
        this.upLogoData.old_logo_url = this.form.website_logo;
      }
      this.upLogoData.token = this.managerInfo.manager_password;
      return isJPG && isLt2M;
    },
    handleLogoSuccess(res, file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.form.website_logo = res.data;
        this.getWebsite();
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    beforeBannerUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      if (this.form.website_banner != "") {
        this.upBannerData.old_banner_url = this.form.website_banner;
      }
      this.upBannerData.token = this.managerInfo.manager_password;
      return isJPG && isLt2M;
    },
    handleBannerSuccess(res, file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.form.website_banner = res.data;
        this.getWebsite();
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    beforeWeboUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      if (this.form.website_webo != "") {
        this.upWeboData.old_webo_url = this.form.website_webo;
      }
      this.upWeboData.token = this.managerInfo.manager_password;
      return isJPG && isLt2M;
    },
    handleWeboSuccess(res, file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.form.website_webo = res.data;
        this.getWebsite();
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    beforeWecharUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      if (this.form.website_wechar != "") {
        this.upWecharData.old_wechar_url = this.form.website_wechar;
      }
      this.upWecharData.token = this.managerInfo.manager_password;
      return isJPG && isLt2M;
    },
    handleWecharSuccess(res, file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.form.website_wechar = res.data;
        this.getWebsite();
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    onEditorAboutChange(value) {
      this.aboutForm.about = value;
    },
    onEditorAgreementChange(value) {
      this.agreementForm.agreement = value;
    },
    onEditorPrivacyChange(value) {
      this.privacyForm.privacy = value;
    },
    beforeApkUpload(file) {
      const isAPK = file.type === "application/vnd.android.package-archive";
      const isLt50M = file.size / 1024 / 1024 < 50;
      if (!isAPK) {
        this.$message.error("上传只能是 APK 格式!");
      }
      if (!isLt50M) {
        this.$message.error("上传大小不能超过 50MB!");
      }
      if (this.form.website_apk != "") {
        this.apk.old_apk_url = this.form.website_apk;
      }
      this.apk.token = this.managerInfo.manager_password;
      return isAPK && isLt50M;
    },
    handleApkSuccess(res,file) {
      if (res.code == 0) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        this.form.website_apk = res.data;
        this.getWebsite();
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
  },
};
</script>

<style>
.websiteBtn {
  margin-top: 15px;
}
.my-label {
  background: #e1f3d8;
  width: 80px;
  height: 30px;
  text-align: left;
}
.textContent {
  margin: 20px;
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
  width: 78px;
  height: 78px;
  line-height: 78px;
  text-align: center;
}
.avatar {
  width: 78px;
  height: 78px;
  display: block;
}
</style>
