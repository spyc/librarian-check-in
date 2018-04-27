const { ipcMain } = require('electron');
const moment = require('moment-timezone');
const openDB = require('../service/db');

const timezone = 'Asia/Hong_Kong';

ipcMain.on('checkin', (event, id) => {
  const sql = 'INSERT INTO record (id, check_in) VALUES (? ,?)';
  openDB()
    .then(db => new Promise((resolve, reject) => {
      db.run('PRAGMA foreign_keys = ON', (err) => {
        if (err) {
          reject(err);
        } else {
          resolve(db);
        }
      });
    }))
    .then(db => new Promise((resolve, reject) => {
      db.get('SELECT COUNT(id) AS row FROM record WHERE id = ? AND check_out IS NULL', id, (err, row) => {
        if (err) {
          reject(err);
        } else if (row.row > 0) {
          resolve(false);
        } else {
          resolve(db);
        }
      });
    }))
    .then(db => new Promise((resolve, reject) => {
      if (!db) {
        db.close();
        resolve();
        return;
      }
      db.run(sql, id, moment().tz(timezone).unix(), (err) => {
        if (err) {
          reject(err);
        } else {
          db.close();
          resolve();
        }
      });
    }))
    .then(() => {
      event.sender.send('checkin.reply', { done: true, error: null });
    })
    .catch((e) => {
      console.error(e);
      event.sender.send('checkin.reply', { done: false, error: e });
    });
});

ipcMain.on('checkout', (event, { id, rank }) => {
  const sql = 'UPDATE record SET check_out = ?, rank = ? WHERE id = ? AND check_out IS NULL';
  openDB()
    .then(db => new Promise((resolve, reject) => {
      db.run('PRAGMA foreign_keys = ON', (err) => {
        if (err) {
          reject(err);
        } else {
          resolve(db);
        }
      });
    }))
    .then(db => new Promise((resolve, reject) => {
      db.get('SELECT COUNT(id) AS row FROM record WHERE id = ? AND check_out IS NULL', id, (err, row) => {
        if (err) {
          reject(err);
        } else if (row.row === 0) {
          resolve(false);
        } else {
          resolve(db);
        }
      });
    }))
    .then(db => new Promise((resolve, reject) => {
      if (!db) {
        db.close();
        resolve();
        return;
      }
      db.run(sql, moment().tz(timezone).unix(), rank, id, (err) => {
        if (err) {
          reject(err);
        } else {
          db.close();
          resolve();
        }
      });
    }))
    .then(() => {
      event.sender.send('checkout.reply', { done: true, error: null });
    })
    .catch((e) => {
      console.error(e);
      event.sender.send('checkout.reply', { done: false, error: e });
    });
});

ipcMain.on('checkout.batch', (event) => {
  const sql = 'UPDATE record SET check_out = ?, rank = ? WHERE check_out IS NULL';
  openDB()
    .then(db => new Promise((resolve, reject) => {
      db.run('PRAGMA foreign_keys = ON', (err) => {
        if (err) {
          reject(err);
        } else {
          resolve(db);
        }
      });
    }))
    .then(db => new Promise((resolve, reject) => {
      if (!db) {
        db.close();
        resolve();
        return;
      }
      db.run(sql, moment().tz(timezone).unix(), 'fair', (err) => {
        if (err) {
          reject(err);
        } else {
          db.close();
          resolve();
        }
      });
    }))
    .then(() => {
      event.sender.send('checkout.reply', { done: true, error: null });
    })
    .catch((e) => {
      console.error(e);
      event.sender.send('checkout.reply', { done: false, error: e });
    });
});

ipcMain.on('checkin.list', (event) => {
  openDB()
    .then(db => new Promise((resolve, reject) => {
      db.run('PRAGMA foreign_keys = ON', (err) => {
        if (err) {
          reject(err);
        } else {
          resolve(db);
        }
      });
    }))
    .then((db) => {
      const sql = `
      SELECT librarian.id, librarian.name, librarian.class, librarian.class_no FROM librarian
      INNER JOIN record ON record.id = librarian.id
      WHERE record.check_out IS NULL
      `;
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
      event.sender.send('checkin.list.reply', result);
    });
});

