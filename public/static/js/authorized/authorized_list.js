import { h, reactive, onMounted } from '/public/static/js/vue/vue3.esm-browser.js';
import naive from '/public/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from './../utils/kiko/Kiko.js';

const useOptions = [
  {
    value: 1,
    label: '启用',
    type: 'success',
  },
  {
    value: -1,
    label: '禁用',
    type: 'warning',
  },
];

//
const defaultColumns = [
  { key: 'id', title: '编号' },
  { key: 'business_key', title: '调用方' },
  { key: 'business_developer', title: '对接人' },
  { key: 'created_at', title: '创建日期' },
  { key: 'updated_at', title: '更新日期' },
  {
    key: 'is_used',
    title: '状态',
    render(row) {
      const { is_used } = row;
      const itm = useOptions.find((m) => m.value == is_used);
      return itm ? h(naive.NTag, { type: itm.type }, itm.label) : is_used;
    },
  },
  {
    key: 'operation',
    title: '操作',
    render(row) {
      const { is_used } = row;
      const op_used = { 1: -1, '-1': 1 }[is_used];
      const itm = useOptions.find((m) => m.value == op_used);
      return h('div', [
        h(naive.NButton, { type: 'info', style: { marginRight: '10px' } }, '详情'),
        itm ? h(naive.NButton, { type: itm.type, style: { marginRight: '10px' } }, itm.label) : undefined,
        h(naive.NButton, { type: 'error' }, '删除'),
      ]);
    },
  },
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
