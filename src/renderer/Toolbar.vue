<template>
    <v-toolbar app color="primary" class="white--text">
        <v-toolbar-title>
            <router-link to="/">Librarian Check-In</router-link>
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items class="hidden-sm-and-down">
            <v-btn
                    v-for="page in pages"
                    flat
                    class="white--text"
                    :key="page.name"
                    exact
                    :to="{name: page.name}"
            >
                {{ page.display }}
            </v-btn>
            <v-btn flat class="white--text">Report</v-btn>
        </v-toolbar-items>
    </v-toolbar>
</template>

<style scoped lang="stylus">
    .toolbar__title a
        color: white
        text-decoration none

</style>

<script>
    export default {
      data() {
        return {
          pages: [
            { display: 'Check In', name: 'checkin' },
            { display: 'Check Out', name: 'checkout' },
            { display: 'Librarian', name: 'librarian' },
          ],
        };
      },
    };
</script>
