<template>
    <v-content>
        <v-container>
            <add></add>
            <list :loading="loading" :librarians="librarian"></list>
        </v-container>
    </v-content>
</template>

<script>
  import { mapState, mapActions } from 'vuex';

  export default {
    components: {
      List: () => import('./List.vue'),
      Add: () => import('./Add.vue'),
    },
    data() {
      return {
        loading: true,
      };
    },
    computed: mapState(['librarian']),
    methods: {
      fetchLibrarian() {
        this.$ipc.send('librarian.list');
        this.loading = true;
      },
      handleFetchLibrarian(e, librarian) {
        console.debug(librarian);
        this.loading = false;
        this.updateLibrarian(librarian);
      },
      ...mapActions(['updateLibrarian']),
    },
    mounted() {
      this.$ipc.on('librarian.list.reply', this.handleFetchLibrarian.bind(this));
      this.fetchLibrarian();
    },
    destroyed() {
      this.$ipc.removeAllListeners('librarian.list.reply');
    },
  };
</script>
