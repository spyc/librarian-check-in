<template>
    <v-content>
        <v-container>
            <v-form @submit.prevent="submit">
                <v-card>
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
                </v-card>
            </v-form>
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
  export default {
    data() {
      return {
        id: '',
        loading: false,
        snackbar: false,
        icon: 'done',
        color: 'success',
        message: '',
      };
    },
    methods: {
      clear() {
        this.id = '';
      },
      submit() {
        this.$ipc.send('checkin', this.id);
        this.loading = true;
      },
      handleCheckInReply(event, result) {
        this.loading = false;
        if (result.done) {
          this.icon = 'done';
          this.color = 'success';
          this.message = 'Finish Check In';
          console.debug('Check In', result);
        } else {
          this.icon = 'close';
          this.color = 'error';
          if (result.error.code === 'SQLITE_CONSTRAINT') {
            this.message = 'Record Not Found';
          } else {
            this.message = result.error.code;
          }
          console.error(result.error);
        }
        this.snackbar = true;
        this.clear();
      },
    },
    mounted() {
      this.$ipc.on('checkin.reply', this.handleCheckInReply.bind(this));
    },
    destroyed() {
      this.$ipc.removeAllListeners('checkin.reply');
    },
  };
</script>