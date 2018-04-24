import Personal from './personal/Query.vue';
import Overview from './overview/Query.vue';
import UnknownQuery from './UnknownQuery.vue';

const elements = {
  personal: Personal,
  overview: Overview,
};

export default {
  name: 'report-query',
  props: {
    report: {
      type: String,
      required: true,
    },
    loading: {
      type: Boolean,
      required: true,
    },
  },
  render(h) {
    if (this.report in elements) {
      return h(elements[this.report], {
        props: {
          loading: this.loading,
        },
        on: this.$listeners,
      });
    }

    return h(UnknownQuery);
  },
};
