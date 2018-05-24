import Vue from 'vue';
import Vuex, { Store } from 'vuex';

Vue.use(Vuex);

const SET_LIBRARIAN = 'SET_LIBRARIAN';

export default function createStore() {
  return new Store({
    state: {
      librarian: [],
    },
    mutations: {
      [SET_LIBRARIAN](state, librarian) {
        state.librarian = librarian;
      },
    },
    actions: {
      updateLibrarian({ commit }, librarian) {
        commit(SET_LIBRARIAN, librarian);
      },
    },
    getters: {},
  });
}
