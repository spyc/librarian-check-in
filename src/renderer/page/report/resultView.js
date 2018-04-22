import Personal from './personal/ReportView.vue';
import UnknownQuery from './UnknownQuery.vue';

const elements = {
  personal: Personal,
};

export default {
  name: 'report-view',
  props: {
    report: {
      type: String,
      required: true,
    },
    rows: {
      type: Array,
      require: true,
    },
  },
  render(h) {
    if (!(this.report in elements)) {
      return h(UnknownQuery);
    }

    return h(elements[this.report], {
      props: {
        rows: this.rows,
      },
    });
  },
};
