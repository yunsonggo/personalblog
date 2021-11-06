<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div class="menuListTable">
      <div class="menuTableButton">
        <el-button type="success" plain @click="addMenuClick"
          >添加菜单</el-button
        >
      </div>
      <vxe-table
        round
        border
        max-height="550"
        resizable
        show-overflow
        keep-source
        row-id="id"
        ref="xTable"
        :tree-config="{ children: 'children' }"
        :data="webMenus"
        highlight-hover-row
        @cell-dblclick="cellDBLClickEvent"
      >
        <vxe-table-column
          field="title"
          title="名称"
          tree-node
        ></vxe-table-column>
        <vxe-table-column field="sort" title="排位"></vxe-table-column>
        <vxe-table-column
          field="state"
          title="状态"
          :formatter="filterStateMethod"
        ></vxe-table-column>
        <vxe-table-column title="操作" show-overflow>
          <template #default="{ row }">
            <el-button
              type="primary"
              icon="el-icon-edit"
              size="mini"
              @click="editEvent(row)"
            ></el-button>
            <el-button
              type="danger"
              icon="el-icon-delete"
              size="mini"
              @click="removeEvent(row)"
            ></el-button>
          </template>
        </vxe-table-column>
      </vxe-table>
    </div>
    <el-dialog
      title="编辑菜单"
      :visible.sync="dialogFormVisible"
      center
      width="50%"
    >
      <el-form :model="form">
        <el-form-item
          v-if="form.parent_id"
          label="上级菜单"
          :label-width="formLabelWidth"
        >
          <el-select v-model="form.parent_id" placeholder="上级菜单">
            <fragment v-for="(menu, index) in webMenus" :key="index">
              <el-option
                :label="menu.title"
                :value="menu.Id"
                :key="menu.Id"
                :selected="form.parent_id == menu.Id"
              ></el-option>
            </fragment>
          </el-select>
        </el-form-item>
        <el-form-item label="名称" :label-width="formLabelWidth">
          <el-input v-model="form.title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="状态" :label-width="formLabelWidth">
          <el-switch
            v-model="stateValue"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="启用"
            inactive-text="关闭"
          >
          </el-switch>
        </el-form-item>
        <el-form-item label="排位" :label-width="formLabelWidth">
          <el-input
            v-model="form.sort"
            autocomplete="off"
            type="number"
          ></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelEdit">取 消</el-button>
        <el-button type="primary" @click="sumbitEdit">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import MainTop from "@/components/common/mainTop.vue";
import { reqWebMenuList, reqWebMenuEdit, reqWebMenuRemove } from "@/api";
export default {
  name: "menuList",
  components: {
    MainTop,
  },
  data() {
    return {
      webMenus: [],
      menuTitle: "",
      dialogTableVisible: false,
      dialogFormVisible: false,
      stateValue: null,
      managerInfo: {},
      form: {
        Id: null,
        hasChildren: null,
        sort: null,
        state: null,
        title: "",
        parent_id: null,
      },
      formLabelWidth: "80px",
    };
  },
  mounted() {
    this.getWebMenusList();
  },
  methods: {
    async getWebMenusList() {
      this.menuTitle = this.$route.query.title;
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
      const result = await reqWebMenuList(this.managerInfo.manager_password);
      if (result.code === 0) {
        this.changeList(result.data);
      }
    },
    changeList(data) {
      data.forEach((item, index) => {
        if (item.WebMenuCategories != null) {
          item.WebMenu.hasChildren = true;
          item.WebMenu.children = [];
          item.WebMenuCategories.forEach((cate, k) => {
            item.WebMenu.children.push(cate);
          });
        } else {
          item.WebMenu.hasChildren = false;
        }
        this.webMenus.push(item.WebMenu);
      });
      this.$store.dispatch("setWebmenu", this.webMenus);
    },
    cellDBLClickEvent({ row }) {
      this.editEvent(row);
    },
    editEvent(row) {
      this.form.Id = row.Id;
      this.form.hasChildren = row.hasChildren;
      this.form.state = row.state;
      this.form.title = row.title;
      this.form.sort = row.sort;
      if (row.state == 0) {
        this.stateValue = true;
      } else {
        this.stateValue = false;
      }
      if (row.parent_id) {
        this.form.parent_id = row.parent_id;
      }
      this.dialogFormVisible = true;
    },
    cancelEdit() {
      this.form = {};
      this.dialogFormVisible = false;
    },
    sumbitEdit() {
      this.EditMenu();
    },
    async EditMenu() {
      if (this.stateValue == true) {
        this.form.state = 0;
      } else {
        this.form.state = 1;
      }
      const result = await reqWebMenuEdit(
        this.managerInfo.manager_password,
        this.form.Id,
        this.form.title,
        Number(this.form.sort),
        this.form.state,
        this.form.parent_id
      );
      if (result.code === 0) {
        this.getWebMenusList();
        this.$toast({
          type: "success",
          message: "修改成功",
        });
        this.dialogFormVisible = false;
        this.$router.push("/home/index");
      } else {
        this.$toast({
          type: "fail",
          message: result.msg,
        });
      }
    },
    filterStateMethod({ cellValue }) {
      if (cellValue == 0) {
        return "启用中";
      } else {
        return "关闭中";
      }
    },
    addMenuClick() {
      this.$router.push("/home/menu/add");
    },
    removeEvent(row) {
      if (row.hasChildren == true) {
        this.$toast({
          type: "fail",
          message: "请先删除二级菜单",
        });
        return;
      }
      var id = row.Id;
      var parentId = row.parent_id
      var token = this.managerInfo.manager_password;
      this.$confirm("确定要删除菜单吗, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.removeMenu(token,id,parentId);
        })
        .catch(() => {
          this.$toast({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    async removeMenu(token,id,parentId) {
      const result = await reqWebMenuRemove(token,id,parentId);
      if (result.code === 0) {
        this.$toast({
          type: "success",
          message: "删除成功",
        });
        this.$router.push("/home/index");
      } else {
        this.$toast({
          type: "fail",
          message: "删除失败",
        });
      }
    },
  },
};
</script>

<style scoped>
.drawer {
  display: flex;
  background-color: yellow;
  width: 500px;
}
.menuTableButton {
  padding: 15px 0;
}
</style>
