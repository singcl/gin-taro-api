import { createElementVNode as _createElementVNode, openBlock as _openBlock, createElementBlock as _createElementBlock, defineComponent } from 'vue'
const _hoisted_1 = {
  xmlns: 'http://www.w3.org/2000/svg',
  'xmlns:xlink': 'http://www.w3.org/1999/xlink',
  viewBox: '0 0 512 512'
}
const _hoisted_2 = /*#__PURE__*/ _createElementVNode(
  'circle',
  {
    cx: '256',
    cy: '256',
    r: '208',
    fill: 'none',
    stroke: 'currentColor',
    'stroke-miterlimit': '10',
    'stroke-width': '32'
  },
  null,
  -1
  /* HOISTED */
)
const _hoisted_3 = /*#__PURE__*/ _createElementVNode(
  'path',
  {
    fill: 'none',
    stroke: 'currentColor',
    'stroke-miterlimit': '10',
    'stroke-width': '32',
    d: 'M108.92 108.92l294.16 294.16'
  },
  null,
  -1
  /* HOISTED */
)
const _hoisted_4 = [_hoisted_2, _hoisted_3]
export default defineComponent({
  name: 'BanOutline',
  render: function render(_ctx, _cache) {
    return _openBlock(), _createElementBlock('svg', _hoisted_1, _hoisted_4)
  }
})
