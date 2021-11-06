<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div class="userList">
      <div class="toolbar">
        <div class="toolbar-input">
          <span>搜索用户</span>
          <el-input
            v-model.trim="inputTitleKeyword"
            type="text"
            placeholder="请输入用户名"
            width="60px"
            @change="inputSearch"
            clearable
          >
            <!-- <template slot="prepend">标题</template> -->
          </el-input>
        </div>
      </div>
      <vxe-table
        show-overflow
        resizable
        height="500"
        row-id="Id"
        :data="showUserList"
        ref="xTable"
      >
        >
        <vxe-table-column type="seq" title="序号" width="60"></vxe-table-column>
        <vxe-table-column
          field="name"
          title="姓名"
          width="100"
        ></vxe-table-column>
        <vxe-table-column
          field="nick_name"
          title="昵称"
          width="100"
        ></vxe-table-column>
        <vxe-table-column
          field="phone"
          title="电话"
          width="110"
        ></vxe-table-column>
        <vxe-table-column
          field="email"
          title="邮箱"
          width="130"
        ></vxe-table-column>
        <vxe-table-column
          field="gender"
          title="性别"
          width="100"
        ></vxe-table-column>
        <vxe-table-column
          field="desc"
          title="签名"
          width="100"
        ></vxe-table-column>
        <vxe-table-column
        title="状态"
        width=150
        >
          <template #default="{ row }">
              <vxe-switch
              v-model="row.is_del"
              open-label="激活"
              :open-value=0
              close-label="关闭"
              :close-value=1
              class="my-switch2"
              @change="changeState(row)"
            ></vxe-switch>
          </template>
        </vxe-table-column>
      </vxe-table>
    </div>
  </div>
</template>

<script>
import { reqConsumerList , reqConsumerState} from "@/api";
import MainTop from "@/components/common/mainTop.vue";
export default {
  name: "user",
  components: {
    MainTop,
  },
  data() {
    return {
      menuTitle: "",
      managerInfo: {},
      userList: [],
      inputTitleKeyword: "",
      showUserList:[],
    };
  },
  mounted() {
    this.getManagerInfo(), this.getUserList();
  },
  methods: {
    getManagerInfo() {
      this.menuTitle = this.$route.query.title;
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
    },
    async getUserList() {
      const result = await reqConsumerList(this.managerInfo.manager_password);
      if (result.code === 0) {
        this.userList = result.data;
        this.showUserList = result.data;
      } else {
        this.$message.error(res.msg);
      }
    },
    inputSearch() {
      if (this.inputTitleKeyword) {
        this.showUserList = []
        this.userList.forEach((user) => {
          if (user.name.indexOf(this.inputTitleKeyword) > -1 || user.nickname.indexOf(this.inputTitleKeyword) > -1) {
            this.showUserList.push(user)
          }
        })
      } else {
        this.showUserList = this.userList
      }
    },
    changeState(row) {
        this.editConsumerState(row.id,row.is_del)
    },
    async editConsumerState(id,state) {
      const result = await reqConsumerState(this.managerInfo.manager_password,id,state)
      if (result.code === 0) {
        this.$message.success(result.data)
        this.getUserList()
      } else {
        this.$message.error(result.msg)
      }
    }
  },
};
</script>

<style scoped>
.toolbar {
  margin: 15px 0;
  width: 100%;
  display: flex;
  align-items: center;
}
.toolbar span {
  color: #808080;
  width: 85px;
  padding: 0 10px;
}
.toolbar .toolbar-input {
  display: flex;
  width: 40%;
  align-items: center;
}
.toolbar .toolbar-select {
  display: flex;
  align-items: center;
  width: 20%;
}
</style>
