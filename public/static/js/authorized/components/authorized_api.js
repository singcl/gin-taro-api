import { h, reactive, toRef, watch } from '/public/static/js/vue/vue3.esm-browser.js';
import naive from '/public/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from '/public/static/js/utils/kiko/Kiko.js';
import { getTagType } from './utils.js';

const { NDrawer, NList, NListItem, /* useMessage, */ NAlert, NThing, NTag, NEmpty } = naive;

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
              header: () => h('h3', { style: { margin: 0 } }, `授权方：${listData.business_key}`),
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
        ]
      );
  },
};
