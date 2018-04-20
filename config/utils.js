const MiniCssExtractPlugin = require('mini-css-extract-plugin');

exports.cssLoaders = function (options) {
  // eslint-disable-next-line no-param-reassign
  options = options || {};

  const cssLoader = {
    loader: 'css-loader',
    options: {
      minimize: process.env.NODE_ENV === 'production',
      sourceMap: options.sourceMap,
    },
  };

  // generate loader string to be used with extract text plugin
  function generateLoaders(loader, loaderOptions) {
    const loaders = [cssLoader];
    if (loader) {
      loaders.push({
        loader: `${loader}-loader`,
        options: Object.assign({}, loaderOptions, {
          sourceMap: options.sourceMap,
        }),
      });
    }

    // Extract CSS when that option is specified
    // (which is the case during production build)
    if (options.extract) {
      return [MiniCssExtractPlugin.loader].concat(loaders);
    }
    return ['style-loader'].concat(loaders);
  }

  return {
    css: generateLoaders(),
    postcss: generateLoaders(),
    less: generateLoaders('less'),
    sass: generateLoaders('sass', { indentedSyntax: true }),
    scss: generateLoaders('sass'),
    stylus: generateLoaders('stylus'),
    styl: generateLoaders('stylus'),
  };
};

exports.styleLoaders = function (options) {
  const loaders = exports.cssLoaders(options);
  return Object.keys(loaders).map((extension) => {
    const loader = loaders[extension];
    return {
      test: new RegExp(`\\.${extension}$`),
      use: loader,
    };
  });
};
