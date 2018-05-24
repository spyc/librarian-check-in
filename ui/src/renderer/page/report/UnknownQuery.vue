<template>
    <v-layout row>
        <v-flex xs12>
            <v-alert type="error" :value="true">
                Unknown Report Type
            </v-alert>
        </v-flex>
    </v-layout>
</template>

<script>
  export default {
    name: 'UnknownQuery',
  };
</script>
