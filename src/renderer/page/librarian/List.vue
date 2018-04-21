<template>
    <v-data-table
            :headers="headers"
            :items="librarians"
            class="elevation-1 my-3"
            :loading="loading"
    >
        <v-progress-linear slot="progress" color="secondary" indeterminate></v-progress-linear>
        <template slot="no-data">
            <v-alert :value="true" color="error" icon="warning">
                No record found
            </v-alert>
        </template>
    </v-data-table>
</template>

<script>
  export default {
    name: 'ListLibrarian',
    data() {
      return {
        headers: [
          { text: 'Student ID', value: 'id' },
          { text: 'Name', value: 'name' },
          { text: 'Class', value: 'class' },
          { text: 'Class No.', value: 'class_no' },
        ],
      };
    },
    props: {
      loading: {
        required: true,
        default: true,
        type: Boolean,
      },
      librarians: {
        required: true,
        default: [],
        type: Array,
      },
    },
  };
</script>
