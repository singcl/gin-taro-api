import { h, reactive, ref, onMounted } from '/views/static/js/vue/vue3.esm-browser.js';
import naive from '/views/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from '/views/static/js/utils/kiko/Kiko.js';

//
import AddDialog from './components/admin_add.js';
// import AuthApiDrawer from './components/authorized_api.js';

const { useMessage, useDialog } = naive;
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

const onlineOptions = [
  {
    value: 1,
    label: '在线',
    type: 'success',
  },
  {
    value: -1,
    label: '不在线',
    type: 'warning',
  },
];

//
const defaultColumnsConfig = [
  { key: 'id', title: '编号' },
  { key: 'hashid', title: 'hashid' },
  { key: 'username', title: '用户名' },
  { key: 'mobile', title: '手机号' },
  { key: 'nickname', title: '昵称' },
  { key: 'created_user', title: '创建人' },
  {
    key: 'created_at',
    title: '创建日期',
    render(row) {
      const { created_at } = row;
      return created_at.split(/\s/)[0];
    },
  },
  { key: 'updated_user', title: '更新人' },
  {
    key: 'updated_at',
    title: '更新日期',
    render(row) {
      const { created_at } = row;
      return created_at.split(/\s/)[0];
    },
  },
  {
    key: 'is_online',
    title: '在线状态',
    render(row) {
      const { is_online } = row;
      const itm = onlineOptions.find((m) => m.value == is_online);
      return itm ? h(naive.NBadge, { type: itm.type, dot: true }, () => itm.label) : is_online;
    },
  },
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
  },
];

