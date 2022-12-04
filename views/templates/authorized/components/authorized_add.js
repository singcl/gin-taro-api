import { h, reactive, ref, toRef } from '/views/static/js/vue/vue3.esm-browser.js';
import naive from '/views/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from '/views/static/js/utils/kiko/Kiko.js';

const { NModal, NCard, NForm, NFormItem, NInput, useMessage } = naive;

export default {
  props: {
    visible: Boolean,
  },
  emits: ['update:visible', 'kiko:conform'],
  // https://staging-cn.vuejs.org/guide/components/attrs.html#attribute-inheritance-on-multiple-root-nodes
  setup(props, ctx) {
    const message = useMessage();
    const visible = toRef(props, 'visible');
    const formData = reactive({
      business_key: undefined,
      business_developer: undefined,
      remark: undefined,
    });
    const formRef = ref(null);
    const rules = reactive({
      business_key: {
        required: true,
        trigger: 'blur',
        message: '请输入调用方标识',
      },
      business_developer: {
        required: true,
        trigger: 'blur',
        message: '请输入调用方对接人',
      },
      remark: {
        required: true,
        trigger: 'blur',
        message: '备注',
      },
    });
    //
    async function handleSure() {
      const response = await formRef.value?.validate();
      try {
        await new Kiko().fetch('/api/authorized', {
          method: 'POST',
          body: formData,
        });
        message.success('创建成功');
        handleCancel();
        ctx.emit('kiko:conform', response);
      } catch (error) {
        message.error(`创建失败:code: ${error.code};message: ${error.message}`);
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
        title: '新增授权',
        preset: 'dialog',
        content: () =>
          h(NCard, () =>
            h(
              NForm,
              { model: formData, rules: rules, ref: formRef, labelPlacement: 'left', labelWidth: '120px' },
              () => [
                //
                h(NFormItem, { label: '调用方', path: 'business_key' }, () =>
                  h(NInput, {
                    value: formData.business_key,
                    'onUpdate:value': (v) => (formData.business_key = v),
                    type: 'text',
                    placeholder: '请输入调用方标识',
                  })
                ),
                //
                h(NFormItem, { label: '调用方对接人', path: 'business_developer' }, () =>
                  h(NInput, {
                    value: formData.business_developer,
                    'onUpdate:value': (v) => (formData.business_developer = v),
                    type: 'text',
                    placeholder: '请输入调用方对接人',
                  })
                ),
                //
                h(NFormItem, { label: '备注', path: 'remark' }, () =>
                  h(NInput, {
                    value: formData.remark,
                    'onUpdate:value': (v) => (formData.remark = v),
                    type: 'text',
                    placeholder: '请输入备注',
                  })
                ),
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
