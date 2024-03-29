import { h, ref, reactive, defineComponent } from 'vue';
import naive from 'naive';
import Kiko from 'kiko';
import EyeOutline from '@vicons/ionicons5/EyeOutline.js';
import EyeOffOutline from '@vicons/ionicons5/EyeOffOutline.js';
import WeChat from '@vicons/ionicons5/WeChat.js';
//
export default defineComponent((props, { emit }) => {
  const nMessage = naive.useMessage();
  const nNotification = naive.useNotification();
  // 账号
  const username = ref('admin');
  function onInputChange(v) {
    username.value = v;
  }
  // 密码
  const password = ref('');
  //
  const onPasswordChange = (v) => (password.value = v);

  //
  const fieldFocused = reactive({
    username: false,
    password: false,
  });
  //
  const hidePassword = ref(true);

  //
  function onFieldReactive(field, res) {
    fieldFocused[field] = res;
  }
  //
  async function handleSubmitClick() {
    try {
      const response = await new Kiko().fetch('/api/login', {
        method: 'POST',
        body: {
          username: username.value,
          password: password.value,
        },
      });
      const token = response && response.token;
      token && localStorage.setItem(Kiko.getTokenName(), token);
      //
      location.href = '/';
    } catch (error) {
      const code = error.code;
      const message = error.message;
      nNotification.error({
        title: '错误信息',
        content: () => h('div', [h('p', [code]), h('p', [message])]),
        // meta: new Date().toLocaleDateString(),
        duration: 2500,
        keepAliveOnHover: true,
      });
      // alert(`code：${code}\r\nmessage：${message}`);
    }
  }

  //
  function handleVerifyCodeLogin() {
    // TODO:待开发
    nMessage.info('紧锣密鼓开发中...');
  }

  function handleForgotPassword() {
    // TODO:待开发
    nMessage.info('紧锣密鼓开发中...');
  }

  function handleLoginWechat() {
    //
    emit('change', 'wechat');
  }

  //
  return () => [
    h('h3', { class: 'self-start pt-4 mb-2 text-2xl' }, ['欢迎使用GIN-TARO-API']),
    h('div', { class: 'w-full flex overflow-hidden flex-col text-sm' }, [
      h('div', { class: 'relative flex flex-none items-center' }, [
        h('div', { class: 'relative flex flex-auto grow' }, [
          h('div', { class: 'relative flex pt-2 pb-2' }, [
            h(
              'div',
              {
                class:
                  'text-green-600 after:absolute after:w-full after:h-[2px] after:left-0 after:bottom-0 after:bg-green-600',
              },
              ['邮箱']
            ),
          ]),
        ]),
      ]),
      h('div', { class: 'flex-auto' }, [
        h('div', { class: 'w-full flex-none outline-none pt-4 pb-4' }, [
          h('div', { class: 'mb-4 flex flex-row flex-wrap' }, [
            h('div', { class: 'flex w-full' }, [
              h(
                'span',
                {
                  class: `login-field ${fieldFocused.username ? 'login-field-focused' : ''}`,
                },
                [
                  h('input', {
                    class: 'inline-block w-full p-0 outline-0 border-0',
                    name: 'username',
                    type: 'text',
                    value: username.value,
                    placeholder: '请输入用户名',
                    onInput: (e) => {
                      // username.value = e.target.value 这样写无法响应式更新，为啥？
                      onInputChange(e.target.value);
                    },
                    onFocus: () => {
                      onFieldReactive('username', true);
                    },
                    onBlur: () => {
                      onFieldReactive('username', false);
                    },
                  }),
                ]
              ),
            ]),
          ]),
          h('div', { class: 'mb-4 flex flex-row flex-wrap' }, [
            h('div', { class: 'flex w-full' }, [
              h(
                'span',
                {
                  class: `login-field ${fieldFocused.password ? 'login-field-focused' : ''}`,
                },
                [
                  h('input', {
                    class: 'inline-block w-full p-0 outline-0 border-0',
                    name: 'password',
                    type: `${hidePassword.value ? 'password' : 'text'}`,
                    value: password.value,
                    placeholder: '请输入密码',
                    onInput: (e) => {
                      // password.value = e.target.value 这样写无法响应式更新，为啥？
                      onPasswordChange(e.target.value);
                    },
                    onFocus: () => {
                      onFieldReactive('password', true);
                    },
                    onBlur: () => {
                      onFieldReactive('password', false);
                    },
                  }),
                  h(naive.NIcon, { size: 18 }, () =>
                    h(
                      'span',
                      {
                        class: 'cursor-pointer text-slate-500',
                        onClick: () => (hidePassword.value = !hidePassword.value),
                      },
                      [hidePassword.value ? h(EyeOffOutline) : h(EyeOutline)]
                    )
                  ),
                ]
              ),
            ]),
          ]),

          h('div', { class: 'w-full' }, [
            h(
              'button',
              {
                class:
                  'w-full inline-flex justify-center items-center mb-2 bg-green-600 hover:bg-green-700 text-white pt-2 pb-2 pl-4 pr-4 rounded outline-0 text-base',
                onClick: handleSubmitClick,
              },
              '登录'
            ),
          ]),

          h('div', { class: 'flex justify-between items-center' }, [
            h(
              'button',
              {
                class: 'inline-flex text-green-600 rounded outline-0 h-8 justify-center items-center',
                onClick: handleVerifyCodeLogin,
              },
              '验证码登陆/注册'
            ),
            h(
              'button',
              {
                class: 'inline-flex text-green-600 rounded outline-0 h-8 justify-center items-center',
                onClick: handleForgotPassword,
              },
              '忘记密码'
            ),
          ]),
        ]),
      ]),
    ]),
    h('div', { class: 'w-full flex flex-col pt-4 mt-0 border-solid border-t border-slate-200' }, [
      h('div', { class: 'py-2' }, [
        h(
          'button',
          {
            class:
              'w-full inline-flex justify-center items-center py-2 text-green-600 rounded outline-0 text-base border-solid border-slate-100 border hover:border-green-600',
            onClick: handleLoginWechat,
          },
          [h(naive.NIcon, { size: 24, class: 'mr-1' }, () => h(WeChat)), '微信登陆']
        ),
      ]),
    ]),
  ];
});
