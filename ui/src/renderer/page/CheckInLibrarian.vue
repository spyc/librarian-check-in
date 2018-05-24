<template>
    <v-card class="my-3">
        <v-card-title primary-title>
            <div>
                <h3 class="headline mb-0">Check In Librarian</h3>
            </div>
        </v-card-title>
        <v-card-actions>
            <v-btn
                    color="primary"
                    @click="getCheckIn"
                    :loading="loading"
            >
                <v-icon left>autorenew</v-icon>
                <span>Update</span>
            </v-btn>
        </v-card-actions>
        <list-librarian
                no-actions
                :loading="loading"
                :librarians="librarians"
        ></list-librarian>
    </v-card>
</template>

<script>
  export default {
    name: 'CheckInLibrarian',
    components: {
      ListLibrarian: () => import('./librarian/List.vue'),
    },
    data() {
      return {
        loading: false,
        librarians: [],
      };
    },
    props: {
      update: {
        type: Boolean,
        default: false,
        required: false,
      },
    },
    watch: {
      update(val) {
        if (val) {
          this.getCheckIn();
        }
      },
    },
    methods: {
      getCheckIn() {
        this.loading = true;
        this.$ipc.send('checkin.list');
      },
      handleCheckInListReply(event, librarians) {
        this.loading = false;
        this.librarians = librarians;
        this.$emit('update:update', false);
      },
    },
    mounted() {
      this.$ipc.on('checkin.list.reply', this.handleCheckInListReply.bind(this));
      this.getCheckIn();
    },
    destroyed() {
      this.$ipc.removeAllListeners('checkin.list.reply');
    },
  };
</script>
