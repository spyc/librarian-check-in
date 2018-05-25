<template>
    <v-content>
        <v-container>
            <v-card class="my-3">
                <v-form @submit.prevent="submit">
                    <v-layout row wrap>
                        <v-flex xs6>
                            <v-subheader>Passcode</v-subheader>
                        </v-flex>
                        <v-flex xs6>
                            <v-text-field
                                    autofocus
                                    label="Passcode"
                                    v-model="passcode"
                                    maxlength="6"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                    <v-btn
                            type="submit"
                            color="success"
                            :loading="loading"
                            :disabled="loading"
                            @click="submit"
                    >Authorize</v-btn>
                    <v-btn
                            color="error"
                            @click="clear"
                            :loading="loading"
                            :disabled="loading"
                    >Clear</v-btn>
                </v-form>
            </v-card>
        </v-container>
        <v-snackbar
                bottom
                v-model="snackbar"
                :timeout="5000"
                color="error"
        >
            <v-icon color="white">error</v-icon>
            Authenticate Fail
            <v-btn flat color="white" @click.native="snackbar = false">Close</v-btn>
        </v-snackbar>
    </v-content>
</template>

<script>
    import axios from 'axios';
    import { createNamespacedHelpers, mapActions } from 'vuex';

    const { mapState } = createNamespacedHelpers('route');

    export default {
      name: 'AuthPage',
      data() {
        return {
          loading: false,
          passcode: '',
          snackbar: false,
        };
      },
      computed: mapState(['query']),
      methods: {
        clear() {
          this.passcode = '';
        },
        submit(e) {
          e.preventDefault();
          this.loading = true;
          axios.post('/api/auth', {
            passcode: this.passcode,
          })
            .then(() => {
              this.loading = false;
              this.updateAuth();
              this.$router.push(this.query.next || '/');
            })
            .catch((err) => {
              this.loading = false;
              console.error(err);
              this.snackbar = true;
            });
        },
        ...mapActions(['updateAuth']),
      },
    };
</script>
