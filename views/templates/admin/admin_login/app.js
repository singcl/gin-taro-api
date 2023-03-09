// import { h } from '/views/static/js/vue/vue3.esm-browser.min.js';
import { h } from 'vue';
// import naive from '/views/static/js/vue/naive.min.js';
import naive from 'naive';
import Page from './admin_login.js';
const { NMessageProvider, NDialogProvider, NNotificationProvider } = naive;

export default {
  setup() {
    return () => h(NNotificationProvider, () => h(NMessageProvider, () => h(NDialogProvider, () => h(Page))));
  },
};
