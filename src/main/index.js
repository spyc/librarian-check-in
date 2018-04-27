const { app, BrowserWindow } = require('electron');
const { format } = require('url');
const path = require('path');

require('./ipc');

const isDevelopment = process.env.NODE_ENV !== 'production';

// global reference to mainWindow (necessary to prevent window from being garbage collected)
let mainWindow;

function createMainWindow() {
  const window = new BrowserWindow({ width: 960, height: 720 });
  if (isDevelopment) {
    window.loadURL('http://localhost:8080');

    BrowserWindow.addDevToolsExtension('./data/VueDevTools');
    mainWindow.webContents.openDevTools();
  } else {
    window.loadURL(format({
      pathname: path.join(app.getAppPath(), 'index.html'),
      protocol: 'file',
      slashes: true,
    }));
  }

  window.on('closed', () => {
    mainWindow = null;
  });
}

// quit application when all windows are closed
app.on('window-all-closed', () => {
  app.quit();
});

app.on('activate', () => {
  // on macOS it is common to re-create a window even after all windows have been closed
  if (mainWindow === null) {
    mainWindow = createMainWindow();
  }
});

// create main BrowserWindow when electron is ready
app.on('ready', () => {
  mainWindow = createMainWindow();
});
