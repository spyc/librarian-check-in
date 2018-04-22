<template>
    <v-content>
        <v-container>
            <v-btn
                    color="primary"
                    @click="fetchLibrarian"
                    :loading="loading"
            >
                <v-icon left>autorenew</v-icon>
                <span>Update</span>
            </v-btn>
            <mutation
                    @update="fetchLibrarian"
                    :display.sync="dialog"
                    :usage.sync="usage"
            ></mutation>
            <list
                    :loading="loading"
                    :librarians="librarian"
                    @click="edit"
            ></list>
        </v-container>
    </v-content>
</template>

<script>
  import { mapState, mapActions } from 'vuex';

  export default {
    components: {
      List: () => import('./List.vue'),
      Mutation: () => import('./Mutation.vue'),
    },
    data() {
      return {
        loading: true,
        dialog: false,
        usage: 'add',
        body: {},
      };
    },
    computed: mapState(['librarian']),
    methods: {
      fetchLibrarian() {
        this.$ipc.send('librarian.list');
        this.loading = true;
      },
      handleFetchLibrarian(e, librarian) {
        console.debug('Fetch Librarian', librarian);
        this.loading = false;
        this.updateLibrarian(librarian);
      },
      edit(body) {
        this.body = body;
        this.dialog = true;
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
