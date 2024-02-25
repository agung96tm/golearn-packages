ALTER TABLE articles
    ADD COLUMN image INT REFERENCES medias(id);