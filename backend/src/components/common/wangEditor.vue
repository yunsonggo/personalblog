<template lang="html">
  <div class="editor">
    <div ref="editorElem" class="editorStyle"></div>
  </div>
</template>

<script>
import E from "wangeditor";
import hljs from "highlight.js";
import staticUrl from "@/components/staticUrl.vue";
export default {
  name: "wangEditor",
  data() {
    return {
      staticUrl: staticUrl.static_url,
      editor: null,
      editorContent: "",
      managerInfo: {},
    };
  },
  model: {
    prop: "value",
    event: "change",
  },
  props: {
    value: {
      type: String,
      default: "",
    },
    isClear: {
      type: Boolean,
      default: false,
    },
    // content:{
    //   type:String,
    //   default:"",
    // }
  },
  watch: {
    isClear(val) {
      if (val) {
        this.editor.txt.clear();
        this.editorContent = null;
      }
    },
    value:function(val) {
      if (val !== this.editor.txt.html()) {
        this.editor.txt.html(val);
      }
    }
  },
  mounted() {
    this.getManagerInfo();
    this.initEditor();
    if (this.value) {
      this.editor.txt.html(this.value);
    }
  },
  methods: {
    getManagerInfo() {
      var managerJson = localStorage.getItem("adminToken");
      this.managerInfo = JSON.parse(managerJson);
    },
    initEditor() {
      const editor = new E(this.$refs.editorElem);
      editor.highlight = hljs;
      editor.config.height = 400;
      editor.config.uploadImgShowBase64 = false;
      editor.config.uploadImgServer = "http://192.168.1.102:8081/backend/auth/editor/upload/img";
      editor.config.uploadImgHeaders = {
        Authorization: this.managerInfo.manager_password,
      };
      editor.config.uploadFileName = "file";
      editor.config.uploadImgMaxSize = 8 * 1024 * 1024; // 2M
      editor.config.uploadImgMaxLength = 6; // 一次最多上传6张
      editor.config.uploadImgTimeout = 3 * 60 * 1000;

      editor.config.onchangeTimeout = 1000; // ms
      editor.config.onchange = (newHtml) => {
        this.editorContent = this.$xss(newHtml);
        this.$emit("onEditor", this.editorContent);
      };
      editor.config.menus = [
        "head",
        "bold",
        "fontSize",
        "fontName",
        "italic",
        "underline",
        "strikeThrough",
        "indent",
        "lineHeight",
        "foreColor",
        "backColor",
        "link",
        "list",
        "todo",
        "justify",
        "quote",
        "emoticon",
        "image",
        "video",
        "table",
        "code",
        "splitLine",
        "undo",
        "redo",
      ];
      editor.config.uploadImgHooks = {
        fail: (xhr, editor, result) => {
          // 插入图片失败
           this.$Message.error("插入图片失败");
        },
        success: (xhr, editor, result) => {
          // 图片上传成功
          if (result.code == 0) {  
                var url = this.staticUrl + '/' + result.data
                insertImg(url)  
            } else {  
                this.$Message.error(result.msg);  
            }  
        },
        timeout: (xhr, editor) => {
          // 超时
           this.$Message.error("超时");
        },
        error: (xhr, editor) => {
          //错误
           this.$Message.error("错误");
        },
        customInsert: (insertImg, result, editor) => {
          // 循环插入图片
            var url = this.staticUrl + '/' + result.data
            insertImg(url);
        },
      };
      editor.create();
      this.editor = editor
    },
  },
  beforeDestroy() {
    // 调用销毁 API 对当前编辑器实例进行销毁
    this.editor.destroy()
    this.editor = null
  },
};
</script>

<style scoped>
.editor {
  width: 100%;
  margin: auto;
}
</style>
