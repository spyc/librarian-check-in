import Vue from 'vue';
import Vuetify from 'vuetify';
import App from './App.vue';
import createStore from './store';

Vue.use(Vuetify);

const store = createStore();

new Vue({
  store,
  render: h => h(App),
}).$mount('#app');

