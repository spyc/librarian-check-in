<template>
    <v-content>
        <v-container>
            <v-card class="my-3">
                <v-form @submit="checkOut">
                    <v-layout row wrap>
                        <v-flex xs6>
                            <v-subheader>Student ID</v-subheader>
                        </v-flex>
                        <v-flex xs6>
                            <v-text-field
                                    autofocus
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
  import axios from 'axios';

  export default {
    components: {
      CheckInLibrarian: () => import('./CheckInLibrarian.vue'),
    },
    data() {
      return {
        id: '',
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
      },
      checkOut(e) {
        if (e) {
          e.preventDefault();
        }
        this.loading = true;
        axios.post(`/api/checkout/${this.id}`)
          .then(({ status }) => {
            this.handleCheckOutReply(status);
          })
          .catch(console.error);
      },
      batchCheckOut() {
        this.loading = true;
        axios.post('/api/checkout')
          .then(({ status }) => {
            this.handleCheckOutReply(status);
          })
          .catch(console.error);
      },
      handleCheckOutReply(status) {
        this.loading = false;
        if (status === 200) {
          this.icon = 'done';
          this.snackbarColor = 'success';
          this.message = 'Finish Check Out';
          console.debug('Check Out');
        } else {
          this.icon = 'close';
          this.color = 'error';
          if (status === 404) {
            this.message = 'Record Not Found';
          } else {
            this.message = 'Error Occur';
          }
          console.error(status);
        }
        this.snackbar = true;
        this.updated = true;
        this.clear();
      },
    },
  };
</script>
