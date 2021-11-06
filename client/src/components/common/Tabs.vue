<template>
  <div>
    <van-tabs v-model="activeMenu" color="#fb7299">
      <van-tab v-for="(menu, index) in menus" :key="index" :title="menu.title">
        <div v-if="menu.category">
          <van-tabs v-model="activeCategory" color="#ff8344">
            <van-tab
              v-for="(cate, index) in menu.category"
              :key="index"
              :title="cate.title"
            >
              <MenuArticle
                :theArticles="theArticles"
                :articleType="1"
                :articleStart="cate.articleStart"
                @getMoreArticle="getMore"
                :resLoading="resLoading"
                :resFinished="resFinished"
              ></MenuArticle>
            </van-tab>
          </van-tabs>
        </div>
        <div v-else-if="activeMenu != 0">
          <MenuArticle
            :theArticles="menuArticles"
            :articleType="0"
            :articleStart="menu.articleStart"
            @getMoreArticle="getMore"
            :resLoading="resLoading"
            :resFinished="resFinished"
          ></MenuArticle>
        </div>
        <div v-else>
          <IndexArticle
            :indexArticles="indexArticles"
            :menuName="menus[0].title"
            menuTag="最新文章"
            :otherIndexArticles="otherIndexArticles"
          ></IndexArticle>
        </div>
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import {
  reqMenu,
  reqCategory,
  reqIndexArticle,
  reqOtherIndexArticle,
  reqMenuCategoryArticle,
  reqMenuArticle,
} from "@/api";
import IndexArticle from "@/views/article/indexArticle.vue";
import MenuArticle from "@/views/article/menuArticle.vue";
export default {
  name: "tabs",
  components: {
    IndexArticle,
    MenuArticle,
  },
  data() {
    return {
      activeMenu: 0,
      activeCategory: 0,
      menus: [],
      category: [],
      count: 6,
      start: 1,
      theCount: 6,
      theStart: 1,
      indexArticles: [],
      otherIndexArticles: [],
      findMenuId: 0,
      findCategoryId: 1,
      theArticles: [],
      menuArticles: [],
      resLoading:false,
      resFinished:false,
    };
  },
  mounted() {
    this.getMenus();
  },
  watch: {
    activeMenu() {
      if (this.activeMenu != 0) {
        this.menus.forEach((menu, index) => {
          if (index === this.activeMenu) {
            this.findMenuId = menu.Id;
            if (menu.category) {
              this.activeCategory = 0;
              this.findCategoryId = menu.category[0].Id;
              menu.category.forEach((cate, i) => {
                cate.articleStart = this.theStart;
              });
              this.resLoading = false;
              this.resFinished = false;
              this.theArticles = [];
              this.getTheArticles();
            } else {
              menu.articleStart = this.theStart;
              this.resLoading = false;
              this.resFinished = false;
              this.menuArticles = [];
              this.getMenuArticles();
            }
          }
        });
      }
    },
    activeCategory() {
      this.menus.forEach((menu, index) => {
        if (this.findMenuId == menu.Id) {
          menu.category.forEach((cate, i) => {
            if (this.activeCategory === i) {
              this.findCategoryId = cate.Id;
              cate.articleStart = this.theStart;
            }
          });
        }
      });
      this.resLoading = false;
      this.resFinished = false;
      this.theArticles = [];
      this.getTheArticles();
    },
  },
  methods: {
    getMore(articleType, newStart) {
      this.resLoading = true
      setTimeout(() =>{
        this.getMoreArticles(articleType, newStart)
      },1000)
    },
    async getMoreArticles(articleType, newStart) {
       if (articleType == 0) {
        this.menus.forEach((menu, index) => {
          if (this.findMenuId == menu.Id) {
            menu.articleStart = newStart;
          }
        });
        var result = await this.getMenuArticles(this.theCount, newStart);
        if(result) {
          this.resLoading = false;
          this.resFinished = false;
        } else {
          this.resLoading = false;
          this.resFinished = true;
        }
      } else if (articleType == 1) {
        this.menus.forEach((menu, index) => {
          if (this.findMenuId == menu.Id) {
            menu.category.forEach((cate, i) => {
              if (this.findCategoryId === cate.Id) {
                cate.articleStart = newStart;
              }
            });
          }
        });
        var result = await this.getTheArticles(this.theCount, newStart);
        if(result) {
          this.resLoading = false;
          this.resFinished = false;
        } else {
          this.resLoading = false;
          this.resFinished = true;
        }
      } 
    },
    getMenus() {
      // this.$store.getters....
      this.menus = this.$store.getters.menu;
      if (!this.menus.length) {
        this.postGetMenus();
      }
    },
    async postGetMenus() {
      const result = await reqMenu();
      if (result.code === 0) {
        this.$store.dispatch("setMenu", result.data);
        this.menus = result.data;
        this.getCategory();
      } else {
        this.$toast({
          type: "fail",
          message: "获取菜单数据失败",
        });
      }
    },
    getCategory() {
      this.category = this.$store.getters.category;
      if (!this.category.length) {
        this.postGetCategory();
      }
    },
    async postGetCategory() {
      const result = await reqCategory();
      if (result.code === 0) {
        this.$store.dispatch("setCategory", result.data);
        this.category = result.data;
        this.addCategoryToMenu();
      } else {
        this.$toast({
          type: "fail",
          message: "获取类别数据失败",
        });
      }
    },
    addCategoryToMenu() {
      this.menus.forEach((menu, index) => {
        var cateList = [];
        this.category.forEach((cate, i) => {
          if (cate.parent_id == menu.Id) {
            cate.loading = false;
            cate.finished = false;
            cate.articleCount = this.theCount;
            cate.articleStart = this.theStart;
            cateList.push(cate);
          }
        });
        if (cateList.length > 0) {
          menu.category = cateList;
        }
        if (!menu.category) {
          menu.loading = false;
          menu.finished = false;
          menu.articleCount = this.theCount;
          menu.articleStart = this.theStart;
        }
      });
      var count = this.count;
      var start = this.start;
      this.getIndexArticle(count, start);
      this.getOtherIndexArticles(count, start);
    },
    // 获取首页展示的文章
    async getIndexArticle(count, start) {
      const result = await reqIndexArticle(count, start);
      if (result.code === 0) {
        this.indexArticles = result.data;
      } else {
        this.$toast({
          type: "fail",
          message: "获取首页文章失败",
        });
      }
    },
    // 获取其他大类的首页展示文章
    async getOtherIndexArticles(count, start) {
      const result = await reqOtherIndexArticle(count, start);
      if (result.code === 0) {
        this.otherIndexArticles = result.data;
      } else {
        this.$toast({
          type: "fail",
          message: "获取类别中最新文章失败",
        });
      }
    },
    // 获取指定大类指定小类文章
    async getTheArticles(
      articleCount = this.theCount,
      articleStart = this.theStart
    ) {
      const result = await reqMenuCategoryArticle(
        this.findMenuId,
        this.findCategoryId,
        articleCount,
        articleStart
      );
      if (result.code === 0) {
        if (result.data !== null) {
          this.theArticles.push(...result.data);
          return true;
        } else {
          return false;
        }
      } else {
        this.$toast({
          type: "fail",
          message: "获取类别中最新文章失败",
        });
        return false;
      }
    },
    // 获取指定大类无小类文章
    async getMenuArticles(
      articleCount = this.theCount,
      articleStart = this.theStart
    ) {
      const result = await reqMenuArticle(
        this.findMenuId,
        articleCount,
        articleStart
      );
      if (result.code === 0) {
        if (result.data !== null) {
          this.menuArticles.push(...result.data);
          return true;
        } else {
          return false;
        }
      } else {
        this.$toast({
          type: "fail",
          message: "获取类别中最新文章失败",
        });
        return false;
      }
    },
  },
};
</script>

<style scoped></style>
