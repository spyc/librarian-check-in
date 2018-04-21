<template>
    <v-form>
        <v-container>
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
            <v-btn color="success">Record</v-btn>
            <v-btn color="error" @click="clear">Clear</v-btn>
        </v-container>
    </v-form>
</template>

<script>
  export default {
    data() {
      return {
        id: '',
      };
    },
    methods: {
      clear() {
        this.id = '';
      },
    },
  };
</script>