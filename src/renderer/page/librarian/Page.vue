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
                    :body="body"
            ></mutation>
            <list
                    :loading="loading"
                    :librarians="librarian"
                    @edit="edit"
                    @remove="removeConfirm"
            ></list>
        </v-container>
        <v-dialog v-model="confirm" max-width="290">
            <v-card>
                <v-card-title class="headline">Confirm to remove librarian?</v-card-title>
                <v-card-text>
                    Are you sure you want to remove {{ removeTarget.name }} ({{ removeTarget.class }}{{removeTarget.class_no}})?
                    All the record related to {{ removeTarget.name }} will be removed.
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="red darken-1" flat="flat" @click.native="remove">Remove</v-btn>
                    <v-btn color="darken-1" flat="flat" @click.native="confirm = false">Cancel</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
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
        confirm: false,
        removeTarget: {},
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
        this.usage = 'edit';
        this.dialog = true;
      },
      removeConfirm(body) {
        this.removeTarget = body;
        this.confirm = true;
      },
      remove() {
        const body = this.removeTarget;
        const query = {
          body,
          action: 'remove',
        };
        console.debug('Remove', query);
        this.$ipc.send('librarian.mutation', query);
        this.confirm = false;
      },
      ...mapActions(['updateLibrarian']),
    },
    watch: {
      dialog(val) {
        if (!val) {
          this.body = {};
        }
      },
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
