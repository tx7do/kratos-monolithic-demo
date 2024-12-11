export default {
  endOfLine: 'auto',
  htmlWhitespaceSensitivity: 'strict',
  overrides: [
    {
      files: ['*.json5'],
      options: {
        quoteProps: 'preserve',
        singleQuote: false,
      },
    },
  ],
  plugins: ['prettier-plugin-tailwindcss'],
  printWidth: 100,
  proseWrap: 'never',
  semi: true,
  singleQuote: true,
  trailingComma: 'all',
  vueIndentScriptAndStyle: true,
};
