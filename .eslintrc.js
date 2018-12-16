module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: [
    'plugin:vue/recommended'
  ],
  // required to lint *.vue files
  plugins: [
    'vue'
  ],
  // add your custom rules here
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    "vue/attribute-hyphenation": ["error", "never"],
    "vue/attributes-order": ["off"],
    "vue/max-attributes-per-line": ["off"],
    "vue/order-in-components": ["error", {
      "order": [
        "el",
        "name",
        "parent",
        ["components", "directives", "filters"],
        ["props", "propsData"],
        "data",
        "watch",
        "methods",
        "template"
      ]
    }],
    "vue/require-default-prop": ["off"]
  }
}
