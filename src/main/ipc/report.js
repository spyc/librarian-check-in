const fs = require('fs');
const { ipcMain } = require('electron');
const openDB = require('../service/db');

ipcMain.on('report.query', (event, { sql, params }) => {
  openDB()
    .then(db => new Promise((resolve, reject) => {
      db.all(sql, params, (err, rows) => {
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

function exportFile(event, { filename, headers, rows }) {
  if (fs.existsSync(filename)) {
    fs.unlink(filename);
  }

  const writer = fs.createWriteStream(filename);
  writer.write(headers.map(field => `"${field}"`).join(','), 'utf8');
  writer.write('\n');

  rows.forEach((row) => {
    if (Array.isArray(row)) {
      writer.write(row.map(field => `"${field}"`).join(','), 'utf8');
    } else {
      writer.write(Object.keys(row).map(key => `"${row[key]}"`).join(','), 'utf8');
    }
    writer.write('\n');
  });

  writer.on('error', (e) => {
    console.error(e);
    event.sender.send('save.file.done', {
      filename,
      success: false,
      error: e,
    });
  });

  writer.on('close', () => {
    event.sender.send('save.file.done', {
      filename,
      error: null,
      success: true,
    });
  });

  writer.end();
}

ipcMain.on('save.file', (event, { filename, headers, rows }) => {
  if (fs.existsSync(filename)) {
    fs.unlink(filename, (error) => {
      if (error) {
        console.error(error);
        event.sender.send({
          filename,
          error,
          success: false,
        });
      } else {
        exportFile(event, { filename, headers, rows });
      }
    });
  } else {
    exportFile(event, { filename, headers, rows });
  }
});
