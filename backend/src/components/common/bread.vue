<template>
  <div>
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item >{{breadParams[0]}}</el-breadcrumb-item>
    </el-breadcrumb>
  </div>
</template>

<script>
export default {
  name: "bread",
  props: {
    breadParams:{
      type:Array,
      default:() => []
    }
  },
};
</script>

<style scoped></style>
