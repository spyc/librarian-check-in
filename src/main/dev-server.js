const express = require('express');
const webpack = require('webpack');
const serveStatic = require('serve-static');
const devConfig = require('../../config/webpack.dev.conf');

const app = express();

app.set('view engine', 'ejs');

const compiler = webpack(devConfig);

const devMiddleware = require('webpack-dev-middleware')(compiler, {
  publicPath: devConfig.output.publicPath,
  quiet: true,
});

const hotMiddleware = require('webpack-hot-middleware')(compiler, {
  log: () => {},
});

app.get('/', (req, res) => {
  res.render('index', {
    production: false,
  });
});

app.use(serveStatic('build/assets'));

app.use(devMiddleware);
app.use(hotMiddleware);

app.listen(8080);
