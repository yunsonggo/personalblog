<template>
  <div class="indexArticle">
    <van-list
      v-model="loading"
      :finished="finished"
      finished-text="没有更多了"
      @load="OnLoad"
      :immediate-check="false"
    >
      <div class="itemParent">
        <div
          v-for="(article, index) in theArticles"
          :key="index"
          class="articleItem"
        >
          <CoverArticle :articleInfo="article"></CoverArticle>
        </div>
      </div>
    </van-list>
  </div>
</template>

<script>
import CoverArticle from "@/views/article/coverArticle.vue";
export default {
  name: "menuArticle",
  components: {
    CoverArticle,
  },
  props: {
    articleType: {
      type: Number,
      default: 0,
    },
    articleStart: {
      type: Number,
      default: 1,
    },
    theArticles: {
      type: Array,
      default: [],
    },
    resLoading: {
      type: Boolean,
      default: false,
    },
    resFinished: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      loading: false,
      finished: false,
    };
  },
  watch:{
    resFinished() {
      this.finished = this.resFinished
    },
    resLoading() {
      this.loading = this.resLoading
    }
  },
  methods: {
    OnLoad() {
      var newStart = this.articleStart + 1;
      this.getMoreArticle(this.articleType, newStart);
      this.loading = this.resLoading
      this.finished = this.resFinished
    },
    getMoreArticle(articleType, newStart) {
      this.$emit("getMoreArticle", articleType, newStart);
    },
  },
};
</script>

<style scoped>
.indexArticle {
  background-color: white;
}
.indexArticle h4 {
  font-size: 3.889vw;
  margin: 0px 5px;
  padding: 10px;
  color: #fb7299;
}
.itemParent {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}
.articleItem {
  width: 45%;
  margin: 5px 0;
}
</style>
