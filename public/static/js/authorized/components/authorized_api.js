import { h, reactive, toRef, watch, ref } from '/public/static/js/vue/vue3.esm-browser.js';
import naive from '/public/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from '/public/static/js/utils/kiko/Kiko.js';
import { getTagType } from './utils.js';
import ApiAddDialog from './authorized_api_add.js';

const { NDrawer, NList, NListItem, /* useMessage, */ NAlert, NThing, NTag, NEmpty, NButton } = naive;

export default {
  props: {
    visible: Boolean,
    detailData: Object,
  },
  emits: ['update:visible', 'kiko:conform'],
  // https://staging-cn.vuejs.org/guide/components/attrs.html#attribute-inheritance-on-multiple-root-nodes
  setup(props /* ctx */) {
    // const message = useMessage();
    const visible = toRef(props, 'visible');
    const detailData = toRef(props, 'detailData');
    const apiModalVisible = ref(false);
    const listData = reactive({
      business_key: undefined,
      list: [],
    });

    watch(visible, async function (newVal) {
      if (newVal) {
        handleSearch();
      }
    });

    //
    async function handleSearch() {
      const response = await new Kiko().fetch(`/api/authorized_api`, {
        method: 'GET',
        body: {
          id: detailData.value.hashid,
        },
      });
      const { business_key, list = [] } = response;
      listData.business_key = business_key;
      listData.list = list;
    }

    function handleAddApiAuth() {
      apiModalVisible.value = true;
    }

    //
    function handleAddModalCancel(v) {
      apiModalVisible.value = Boolean(v);
    }
    //
    function handleConfirm(v) {
      console.log(v);
      handleSearch();
    }

    //
    return () =>
      h(
        NDrawer,
        {
          show: visible.value,
          style: { width: '900px' },
        },
        () => [
          h(
            NAlert,
            { type: 'info', title: 'API授权详情' },
            () => '接口地址支持通配符(*)，其中 * 表示 1 级，** 表示 n 级。'
          ),
          h(
            NList,
            { style: { margin: '0 15px' } },
            {
              header: () =>
                h(
                  'h3',
                  { style: { margin: 0, display: 'flex', alignItems: 'center', justifyContent: 'space-between' } },
                  [
                    h('span', null, `授权方：${listData.business_key}`),
                    h(NButton, { type: 'success', onClick: handleAddApiAuth }, () => '新增API授权'),
                  ]
                ),
              default: () =>
                listData.list && listData.list.length > 0
                  ? listData.list.map((item) =>
                      h(NListItem, null, {
                        default: () =>
                          h(NThing, null, {
                            header: () => h(NTag, { type: getTagType(item.method), round: true }, () => item.method),
                            description: () => item.api,
                          }),
                      })
                    )
                  : h(NEmpty, { description: '没有任何已授权接口' }),
            }
          ),

          // 新增接口授权
          h(ApiAddDialog, {
            visible: apiModalVisible.value,
            'onUpdate:visible': handleAddModalCancel,
            onClose: handleAddModalCancel,
            'onKiko:conform': handleConfirm,
            detailData: detailData,
          }),
        ]
      );
  },
};
