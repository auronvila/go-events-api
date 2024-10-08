CREATE TABLE IF NOT EXISTS users (
                                     id TEXT PRIMARY KEY NOT NULL,
                                     email TEXT NOT NULL UNIQUE,
                                     password TEXT NOT NULL,
                                     sex TINYINT NOT NULL DEFAULT 0,
                                     username TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS events (
                                      id TEXT PRIMARY KEY NOT NULL,
                                      name TEXT NOT NULL,
                                      description TEXT NOT NULL,
                                      location TEXT NOT NULL,
                                      dateTime DATETIME NOT NULL,
                                      userId TEXT,
                                      FOREIGN KEY(userId) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS user_events (
                                             id TEXT PRIMARY KEY NOT NULL,
                                             event_id TEXT,
                                             user_id TEXT,
                                             FOREIGN KEY(event_id) REFERENCES events(id),
                                            FOREIGN KEY(user_id) REFERENCES users(id)
);
