const { ipcMain } = require('electron');
const openDB = require('../service/db');

ipcMain.on('librarian.list', (event) => {
  openDB()
    .then((db) => {
      const sql = 'SELECT id, name, class, class_no FROM librarian';
      return new Promise((resolve, reject) => {
        db.all(sql, (err, rows) => {
          if (err) {
            reject(err);
          } else {
            db.close();
            resolve(rows);
          }
        });
      });
    })
    .then((result) => {
      event.sender.send('librarian.list.reply', result);
    });
});

ipcMain.on('librarian.add', (event, arg) => {
  openDB()
    .then((db) => {
      const sql = 'INSERT INTO librarian (id, name, class, class_no) VALUES (?, ?, ?, ?)';
      return new Promise((resolve, reject) => {
        db.run(sql, arg.id, arg.name, arg.class, arg.class_no, (err) => {
          if (err) {
            reject(err);
          } else {
            db.close();
            resolve();
          }
        });
      });
    })
    .then(() => {
      event.sender.send('librarian.add.reply', { done: true, error: null });
    })
    .catch((e) => {
      console.error(e);
      event.sender.send('librarian.add.reply', { done: false, error: e });
    });
});
