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
        test: /\.js$/,
        loader: 'eslint-loader',
        enforce: 'pre',
        options: {
          // eslint-disable-next-line global-require
          formatter: require('eslint-friendly-formatter'),
        },
      },
      {
        test: /\.js$/,
        loader: 'babel-loader',
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
