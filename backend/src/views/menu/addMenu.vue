<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div class="addMenu">
      <el-form
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"
        label-width="100px"
        class="demo-ruleForm"
      >
        <el-form-item label="菜单级别" prop="hasChildren">
          <el-select
            v-model="hasChildren"
            placeholder="请选择菜单级别"
            @change="changeSelect"
          >
            <el-option label="一级菜单" :key="1" :value="true"></el-option>
            <el-option label="二级菜单" :key="2" :value="false"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item  label="父级菜单" prop="parent_id">
          <el-select
            v-model="ruleForm.parent_id"
            placeholder="请选择父级菜单"
            :disabled="hasChildren"
          >
            <el-option v-for="(menu,index) in webMenus" :key="index" :label="menu.title"  :value="menu.Id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="菜单名称" prop="title">
          <el-input v-model="ruleForm.title"></el-input>
        </el-form-item>
        <el-form-item label="状态" prop="state">
          <el-switch v-model="stateValue"></el-switch>
        </el-form-item>
        <el-form-item label="排位" prop="sort">
          <el-input v-model="ruleForm.sort" type="number"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')"
            >提交</el-button
          >
          <el-button @click="resetForm('ruleForm')">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import MainTop from "@/components/common/mainTop.vue";
import {reqWebMenuAdd} from "@/api"
export default {
  name: "addMenu",
  components: {
    MainTop,
  },
  data() {
    return {
      webMenus: [],
      menuTitle: "",
      stateValue: true,
      hasChildren: true,
      ruleForm: {
        hasChildren: true,
        title: "",
        state: 0,
        sort: null,
        parent_id: null,
      },
      rules: {
          hasChildren: [{ required: true, message: '请选择菜单级别', trigger: 'blur' }],
          title: [{ required: true, message: '请输入菜单名称', trigger: 'blur' }]
      },
    };
  },
  mounted() {
    this.menuTitle = this.$route.query.title;
    this.getWebMenus();
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.stateValue == true) {
              this.ruleForm.state = 0
          } else {
              this.ruleForm.state = 1
          }
          this.ruleForm.sort = Number(this.ruleForm.sort)
          this.addMenu()
        } else {
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    changeSelect() {
      this.ruleForm.hasChildren = this.hasChildren;
    },
    getWebMenus() {
      this.webMenus = this.$store.getters.webMenu;
    },
    async addMenu() {
        var managerJson = localStorage.getItem("adminToken")
        var manager = JSON.parse(managerJson)
        const result = await reqWebMenuAdd(manager.manager_password,this.ruleForm.title,this.ruleForm.sort,this.ruleForm.state,this.ruleForm.parent_id)
        if (result.code === 0) {
            this.$toast({
                type:'success',
                message:"添加成功"
            })
            this.$router.push('/home/index')
        } else {
            this.$toast({
                type:'fail',
                message:"添加失败"
            })
        }
    }
  },
};
</script>

<style scoped>
.addMenu {
  width: 500px;
  margin: 15px;
}
</style>
