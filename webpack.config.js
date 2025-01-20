const path = require("path");

module.exports = {
  entry: "./examples/assets/js/app.js",
  output: {
    path: path.resolve(__dirname, "examples/assets/dist/"),
    filename: "bundle.js",
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",
          options: {
            presets: ["@babel/preset-react"],
          },
        },
      },
    ],
  },
  resolve: {
    extensions: [".js", ".jsx"],
  },
};
