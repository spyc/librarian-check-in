import Vue from 'vue';
import Vuetify from 'vuetify';
import colors from 'vuetify/es5/util/colors';
import { sync } from 'vuex-router-sync';
import App from './App.vue';
import createStore from './store';
import createRouter from './router';

Vue.use(Vuetify, {
  theme: {
    primary: '#1a9eed',
    secondary: colors.amber.accent4,
  },
});

const store = createStore();
const router = createRouter();

sync(store, router);

const vm = new Vue({
  store,
  router,
  render: h => h(App),
});
vm.$mount('#app');

