<template>
    <v-content>
        <v-container>
            <v-card class="my-3">
                <v-form>
                    <v-layout row wrap>

                        <v-flex xs6>
                            <v-subheader>Review</v-subheader>
                        </v-flex>
                        <v-flex xs6>
                            <v-select
                                    :items="ranks"
                                    v-model="rank"
                                    label="Rank"
                                    single-line
                            ></v-select>
                        </v-flex>

                        <v-flex xs6>
                            <v-subheader>Student ID</v-subheader>
                        </v-flex>
                        <v-flex xs6>
                            <v-text-field
                                    autofocus
                                    name="pycid"
                                    label="Student ID"
                                    v-model="id"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                    <v-btn color="success" :loading="loading" :disabled="loading" @click="checkOut">Record</v-btn>
                    <v-btn color="error" :loading="loading" :disabled="loading" @click="clear">Clear</v-btn>
                </v-form>
                <v-btn color="secondary" :loading="loading" :disabled="loading" @click="batchCheckOut">Check Out All</v-btn>
            </v-card>
            <check-in-librarian :update.sync="updated"></check-in-librarian>
        </v-container>
        <v-snackbar
                v-model="snackbar"
                :color="snackbarColor"
        >
            <v-icon color="white">{{ icon }}</v-icon>
            {{ message }}
            <v-btn flat color="white" @click.native="snackbar = false">Close</v-btn>
        </v-snackbar>
    </v-content>
</template>

<script>
  const ranks = [
    'Good',
    'Fair',
    'Fail',
  ];

  export default {
    components: {
      CheckInLibrarian: () => import('./CheckInLibrarian.vue'),
    },
    data() {
      return {
        ranks,
        id: '',
        rank: 'Fair',
        snackbar: false,
        snackbarColor: 'success',
        message: '',
        icon: '',
        loading: false,
        updated: false,
      };
    },
    methods: {
      clear() {
        this.id = '';
        this.rank = 'Fair';
      },
      checkOut() {
        this.loading = true;
        this.$ipc.send('checkout', { id: this.id, rank: this.rank });
      },
      batchCheckOut() {
        this.loading = true;
        this.$ipc.send('checkout.batch');
      },
      handleCheckOutReply(event, { done, error }) {
        this.loading = false;
        if (done) {
          this.icon = 'done';
          this.snackbarColor = 'success';
          this.message = 'Finish Check Out';
          console.debug('Check Out');
        } else {
          this.icon = 'close';
          this.color = 'error';
          if (error.code === 'SQLITE_CONSTRAINT') {
            this.message = 'Record Not Found';
          } else {
            this.message = error.code;
          }
          console.error(error);
        }
        this.snackbar = true;
        this.updated = true;
        this.clear();
      },
    },
    mounted() {
      this.$ipc.on('checkout.reply', this.handleCheckOutReply.bind(this));
    },
  };
</script>