// App
export default {
  setup() {
    const defaultColumns = defaultColumnsConfig.map((item) => {
      let newItem = item;
      switch (item.key) {
        case 'operation':
          newItem = {
            ...newItem,
            render(row) {
              return h(
                naive.NDropdown,
                {
                  options: dropdownOptionsRender(row),
                  trigger: 'click',
                  width: 'trigger',
                },
                () => h(naive.NButton, { type: 'info', size: 'small' }, () => '批量操作')
              );
            },
          };
          break;
        default:
          break;
      }
      return newItem;
    });
    // 下拉菜单 @see https://www.naiveui.com/zh-CN/os-theme/components/dropdown

    function dropdownOptionsRender(row) {
      const { is_used } = row;
      const op_used = { 1: -1, '-1': 1 }[is_used];
      const itm = useOptions.find((m) => m.value == op_used);
      return [
        {
          type: 'render',
          key: 'interface',
          render() {
            return h(
              naive.NButton,
              {
                type: 'info',
                text: true,
                style: { marginRight: '10px', width: '100%', marginBottom: '6px' },
                onClick: () => handleAuthApiDetail(row),
              },
              () => '接口'
            );
          },
        },
        {
          type: 'render',
          key: 'used',
          render() {
            return itm
              ? h(
                  naive.NButton,
                  {
                    type: itm.type,
                    text: true,
                    style: { width: '100%', marginBottom: '6px' },
                    onClick: () => handleUpdateUsed(row),
                  },
                  () => itm.label
                )
              : undefined;
          },
        },
        {
          type: 'render',
          key: 'delete',
          render() {
            return h(
              naive.NButton,
              {
                type: 'error',
                style: { width: '100%', marginBottom: '6px' },
                text: true,
                onClick: () => handleDelete(row),
              },
              () => '删除'
            );
          },
        },
        {
          type: 'render',
          key: 'offline',
          render() {
            return h(
              naive.NButton,
              {
                type: 'error',
                text: true,
                style: { width: '100%', marginBottom: '6px' },
                onClick: () => handleOffline(row),
              },
              () => '下线'
            );
          },
        },
        {
          type: 'render',
          key: 'reset',
          render() {
            return h(
              naive.NButton,
              {
                type: 'warning',
                text: true,
                style: { width: '100%', marginBottom: '6px' },
                onClick: () => handleResetPassword(row),
              },
              () => '重置密码'
            );
          },
        },
      ];
    }

    const columns = reactive(defaultColumns);
    const tableData = ref([]);
    const adModalVisible = ref(false);
    const apiDrawerVisible = ref(false);
    const detailData = ref(null);
    const message = useMessage();
    const dialog = useDialog();
    //
    onMounted(async () => {
      await handleSearch();
    });
    //
    async function handleSearch() {
      try {
        const response = await new Kiko().fetch('/api/admin', {
          method: 'get',
        });
        tableData.value = response.list;
      } catch (error) {
        message.error(`code: ${error.code};message: ${error.message}`);
      }
    }
    //
    function handleAddAuth() {
      adModalVisible.value = true;
    }
    //
    function handleAddModalCancel(v) {
      adModalVisible.value = Boolean(v);
    }
    // //
    // function handleApiDrawerCancel(v) {
    //   apiDrawerVisible.value = Boolean(v);
    // }
    //
    function handleConfirm(v) {
      console.log(v);
      handleSearch();
    }

    async function handleUpdateUsed(row) {
      const { hashid, is_used } = row;
      const used = { 1: -1, '-1': 1 }[is_used];
      const itm = useOptions.find((item) => item.value == used);
      dialog.warning({
        title: '警告',
        content: `确定${itm.label}当前用户吗？`,
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
          try {
            await new Kiko().fetch('/api/admin/used', {
              method: 'PATCH',
              body: {
                id: hashid,
                used: used,
              },
            });
            message.success(`${itm.label}成功！`);
            handleSearch();
          } catch (error) {
            message.error(`code: ${error.code};message: ${error.message}`);
          }
        },
      });
    }

    //
    async function handleAuthApiDetail(row) {
      apiDrawerVisible.value = true;
      detailData.value = row;
    }

    //
    async function handleDelete(row) {
      const { hashid } = row;

      dialog.warning({
        title: '警告',
        content: `确定删除当前用户吗？`,
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
          try {
            await new Kiko().fetch(`/api/admin/${hashid}`, {
              method: 'DELETE',
            });
            message.success(`删除成功！`);
            handleSearch();
          } catch (error) {
            message.error(`code: ${error.code};message: ${error.message}`);
          }
        },
      });
    }
    //
    async function handleOffline(row) {
      const { hashid } = row;

      dialog.warning({
        title: '警告',
        content: `确定下线当前用户吗？`,
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
          try {
            await new Kiko().fetch(`/api/admin/offline`, {
              method: 'PATCH',
              body: {
                id: hashid,
              },
            });
            message.success(`下线成功！`);
            handleSearch();
          } catch (error) {
            message.error(`code: ${error.code};message: ${error.message}`);
          }
        },
      });
    }

    //
    async function handleResetPassword(row) {
      const { hashid } = row;

      dialog.warning({
        title: '警告',
        content: `确定重置当前用户密码吗？`,
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
          try {
            await new Kiko().fetch(`/api/admin/reset_password/${hashid}`, {
              method: 'PATCH',
            });
            message.success(`重置成功！`);
            handleSearch();
          } catch (error) {
            message.error(`code: ${error.code};message: ${error.message}`);
          }
        },
      });
    }
    //
    return () =>
      h('div', { style: { maxWidth: '1200px', margin: '0 auto' } }, [
        h(
          naive.NAlert,
          { type: 'success', title: '管理员信息', style: 'margin-bottom: 10px' },
          () => '已授权管理员列表'
        ),
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

        // 新增管理员 dialog
        h(naive.NMessageProvider, () =>
          h(AddDialog, {
            visible: adModalVisible.value,
            'onUpdate:visible': handleAddModalCancel,
            onClose: handleAddModalCancel,
            'onKiko:conform': handleConfirm,
          })
        ),

        // // 已授权API drawer
        // h(naive.NMessageProvider, () =>
        //   h(AuthApiDrawer, {
        //     detailData: detailData,
        //     visible: apiDrawerVisible.value,
        //     'onUpdate:visible': handleApiDrawerCancel,
        //     'onUpdate:show': handleApiDrawerCancel,
        //   })
        // ),
      ]);
  },
};
