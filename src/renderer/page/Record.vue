<template>
    <v-content app>
        <v-container>
            <v-card>
                <v-tabs v-model="active" color="primary" dark slider-color="secondary">
                    <v-tab ripple>Check In</v-tab>
                    <v-tab-item>
                        <check-in></check-in>
                    </v-tab-item>
                </v-tabs>
            </v-card>
        </v-container>
    </v-content>
</template>

<script>
  export default {
    components: {
      CheckIn: () => import('./CheckIn.vue'),
    },
    data() {
      return {
        active: null,
      };
    },
  };
</script>
