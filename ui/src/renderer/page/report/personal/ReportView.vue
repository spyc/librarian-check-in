<template>
    <v-card class="my-3">
        <v-card-title primary-title>
            <div>
                <h3 class="headline mb-0">Report</h3>
                <div>ID: {{ id }}</div>
                <div>Name: {{ name }}</div>
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
                <td>{{ props.item.check_in }}</td>
                <td>{{ props.item.check_out }}</td>
                <td>{{ props.item.rank }}</td>
                <td>{{ props.item.total }}</td>
            </template>
            <template slot="footer">
                <td colspan="3">
                    <strong>Total</strong>
                </td>
                <td>{{ total }}</td>
            </template>
        </v-data-table>
        <v-card-actions>
            <v-btn
                    flat
                    color="primary"
                    :loading="exporting"
                    :disabled="exporting"
                    :href="csvLink"
                    download="record.csv"
            >Export</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
  import moment from 'moment-timezone';

  const headers = [
    { text: 'Check In', value: 'check_in' },
    { text: 'Check Out', value: 'check_out' },
    { text: 'Rank', value: 'rank' },
    { text: 'Total (min)', value: 'total', sortable: false },
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
      id() {
        return this.result ? this.result.librarian.pycid : '';
      },
      name() {
        return this.result ? this.result.librarian.name : '';
      },
      total() {
        if (!this.result) {
          return 0;
        }
        return this.result.records.reduce((total, row) => {
          const start = this.parseDate(row.check_in);
          const end = this.parseDate(row.check_out);
          return total + this.duration(start, end, row.rank);
        }, 0);
      },
      dataRows() {
        if (!this.result || !this.result.records) {
          return [];
        }
        return this.result.records.map((row) => {
          const start = this.parseDate(row.check_in);
          const end = this.parseDate(row.check_out);
          return {
            check_in: this.formatDate(start),
            check_out: this.formatDate(end),
            rank: this.ucfirst(row.rank),
            total: this.duration(start, end, row.rank),
          };
        });
      },
      csvContent() {
        if (this.dataRows.length === 0) {
          return '';
        }

        return this.dataRows.map(row => `"${row.check_in}",${row.check_out},${row.rank},${row.total}`).join('\r\n');
      },
      csvLink() {
        return `data:text/csv;charset=utf-8,"Check In","Check Out",Rank,Total\r\n${this.csvContent}`;
      },
    },
    methods: {
      parseDate(time) {
        return moment(time).tz('Asia/Hong_Kong');
      },
      formatDate(time) {
        return time.format('Do MMM, YYYY HH:mm');
      },
      ucfirst(rank) {
        return `${rank.charAt(0).toUpperCase()}${rank.substr(1)}`;
      },
      duration(startTime, endTime, rank) {
        const diff = endTime.diff(startTime, 'minutes');
        if (rank === 'good') {
          return diff * 2;
        } else if (rank === 'fair') {
          return diff * 1;
        }
        return diff * 0.5;
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
