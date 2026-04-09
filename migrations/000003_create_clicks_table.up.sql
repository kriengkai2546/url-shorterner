CREATE TABLE clicks (
    id         SERIAL PRIMARY KEY,
    url_id     INTEGER REFERENCES urls(id) ON DELETE CASCADE,
    clicked_at TIMESTAMP DEFAULT NOW()
);