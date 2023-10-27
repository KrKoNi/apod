CREATE TABLE PICS (
    copyright       TEXT,
    date            DATE NOT NULL UNIQUE,
    explanation     TEXT,
    hdurl           TEXT,
    media_type      TEXT,
    service_version TEXT,
    title           TEXT,
    url             TEXT,
    content         BYTEA
);