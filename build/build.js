const webpack = require('webpack');
const shell = require('shelljs');
const fs = require('fs');
const ejs = require('ejs');

process.env.NODE_ENV = 'production';

const rendererConfig = require('../config/webpack.prod.conf');
const mainConfig = require('../config/webpack.main.conf');

shell.rm('-rf', 'dist');

shell.cp('-r', 'build/assets', 'dist');
shell.cp('yarn.lock', 'dist');
shell.cp('package.json', 'dist');

ejs.renderFile('views/index.ejs', { production: true }, (err, str) => {
  if (err) {
    throw err;
  } else {
    fs.writeFile('dist/index.html', str, 'utf8', (e) => {
      if (e) {
        console.error(e);
      }
    });
  }
});

webpack([mainConfig, rendererConfig], (err, stats) => {
  if (err) {
    console.error(err);
    throw err;
  }
  console.log(stats.toString({
    colors: true,
    modules: false,
    children: false,
    chunks: false,
    chunkModules: false,
  }));
});
