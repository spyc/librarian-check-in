PRAGMA foreign_keys = ON;

CREATE TABLE librarian (
    id CHARACTER(8) NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    class CHARACTER(2) NOT NULL,
    class_no VARCHAR(2) NOT NULL
);

CREATE TABLE record(
    id CHARACTER(8) NOT NULL,
    check_in INTEGER NOT NULL,
    check_out INTEGER DEFAULT NULL,
    rank TEXT DEFAULT NULL,
    FOREIGN KEY (id) REFERENCES librarian(id)
);