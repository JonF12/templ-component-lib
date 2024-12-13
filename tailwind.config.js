module.exports = {
  content: [
    './src/**/*.{templ,go}',
    './examples/**/*.{templ,go}',
    './examples/*.{templ,go}'
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
