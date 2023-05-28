-- Local database definition.

DROP DATABASE IF EXISTS local_db;

CREATE DATABASE local_db;

USE local_db;

DROP TABLE IF EXISTS books;

CREATE TABLE books (
    id int unsigned NOT NULL AUTO_INCREMENT,
    title varchar(60) NOT NULL DEFAULT '',
    author varchar(60) NOT NULL DEFAULT '',
    PRIMARY KEY(id)
);

INSERT INTO books VALUES(1, '1984', 'George Orwell');
INSERT INTO books VALUES(2, 'Creativity, Inc', 'Ed Catmull');
INSERT INTO books VALUES(3, 'The Brothers Karamazov', 'Fyodor Dostoevsky');