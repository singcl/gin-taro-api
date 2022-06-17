import { h, reactive, onMounted } from '/public/static/js/vue/vue3.esm-browser.js';
import naive from '/public/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from './../utils/kiko/Kiko.js';

//
const defaultColumns = [
  { key: 'id', title: '编号' },
  { key: 'business_key', title: '调用方' },
  { key: 'business_developer', title: '对接人' },
  { key: 'created_at', title: '创建日期' },
  { key: 'updated_at', title: '更新日期' },
  { key: 'is_used', title: '状态' },
  { key: 'operation', title: '操作' },
];

// App
export default {
  setup() {
    const columns = reactive(defaultColumns);
    let tableData = reactive([]);
    //
    onMounted(async () => {
      const response = await new Kiko().fetch('/api/authorized');
      tableData.splice(0, tableData.length);
      tableData.push(...response.list);
    });
    //
    return () =>
      h('div', [
        h(naive.NAlert, { type: 'success', title: '授权信息' }, '已授权信息表格'),
        h('div', [
          h(naive.NDataTable, {
            columns: columns,
            bordered: false,
            data: tableData,
          }),
        ]),
      ]);
  },
};
