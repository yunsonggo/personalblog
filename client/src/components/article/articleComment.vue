<template>
  <div class="comment">
    <div class="commentTitle" id="commentTop">
      <span>评论</span>
      <span>({{ commentNum }})</span>
    </div>
    <div class="commentUserInfo" v-if="userInfo.id">
      <img :src="staticUrl + userInfo.icon" alt="" />
      <span>{{ userInfo.nickname }}</span>
      <div class="commentInputBtn" @click="editComment">
        <span><van-icon name="edit"/></span>
        <span>说点什么...</span>
      </div>
    </div>
    <div class="commentUserInfo" v-else>
      <img src="@/assets/default_avatar.jpg" alt="" />
      <span @touchstart="toLogin">未登录</span>
    </div>
    <van-dialog
      v-model="showInput"
      :title="commentTitle"
      confirmButtonText="发表"
      show-cancel-button
      width="90%"
      @confirm="submitComment"
    >
      <p class="commentWarn">技术文章讨论,严禁发表敏感言论</p>
      <p class="commentWarn">和谐社会 文明用语</p>
      <van-field
        v-model="oneLevelComment"
        rows="3"
        autosize
        type="textarea"
        maxlength="300"
        placeholder="请输入..."
        show-word-limit
        autofocus
        clearable
      />
    </van-dialog>
    <div v-for="(item, index) in commentData" :key="index">
      <div class="commentParent">
        <div class="commentAuthor">
          <img
            v-if="item.Comment.author_icon"
            :src="staticUrl + item.Comment.author_icon"
            alt=""
          />
          <img v-else src="@/assets/default_avatar.jpg" alt="" />
          <p>
            <span>{{ item.Comment.author_nickname }}</span>
            <span>{{ item.Comment.create_time }}</span>
          </p>
        </div>
        <div class="commentContent">
          {{ item.Comment.content }}
          <span @click="askComment(item.Comment.Id,)">回复</span>
        </div>
        <div v-if="item.CommentChildList" class="commentChild">
          <div v-for="(commentChild, i) in item.CommentChildList" :key="i">
            <div class="commentChildAuthor">
              <img v-if="commentChild.author_icon" :src="staticUrl + commentChild.author_icon" alt="" />
              <img v-else src="@/assets/default_avatar.jpg" alt="" />
              <p>
                <span>{{ commentChild.author_nickname }}</span>
                <span>{{ commentChild.create_time }}</span>
              </p>
            </div>
            <div class="commentChildContent">
              <span>回复:</span>
              {{commentChild.content}}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import WebStaticUrl from "@/components/common/WebStaticUrl.vue";
import MgWords from "@/components/mingan.js";
// import XLSX from 'xlsx'
// import Axios from 'axios'
export default {
  name: "articleComment",
  props: {
    userInfo: {
      type: Object,
      default: () => ({}),
    },
    commentData: {
      type: Array,
      default: () => [],
    },
  },
  computed: {
    commentNum () {
      if (this.commentData) {
        this.$emit("commentNum",this.commentData.length)
        return this.commentData.length
      } else {
        this.$emit("commentNum",0)
        return 0
      }
    }
  },
  data() {
    return {
      staticUrl: WebStaticUrl.static_url,
      showInput: false,
      oneLevelComment: "",
      mgWords: MgWords,
      commentTitle:"发表评论",
      commentTag:0,
      askCommentId:0,
    };
  },
  methods: {
    toLogin() {
      this.$router.push("/personal/login");
    },
    editComment() {
      if (localStorage.getItem("token") && this.userInfo.id) {
        this.oneLevelComment = "";
        this.commentTitle = "发表评论"
        this.commentTag = 0;
        this.showInput = true;
        // this.readExcelFile()
      } else {
        this.$toast({
          type: "fail",
          message: "登录后才能参与评论哦",
        });
      }
    },
    // formatter(value) {
    //     var resValue = this.handleFilter(value);
    //     return resValue
    // },
    submitComment() {
      if (this.oneLevelComment != "") {
        var resValue = this.handleFilter(this.oneLevelComment);
        if (resValue) {
          if (this.commentTag === 0) {
            this.$emit("submitComment", resValue);
          } else if (this.commentTag === 1) {
            this.$emit("submitCommentChild",resValue,this.askCommentId)
          }
        }
      } else {
        this.$toast({
          type: "fail",
          message: "没有输入内容",
        });
      }
      //   this.$emit("submitComment",this.oneLevelComment)
    },
    handleFilter(value) {
      var v = value.trim();
      this.mgWords.forEach((word) => {
        if (v.indexOf(word) != -1) {
          var reg = new RegExp(word, "gi");
          v = v.replace(reg, "*");
        }
      });
      return v;
    },
    askComment(comment_id) {
       if (localStorage.getItem("token") && this.userInfo.id) {
        this.oneLevelComment = "";
        this.commentTitle = "回复评论"
        this.commentTag = 1;
        this.askCommentId = comment_id;
        this.showInput = true;
        // this.readExcelFile()
      } else {
        this.$toast({
          type: "fail",
          message: "登录后才能参与评论哦",
        });
      }
    }
    // readExcelFile() {
    //     var excelUrl = "/mg.xlsx"
    //     Axios.get(excelUrl,{responseType:'arraybuffer'})
    //     .then((res) => {
    //         var data = new Uint8Array(res.data)
    //         var wb = XLSX.read(data,{type:"array"})
    //         var sheets = wb.Sheets
    //         this.transFromSheets(sheets)
    //     }).catch(err => {
    //         console.log(err)
    //     })
    // },
    // transFromSheets(sheets) {
    //     var content = []
    //     var templist = []
    //     for (var key in sheets) {
    //         templist.push(XLSX.utils.sheet_to_json(sheets[key].length))
    //         content.push(XLSX.utils.sheet_to_json(sheets[key]))
    //     }
    //     content[0].forEach((c) => {
    //         this.mgWords.push(c.word)
    //     })
    //     // this.mgWords = content[0]
    //     // console.log(this.mgWords);
    //     // var jsonMg = JSON.stringify(content[0])
    //     // console.log(templist)
    //     // console.log(jsonMg)
    //     // content[0].forEach((key,item) => {
    //     //     console.log(key.WORDS)
    //     // })
    // }
  },
};
</script>

