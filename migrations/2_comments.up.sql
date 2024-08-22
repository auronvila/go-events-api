CREATE TABLE IF NOT EXISTS comments
(
    id        TEXT PRIMARY KEY NOT NULL,
    userId    TEXT             NOT NULL,
    eventId   TEXT             NOT NULL,
    text      TEXT,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES users (id),
    FOREIGN KEY (eventId) REFERENCES events (id)
);
