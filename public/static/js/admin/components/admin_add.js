import { h, reactive, ref, toRef } from '/public/static/js/vue/vue3.esm-browser.js';
import naive from '/public/static/js/vue/naive.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from '/public/static/js/utils/kiko/Kiko.js';

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
      username: undefined,
      nickname: undefined,
      mobile: undefined,
      password: undefined,
    });
    const formRef = ref(null);
    const rules = reactive({
      username: [
        {
          required: true,
          trigger: 'blur',
          message: '请输入用户名',
        },
        // {
        //   min: 6,
        //   max: 12,
        //   message: '用户名长度必须在6-12之间',
        // trigger: 'blur',
        // },
        {
          pattern: /^[a-z][a-z0-9_]{5,11}$/,
          message: '用户名长度必须在6-12位之间，以小写字母开头且只能包含字母数字下划线',
          trigger: 'blur',
        },
      ],
      nickname: [
        {
          required: true,
          trigger: 'blur',
          message: '请输入昵称',
        },
        {
          pattern: /^[^0-9-][a-zA-Z0-9\u4e00-\u9fa5_-]{7,17}$/,
          message: '昵称长度必须在8-18位之间，只能包含字母数字下划线中划线中文',
          trigger: 'blur',
        },
      ],
      mobile: [
        {
          required: true,
          trigger: 'blur',
          message: '手机号',
        },
        {
          pattern: /^1[34578]\d{9}$/,
          message: '请输入正确的手机号码',
          trigger: 'blur',
        },
      ],
      password: [
        {
          required: true,
          trigger: 'blur',
          message: '密码',
        },
        {
          pattern: /^(?![a-zA-Z]+$)(?![A-Z0-9]+$)(?![A-Z\W_]+$)(?![a-z0-9]+$)(?![a-z\W_]+$)(?![0-9\W_]+$)[a-zA-Z0-9\W_]{8,20}$/,
          message: '请输入8-20位字符，必须包含大写字母小写字母和数字',
          trigger: 'blur',
        },
      ],
    });
    //
    async function handleSure() {
      const response = await formRef.value?.validate();
      try {
        await new Kiko().fetch('/api/admin', {
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
        title: '新增管理员',
        preset: 'dialog',
        content: () =>
          h(NCard, () =>
            h(
              NForm,
              { model: formData, rules: rules, ref: formRef, labelPlacement: 'left', labelWidth: '120px' },
              () => [
                //
                h(NFormItem, { label: '用户名', path: 'username' }, () =>
                  h(NInput, {
                    value: formData.username,
                    'onUpdate:value': (v) => (formData.username = v),
                    type: 'text',
                    placeholder: '请输入用户名',
                  })
                ),
                //
                h(NFormItem, { label: '昵称', path: 'nickname' }, () =>
                  h(NInput, {
                    value: formData.nickname,
                    'onUpdate:value': (v) => (formData.nickname = v),
                    type: 'text',
                    placeholder: '请输入昵称',
                  })
                ),
                //
                h(NFormItem, { label: '手机号', path: 'mobile' }, () =>
                  h(NInput, {
                    value: formData.mobile,
                    'onUpdate:value': (v) => (formData.mobile = v),
                    type: 'text',
                    placeholder: '请输入手机号',
                  })
                ),
                //
                h(NFormItem, { label: '密码', path: 'password' }, () =>
                  h(NInput, {
                    value: formData.password,
                    'onUpdate:value': (v) => (formData.password = v),
                    type: 'password',
                    placeholder: '请输入密码',
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
