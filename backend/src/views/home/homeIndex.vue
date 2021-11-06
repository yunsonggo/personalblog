<template>
  <div>
    <MainTop :title="menuTitle"></MainTop>
    <div class="serverInfo" v-if="serverInfo.cpu_mode">
      <div class="infoItem">
        <span>CPU模块:</span>
        <span>{{ serverInfo.cpu_mode }}</span>
      </div>
      <div class="infoItem">
        <span>CPU核心:</span>
        <span>{{ serverInfo.cpu_cores }}</span>
      </div>
      <div class="infoItem">
        <span>系统版本:</span>
        <span>{{ serverInfo.host_platform }}</span>
      </div>
      <div class="infoItem">
        <span>主机名:</span>
        <span>{{ serverInfo.host_name }}</span>
      </div>
      <div class="infoItem">
        <span>主机IP:</span>
        <span>{{ serverInfo.host_ip }}</span>
      </div>
      <div class="infoItem">
        <span>内存大小:</span>
        <span>{{ (serverInfo.mem_total / 1024 / 1024).toFixed(2) }} MB</span>
      </div>
      <div class="infoItem">
        <span>使用内存:</span>
        <span>{{ (serverInfo.mem_used / 1024 / 1024).toFixed(2) }} MB</span>
      </div>
      <div class="infoItem">
        <span>剩余内存:</span>
        <span>{{ (serverInfo.mem_free / 1024 / 1024).toFixed(2) }} MB</span>
      </div>
      <div class="infoItem">
        <span>磁盘路径:</span>
        <span>{{ serverInfo.disk_path }}</span>
      </div>
      <div class="infoItem">
        <span>磁盘大小:</span>
        <span>{{ (serverInfo.disk_total / 1024 / 1024).toFixed(2) }} G</span>
      </div>
      <div class="infoItem">
        <span>磁盘使用:</span>
        <span>{{ (serverInfo.disk_used / 1024 / 1024).toFixed(2) }} G</span>
      </div>
      <div class="infoItem">
        <span>使用百分比:</span>
        <span>{{ serverInfo.disk_used_percent.toFixed(2) }} %</span>
      </div>
    </div>
    <div v-else>
        <el-table v-loading="loading"></el-table>
    </div>
  </div>
</template>

<script>
import { reqServerInfo } from "@/api";
import MainTop from "@/components/common/mainTop.vue";
export default {
  name: "homeIndex",
  components: {
    MainTop,
  },
  data() {
    return {
      serverInfo: {},
      menuTitle: "",
      loading:true
    };
  },
  mounted() {
    this.menuTitle = this.$route.query.title;
    this.getServerInfo();
  },
  methods: {
    async getServerInfo() {
      const result = await reqServerInfo();
      if (result.code === 0) {
        this.serverInfo = result.data;
      }
    },
  },
};
</script>

<style scoped>
.serverInfo {
  margin-top: 15px;
}
.infoItem {
  display: flex;
  align-items: center;
  font-size: 14px;
  margin-top:5px;
}
.infoItem span:nth-child(1) {
    width: 75px;
    margin-right: 15px;
    padding:5px 10px;
    background-color: #ecf5ff;
}
</style>
