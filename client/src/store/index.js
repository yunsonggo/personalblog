import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const types = {
  SET_MENU:"SET_MENU",
  SET_CATEGORY:"SET_CATEGORY",
  SET_WEBSITE:"SET_WEBSITE",
};

const state = {
  menu:{},
  category:{},
  website:{},
};

const getters = {
  menu:(state) => state.menu,
  category:(state) => state.category,
  website:(state) => state.website,
};

const mutations = {
  [types.SET_MENU](state,menu) {
    if (menu) {
      state.menu = menu;
    } else {
      state.menu = null;
    }
  },
  [types.SET_CATEGORY](state,category) {
    if (category) {
      state.category = category;
    } else {
      state.category = null;
    }
  },
  [types.SET_WEBSITE](state,website) {
    if (website) {
      state.website = website;
    } else {
      state.website = null;
    } 
  } 
};

const actions = {
  setMenu:({commit},menu) => {
    commit(types.SET_MENU,menu);
  },
  setCategory:({commit},category) => {
    commit(types.SET_CATEGORY,category)
  },
  setWebsite:({commit},website) => {
    commit(types.SET_WEBSITE,website)
  },
};

export default new Vuex.Store({
  state,
  getters,
  mutations,
  actions,
});

// this.$store.dispatch("setExample",data)
// this.$store.getters....