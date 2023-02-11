CREATE TABLE users (
    id bigserial not null primary key,
    login varchar not null,
    password_hash varchar not null
);

CREATE TABLE articles (
    id bigserial not null primary key,
    title varchar not null,
    author varchar not null,
    content varchar not null
);