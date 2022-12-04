import { h /* reactive, ref, onMounted */ } from '/views/static/js/vue/vue3.esm-browser.js';
// import naive from '/public/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
// import Kiko from './../utils/kiko/Kiko.js';

export default {
  setup() {
    return () => h('div', () => '菜单');
  },
};
