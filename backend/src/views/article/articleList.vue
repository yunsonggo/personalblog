<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div>
      <div class="articleTableButton">
        <el-button type="success" plain @click="addArticle">添加文章</el-button>
        <el-button type="danger" plain @click="removeArticleList"
          >批量删除</el-button
        >
      </div>
      <div class="toolbar">
        <div class="toolbar-input">
          <span>标题</span>
          <el-input
            v-model="inputTitleKeyword"
            type="text"
            placeholder="请输入关键词"
            width="60px"
            @change="inputSearch"
            clearable
          >
            <!-- <template slot="prepend">标题</template> -->
          </el-input>
        </div>
        <div class="toolbar-select">
          <span>大类</span>
          <el-select
            v-model="selectMenuKeywords"
            clearable
            placeholder="请选择"
            @change="handleSelectMenu"
          >
            <el-option-group
              v-for="group in menuFilters"
              :key="group.label"
              :label="group.label"
            >
              <el-option
                v-for="item in group.options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-option-group>
          </el-select>
        </div>
        <div class="toolbar-select">
          <span> 小类</span>
          <el-select
            v-model="selectCategoryKeywords"
            clearable
            placeholder="请选择"
            @change="handleSelectCategory"
          >
            <el-option
              v-for="item in categoryFilters"
              :key="item.Id"
              :label="item.title"
              :value="item.Id"
            >
            </el-option>
          </el-select>
        </div>
        <div class="toolbar-select">
          <span> 位置</span>
          <el-select
            v-model="selectPositionKeywords"
            clearable
            placeholder="请选择"
            @change="handleSelectPosition"
          >
            <el-option
              v-for="item in positionFilters"
              :key="item.label"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </div>
        <div class="toolbar-select">
          <el-button type="primary" @click="resetAll" style="margin-left:10px"
            >重置</el-button
          >
        </div>
      </div>
      <vxe-table
        show-overflow
        resizable
        height="500"
        row-id="Id"
        :loading="loading"
        :data="articleList"
        ref="xTable"
        @checkbox-change="checkboxChangeEvent"
        @checkbox-all="checkboxChangeEvent"
      >
        >
        <vxe-table-column type="checkbox" width="60"></vxe-table-column>
        <vxe-table-column type="seq" title="序号" width="60"></vxe-table-column>
        <vxe-table-column field="title" title="标题" width="400"></vxe-table-column>
        <vxe-table-column
          field="menu_id"
          title="大类"
          width="100"
          :formatter="['formatMenu', menuList]"
        ></vxe-table-column>
        <vxe-table-column
          field="category_id"
          title="小类"
          width="110"
          :formatter="['formatCategory', categoryList]"
        ></vxe-table-column>
        <vxe-table-column
          field="is_original"
          title="出处"
          width="110"
          :formatter="['formatOriginal']"
          :filters="[
            { label: '原创', value: 0 },
            { label: '引用', value: 1 },
          ]"
          :filter-multiple="false"
        ></vxe-table-column>
        <vxe-table-column
        title="操作"
        width=150
        >
          <template #default="{ row }">
            <el-button type="primary" icon="el-icon-edit" size="small" @click="editArticle(row)"></el-button>
            <!-- <el-button type="danger" icon="el-icon-delete" size="small" @click="deleteArticle(row)"></el-button> -->
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
import { reqArticleAll, reqWebMenuList, reqRemoveArticleArr } from "@/api";
export default {
  name: "articleList",
  components: {
    MainTop,
  },
  data() {
    return {
      managerInfo: {},
      articleList: [],
      loading: false,
      tablePage: {
        currentPage: 1,
        pageSize: 10,
        totalResult: 0,
      },
      menuTitle: "",
      articleData: [],
      baseArticleData: [],
      baseMenuArticleData: [],
      basePositionData: [],
      menuList: [],
      categoryList: [],
      inputSearchData: [],
      menuFilters: [
        {
          label: "选择",
          options: [
            {
              label: "全选",
              value: "99999",
            },
          ],
        },
      ],
      positionFilters: [
        { label: "主页展示", value: "is_index" },
        { label: "本类置顶", value: "is_top" },
        { label: "详情底面", value: "is_bottom" },
      ],
      categoryFilters: [],
      inputTitleKeyword: "",
      selectMenuKeywords: "",
      selectCategoryKeywords: "",
      selectPositionKeywords: "",
      selectRecords: [],
      isAllChecked: false,
      isIndeterminate: false,
    };
  },
  mounted() {
    this.getManagerInfo();
    this.getMenuData();
    this.getArticleAll();
  },
  methods: {
    getManagerInfo() {
      this.menuTitle = this.$route.query.title;
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
    },
    async getMenuData() {
      const result = await reqWebMenuList(this.managerInfo.manager_password);
      if (result.code === 0) {
        result.data.forEach((item, index) => {
          if (item.WebMenuCategories) {
            item.WebMenu.hasChildren = true;
            item.WebMenu.children = [];
            item.WebMenu.children.push(...item.WebMenuCategories);
            this.categoryList.push(...item.WebMenuCategories);
          } else {
            item.WebMenu.hasChildren = false;
          }
          this.menuList.push(item.WebMenu);
        });
        this.$store.dispatch("setWebmenu", this.menuList);
        this.handleMenuFileters(this.menuList);
      }
    },
    handleMenuFileters(menuList) {
      var pushMenuFilter = { label: "大类", options: [] };
      menuList.forEach((menu) => {
        if (menu.children) {
          var filter = { label: "", value: "", children: [] };
          filter.label = menu.title;
          filter.value = String(menu.Id);
          filter.children.push(...menu.children);
          pushMenuFilter.options.push(filter);
        } else {
          var filter = { label: "", value: "" };
          filter.label = menu.title;
          filter.value = String(menu.Id);
          pushMenuFilter.options.push(filter);
        }
      });
      this.menuFilters.push(pushMenuFilter);
    },
    async getArticleAll() {
      this.loading = true;
      const result = await reqArticleAll(this.managerInfo.manager_password);
      if (result.code === 0) {
        this.loading = false;
        this.articleData = result.data;
        this.baseArticleData = result.data;
        this.basePositionData = result.data;
        this.inputSearchData = result.data;
        this.tablePage.totalResult = this.articleData.length;
        this.findList();
      }
    },
    handlePageChange({ currentPage, pageSize }) {
      this.tablePage.currentPage = currentPage;
      this.tablePage.pageSize = pageSize;
      this.findList();
    },
    findList() {
      this.articleList = [];
      this.articleList = this.articleData.slice(
        (this.tablePage.currentPage - 1) * this.tablePage.pageSize,
        this.tablePage.currentPage * this.tablePage.pageSize
      );
    },
    handleSelectMenu(value) {
      var articleForArr = [];
      if (this.baseArticleData.length >= this.inputSearchData.length) {
        this.articleData = this.inputSearchData;
        articleForArr = this.inputSearchData;
      } else {
        this.articleData = this.baseArticleData;
        articleForArr = this.baseArticleData;
      }
      this.categoryFilters = [];
      this.baseMenuArticleData = [];
      this.selectCategoryKeywords = "";
      this.selectPositionKeywords = "";
      this.tablePage.totalResult = articleForArr.length;
      if (value == "") {
        this.articleData = articleForArr;
        this.categoryFilters = [];
        this.baseMenuArticleData = [];
        this.tablePage.totalResult = articleForArr.length;
      } else if (value == 99999) {
        this.articleData = articleForArr;
        this.categoryFilters = this.categoryList;
        this.tablePage.totalResult = articleForArr.length;
        this.baseMenuArticleData = articleForArr;
      } else {
        this.menuFilters[1].options.forEach((item) => {
          if (item.children) {
            item.children.forEach((cate) => {
              if (cate.parent_id == value) {
                this.categoryFilters.push(cate);
              }
            });
          }
        });
        this.articleData = [];
        articleForArr.forEach((article) => {
          if (article.menu_id == value) {
            this.articleData.push(article);
          }
        });
        this.tablePage.totalResult = this.articleData.length;
        this.baseMenuArticleData = this.articleData;
      }
      this.basePositionData = this.articleData;
      var xTable = this.$refs.xTable;
      xTable.updateData();
      this.findList();
    },
    handleSelectCategory(value) {
      this.articleData = this.baseMenuArticleData;
      this.basePositionData = this.baseMenuArticleData;
      this.tablePage.totalResult = this.baseMenuArticleData.length;
      this.selectPositionKeywords = "";
      if (value != "") {
        this.articleList = [];
        this.articleData = [];
        this.baseMenuArticleData.forEach((article) => {
          if (article.category_id == value) {
            this.articleData.push(article);
          }
        });
        this.tablePage.totalResult = this.articleData.length;
      }
      this.basePositionData = this.articleData;
      this.findList();
    },
    handleSelectPosition(value) {
      this.articleData = [];
      if (value == "is_index") {
        this.basePositionData.forEach((article) => {
          if (article.is_index == 1) {
            this.articleData.push(article);
          }
        });
      } else if (value == "is_top") {
        this.basePositionData.forEach((article) => {
          if (article.is_top == 1) {
            this.articleData.push(article);
          }
        });
      } else if (value == "is_bottom") {
        this.basePositionData.forEach((article) => {
          if (article.is_bottom == 1) {
            this.articleData.push(article);
          }
        });
      } else {
        this.articleData = this.basePositionData;
      }
      this.tablePage.totalResult = this.articleData.length;
      this.findList();
    },
    inputSearch() {
      this.articleData = [];
      var inputSearchArr = [];
      if (this.baseMenuArticleData.length > 0) {
        inputSearchArr = this.baseMenuArticleData;
      } else {
        inputSearchArr = this.baseArticleData;
      }
      if (this.inputTitleKeyword == "") {
        this.articleData = inputSearchArr;
      } else {
        inputSearchArr.forEach((article) => {
          if (article.title.indexOf(this.inputTitleKeyword) > -1) {
            this.articleData.push(article);
          }
        });
      }

      this.inputSearchData = this.articleData;
      this.tablePage.totalResult = this.inputSearchData.length;
      this.findList();
    },
    // 添加文章
    addArticle() {
      this.$router.push({
        path:"/home/article/add",
        query:{
                title:"添加文章",
            }
      })
    },
    // 批量删除
    removeArticleList() {
      if (this.selectRecords.length == 0) {
        this.$toast({
          type: "fail",
          message: "没有选择文章",
        });
        return;
      } else {
        this.$dialog
          .confirm({
            title: "删除文章",
            message: "确认要删除选择的文章吗",
          })
          .then(() => {
            // on confirm
            var ids = ""
            this.selectRecords.forEach((item) => {
              ids += item.Id + ','
            })
            this.handleRemoveArticleArr(ids)
          })
          .catch(() => {
            // on cancel
            this.$toast({
              type:"fail",
              message:"已经取消删除"
            })
            return
          });
      }
    },
    // 处理删除
    async handleRemoveArticleArr(ids) {
      const result = await reqRemoveArticleArr(this.managerInfo.manager_password,ids)
      if (result.code == 0) {
        this.$toast({
          type:"success",
          message:"删除成功"
        })
        this.$router.push({
          path:"/home/index",
            query:{}
          })
      } else {
        this.$toast({
          type:"fail",
          message:"删除失败"
        })
      }
    },
    resetAll() {
      this.inputTitleKeyword = "";
      this.selectMenuKeywords = "";
      this.selectCategoryKeywords = "";
      this.selectPositionKeywords = "";
      this.articleData = this.baseArticleData;
      this.tablePage.totalResult = this.articleData.length;
      this.findList();
    },
    checkboxChangeEvent({ records }) {
      const $grid = this.$refs.xTable;
      this.isAllChecked = $grid.isAllCheckboxChecked();
      this.isIndeterminate = $grid.isAllCheckboxIndeterminate();
      this.selectRecords = records;
    },
    changeAllEvent() {
      const $grid = this.$refs.xTable;
      $grid.setAllCheckboxRow(this.isAllChecked);
      this.selectRecords = $grid.getCheckboxRecords();
    },
    editArticle(row) {
      // this.$router.push({
      //   name:"editArticle",
      //   params:{
      //     articleInfo:row,
      //     title:"编辑文章"
      //   }
      // })
      this.$router.push({
        path:"/home/article/edit",
        query:{id:row.Id,title:"编辑文章"}
      })
    },
    // deleteArticle(row) {
    // }
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
  width: 55px;
  padding: 0 10px;
}
.toolbar .toolbar-input {
  display: flex;
  width: 30%;
  align-items: center;
}
.toolbar .toolbar-select {
  display: flex;
  align-items: center;
  width: 20%;
}
.articleTableButton {
  padding: 15px 0;
}
</style>
