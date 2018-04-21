import Vue from 'vue';
import Vuetify from 'vuetify';
import colors from 'vuetify/es5/util/colors';
import App from './App.vue';
import createStore from './store';

Vue.use(Vuetify, {
  theme: {
    primary: '#1a9eed',
    secondary: colors.amber.accent4,
  },
});

const store = createStore();

const vm = new Vue({
  store,
  render: h => h(App),
});
vm.$mount('#app');

