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
                    @click="exportFile"
            >Export</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
  import moment from 'moment-timezone';
  // import { remote } from 'electron';

  // const { dialog } = remote;

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
      rows: {
        type: Array,
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
        return this.rows.map(row => ({
          id: row.id,
          name: row.name,
          fair_time: row.fair_time / 60,
          good_time: row.good_time / 60,
          fail_time: row.fail_time / 60,
        }));
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
      exportFile() {
        // const filename = dialog.showSaveDialog({
        //   title: 'Export file',
        //   filters: [
        //     { name: 'CSV files', extensions: ['csv'] },
        //   ],
        //   properties: ['createDirectory'],
        // });
        // if (!filename) {
        //   return;
        // }
        // console.debug('Export to ', filename);
        // this.exporting = true;
        // this.$ipc.send('save.file', {
        //   filename,
        //   headers: this.headers.map(header => header.text),
        //   rows: this.dataRows,
        // });
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
    mounted() {
      this.$ipc.on('save.file.done', this.handleExportDone.bind(this));
    },
    destroyed() {
      this.$ipc.removeAllListeners('save.file.done');
    },
  };
</script>
