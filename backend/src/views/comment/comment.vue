<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div class="commentBody">
      <vxe-table
        border
        resizable
        show-overflow
        row-id="id"
        :data="commentList"
        :tree-config="{ children: 'CommentChildren' }"
        round
        max-height="550"
        keep-source
        ref="xTable"
        highlight-hover-row
        @cell-dblclick="cellDBLClickEvent"
        :loading="loading"
      >
        <vxe-column type="seq" width="100" title="序号" tree-node></vxe-column>
        <vxe-column
          field="AuthorNickname"
          title="作者"
          width="120"
        ></vxe-column>
        <vxe-column
          field="CommentArticle.title"
          title="文章"
          width="160"
        ></vxe-column>
        <vxe-column field="Content" title="内容"></vxe-column>
        <vxe-column
          field="CreateTime"
          title="时间"
          width="160"
          sortable
        ></vxe-column>
        <vxe-table-column title="操作" width="80">
          <template #default="{ row }">
            <el-button
              type="danger"
              icon="el-icon-delete"
              size="small"
              @click="deleteComment(row)"
            ></el-button>
          </template>
        </vxe-table-column>
      </vxe-table>
      <vxe-pager
        :loading="loading"
        :current-page="tablePage.currentPage"
        :page-size="tablePage.pageSize"
        :total="tablePage.totalResult"
        :layouts="[
          'PrevPage',
          'JumpNumber',
          'NextPage',
          'FullJump',
          'Sizes',
          'Total',
        ]"
        @page-change="handlePageChange"
      >
      </vxe-pager>
    </div>
  </div>
</template>

<script>
import MainTop from "@/components/common/mainTop.vue";
import staticUrl from "@/components/staticUrl.vue";
import { reqCommentAll,reqCommentRemove } from "@/api";
export default {
  name: "comment",
  components: {
    MainTop,
  },
  data() {
    return {
      staticUrl: staticUrl.static_url,
      managerInfo: {},
      commentData: [],
      menuTitle: "",
      loading: false,
      tablePage: {
        currentPage: 1,
        pageSize: 10,
        totalResult: 0,
      },
      commentList: [],
    };
  },
  mounted() {
    this.getManagerInfo();
    this.getCommentData();
  },
  methods: {
    getManagerInfo() {
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
      this.menuTitle = this.$route.query.title;
    },
    async getCommentData() {
      this.loading = true;
      const result = await reqCommentAll(this.managerInfo.manager_password);
      if (result.code === 0) {
        this.loading = false;
        this.commentData = result.data;
        this.tablePage.totalResult = this.commentData.length;
        this.findList();
      } else {
        this.$message.error(result.msg);
      }
    },
    handlePageChange({ currentPage, pageSize }) {
      this.tablePage.currentPage = currentPage;
      this.tablePage.pageSize = pageSize;
      this.findList();
    },
    findList() {
      this.commentList = [];
      this.commentList = this.commentData.slice(
        (this.tablePage.currentPage - 1) * this.tablePage.pageSize,
        this.tablePage.currentPage * this.tablePage.pageSize
      );
    },
    cellDBLClickEvent({ row }) {
      console.log(row);
    },
    deleteComment(row) { 
      var id = 0
      var comment_id = 0
      if (row.CommentId) {
          id = row.Id
          comment_id = row.CommentId
      } else {
          id = row.Id
      }
      this.RemoveComment(id,comment_id)
    },
    async RemoveComment(id,comment_id) {
        const result = await reqCommentRemove(this.managerInfo.manager_password,id,comment_id)
        
        if (result.code === 0) {
            this.$message.success("删除成功")
            this.getCommentData()
        } else {
            this.$message.error(result.msg)
        }
    }
  },
};
</script>

<style scoped>
.commentBody {
  margin-top: 15px;
}
</style>
