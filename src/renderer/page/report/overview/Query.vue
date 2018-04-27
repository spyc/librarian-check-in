<template>
    <v-card class="my-3">
        <v-form ref="form" lazy-validation @submit.prevent="submit">
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
  SELECT
    librarian.id,
    name,
    SUM(CASE rank WHEN 'fair' THEN 1 * (check_out - check_in) ELSE 0 END) AS fair_time,
    SUM(CASE rank WHEN 'good' THEN 2 * (check_out - check_in) ELSE 0 END) AS good_time,
    SUM(CASE rank WHEN 'fail' THEN 0.5 * (check_out - check_in) ELSE 0 END) AS fail_time
  FROM librarian
  INNER JOIN record ON record.id = librarian.id
  WHERE (check_in BETwEEN ? AND ?) AND check_out IS NOT NULL
  GROUP BY librarian.id
  ORDER BY check_in ASC
  `;

  export default {
    name: 'OverviewReportQuery',
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
        startDate: null,
        endDate: null,
      };
    },
    methods: {
      submit() {
        if (this.startDate && this.endDate && this.$refs.form.validate()) {
          this.$emit('query', {
            sql,
            params: [
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
