
CREATE DATABASE BackendCinemaInterface;
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

/* movies in cinema */

CREATE TABLE movies_in_cinema
(
    ID serial NOT NULL UNIQUE PRIMARY KEY,
    Movie_id INTEGER NOT NULL REFERENCES movie (ID),
    Cinema_id INTEGER NOT NULL REFERENCES cinema (ID),
    Rental_start DATE NOT NULL,
    Rental_end DATE NOT NULL
);
/* TODO cinema halls structure */


/* profile ticket purchases */



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

INSERT INTO movies_in_cinema (Movie_id, Cinema_id, Rental_start, Rental_end)
VALUES (1,2,'2020-09-03','2020-11-21'),
       (1,3,'2020-09-03','2020-11-03'),
       (1,4,'2020-05-07','2020-11-29'),
       (1,1,'2020-09-03','2020-11-15'),
       (2,2,'2020-09-03','2020-11-03'),
       (4,3,'2020-09-03','2020-11-29'),
       (3,4,'2020-09-03','2020-11-21'),
       (5,1,'2020-09-03','2020-11-30'),
       (8,2,'2020-11-03','2020-12-18'),
       (7,3,'2020-09-03','2020-12-31'),
       (3,1,'2020-09-03','2020-11-28'),
       (2,1,'2020-09-03','2020-12-29'),
       (6,1,'2020-05-11','2020-06-07'),
       (7,1,'2020-09-03','2020-11-29')