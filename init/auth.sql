set timezone = '+3';
CREATE DATABASE BackendCinemaInterfaceUser;

/* users table */
CREATE TABLE if not exists users
(
    ID serial NOT NULL PRIMARY KEY,
    Login VARCHAR(32) NOT NULL UNIQUE,
    Password VARCHAR(64) NOT NULL
);