<style scoped>
.comment {
  /* background-color: aquamarine; */
  padding: 10px 15px;
}
.commentTitle {
  display: flex;
  align-items: center;
  justify-items: center;
  width: 100%;
  height: 30px;
}
.commentTitle span {
  font-size: 3.889vw;
}
.commentTitle span:nth-child(2) {
  margin-left: 1.389vw;
  color: #808080;
}
.commentUserInfo {
  padding: 10px 0;
  display: flex;
  align-items: center;
  justify-items: center;
}
.commentUserInfo img {
  height: 7.944vw;
  width: 7.944vw;
  border-radius: 2.778vw;
}
.commentUserInfo span:nth-child(2) {
  font-size: 3.889vw;
  color: #fb7299;
  margin-left: 1.389vw;
}
.commentUserInfo .commentInputBtn {
  display: flex;
  align-items: center;
  justify-items: center;
  padding: 0.633vw 0.633vw;
  margin-left: 2.778vw;
  background-color: #f8f9f9;
  border-radius: 2.222vw;
  color: #808080;
  height: 5.556vw;
  width: 50%;
  position: relative;
}
.commentUserInfo .commentInputBtn :nth-child(1) {
  color: #808080;
  padding-left: 1.389vw;
  position: absolute;
  bottom: 0.733vw;
}
.commentUserInfo .commentInputBtn :nth-child(2) {
  color: #808080;
  position: absolute;
  left: 15%;
  font-size: 3.333vw;
}
.commentWarn {
  font-size: 3.896vw;
  color: tomato;
  text-align: center;
  font-weight: bold;
  margin: 10px 0;
}
.commentParent {
  margin-top: 10px;
}
.commentAuthor {
  display: flex;
  align-items: center;
}
.commentAuthor img {
  width: 6.944vw;
  height: 6.944vw;
  border-radius: 2.778vw;
}
.commentAuthor p {
  margin-left: 2.778vw;
  font-size: 3.333vw;
  color: #808080;
}
.commentAuthor p span {
  margin-left: 1.389vw;
}
.commentContent {
  background-color: #f4f4f4;
  margin-top: 1.389vw;
  padding: 2.222vw;
  border-radius: 4.167vw;
  font-size: 3.333vw;
}
.commentContent span {
  color:#fb7299;
  padding-left: 2.778vw;
}
.commentChild {
  margin: 15px;
  font-style: italic;
}
.commentChildAuthor {
  display: flex;
  align-items: center;
}
.commentChildAuthor img{
  width: 6.544vw;
  height: 6.544vw;
  border-radius: 2.778vw;
}
.commentChildAuthor p {
  margin-left: 2.778vw;
  font-size: 3.333vw;
  color: #808080;
}
.commentChildAuthor p span {
  margin-left: 1.389vw;
}
.commentChildContent {
  background-color: #f4f4f4;
  margin-top: 1.389vw;
  padding: 1.222vw;
  border-radius: 4.167vw;
  font-size: 3.333vw;
}
</style>
