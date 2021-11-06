import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const types = {
  SET_WEBMENU:"SET_WEBMENU",
  // SET_MENULIST:"SET_MENULIST",
  // SET_CATEGORYLIST:"SET_CATEGORYLIST",
};

const state = {
  webMenu:{},
  // menuList:{},
  // categoryList:{},
};

const getters = {
  webMenu:(state) => state.webMenu,
  // menuList:(state) => state.menuList,
  // categoryList:(state) => state.categoryList,
};

const mutations = {
  [types.SET_WEBMENU](state,webMenu) {
    if (webMenu) {
      state.webMenu = webMenu;
    } else {
      state.webMenu = null;
    }
  },
  // [types.SET_MENULIST](state,menuList) {
  //   if (menuList) {
  //     state.menuList = menuList
  //   } else {
  //     state.menuList = null;
  //   } 
  // },
  // [types.SET_CATEGORYLIST](state,categoryList) {
  //   if (categoryList) {
  //     state.categoryList = categoryList
  //   } else {
  //     state.categoryList = null;
  //   }
  // }
};

const actions = {
  setWebmenu:({commit},webMenu) => {
    commit(types.SET_WEBMENU,webMenu);
  },
  // setMenuList:({commit},menuList) => {
  //   commit(types.SET_MENULIST,menuList);
  // },
  // setCategoryList:({commit},categoryList) => {
  //   commit(types.SET_CATEGORYLIST,categoryList)
  // }, 
};

export default new Vuex.Store({
  state,
  getters,
  mutations,
  actions,
});

// this.$store.dispatch("setExample",data)
// this.$store.getters....