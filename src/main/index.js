const { app, BrowserWindow } = require('electron');
const { format } = require('url');
const path = require('path');

const isDevelopment = process.env.NODE_ENV !== 'production';

if (isDevelopment) {
  // eslint-disable-next-line global-require
  require('./dev-server');
}

// global reference to mainWindow (necessary to prevent window from being garbage collected)
let mainWindow;

function createMainWindow() {
  const window = new BrowserWindow();

  if (isDevelopment) {
    window.loadURL('http://localhost:8080');
  } else {
    window.loadURL(format({
      pathname: path.join(__dirname, 'index.html'),
      protocol: 'file',
      slashes: true,
    }));
  }

  window.on('closed', () => {
    mainWindow = null;
  });

  return window;
}

// quit application when all windows are closed
app.on('window-all-closed', () => {
  // on macOS it is common for applications to stay open until the user explicitly quits
  if (process.platform !== 'darwin') {
    app.quit();
  }
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
  if (isDevelopment) {
    BrowserWindow.addDevToolsExtension('./data/VueDevTools');
    mainWindow.webContents.openDevTools();
  }
});
