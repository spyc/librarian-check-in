const merge = require('webpack-merge');
const { HotModuleReplacementPlugin } = require('webpack');
const FriendlyErrorsPlugin = require('friendly-errors-webpack-plugin');
const base = require('./webpack.base.conf');
const utils = require('./utils');

module.exports = merge.smart(base, {
  mode: 'development',
  module: {
    rules: utils.styleLoaders({ sourceMap: true }),
  },
  plugins: [
    new HotModuleReplacementPlugin(),
    new FriendlyErrorsPlugin(),
  ],
  optimization: {
    noEmitOnErrors: true,
  },
});
