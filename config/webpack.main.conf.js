const nodeExternals = require('webpack-node-externals');

module.exports = {
  mode: 'production',
  devtool: '#source-map',
  entry: {
    main: './src/main/index.js',
  },
  externals: [nodeExternals()],
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
