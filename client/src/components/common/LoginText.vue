<template>
  <div>
    <!-- 输入任意文本 -->
    <van-field
      v-model="content"
      :label="label"
      :type="type"
      :placeholder="placeholder"
      :rule="rule"
      :required="required"
      :error-message="errorMessage"
    />
  </div>
</template>

<script>
export default {
  name: "LoginText",
  props: {
    label: {
      type: String,
      default: "请输入...",
    },
    type: {
      type: String,
      default: "text",
    },
    placeholder: {
      type: String,
      default: "",
    },
    rule: {
      type: String,
      default: "",
    },
    required: {
      type: Boolean,
      default: false,
    },
    errorInfo: {
      type: String,
      default: "",
    },
    clearErrorInfo: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      content: "",
      errorMessage: "",
    };
  },
  watch: {
    content() {
      this.regxHandler();
    },
    clearErrorInfo() {
      console.log("clear" + this.clearErrorInfo);
      this.errorMessage = "";
      this.content = "";
    },
  },
  methods: {
    regxHandler() {
      const regRule = new RegExp(this.rule);
      if (regRule.test(this.content)) {
        this.errorMessage = "";
        this.$emit("inputChange", this.content);
      } else {
        this.errorMessage = this.errorInfo;
      }
    },
  },
};
</script>

<style scoped></style>
