<template>
    <v-content>
        <v-container>
            <add></add>
            <list :loading="loading" :librarians="items"></list>
        </v-container>
    </v-content>
</template>

<script>
  export default {
    components: {
      List: () => import('./List.vue'),
      Add: () => import('./Add.vue'),
    },
    data() {
      return {
        loading: true,
        items: [],
      };
    },
  };
</script>
