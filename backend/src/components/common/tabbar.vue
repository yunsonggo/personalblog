<template>
  <div class="menusInfo">
    <el-radio-group v-model="isCollapse" style="margin-bottom: 20px;">
      <el-radio-button :label="false">展开</el-radio-button>
      <el-radio-button :label="true">收起</el-radio-button>
    </el-radio-group>
    <el-menu
      :collapse="isCollapse"
      class="el-menu-vertical-demo"
      mode="vertical"
    >
      <fragment v-for="(item, i) in menusInfo" :key="i">
        <el-menu-item 
        v-if="item.MenuCategory == null" 
        :index="item.Menu.menu_path" 
        @click="handelClick(item.Menu.menu_title,item.Menu.menu_path)">
          <i :class="item.Menu.menu_icon"></i>
          <span slot="title">{{ item.Menu.menu_title }}</span>
        </el-menu-item>
        <el-submenu v-else :index="item.Menu.menu_path">
          <template slot="title">
            <i :class="item.Menu.menu_icon"></i>
            <span slot="title">{{ item.Menu.menu_title }}</span>
          </template>
          <el-menu-item
            v-for="(cate, k) in item.MenuCategory"
            :key="k"
            :index="cate.category_path"
            @click="handelClick(cate.category_title,cate.category_path)"
            >{{ cate.category_title }}
          </el-menu-item>
        </el-submenu>
      </fragment>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: "tabbar",
  props: {
    menusInfo: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      isCollapse: false,
    };
  },
  methods: {
      handelClick(title,path) {
        var breadParamArr = [
          title,
          path
        ]
        this.$emit("breadParam",breadParamArr)
        this.$router.push({
            path:path,
            query:{
                title:title,
            }
        })
      }
  },
};
</script>

<style scoped>
.menusInfo {
  width: 260px;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
