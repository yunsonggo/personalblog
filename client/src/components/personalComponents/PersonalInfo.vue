<template>
  <div style=" background-color: white;">
    <div class="personal_info">
      <div>
        <van-image
        v-if="userInfo.icon"
          :src="staticUrl+userInfo.icon"
          round
          fit="fill"
          height="85px"
          width="85px"
        />
        <van-image
        v-else
          :src="require('@/assets/avatar.jpg')"
          round
          fit="fill"
          height="85px"
          width="85px"
        />
      </div>
      <div class="personal_edit">
        <div>
          <p>
            <span>0</span>
            <span>粉丝</span>
          </p>
          <p>
            <span>0</span>
            <span>关注</span>
          </p>
          <p>
            <span>0</span>
            <span>获赞</span>
          </p>
        </div>
        <div>
          <div @touchstart="toEdit">编辑资料</div>
        </div>
      </div>
    </div>
    <div class="pensonal_desc">
      <h3>{{userInfo.nickname}}</h3>
      <p>{{userInfo.desc}}</p>
      <van-divider hairline/>
      <div>
        <span>动态</span>
        <span>视频</span>
      </div>
      <div style="margin:0 15px">
        没有开放文章发布...
      </div>
    </div>
  </div>
</template>

<script>
import WebStaticUrl from "@/components/common/WebStaticUrl.vue"
export default {
  name: "PersonalInfo",
  components: {
    WebStaticUrl
  },
  props:{
    userInfo:{
      type:Object,
      default:{
        name:"",
        desc:""
      }
    }
  },
  data() {
    return {
      staticUrl:WebStaticUrl.static_url
    }
  },
  methods:{
    toEdit() {
      this.$router.push({
        name:"personalEdit",
        params:{
          userInfo:this.userInfo
        }
      })
    }
  }
};
</script>

<style scoped>
.personal_info {
  background-color: white;
  display: flex;
}
.personal_info div {
  justify-content: center;
  align-items: center;
}
.personal_info div:nth-child(1) {
  display: flex;
  padding: 5px 5px;
  width: 95px;
}
.personal_edit {
  width: 100%;
  display: flex;
  flex-direction: column;
}
.personal_edit div:nth-child(1) {
  flex: 1;
  justify-content: center;
  align-items: center;
  display: flex;
  flex-grow: 1;
  width: 100%;
}
.personal_edit div:nth-child(1) p {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  font-size: 13px;
  flex: 1;
}
.personal_edit div:nth-child(1) p:nth-child(1),
p:nth-child(2) {
  border-right: 1px solid #ccc;
}
.personal_edit div:nth-child(1) p span:nth-child(1) {
  color: #fb7299;
}
.personal_edit div:nth-child(2) {
  flex: 1;
  display: flex;
}
.personal_edit div:nth-child(2) div {
  width: 55.556vw;
  border: 1px solid #fb7a9f;
  padding: 3px;
  color: #fb7a9f;
  border-radius: 3px;
}
.pensonal_desc {
    background-color: white;
}
.pensonal_desc h3{
    background-color: white;
    margin:5px 25px;
    color:#5f6368;
}
.pensonal_desc p{
    margin:5px 10px;
    padding:5px 0;
    font-size: 13px;
    border: none;
    color:#5f6368;
}
.pensonal_desc div{
    margin:0 5px;
    display: flex;
}
.pensonal_desc div span{
    margin:5px 5px;
    font-size: 15px;
    color:#5f6368;
}
</style>
