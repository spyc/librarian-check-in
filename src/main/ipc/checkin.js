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
