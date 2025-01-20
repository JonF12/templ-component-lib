module.exports = {
  content: [
    "./src/**/*.{templ,go,jsx}",
    "./examples/**/*.{templ,go,jsx}",
    "./examples/*.{templ,go,jsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography")],
};
