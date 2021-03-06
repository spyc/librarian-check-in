<template>
    <v-content>
        <v-container>
            <v-card class="my-3">
                <v-card-title primary-title>
                    <div>
                        <h3 class="headline mb-0">Report Type</h3>
                    </div>
                </v-card-title>
                <v-card-text>
                    <v-layout row>
                        <v-flex xs2>
                            <v-subheader>Report Type</v-subheader>
                        </v-flex>
                        <v-flex xs10>
                            <v-select
                                    :items="reportTypes"
                                    v-model="report"
                                    label="Select"
                                    single-line
                            ></v-select>
                        </v-flex>
                    </v-layout>
                </v-card-text>
            </v-card>
            <query :report="report" :loading="loading" @query="query"></query>
            <v-progress-linear v-if="loading" indeterminate color="primary"></v-progress-linear>
            <v-alert v-else-if="result.error" type="error" :value="true">{{ result.error.code }}</v-alert>
            <report-view
                    v-else-if="hasResult"
                    :report="report"
                    :result="result"
                    @export="handleFinishExport"
                    @error="handleError"
            ></report-view>
        </v-container>
        <v-snackbar
                bottom
                v-model="snackbar"
                :color="snackbarColor"
                :timeout="5000"
        >
            {{ message }}
            <v-btn flat color="white" @click.native="snackbar = false">Close</v-btn>
        </v-snackbar>
    </v-content>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'report',
    components: {
      Query: () => import('./query'),
      ReportView: () => import('./resultView'),
    },
    data() {
      return {
        loading: false,
        report: 'personal',
        reportTypes: [
          { value: 'personal', text: 'Report By Person' },
          { value: 'overview', text: 'Overview Report' },
        ],
        hasResult: false,
        result: { error: null },
        message: '',
        snackbar: false,
        snackbarColor: 'success',
      };
    },
    watch: {
      report(newVal, oldVal) {
        if (newVal !== oldVal) {
          this.hasResult = false;
          this.result = { error: null };
        }
      },
    },
    methods: {
      query(q) {
        console.log('Query', q);
        this.loading = true;
        axios.get(`/api/report/${this.report}`, { params: q })
          .then(({ data, status }) => {
            this.loading = false;
            if (status === 200) {
              this.hasResult = true;
            }
            console.debug(data);
            this.result = data;
          });
      },
      handleFinishExport() {
        this.snackbarColor = 'success';
        this.message = 'Finish Export File';
        this.snackbar = true;
      },
      handleError(err) {
        this.snackbarColor = 'error';
        this.message = `Error: ${err}`;
        this.snackbar = true;
      },
    },
  };
</script>
