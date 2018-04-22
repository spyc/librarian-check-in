const sqlite3 = require('sqlite3').verbose();

module.exports = function openDB() {
  return new Promise((resolve, reject) => {
    const db = new sqlite3.Database(`${process.cwd()}/data/record.sqlite`, sqlite3.OPEN_READWRITE, (e) => {
      if (e) {
        reject(e);
      } else {
        resolve(db);
      }
    });
  });
};
