module.exports = {
  root: true,
  env: {
    node: true,
    browser: true,
  },
  parserOptions: {
    ecmaVersion: 13,
    sourceType: 'module',
  },
  extends: ['eslint:recommended'],
  rules: {
    // "no-unused-vars": [2, {"vars": "all", "args": "after-used"}]
    // 'no-unused-vars': 0,
  },
};
