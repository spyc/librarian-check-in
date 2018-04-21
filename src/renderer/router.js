import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

export default function createRouter() {
  return new VueRouter({
    routes: [
      { name: 'home', path: '/', component: () => import('./page/Home.vue') },
      { name: 'record', path: '/record', component: () => import('./page/Record.vue') },
    ],
  });
}
