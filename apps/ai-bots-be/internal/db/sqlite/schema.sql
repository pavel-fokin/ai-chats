CREATE TABLE IF NOT EXISTS user (
  pk INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  id TEXT NOT NULL UNIQUE,
  username TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS chat_user (
    pk INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    FOREIGN KEY (chat_id) REFERENCES chat(id),
    FOREIGN KEY (user_id) REFERENCES user(id),
    UNIQUE(chat_id, user_id)
);

CREATE TABLE IF NOT EXISTS chat (
    pk INTEGER PRIMARY KEY AUTOINCREMENT,
    id TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS message (
    pk INTEGER PRIMARY KEY AUTOINCREMENT,
    id TEXT NOT NULL UNIQUE,
    chat_id TEXT NOT NULL,
    sender TEXT NOT NULL,
    text TEXT NOT NULL,
    -- created_at TEXT NOT NULL,
    FOREIGN KEY (chat_id) REFERENCES chat(id)
);