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
  import axios from 'axios';

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
        axios.get('/api/checkin')
          .then(({ data }) => {
            this.loading = false;
            this.librarians = data;
            this.$emit('update:update', false);
          });
      },
    },
    mounted() {
      this.getCheckIn();
    },
  };
</script>
