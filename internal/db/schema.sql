-- Some important differences between the schema here and the version
-- shown in the book:
--
-- 1) SQLite doesn't support the VARCHAR length restrictions (VARCHAR
--    is translated into TEXT).
-- 2) SQLite doesn't support user accounts, which means granular
--    permissions cannot be set with GRANT and REVOKE statements.
--    While prepared statements eliminate the possibility of SQL
--    injection, the principle of least privilege can't be applied
--    here in the same way as the book.

CREATE TABLE spaces(
       space_id INTEGER PRIMARY KEY AUTOINCREMENT,
       name VARCHAR(255) NOT NULL,
       owner VARCHAR(30) NOT NULL
);

CREATE TABLE messages(
       space_id INTEGER NOT NULL REFERENCES spaces(space_id),
       msg_id INTEGER PRIMARY KEY AUTOINCREMENT,
       author VARCHAR(30) NOT NULL,
       msg_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       msg_text VARCHAR(1024) NOT NULL
);

CREATE INDEX msg_timestamp_idx ON messages(msg_time);

CREATE UNIQUE INDEX space_name_idx ON spaces(name);
