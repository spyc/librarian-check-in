<template>
    <v-dialog
            persistent
            v-model="modal"
            lazy
            full-width
            width="290px"
    >
        <v-text-field
                slot="activator"
                :label="label"
                v-model="val"
                prepend-icon="event"
                readonly
        ></v-text-field>
        <v-date-picker v-model="val" scrollable :allowed-dates="allowedDates">
            <v-spacer></v-spacer>
            <v-btn flat color="primary" @click="modal = false">OK</v-btn>
        </v-date-picker>
    </v-dialog>
</template>

<script>
  import moment from 'moment-timezone';

  export default {
    name: 'DateField',
    props: {
      date: {
        required: false,
        default: null,
      },
      label: {
        type: String,
        required: true,
      },
      validateDate: {
        type: Function,
        required: false,
      },
    },
    computed: {
      val: {
        get() {
          return this.date;
        },
        set(val) {
          this.$emit('update:date', val);
        },
      },
    },
    data() {
      return {
        modal: false,
      };
    },
    methods: {
      allowedDates(val) {
        const target = moment(val).tz('Asia/Hong_Kong');
        const result = target.isBefore(moment());
        if (this.validateDate) {
          return result && this.validateDate(target);
        }
        return result;
      },
    },
  };
</script>