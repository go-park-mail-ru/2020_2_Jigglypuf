
CREATE DATABASE cinema_interface;
/* users table */
CREATE TABLE users
(
    ID serial NOT NULL PRIMARY KEY,
    Username VARCHAR(32) NOT NULL UNIQUE ,
    Password VARCHAR(64) NOT NULL
);

/* profile table */

CREATE TABLE profile
(
    user_id integer NOT NULL PRIMARY KEY REFERENCES users (ID),
    ProfileName VARCHAR(32),
    ProfileSurname VARCHAR(32),
    AvatarPath VARCHAR(64)
);

/* cinema table */

CREATE TABLE cinema
(
    ID serial NOT NULL PRIMARY KEY,
    CinemaName VARCHAR(32) NOT NULL,
    Address VARCHAR(64) NOT NULL,
    Author_ID integer REFERENCES users (ID)
);

/* movie table */

CREATE TABLE movie
(
    ID serial NOT NULL PRIMARY KEY ,
    MovieName TEXT NOT NULL UNIQUE,
    Description TEXT,
    Rating FLOAT DEFAULT 0.0,
    Rating_count INTEGER DEFAULT 0,
    PathToAvatar VARCHAR(64)
);

/* rating table */

CREATE TABLE rating_history
(
    ID serial NOT NULL PRIMARY KEY ,
    user_id integer references users (ID),
    movie_id integer references movie (ID),
    movie_rating integer NOT NULL
);

INSERT INTO cinema (CinemaName, Address)
VALUES  ('CinemaScope1','Москва, Первая улица, д.1'),
        ('CinemaScope2','Москва, Первая улица, д.2'),
        ('CinemaScope3','Москва, Первая улица, д.3'),
        ('CinemaScope4','Москва, Первая улица, д.4');

INSERT INTO movie (MovieName,Description,PathToAvatar)
VALUES  ('Гренландия','Greenland description','/media/greenland.jpg'),
        ('Антибеллум','Антибеллум description','/media/antibellum.jpg'),
        ('Довод','Довод description','/media/dovod.jpg'),
        ('Гнездо','Гнездо description','/media/gnezdo.jpg'),
        ('Сделано в Италии','Италиан description','/media/italian.jpg'),
        ('Мулан','Мулан description','/media/mulan.jpg'),
        ('Никогда всегда всегда никогда','Никогда description','/media/nikogda.jpg'),
        ('После','После description','/media/posle.jpg'),
        ('Стрельцов','Стрельцов description','/media/strelcov.jpg');