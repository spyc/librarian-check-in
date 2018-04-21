import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

export default function createRouter() {
  return new VueRouter({
    routes: [
      { name: 'record', path: '/', component: () => import('./page/Record.vue') },
    ],
  });
}
