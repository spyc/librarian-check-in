const vueConf = require('./vue-loader.conf');

module.exports = {
  entry: {
    renderer: './src/renderer/index.js',
  },
  devtool: '#source-map',
  externals: [
    'source-map-support',
    'electron',
    'webpack',
    'electron-devtools-installer',
  ],
  output: {
    filename: '[name].js',
    path: `${__dirname}/../dist`,
  },
  target: 'electron-renderer',
  module: {
    rules: [
      {
        test: /\.(js|vue)$/,
        loader: 'eslint-loader',
        enforce: 'pre',
        options: {
          // eslint-disable-next-line global-require
          formatter: require('eslint-friendly-formatter'),
        },
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader',
        options: vueConf,
      },
      {
        test: /\.js$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
      },
      {
        test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
        loader: 'url-loader',
        options: {
          limit: 10000,
        },
      },
    ],
  },
};
