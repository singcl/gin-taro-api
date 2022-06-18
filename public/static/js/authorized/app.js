import { h } from '/public/static/js/vue/vue3.esm-browser.js';
import naive from '/public/static/js/vue/naive.js';
import Page from './authorized_list.js';
const { NMessageProvider } = naive;

export default {
  setup() {
    return () => h(NMessageProvider, () => h(Page));
  },
};
