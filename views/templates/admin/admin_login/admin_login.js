import { h, ref, reactive } from 'vue';
import naive from 'naive';
import Kiko from 'kiko';
import EyeOutline from '@vicons/ionicons5/EyeOutline.js';

// App
export default {
  setup() {
    const nMessage = naive.useMessage();
    const nNotification = naive.useNotification();
    // 账号
    const username = ref('admin');
    function onInputChange(v) {
      username.value = v;
    }
    // 密码
    const password = ref('123456');
    //
    const onPasswordChange = (v) => (password.value = v);

    //
    const fieldFocused = reactive({
      username: false,
      password: false,
    });
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
      // TODO:
      nMessage.info('紧锣密鼓开发中...');
    }

    function handleForgotPassword() {
      // TODO:
      nMessage.info('紧锣密鼓开发中...');
    }

    // 渲染函数
    return () =>
      h(
        'div',
        {
          class:
            "w-full h-screen bg-cover bg-no-repeat bg-center bg-[url('https://cdn.apifox.cn/mirror-www/web/static/bg-texture.c61f6dbd.svg')] relative flex flex-col justify-center overflow-auto",
        },
        [
          h('div', { class: 'flex justify-start px-2 fixed top-16 left-16 text-green-600 text-2xl items-center' }, [
            h(naive.NIcon, { size: 48, class: 'mr-1' }, () =>
              h(
                'svg',
                {
                  viewBox: '64 64 896 896',
                  focusable: false,
                  dataIcon: 'wechat',
                  width: '1em',
                  height: '1em',
                  fill: 'currentColor',
                },
                [
                  h('path', {
                    d: 'M690.1 377.4c5.9 0 11.8.2 17.6.5-24.4-128.7-158.3-227.1-319.9-227.1C209 150.8 64 271.4 64 420.2c0 81.1 43.6 154.2 111.9 203.6a21.5 21.5 0 019.1 17.6c0 2.4-.5 4.6-1.1 6.9-5.5 20.3-14.2 52.8-14.6 54.3-.7 2.6-1.7 5.2-1.7 7.9 0 5.9 4.8 10.8 10.8 10.8 2.3 0 4.2-.9 6.2-2l70.9-40.9c5.3-3.1 11-5 17.2-5 3.2 0 6.4.5 9.5 1.4 33.1 9.5 68.8 14.8 105.7 14.8 6 0 11.9-.1 17.8-.4-7.1-21-10.9-43.1-10.9-66 0-135.8 132.2-245.8 295.3-245.8zm-194.3-86.5c23.8 0 43.2 19.3 43.2 43.1s-19.3 43.1-43.2 43.1c-23.8 0-43.2-19.3-43.2-43.1s19.4-43.1 43.2-43.1zm-215.9 86.2c-23.8 0-43.2-19.3-43.2-43.1s19.3-43.1 43.2-43.1 43.2 19.3 43.2 43.1-19.4 43.1-43.2 43.1zm586.8 415.6c56.9-41.2 93.2-102 93.2-169.7 0-124-120.8-224.5-269.9-224.5-149 0-269.9 100.5-269.9 224.5S540.9 847.5 690 847.5c30.8 0 60.6-4.4 88.1-12.3 2.6-.8 5.2-1.2 7.9-1.2 5.2 0 9.9 1.6 14.3 4.1l59.1 34c1.7 1 3.3 1.7 5.2 1.7a9 9 0 006.4-2.6 9 9 0 002.6-6.4c0-2.2-.9-4.4-1.4-6.6-.3-1.2-7.6-28.3-12.2-45.3-.5-1.9-.9-3.8-.9-5.7.1-5.9 3.1-11.2 7.6-14.5zM600.2 587.2c-19.9 0-36-16.1-36-35.9 0-19.8 16.1-35.9 36-35.9s36 16.1 36 35.9c0 19.8-16.2 35.9-36 35.9zm179.9 0c-19.9 0-36-16.1-36-35.9 0-19.8 16.1-35.9 36-35.9s36 16.1 36 35.9a36.08 36.08 0 01-36 35.9z',
                  }),
                ]
              )
            ),
            'Taro',
            h(naive.NIcon, { size: 48 }, () => h(EyeOutline)),
          ]),
          h('div', { class: 'mx-auto my-0 px-6 flex flex-col justify-center items-center' }, [
            h(
              'div',
              {
                class: 'w-[400px] rounded-lg bg-[#fff] p-6 flex flex-col justify-center items-center text-slate-900',
              },
              [
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
                                type: 'password',
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
                      },
                      [
                        h(naive.NIcon, { size: 24, class: 'mr-1' }, () =>
                          h(
                            'svg',
                            {
                              viewBox: '64 64 896 896',
                              focusable: false,
                              dataIcon: 'wechat',
                              width: '1em',
                              height: '1em',
                              fill: 'currentColor',
                            },
                            [
                              h('path', {
                                d: 'M690.1 377.4c5.9 0 11.8.2 17.6.5-24.4-128.7-158.3-227.1-319.9-227.1C209 150.8 64 271.4 64 420.2c0 81.1 43.6 154.2 111.9 203.6a21.5 21.5 0 019.1 17.6c0 2.4-.5 4.6-1.1 6.9-5.5 20.3-14.2 52.8-14.6 54.3-.7 2.6-1.7 5.2-1.7 7.9 0 5.9 4.8 10.8 10.8 10.8 2.3 0 4.2-.9 6.2-2l70.9-40.9c5.3-3.1 11-5 17.2-5 3.2 0 6.4.5 9.5 1.4 33.1 9.5 68.8 14.8 105.7 14.8 6 0 11.9-.1 17.8-.4-7.1-21-10.9-43.1-10.9-66 0-135.8 132.2-245.8 295.3-245.8zm-194.3-86.5c23.8 0 43.2 19.3 43.2 43.1s-19.3 43.1-43.2 43.1c-23.8 0-43.2-19.3-43.2-43.1s19.4-43.1 43.2-43.1zm-215.9 86.2c-23.8 0-43.2-19.3-43.2-43.1s19.3-43.1 43.2-43.1 43.2 19.3 43.2 43.1-19.4 43.1-43.2 43.1zm586.8 415.6c56.9-41.2 93.2-102 93.2-169.7 0-124-120.8-224.5-269.9-224.5-149 0-269.9 100.5-269.9 224.5S540.9 847.5 690 847.5c30.8 0 60.6-4.4 88.1-12.3 2.6-.8 5.2-1.2 7.9-1.2 5.2 0 9.9 1.6 14.3 4.1l59.1 34c1.7 1 3.3 1.7 5.2 1.7a9 9 0 006.4-2.6 9 9 0 002.6-6.4c0-2.2-.9-4.4-1.4-6.6-.3-1.2-7.6-28.3-12.2-45.3-.5-1.9-.9-3.8-.9-5.7.1-5.9 3.1-11.2 7.6-14.5zM600.2 587.2c-19.9 0-36-16.1-36-35.9 0-19.8 16.1-35.9 36-35.9s36 16.1 36 35.9c0 19.8-16.2 35.9-36 35.9zm179.9 0c-19.9 0-36-16.1-36-35.9 0-19.8 16.1-35.9 36-35.9s36 16.1 36 35.9a36.08 36.08 0 01-36 35.9z',
                              }),
                            ]
                          )
                        ),
                        '微信登陆',
                      ]
                    ),
                  ]),
                ]),
              ]
            ),
          ]),
        ]
      );
  },
};

// // App
// export default {
//     data() {
//       return { count: 0 }
//     },
//     render() {
//         return h("div", null, this.count)
//     }
//   }
