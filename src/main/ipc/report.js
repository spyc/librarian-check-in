const { ipcMain } = require('electron');
const openDB = require('../service/db');

ipcMain.on('report.query', (event, query) => {
  openDB()
    .then(db => new Promise((resolve, reject) => {
      db.all(query.sql, query.param, (err, rows) => {
        if (err) {
          reject(err);
        } else {
          db.close();
          resolve(rows);
        }
      });
    }))
    .then((rows) => {
      event.sender.send('report.query.reply', {
        rows,
        success: true,
        error: null,
      });
    })
    .catch((error) => {
      console.error(error);
      event.sender.send('report.query.reply', {
        error,
        rows: null,
        success: false,
      });
    });
});
