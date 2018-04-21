import Vue from 'vue';
import Vuex, { Store } from 'vuex';

Vue.use(Vuex);

export default function createStore() {
  return new Store({
    state: {},
    mutations: {},
    actions: {},
    getters: {},
  });
}
