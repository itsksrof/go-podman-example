-- Local database definition.

DROP DATABASE IF EXISTS local_db;

CREATE DATABASE local_db;

USE local_db;

DROP TABLE IF EXISTS channels;

CREATE TABLE channels (
    id int unsigned NOT NULL AUTO_INCREMENT,
    name varchar(30) NOT NULL DEFAULT '',
    PRIMARY KEY(id)
);

INSERT INTO channels VALUES(1, 'ThePrimeagen');
INSERT INTO channels VALUES(2, 'Randy');
INSERT INTO channels VALUES(3, 'CS Jackie');