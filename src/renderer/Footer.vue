<template>
    <v-footer app>
        <v-btn flat color="primary" small dark @click="reportIssue">
            <v-icon left color="primary">bug_report</v-icon>
            Report Bug
        </v-btn>
        <v-spacer></v-spacer>
        <div>&copy; 2018 Tony Yip</div>
    </v-footer>
</template>

<style scoped lang="stylus">
    .footer.footer--fixed
        height: auto !important
</style>

<script>
  import { shell } from 'electron';

  export default {
    methods: {
      reportIssue() {
        shell.openExternal('https://github.com/spyc/librarian-check-in/issues/new');
      },
    },
  };
</script>
