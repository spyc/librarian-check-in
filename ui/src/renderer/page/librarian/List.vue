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
        <template slot="items" slot-scope="props">
            <td>{{ props.item.pycid }}</td>
            <td>{{ props.item.name }}</td>
            <td>{{ props.item.class }}</td>
            <td>{{ props.item.class_no }}</td>
            <td class="justify-center layout px-0" v-if="!noActions">
                <v-btn icon class="mx-0" @click="() => edit(props.item)">
                    <v-icon color="teal">edit</v-icon>
                </v-btn>
                <v-btn icon class="mx-0" @click="() => remove(props.item)">
                    <v-icon color="pink">delete</v-icon>
                </v-btn>
            </td>
        </template>
    </v-data-table>
</template>

<script>
  export default {
    name: 'ListLibrarian',
    computed: {
      headers() {
        if (this.noActions) {
          return [
            { text: 'Student ID', value: 'id' },
            { text: 'Name', value: 'name' },
            { text: 'Class', value: 'class' },
            { text: 'Class No.', value: 'class_no' },
          ];
        }
        return [
          { text: 'Student ID', value: 'id' },
          { text: 'Name', value: 'name' },
          { text: 'Class', value: 'class' },
          { text: 'Class No.', value: 'class_no' },
          { text: 'Actions', value: 'action', sortable: false },
        ];
      },
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
      noActions: {
        required: false,
        default: false,
        type: Boolean,
      },
    },
    methods: {
      edit(item) {
        this.$emit('edit', item);
      },
      remove(item) {
        this.$emit('remove', item);
      },
    },
  };
</script>
