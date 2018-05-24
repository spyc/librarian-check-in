<template>
    <v-card class="my-3">
        <v-card-title primary-title>
            <div>
                <h3 class="headline mb-0">Report</h3>
            </div>
        </v-card-title>
        <v-data-table
                :headers="headers"
                :items="dataRows"
                class="my-3"
        >
            <template slot="no-data">
                <v-alert :value="true" color="warning" icon="warning">
                    No record found
                </v-alert>
            </template>

            <template slot="items" slot-scope="props">
                <td>{{ props.item.id }}</td>
                <td>{{ props.item.name }}</td>
                <td>{{ props.item.fair_time }}</td>
                <td>{{ props.item.good_time }}</td>
                <td>{{ props.item.fail_time }}</td>
            </template>
        </v-data-table>
        <v-card-actions>
            <v-btn
                    flat
                    color="primary"
                    :loading="exporting"
                    :disabled="exporting"
                    :href="csvLink"
                    download="report.csv"
            >Export</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
  import moment from 'moment-timezone';

  const headers = [
    { text: 'Student ID', value: 'id' },
    { text: 'Name', value: 'name' },
    { text: 'Fair (min)', value: 'fair_time' },
    { text: 'Good (min)', value: 'good_time' },
    { text: 'Fail (min)', value: 'fail_time' },
  ];

  export default {
    name: 'PersonalReportView',
    props: {
      result: {
        required: true,
      },
    },
    data() {
      return {
        headers,
        exporting: false,
      };
    },
    computed: {
      dataRows() {
        return this.result.map(row => ({
          id: row.pycid,
          name: row.name,
          fair_time: row.fair_time,
          good_time: row.good_time,
          fail_time: row.fail_time,
        }));
      },
      csvContent() {
        return this.dataRows
          .map(row => `${row.id},${row.name},${row.fair_time},${row.good_time},${row.fail_time}`)
          .join('\r\n');
      },
      csvLink() {
        return `data:text/csv;charset=utf-8,"Student ID",Name,"Fair (min)","Good (min)","Fail (min)"\r\n${this.csvContent}`;
      },
    },
    methods: {
      parseDate(time) {
        return moment.unix(time).tz('Asia/Hong_Kong');
      },
      formatDate(time) {
        return time.format('Do MMM, YYYY HH:mm');
      },
      ucfirst(rank) {
        return `${rank.charAt(0).toUpperCase()}${rank.substr(1)}`;
      },
      handleExportDone(event, { success, filename, error }) {
        this.exporting = false;
        if (success) {
          console.debug('Finish Export', filename);
          this.$emit('export', filename);
        } else {
          console.error('Finish Export', error);
          this.$emit('error', error);
        }
      },
    },
  };
</script>
