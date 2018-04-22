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

ipcMain.on('librarian.mutation', (event, { action, body }) => {
  openDB()
    .then((db) => {
      const sql = (
        action === 'add' ?
          'INSERT INTO librarian (name, class, class_no, id) VALUES (?, ?, ?, ?)' :
          'UPDATE librarian SET name = ?, class = ?, class_no = ? WHERE id = ?'
      );
      return new Promise((resolve, reject) => {
        db.run(sql, body.name, body.class, body.class_no, body.id, (err) => {
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
      event.sender.send('librarian.mutation.reply', { done: true, error: null });
    })
    .catch((e) => {
      console.error(e);
      event.sender.send('librarian.mutation.reply', { done: false, error: e });
    });
});
