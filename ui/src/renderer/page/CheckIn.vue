<template>
    <v-content>
        <v-container>
            <v-card class="my-3">
                <v-form @submit.prevent="submit">
                    <v-layout row wrap>
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
                    <v-btn
                            type="submit"
                            color="success"
                            :loading="loading"
                            :disabled="loading"
                    >Record</v-btn>
                    <v-btn
                            color="error"
                            @click="clear"
                            :loading="loading"
                            :disabled="loading"
                    >Clear</v-btn>
                </v-form>
            </v-card>
            <check-in-librarian :update.sync="updated"></check-in-librarian>
        </v-container>
        <v-snackbar
                bottom
                v-model="snackbar"
                :timeout="5000"
                :color="color"
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
        loading: false,
        snackbar: false,
        icon: 'done',
        color: 'success',
        message: '',
        updated: false,
      };
    },
    methods: {
      clear() {
        this.id = '';
      },
      submit() {
        this.loading = true;
        axios.post(`/api/checkin/${this.id}`, {}, {
          validateStatus: status => (status >= 200 && status < 500),
        })
          .then(({ status }) => {
            this.loading = false;
            if (status === 200) {
              this.icon = 'done';
              this.color = 'success';
              this.message = 'Finish Check In';
              console.debug('Check In');
            } else {
              this.icon = 'close';
              this.color = 'error';
              if (status === 404) {
                this.message = 'Record Not Found';
              } else if (status === 409) {
                this.message = 'Already Check In';
              } else {
                this.message = 'Error occur';
              }

              console.error(this.message);
            }

            this.snackbar = true;
            this.updated = true;
            this.clear();
          });
      },
    },
  };
</script>
