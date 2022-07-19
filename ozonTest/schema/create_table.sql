CREATE TABLE IF NOT EXISTS linksDB (
    id serial not null unique,
    long_url  varchar(1024) NOT NULL,
    short_url varchar(12) NOT NULL unique
);
