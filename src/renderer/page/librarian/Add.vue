<template>
    <v-layout row>
        <v-dialog v-model="display" fullscreen hide-overlay transition="dialog-bottom-transition">
            <template slot="activator">
                <v-btn color="primary" dark @click="display = true">
                    <v-icon left dark>add</v-icon>
                    Add Librarian
                </v-btn>
            </template>
            <v-card>
                <v-toolbar dark color="primary">
                    <v-btn icon @click.native="display = false" dark>
                        <v-icon>close</v-icon>
                    </v-btn>
                    <v-toolbar-title>Add Librarian</v-toolbar-title>
                    <v-spacer></v-spacer>
                    <v-toolbar-items>
                        <v-btn dark flat @click.native="display = false">Add</v-btn>
                    </v-toolbar-items>
                </v-toolbar>

                <v-container>
                    <v-layout row>
                        <v-flex xs2>Student ID</v-flex>
                        <v-flex xs10>
                            <v-text-field
                                    v-model="id"
                                    label="Student ID"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                    <v-layout row>
                        <v-flex xs2>Name</v-flex>
                        <v-flex xs10>
                            <v-text-field
                                    v-model="name"
                                    label="Name"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                    <v-layout row>
                        <v-flex xs2>Class</v-flex>
                        <v-flex xs10>
                            <v-text-field
                                    v-model="studentClass"
                                    label="Class"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                    <v-layout row>
                        <v-flex xs2>Class No</v-flex>
                        <v-flex xs10>
                            <v-text-field
                                    v-model="classNo"
                                    label="Class No"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                </v-container>
            </v-card>
        </v-dialog>
    </v-layout>
</template>

<style scoped lang="stylus">
    .layout.row
        justify-content flex-end
</style>

<script>
  export default {
    name: 'AddLibrarian',
    data() {
      return {
        display: false,
        id: '',
        name: '',
        studentClass: '',
        classNo: '',
      };
    },
    methods: {
      clear() {
        this.id = '';
        this.name = '';
        this.studentClass = '';
        this.classNo = '';
      },
    },
  };
</script>
