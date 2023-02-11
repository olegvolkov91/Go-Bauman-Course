CREATE TABLE users (
    id bigserial not null primary key,
    login varchar not null unique,
    password_hash varchar not null
);

CREATE TABLE articles (
    id bigserial not null primary key,
    title varchar not null unique,
    author varchar not null FOREIGN KEY REFERENCES users(id),
    content varchar not null,
)