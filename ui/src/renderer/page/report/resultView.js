import Personal from './personal/ReportView.vue';
import Overview from './overview/ReportView.vue';
import UnknownQuery from './UnknownQuery.vue';

const elements = {
  personal: Personal,
  overview: Overview,
};

export default {
  name: 'report-view',
  props: {
    report: {
      type: String,
      required: true,
    },
    result: {
      require: true,
    },
  },
  render(h) {
    if (!(this.report in elements)) {
      return h(UnknownQuery);
    }

    return h(elements[this.report], {
      on: this.$listeners,
      props: {
        result: this.result,
      },
    });
  },
};
