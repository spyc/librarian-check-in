require('./dev-server');

require('child_process').spawn('yarn', ['dev:main'], { stdio: 'inherit' });
