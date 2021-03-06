import { h, ref } from '/public/static/js/vue/vue3.esm-browser.js';
// 绝对路径导入无法识别类型。 不知道怎么配置？
import Kiko from './../../utils/kiko/Kiko.js';
import naive from '/public/static/js/vue/naive.js';

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
      h('div', { class: 'admin-login' }, [
        h('div', { class: 'admin-login-form' }, [
          h('div', { class: 'admin-login-form__item' }, [
            h('label', { for: 'username' }, ['用户名:']),
            h('input', {
              name: 'username',
              type: 'text',
              value: username.value,
              placeholder: '请输入用户名',
              onInput: (e) => {
                // username.value = e.target.value 这样写无法响应式更新，为啥？
                onInputChange(e.target.value);
              },
            }),
          ]),
          h('div', { class: 'admin-login-form__item' }, [
            h('label', { for: 'password' }, ['密码:']),
            h('input', {
              name: 'password',
              type: 'password',
              value: password.value,
              placeholder: '请输入密码',
              onInput: (e) => {
                // password.value = e.target.value 这样写无法响应式更新，为啥？
                onPasswordChange(e.target.value);
              },
            }),
          ]),

          h('div', { class: 'admin-login-form__item' }, [
            h('label'),
            h(naive.NButton, {type: 'info', onClick: handleSubmitClick }, '登录'),
          ]),
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
