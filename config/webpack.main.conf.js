module.exports = {
  mode: 'production',
  devtool: '#source-map',
  entry: {
    main: './src/main/index.js',
  },
  externals: [
    'source-map-support',
    'electron',
    'webpack',
    'electron-devtools-installer',
    'webpack/hot/log-apply-result',
    'source-map-support/source-map-support.js',
  ],
  output: {
    filename: '[name].js',
    chunkFilename: '[id].bundle.js',
    path: `${__dirname}/../dist`,
  },
  target: 'electron-main',
  optimization: {
    minimize: false,
  },
};
