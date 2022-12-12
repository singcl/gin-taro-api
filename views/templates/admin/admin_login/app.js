import { h } from '/views/static/js/vue/vue3.esm-browser.js';
import naive from '/views/static/js/vue/naive.js';
import Page from './admin_login.js';
const { NMessageProvider, NDialogProvider, NNotificationProvider } = naive;

export default {
  setup() {
    return () => h(NNotificationProvider, () => h(NMessageProvider, () => h(NDialogProvider, () => h(Page))));
  },
};
