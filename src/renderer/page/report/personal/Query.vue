<template>
    <v-card class="my-3">
        <v-form ref="form" lazy-validation @submit.prevent="submit">
            <v-layout row>
                <v-flex xs2>
                    <v-subheader>Student ID</v-subheader>
                </v-flex>
                <v-flex xs10>
                    <v-text-field
                            name="pycid"
                            label="Student ID"
                            v-model="id"
                            required
                            :rules="[v => !!v || 'Required ID']"
                    ></v-text-field>
                </v-flex>
            </v-layout>
            <v-layout row>
                <v-flex xs6>
                    <date-field label="Start Date" :date.sync="startDate"></date-field>
                </v-flex>
                <v-flex xs6>
                    <date-field
                            label="End Date"
                            :date.sync="endDate"
                            :validateDate="checkDate"
                    ></date-field>
                </v-flex>
            </v-layout>
            <v-card-actions>
                <v-btn
                        color="primary"
                        type="submit"
                        :loading="loading"
                        :disabled="loading"
                >Query</v-btn>
            </v-card-actions>
        </v-form>
    </v-card>
</template>

<script>
  import moment from 'moment-timezone';

  const sql = `
  SELECT librarian.id, name, check_in, check_out, rank FROM librarian
  INNER JOIN record ON record.id = librarian.id
  WHERE librarian.id = ? AND check_in BETwEEN ? AND ?
  ORDER BY check_in ASC
  `;

  export default {
    name: 'PersonalReportQuery',
    components: {
      DateField: () => import('../common/DateField.vue'),
    },
    props: {
      loading: {
        type: Boolean,
        required: true,
      },
    },
    data() {
      return {
        id: '',
        startDate: null,
        endDate: null,
      };
    },
    methods: {
      submit() {
        if (this.startDate && this.endDate && this.$refs.form.validate()) {
          this.$emit('query', {
            sql,
            param: [
              this.id,
              moment(`${this.startDate}T00:00:00+08:00`).tz('Asia/Hong_Kong').unix(),
              moment(`${this.endDate}T23:59:59+08:00`).tz('Asia/Hong_Kong').unix(),
            ],
          });
        }
      },
      checkDate(val) {
        return val.isSameOrAfter(moment(this.startDate).tz('Asia/Hong_Kong'));
      },
    },
  };
</script>
