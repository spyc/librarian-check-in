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
                :items="rows"
                class="my-3"
        >
            <template slot="no-data">
                <v-alert :value="true" color="warning" icon="warning">
                    No record found
                </v-alert>
            </template>

            <template slot="items" slot-scope="props">
                <td>{{ formatDate(props.item.check_in) }}</td>
                <td>{{ formatDate(props.item.check_out) }}</td>
                <td>{{ ucfirst(props.item.rank) }}</td>
                <td>{{ duration(props.item.check_in, props.item.check_out, props.item.rank) }}</td>
            </template>
            <template slot="footer">
                <td colspan="3">
                    <strong>Total</strong>
                </td>
                <td>{{ total }}</td>
            </template>
        </v-data-table>
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
      rows: {
        type: Array,
        required: true,
      },
    },
    data() {
      return {
        headers,
      };
    },
    computed: {
      id() {
        return this.rows.length > 0 ? this.rows[0].id : '';
      },
      name() {
        return this.rows.length > 0 ? this.rows[0].name : '';
      },
      total() {
        return this.rows.reduce((total, row) => total + this.duration(row.check_in, row.check_out, row.rank), 0);
      },
    },
    methods: {
      formatDate(time) {
        return moment.unix(time).tz('Asia/Hong_Kong').format('Do MMM, YYYY HH:mm');
      },
      ucfirst(rank) {
        return `${rank.charAt(0).toUpperCase()}${rank.substr(1)}`;
      },
      duration(start, end, rank) {
        const endTime = moment.unix(end).tz('Asia/Hong_Kong');
        const startTime = moment.unix(start).tz('Asia/Hong_Kong');
        const diff = endTime.diff(startTime, 'minutes');
        if (rank === 'good') {
          return diff * 2;
        } else if (rank === 'fair') {
          return diff * 1;
        }
        return diff * 0.5;
      },
    },
  };
</script>
