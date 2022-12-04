import { h } from '/views/static/js/vue/vue3.esm-browser.js';
import naive from '/views/static/js/vue/naive.js';
import Page from './menu_view.js';
const { NMessageProvider, NDialogProvider } = naive;

export default {
  setup() {
    return () => h(NMessageProvider, () => h(NDialogProvider, () => h(Page)));
  },
};
