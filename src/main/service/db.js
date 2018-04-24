const sqlite3 = require('sqlite3').verbose();
const path = require('path');

module.exports = function openDB() {
  return new Promise((resolve, reject) => {
    const file = path.join(process.cwd(), 'data/record.sqlite');
    const db = new sqlite3.Database(file, sqlite3.OPEN_READWRITE, (e) => {
      if (e) {
        reject(e);
      } else {
        resolve(db);
      }
    });
  })
    .catch(console.error);
};
