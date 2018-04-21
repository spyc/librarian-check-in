import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

export default function createRouter() {
  return new VueRouter({
    routes: [
      { name: 'home', path: '/', component: () => import('./page/Home.vue') },
      { name: 'checkout', path: '/checkout', component: () => import('./page/CheckOut.vue') },
      { name: 'checkin', path: '/checkin', component: () => import('./page/CheckIn.vue') },
    ],
  });
}
