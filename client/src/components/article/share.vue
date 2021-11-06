<template>
  <div>
    <van-share-sheet
      v-model="showShare"
      title="立即分享"
      :options="options"
      @select="onSelect"
      @click-overlay="closeShare"
      @cancel="closeShare"
    />
  </div>
</template>
<script>
import WebStaticUrl from "@/components/common/WebStaticUrl.vue";
export default {
  name: "share",
  props: {
    showShareBtn: {
      type: Boolean,
      default: false,
    },
    articleUrlId: {
      type: Number,
      default: 0,
    },
    articleInfo:{
        type:Object,
        default:() => ({})
    }
  },
  watch: {
    showShareBtn() {
      this.showShare = this.showShareBtn;
    },
  },
  data() {
    return {
      options: [
        { name: "微信", icon: "wechat" },
        { name: "微博", icon: "weibo" },
        { name: "复制链接", icon: "link" },
        { name: "二维码", icon: "qrcode" },
      ],
      showShare: false,
      webUrl: WebStaticUrl.web_url,
    };
  },
  methods: {
    onSelect(option) {
      var shareUrl = this.webUrl + "article/content/" + this.articleUrlId;
      if (option.name == "复制链接") {
        this.copyInput(shareUrl);
      } else if (option.name == "微信") {
        this.$notify({
          type: "warning",
          duration: 5000,
          message:
            "暂未开通微信公众号，无法调用微信分享,将复制链接地址,请手动分享",
        });
        setTimeout(() => {
          this.copyInput(shareUrl);
        }, 1000);
      } else if (option.name == "微博") {
          this.shareWeiBo(shareUrl) 
      } else if (option.name == "二维码") {
        this.$emit("qrcode",shareUrl)
      }
      this.closeShare();
    },
    closeShare() {
      this.$emit("closeShare");
      this.showShare = false;
    },
    // 复制链接
    copyInput(url) {
      var input = document.createElement("input");
      input.value = url;
      document.body.appendChild(input);
      input.select();
      document.execCommand("Copy");
      document.body.removeChild(input);
      this.$toast({
        type: "success",
        message: "复制成功",
      });
    },
    //分享微博
    shareWeiBo(url) {
        console.log(this.articleInfo);
        var p = this.webUrl + this.articleInfo.banner
        var f = "http://v.t.sina.com.cn/share/share.php?url=" + url+"&sharesource=weibo&title="+this.articleInfo.title+"&pic=" + p
        window.open(f,'_blank',"",true)
    }
  },
};
</script>

<style scoped></style>
