CREATE TABLE IF NOT EXISTS user (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL UNIQUE CHECK (length(username) > 0),
    password_hash TEXT NOT NULL CHECK (length(password_hash) > 0)
);

CREATE TABLE IF NOT EXISTS chat (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL CHECK (length(title) > 0),
    user_id TEXT NOT NULL CHECK (length(user_id) > 0),
    default_model_id text NOT NULL CHECK (length(default_model_id) > 0),
    created_at TEXT NOT NULL CHECK (length(created_at) > 0),
    updated_at TEXT NOT NULL CHECK (length(updated_at) > 0),
    deleted_at TEXT,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS message (
    id TEXT PRIMARY KEY,
    chat_id TEXT NOT NULL CHECK (length(chat_id) > 0),
    sender TEXT NOT NULL CHECK (length(sender) > 0),
    text TEXT NOT NULL CHECK (length(text) > 0),
    created_at TEXT NOT NULL CHECK (length(created_at) > 0),
    FOREIGN KEY (chat_id) REFERENCES chat(id)
);

CREATE TABLE IF NOT EXISTS model_description (
    name TEXT NOT NULL CHECK (length(name) > 0),
    description TEXT NOT NULL CHECK (length(description) > 0),
    PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS ollama_model_pull_event (
    id TEXT PRIMARY KEY,
    occurred_at TEXT NOT NULL CHECK (length(occurred_at) > 0),
    type TEXT NOT NULL CHECK (length(type) > 0),
    model TEXT NOT NULL CHECK (length(model) > 0)
);

PRAGMA foreign_keys = ON;