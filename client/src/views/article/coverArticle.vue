<template>
  <div class="coverArticle" @click="toArticleContent">
    <video v-if="articleInfo.video" :src="staticUrl + articleInfo.video"></video>
    <img v-else :src="staticUrl + articleInfo.cover" alt="" />
    <div class="articleComment">
      <div>
        <van-icon name="browsing-history-o" /><span
          >阅读:{{ articleInfo.read_num }}</span
        >
      </div>
      <div>
        <van-icon name="chat-o" /><span
          >评论:{{ articleInfo.comment_num }}</span
        >
      </div>
    </div>
    <div class="articleTitle">
      <p>{{ articleInfo.title }}</p>
    </div>
    <div class="articleGoodJob">
      <div>
        <van-icon name="good-job-o" /><span>{{ articleInfo.star }}</span>
      </div>
      <div>
        <van-icon name="closed-eye" /><span>{{ articleInfo.tread }}</span>
      </div>
      <div><span>2021-08-01</span></div>
    </div>
  </div>
</template>

<script>
import WebStaticUrl from "@/components/common/WebStaticUrl.vue";
export default {
  name: "coverArticle",
  props: {
    articleInfo: {
      type: Object,
      default: () => ({}),
    },
  },
  data() {
    return {
      staticUrl: WebStaticUrl.static_url,
    };
  },
  methods:{
    toArticleContent() {
      this.$router.push({
        name:"contentArticle",
        params:{
          id:this.articleInfo.Id,
          articleInfo:this.articleInfo
        }
      })
    }
  }
};
</script>

<style scoped>
.coverArticle {
  position: relative;
}
.coverArticle video {
  width: 100%;
  height: 39.444vw;
  background-color: black;
}
.coverArticle img {
  width: 100%;
  height: 39.444vw;
}
.articleComment {
  display: flex;
  position: absolute;
  top:34.6vw;
  color: white;
  left: 0;
  right: 0;
  justify-content: space-between;
  background-color: rgba(0, 0, 0, 0.6);
}
.articleComment div {
  display: flex;
  align-items: center;
  padding: 0 5px;
}
.articleComment span {
  font-size: 3.333vw;
}

.articleTitle p {
  font-size: 3.889vw;
  padding:1.389vw 0;
}
.articleGoodJob {
  /* position: absolute; */
  display: flex;
  align-items: center;
  width: 100%;
  height: 4.167vw;
  justify-content: space-between;
}
.articleGoodJob div {
  display: flex;
  align-items: center;
}
.articleGoodJob span {
  color: grey;
  font-size: 3.333vw;
}
</style>
