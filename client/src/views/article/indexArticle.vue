<template>
<div>
   <div class="indexArticle">
    <h4>{{ menuName }}{{menuTag}}</h4>
    <div class="itemParent">
      <div
        v-for="(article, index) in indexArticles"
        :key="index"
        class="articleItem"
      >
        <CoverArticle :articleInfo="article"></CoverArticle>
      </div>
    </div>
  </div> 
  <div class="indexArticle" v-for="(otherItem,itemIndex) in otherIndexArticles" :key="itemIndex">
    <h4>{{ otherItem.MenuTitle }}{{menuTag}}</h4>
    <div class="itemParent">
      <div
        v-for="(otherArticle, otherIndex) in otherItem.Articles"
        :key="otherIndex"
        class="articleItem"
      >
        <CoverArticle :articleInfo="otherArticle"></CoverArticle>
      </div>
    </div>
  </div> 
</div>
</template>

<script>
import CoverArticle from "@/views/article/coverArticle.vue";
export default {
  name: "indexArticle",
  components: {
    CoverArticle,
  },
  props: {
    menuName: {
      type: String,
      default: "首页",
    },
    menuTag:{
        type:String,
        default:"最新"
    },
    indexArticles: {
      type: Array,
      default: [],
    },
    otherIndexArticles:{
        type:Array,
        default:[],
    }
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
