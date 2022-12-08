import { h, ref } from '/views/static/js/vue/vue3.esm-browser.js';
// import naive from '/views/static/js/vue/naive.js';
import Kiko from '/views/static/js/utils/kiko/Kiko.js';

// App
export default {
  setup() {
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
        alert(`code：${code}\r\nmessage：${message}`);
      }
    }

    // 渲染函数
    return () =>
      h('div', null, [
        h('div', { class: 'mx-auto my-0 px-6 flex flex-col justify-center items-center' }, [
          h(
            'div',
            { class: 'w-[400px] rounded-lg bg-[#fff] p-6 flex flex-col justify-center items-center text-slate-900' },
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
                            class:
                              'w-full inline-flex pt-2 pb-2 pl-[10px] pr-[10px] border-solid border-slate-100 border hover:border-green-600 focus:border-green-600 rounded',
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
                            class:
                              'w-full inline-flex pt-2 pb-2 pl-[10px] pr-[10px] border-solid border-slate-100 border hover:border-green-600 focus:border-green-600 rounded',
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
                  ]),
                ]),
              ]),
            ]
          ),
        ]),
      ]);
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
