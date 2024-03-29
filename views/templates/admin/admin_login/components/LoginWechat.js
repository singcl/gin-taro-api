import { h, defineComponent } from 'vue';
import naive from 'naive';
// import Kiko from 'kiko';
import Mail from '@vicons/ionicons5/Mail.js';
//
export default defineComponent((props, { emit }) => {
  function handleLoginNormal() {
    //
    emit('change', 'normal');
  }
  return () =>
    h('div', { class: 'flex flex-col items-center justify-content text-center w-full' }, [
      h('h3', { class: 'my-4 text-xl' }, ['微信登录']),
      h('p', { class: 'm-0' }, ['请使用', h('a', { class: 'text-green-600' }, ['微信扫码']), '关注公众号即可安全登录']),
      h('div', { class: 'w-full flex flex-col pt-4 mt-0 border-solid border-t border-slate-200' }, [
        h('div', { class: 'py-2' }, [
          h(
            'button',
            {
              class:
                'w-full inline-flex justify-center items-center py-2 text-green-600 rounded outline-0 text-base border-solid border-green-600 border hover:border-green-600',
              onClick: handleLoginNormal,
            },
            [h(naive.NIcon, { size: 24, class: 'mr-1' }, () => h(Mail)), '手机/邮箱登录']
          ),
        ]),
      ]),
    ]);
});
