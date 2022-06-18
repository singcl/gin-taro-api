import { h, reactive, ref, toRef } from '/public/static/js/vue/vue3.esm-browser.js';
import naive from '/public/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from '/public/static/js/utils/kiko/Kiko.js';

const { NModal, NCard, NForm, NFormItem, NInput, NSelect, useMessage } = naive;
const methodOptionsConfig = [
  {
    label: 'POST',
    value: 'POST',
  },
  {
    label: 'GET',
    value: 'GET',
  },
  {
    label: 'PUT',
    value: 'PUT',
  },
  {
    label: 'DELETE',
    value: 'DELETE',
  },
  {
    label: 'PATCH',
    value: 'PATCH',
  },
];

export default {
  props: {
    visible: Boolean,
    detailData: Object,
  },
  emits: ['update:visible', 'kiko:conform'],
  // https://staging-cn.vuejs.org/guide/components/attrs.html#attribute-inheritance-on-multiple-root-nodes
  setup(props, ctx) {
    const message = useMessage();
    const visible = toRef(props, 'visible');
    const detailData = toRef(props, 'detailData');
    const methodOptions = reactive(methodOptionsConfig);
    const formData = reactive({
      method: undefined,
      api: undefined,
    });
    const formRef = ref(null);
    const rules = reactive({
      method: {
        required: true,
        trigger: 'blur',
        message: '请选择请求方式',
      },
      api: {
        required: true,
        trigger: 'blur',
        message: '请输入接口地址',
      },
    });
    //
    async function handleSure() {
      const response = await formRef.value?.validate();
      try {
        await new Kiko().fetch('/api/authorized_api', {
          method: 'POST',
          body: {
            ...formData,
            id: detailData.value.hashid,
          },
        });
        message.success('创建成功');
        handleCancel();
        ctx.emit('kiko:conform', response);
      } catch (error) {
        message.error(`API授权失败:code: ${error.code};message: ${error.message}`);
      }
    }
    //
    const handleCancel = () => {
      ctx.emit('update:visible', false);
    };

    //
    return () =>
      h(NModal, {
        show: visible.value,
        title: '新增API授权',
        preset: 'dialog',
        content: () =>
          h(NCard, () =>
            h(
              NForm,
              { model: formData, rules: rules, ref: formRef, labelPlacement: 'left', labelWidth: '120px' },
              () => [
                //
                h(NFormItem, { label: '请求方式', path: 'method' }, () =>
                  h(NSelect, {
                    value: formData.method,
                    'onUpdate:value': (v) => (formData.method = v),
                    options: methodOptions,
                    placeholder: '请选择请求方式',
                  })
                ),
                //
                h(NFormItem, { label: '接口地址', path: 'api' }, () =>
                  h(NInput, {
                    value: formData.api,
                    'onUpdate:value': (v) => (formData.api = v),
                    type: 'text',
                    placeholder: '请输入接口地址',
                  })
                ),
                //
              ]
            )
          ),
        positiveText: '确认',
        negativeText: '取消',
        onNegativeClick: handleCancel,
        onPositiveClick: handleSure,
        style: { width: '900px' },
      });
  },
};
