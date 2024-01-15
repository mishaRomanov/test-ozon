CREATE TABLE IF NOT EXISTS links(
    link_id SERIAL PRIMARY KEY,
    old_link TEXT NOT NULL,
    new_link TEXT NOT NULL
);