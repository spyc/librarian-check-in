<template>
    <v-app id="app">
        <v-toolbar app color="primary" class="white--text">
            <v-toolbar-title>Librarian Check-In</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items class="hidden-sm-and-down">
                <v-btn flat class="white--text" exact :to="{name: 'record'}">Record</v-btn>
                <v-btn flat class="white--text">Personal Report</v-btn>
                <v-btn flat class="white--text">Summary Report</v-btn>
            </v-toolbar-items>
        </v-toolbar>
        <router-view></router-view>
        <v-footer app>
            <v-btn color="primary" small dark @click="reportIssue">
                <v-icon left dark>bug_report</v-icon>
                Report Bug
            </v-btn>
            <v-spacer></v-spacer>
            <div>&copy; 2018 Tony Yip</div>
        </v-footer>
    </v-app>
</template>

<style lang="stylus">
    @import '~vuetify/src/stylus/main';

    .footer.footer--fixed
        height: auto !important
</style>

<script>
  import { shell } from 'electron';

  export default {
    data() {
      const numbers = [];
      for (let i = 0; i < 100; i += 1) {
        numbers.push(i);
      }
      return {
        numbers,
      };
    },
    methods: {
      reportIssue() {
        shell.openExternal('https://github.com/spyc/library-checker/issues/new');
      },
    },
  };
</script>