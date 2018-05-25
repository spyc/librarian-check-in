import Vue from 'vue';
import Vuex, { Store } from 'vuex';

Vue.use(Vuex);

const SET_LIBRARIAN = 'SET_LIBRARIAN';
const SET_AUTH = 'SET_AUTH';

export default function createStore() {
  return new Store({
    state: {
      librarian: [],
      auth: false,
    },
    mutations: {
      [SET_LIBRARIAN](state, librarian) {
        state.librarian = librarian;
      },
      [SET_AUTH](state, auth) {
        state.auth = auth;
      },
    },
    actions: {
      updateLibrarian({ commit }, librarian) {
        commit(SET_LIBRARIAN, librarian);
      },
      updateAuth({ commit }) {
        commit(SET_AUTH, true);
      },
    },
    getters: {},
  });
}
