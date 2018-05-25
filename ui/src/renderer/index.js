import Vue from 'vue';
import Vuetify from 'vuetify';
import colors from 'vuetify/es5/util/colors';
import { sync } from 'vuex-router-sync';
import App from './App.vue';
import createStore from './store';
import createRouter from './router';

import './plugins';

Vue.use(Vuetify, {
  theme: {
    primary: '#1a9eed',
    secondary: colors.amber.accent4,
  },
});

const store = createStore();
const router = createRouter();

sync(store, router);

router.beforeEach((to, from, next) => {
  if (store.state.auth && to.name === 'auth') {
    next(false);
  } else if (!store.state.auth && (to.name === 'librarian' || to.name === 'report')) {
    next(`/auth?next=${to.fullPath}`);
  } else {
    next();
  }
});

const vm = new Vue({
  store,
  router,
  render: h => h(App),
});
vm.$mount('#app');

