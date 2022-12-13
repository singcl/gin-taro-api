import { h, ref } from 'vue';
import naive from 'naive';
import WeChat from '@vicons/ionicons5/WeChat.js';
import LoginNormal from './components/LoginNormal.js';
import LoginWechat from './components/LoginWechat.js';

// App
export default {
  setup() {
    const loginType = ref('normal');
    function changeLoginType(v) {
      loginType.value = v;
    }
    function loginTypeRender() {
      switch (loginType.value) {
        case 'wechat':
          return h(LoginWechat, { onChange: changeLoginType });
        case 'normal':
        default:
          return h(LoginNormal, {
            onChange: changeLoginType,
          });
      }
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
            h(naive.NIcon, { size: 48, class: 'mr-1' }, () => h(WeChat)),
            'Taro',
          ]),
          h('div', { class: 'mx-auto my-0 px-6 flex flex-col justify-center items-center' }, [
            h(
              'div',
              {
                class: 'w-[400px] rounded-lg bg-[#fff] p-6 flex flex-col justify-center items-center text-slate-900',
              },
              [loginTypeRender()]
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
