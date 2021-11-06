import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import Vant from 'vant';
import 'vant/lib/index.css';
import Fragment from 'vue-fragment'
import 'xe-utils'
import VXETable from 'vxe-table'
import 'vxe-table/lib/style.css'
import VTFormats from '@/components/vxeTableFormats.js'
import VueXss from 'vue-xss'
import 'highlight.js/styles/atom-one-light.css'

Vue.config.productionTip = false;
Vue.use(ElementUI);
Vue.use(Vant);
Vue.use(Fragment.Plugin)
Vue.use(VXETable)
Vue.use(VTFormats)
Vue.use(VueXss)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
