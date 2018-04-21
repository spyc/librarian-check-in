<template>
    <v-content>
        <v-container>
            <v-card>
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
                    <v-btn color="success">Record</v-btn>
                    <v-btn color="error" @click="clear">Clear</v-btn>
                </v-form>
                <v-btn color="secondary">Check Out All</v-btn>
            </v-card>
        </v-container>
    </v-content>
</template>

<script>
  export default {
    data() {
      return {
        id: '',
        ranks: [
          'Good',
          'Fair',
          'Fail',
        ],
        rank: 'Fair',
      };
    },
    methods: {
      clear() {
        this.id = '';
        this.rank = 'Fair';
      },
    },
  };
</script>
