<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div>
      <div class="articleTableButton">
        <el-button type="success" plain @click="addLink">添加链接</el-button>
      </div>
      <vxe-table
        show-overflow
        keep-source
        resizable
        height="500"
        row-id="Id"
        :data="linkData"
        ref="xTable"
        :edit-config="{ trigger: 'click', mode: 'row' }"
        @edit-closed="editClosedEvent"
      >
        >
        <vxe-table-column
          field="title"
          title="标题"
          width="110"
          :edit-render="{ name: 'input', attrs: { type: 'text' } }"
        ></vxe-table-column>
        <vxe-table-column
          field="link_url"
          title="地址"
          width="300"
          :edit-render="{ name: 'input', attrs: { type: 'text' } }"
        ></vxe-table-column>
        <vxe-table-column
          field="sort"
          title="排序"
          width="90"
          :edit-render="{ name: 'input', attrs: { type: 'number' } }"
        ></vxe-table-column>
        <vxe-column field="icon" title="LOGO" width="160">
          <template #default="{ row }">
            <img
              v-if="row.icon"
              :src="staticUrl + '/' + row.icon"
              style="width: 120px;"
              @click="editLinkIcon(row)"
            />
            <span v-else>无</span>
          </template>
        </vxe-column>
        <vxe-table-column title="操作" width="150">
          <template #default="{ row }">
            <el-button
              type="danger"
              icon="el-icon-delete"
              size="small"
              @click="deleteLink(row)"
            ></el-button>
          </template>
        </vxe-table-column>
      </vxe-table>
      <el-dialog
        title="点击图片更新"
        :visible.sync="showEditIcon"
        width="30%"
        center
      >
        <el-form
          :model="ruleForm"
          ref="ruleForm"
          label-width="100px"
          class="demo-ruleForm"
        >
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
              v-if="ruleForm.link_icon"
              :src="staticUrl + '/' + ruleForm.link_icon"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="showEditIcon = false">取 消</el-button>
          <el-button type="primary" @click="showEditIcon = false"
            >确 定</el-button
          >
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import MainTop from "@/components/common/mainTop.vue";
import staticUrl from "@/components/staticUrl.vue";
import { reqLinkList, reqLinkText, reqRemoveLink } from "@/api";
export default {
  name: "linkList",
  components: {
    MainTop,
  },
  data() {
    return {
      staticUrl: staticUrl.static_url,
      managerInfo: {},
      menuTitle: "",
      linkData: [],
      showEditIcon: false,
      upIconParamData: {
        token: "",
        old_icon_url: "",
        link_id: 0,
      },
      ruleForm: {
        link_icon: "",
      },
    };
  },
  mounted() {
    this.getManagerInfo();
    this.getLinkList();
  },
  methods: {
    getManagerInfo() {
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
      this.menuTitle = this.$route.query.title;
    },
    async getLinkList() {
      const result = await reqLinkList(this.managerInfo.manager_password);
      if (result.code === 0) {
        this.linkData = result.data;
      } else {
        this.$message.error(res.msg);
      }
    },
    addLink() {
      this.$router.push({
        path: "/home/link/add",
        query: {
          title: "添加链接",
        },
      });
    },
    deleteLink(row) {
      this.$confirm("确定要删除该链接?, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          var token = this.managerInfo.manager_password;
          var id = row.Id;
          this.removeLinkById(token, id);
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    editClosedEvent({ row, column }) {
      const $table = this.$refs.xTable;
      const field = column.property;
      const cellValue = row[field];
      // 判断单元格值是否被修改
      if ($table.isUpdateByRow(row, field)) {
        var token = this.managerInfo.manager_password;
        var id = row.Id;
        var icon = row.icon;
        var link_url = row.link_url;
        var sort = Number(row.sort);
        var title = row.title;
        this.editLinkText(token, id, icon, link_url, sort, title, row, field);
      }
    },
    async editLinkText(token, id, icon, link_url, sort, title, row, field) {
      const $table = this.$refs.xTable;
      const result = await reqLinkText(token, id, icon, link_url, sort, title);
      if (result.code === 0) {
        $table.reloadRow(row, null, field);
        this.$message.success("更新成功");
      } else {
        this.$message.error("更新失败");
      }
    },
    editLinkIcon(row) {
      this.showEditIcon = true;
      this.upIconParamData.link_id = row.Id;
      this.ruleForm.link_icon = row.icon;
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
      if (this.ruleForm.link_icon != "") {
        this.upIconParamData.old_icon_url = this.ruleForm.link_icon;
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
        this.ruleForm.link_icon = res.data;
        this.getLinkList();
      } else {
        this.$message.error("上传失败", res.msg);
      }
    },
    async removeLinkById(token, id) {
      const result = await reqRemoveLink(token, id);
      if (result.code === 0) {
        this.$message.success(result.msg);
        this.getLinkList();
      } else {
        this.$message.error(result.msg);
      }
    },
  },
};
</script>

<style scoped>
.articleTableButton {
  padding: 15px 0;
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
