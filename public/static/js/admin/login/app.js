import { h, ref } from '/public/static/js/vue/vue3.esm-browser.js';

// // App
// export default {
//     data() {
//       return { count: 0 }
//     },
//     render() {
//         return h("div", null, this.count)
//     }
//   }

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
        const onPasswordChange = (v) => (password.value = v);

        // 渲染函数
        return () =>
            h('div', { class: 'admin-login' }, [
                h('div', { class: 'admin-login-form' }, [
                    h('div', { class: 'admin-login-form__item' }, [
                        h('label', [
                            '用户名:',
                            h('input', {
                                type: 'text',
                                value: username.value,
                                onInput: (e) => {
                                    // username.value = e.target.value 这样写无法响应式更新，为啥？
                                    onInputChange(e.target.value);
                                },
                            }),
                        ]),
                    ]),
                    h('div', { class: 'admin-login-form__item' }, [
                        h('label', [
                            '密码:',
                            h('input', {
                                type: 'password',
                                value: password.value,
                                onInput: (e) => {
                                    // password.value = e.target.value 这样写无法响应式更新，为啥？
                                    onPasswordChange(e.target.value);
                                },
                            }),
                        ]),
                    ]),
                ]),
            ]);
    },
};
