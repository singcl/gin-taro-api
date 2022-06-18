import { h, reactive, ref, onMounted } from '/public/static/js/vue/vue3.esm-browser.js';
import naive from '/public/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from './../utils/kiko/Kiko.js';

//
import AddDialog from './components/authorized_add.js';

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
  { key: 'hashid', title: 'hashid' },
  { key: 'business_key', title: '调用方' },
  { key: 'business_secret', title: '调用方Secret' },
  { key: 'business_developer', title: '对接人' },
  {
    key: 'created_at',
    title: '创建日期',
    render(row) {
      const { created_at } = row;
      return created_at.split(/\s/)[0];
    },
  },
  {
    key: 'updated_at',
    title: '更新日期',
    render(row) {
      const { created_at } = row;
      return created_at.split(/\s/)[0];
    },
  },
  { key: 'remark', title: '备注' },
  {
    key: 'is_used',
    title: '状态',
    render(row) {
      const { is_used } = row;
      const itm = useOptions.find((m) => m.value == is_used);
      return itm ? h(naive.NTag, { type: itm.type }, () => itm.label) : is_used;
    },
  },
  {
    key: 'operation',
    title: '操作',
    align: 'right',
    render(row) {
      const { is_used } = row;
      const op_used = { 1: -1, '-1': 1 }[is_used];
      const itm = useOptions.find((m) => m.value == op_used);
      return h('div', [
        h(naive.NButton, { type: 'info', style: { marginRight: '10px' } }, () => '详情'),
        itm ? h(naive.NButton, { type: itm.type, style: { marginRight: '10px' } }, () => itm.label) : undefined,
        h(naive.NButton, { type: 'error' }, () => '删除'),
      ]);
    },
  },
];

// App
export default {
  setup() {
    const columns = reactive(defaultColumns);
    const tableData = ref([]);
    const adModalVisible = ref(false);
    //
    onMounted(async () => {
      await handleSearch();
    });
    //
    async function handleSearch() {
      const response = await new Kiko().fetch('/api/authorized');
      tableData.value = response.list;
    }
    //
    function handleAddAuth() {
      adModalVisible.value = true;
    }
    //
    function handleAddModalCancel(v) {
      adModalVisible.value = Boolean(v);
    }
    //
    function handleConfirm(v) {
      console.log(v);
      handleSearch();
    }
    //
    return () =>
      h('div', { style: { maxWidth: '1200px', margin: '0 auto' } }, [
        h(naive.NAlert, { type: 'success', title: '授权信息', style: 'margin-bottom: 10px' }, () => '已授权信息表格'),
        h('div', [
          h('div', { style: { marginBottom: '10px', textAlign: 'right', padding: '0 12px' } }, [
            h(naive.NButton, { type: 'info', onClick: handleAddAuth }, () => '新增'),
          ]),
          h(naive.NDataTable, {
            columns: columns,
            bordered: false,
            data: tableData.value,
          }),
        ]),

        // 新增授权 dialog
        h(naive.NMessageProvider, () =>
          h(AddDialog, {
            visible: adModalVisible.value,
            'onUpdate:visible': handleAddModalCancel,
            onClose: handleAddModalCancel,
            'onKiko:conform': handleConfirm,
          })
        ),
      ]);
  },
};
