<template>
    <v-app id="app">
        <app-toolbar></app-toolbar>
        <router-view></router-view>
        <app-footer></app-footer>
    </v-app>
</template>

<style lang="stylus">
    @import '~vuetify/src/stylus/main';
</style>

<script>
  export default {
    components: {
      AppFooter: () => import('./Footer.vue'),
      AppToolbar: () => import('./Toolbar.vue'),
    },
  };
</script>