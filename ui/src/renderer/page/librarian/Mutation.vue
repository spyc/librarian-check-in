<template>
    <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
        <template slot="activator">
            <v-btn color="primary" dark @click="() => {this.$emit('update:usage', 'add');this.open();}">
                <v-icon left dark>add</v-icon>
                Add Librarian
            </v-btn>
        </template>
        <v-form lazy-validation v-model="validate" ref="form" @submit.prevent="submit">
            <v-card>
                <v-toolbar dark color="primary">
                    <v-btn icon @click.native="close" dark>
                        <v-icon>close</v-icon>
                    </v-btn>
                    <v-toolbar-title>{{ title }} Librarian</v-toolbar-title>
                    <v-spacer></v-spacer>
                    <v-toolbar-items>
                        <v-btn type="submit" dark flat :disabled="!validate">{{ title }}</v-btn>
                    </v-toolbar-items>
                </v-toolbar>

                <v-container>
                    <v-layout row>
                        <v-flex xs2>Student ID</v-flex>
                        <v-flex xs10>
                            <v-text-field
                                    v-model="pycid"
                                    label="Student ID"
                                    required
                                    :rules="rules.id"
                                    :readonly="usage === 'edit'"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                    <v-layout row>
                        <v-flex xs2>Name</v-flex>
                        <v-flex xs10>
                            <v-text-field
                                    v-model="name"
                                    label="Name"
                                    required
                                    :rules="rules.name"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                    <v-layout row>
                        <v-flex xs2>Class</v-flex>
                        <v-flex xs10>
                            <v-text-field
                                    v-model="studentClass"
                                    label="Class"
                                    required
                                    :rules="rules.class"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                    <v-layout row>
                        <v-flex xs2>Class No</v-flex>
                        <v-flex xs10>
                            <v-text-field
                                    v-model="classNo"
                                    label="Class No"
                                    required
                                    type="Number"
                                    min="1"
                                    step="1"
                                    :rules="rules.class_no"
                            ></v-text-field>
                        </v-flex>
                    </v-layout>
                </v-container>
            </v-card>
            <v-snackbar :timeout="5000" bottom v-model="snackbar">
                {{ message }}
                <v-btn flat color="pink" @click.native="snackbar = false">Close</v-btn>
            </v-snackbar>
        </v-form>
    </v-dialog>
</template>
]

<script>
  import axios from 'axios';

  const classRegexp = /^[1-6][A-F]$/;

  const usage = ['add', 'edit'];

  const rules = {
    id: [
      v => !!v || 'You must input Student ID!',
      v => v.startsWith('pyc') || 'Invalid ID',
    ],
    name: [
      v => !!v || 'You must input Name!',
    ],
    class: [
      v => !!v || 'You must input Class!',
      v => v.length === 2 || 'You can only input 2 character!',
      v => classRegexp.test(v) || 'You can only input 2 character!',
    ],
    class_no: [
      v => !!v || 'You must input Class No!',
    ],
  };

  export default {
    name: 'LibrarianMutation',
    props: {
      display: {
        type: Boolean,
        required: true,
        default: false,
      },
      usage: {
        type: String,
        required: true,
        validator(value) {
          return usage.indexOf(value) !== -1;
        },
      },
      body: {
        type: Object,
        required: false,
      },
    },
    watch: {
      body(newVal) {
        console.log('Body Change');
        if (!newVal) {
          return;
        }
        this.pycid = newVal.pycid;
        this.name = newVal.name;
        this.studentClass = newVal.class;
        this.classNo = newVal.class_no;
      },
    },
    data() {
      return {
        rules,
        pycid: '',
        name: '',
        studentClass: '',
        classNo: '',
        validate: true,
        snackbar: false,
        message: '',
      };
    },
    computed: {
      dialog: {
        get() {
          return this.display;
        },
        set(val) {
          this.$emit('update:display', val);
        },
      },
      title() {
        return this.usage.charAt(0).toUpperCase() + this.usage.substr(1);
      },
    },
    methods: {
      clear() {
        this.pycid = '';
        this.name = '';
        this.studentClass = '';
        this.classNo = '';
        this.validate = true;
      },
      open() {
        this.$emit('update:display', true);
      },
      close() {
        this.$emit('update:display', false);
        this.clear();
      },
      submit() {
        if (this.$refs.form.validate()) {
          const body = {
            pycid: this.pycid,
            name: this.name,
            class: this.studentClass,
            class_no: this.classNo,
          };
          console.debug('Mutation', this.usage, body);
          if (this.usage === 'edit') {
            axios.put(`/api/librarian/${body.pycid}`, body)
              .then((response) => {
                console.debug('librarian.mutation.reply', response);
                if (response.status === 200) {
                  this.close();
                  this.$emit('update');
                } else {
                  this.message = response.data;
                  this.snackbar = true;
                }
              });
          } else {
            axios.post('/api/librarian', body)
              .then((response) => {
                console.debug('librarian.mutation.reply', response);
                if (response.status === 201) {
                  this.close();
                  this.$emit('update');
                } else {
                  this.message = response.data;
                  this.snackbar = true;
                }
              });
          }
        } else {
          this.message = 'Incorrect form';
          this.snackbar = true;
        }
      },
    },
  };
</script>
