<template>
  <div class="articleBottom">
    <div class="pageUpDown">
      <div>
        <span>
          <van-button
            type="default"
            icon="arrow-left"
            style="border:0"
            :disabled="btnDisabled"
            @touchstart="changePageUp"
            >{{ leftTitle }}
          </van-button>
        </span>
      </div>
      <div>
        <span>
          <van-button
            type="default"
            style="border:0"
            @touchstart="changePageBack"
          >
            返回首页
          </van-button>
        </span>
      </div>
      <div>
        <span>
          <van-button
            type="default"
            icon="arrow"
            icon-position="right"
            style="border:0"
            @touchstart="changePageDown"
            >{{ rightTitle }}</van-button
          ></span
        >
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "bottomArticle",
  props: {
    bottomArticle: {
      type: Array,
      default: () => [],
    },
    articleId: {
      type: Number,
      default: 0,
    },
    articleMenuId: {
      type: Number,
      default: 0,
    },
    articleCategoryId: {
      type: Number,
      default: 0,
    },
  },
  data() {
    return {
      leftTitle: "上一篇",
      rightTitle: "下一篇",
      btnDisabled: false,
    };
  },
  methods: {
    changePageUp() {
      if (this.articleId <= 1) {
        this.btnDisabled = true;
        this.leftTitle = "第一篇";
        this.$toast({
          type: "fail",
          message: "已经是第一篇了",
        });
      } else {
        this.$emit("changeArticleId", "pageUp");
      }
    },
    changePageBack() {
      this.btnDisabled = false;
      this.$emit("changeArticleId", "pageBack");
    },
    changePageDown() {
      this.$emit("changeArticleId", "pageDown");
    },
  },
};
</script>

<style scoped>
.articleBottom {
  background-color: white;
}
.pageUpDown {
  display: flex;
  justify-content: space-between;
}
.pageUpDown div {
  display: flex;
  align-content: center;
  align-items: center;
  color:  #fb7299;
}
.pageUpDown div span {
  font-size: 3.889vw;
}
</style>
