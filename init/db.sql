SET timezone ='+3';
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
    Hall_count integer NOT NULL,
    Author_ID integer REFERENCES users (ID)
);

/* movie table */

CREATE TABLE movie
(
    ID serial NOT NULL PRIMARY KEY,
    MovieName TEXT NOT NULL UNIQUE,
    Description TEXT,
    Genre VARCHAR(64),
    Duration integer,
    Producer VARCHAR(64),
    Country VARCHAR(64),
    Release_Year integer,
    Age_group integer,
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
/* cinema halls structure */
CREATE TABLE cinema_hall
(
    ID serial NOT NULL UNIQUE PRIMARY KEY,
    Place_amount integer not null,
    Hall_params jsonb not null
);

/* schedule table */

CREATE TABLE schedule
(
    ID serial NOT NULL UNIQUE PRIMARY KEY,
    Movie_ID INTEGER NOT NULL REFERENCES movie (ID),
    Cinema_ID INTEGER NOT NULL REFERENCES cinema (ID),
    Hall_ID INTEGER NOT NULL REFERENCES cinema_hall (ID),
    Premiere_time timestamp NOT NULL,
    UNIQUE(Cinema_ID,Hall_ID,Premiere_time)
);

/* tickets */
CREATE TABLE ticket
(
    ID serial not null unique primary key,
    User_login VARCHAR(32) not null references users (Username),
    schedule_id integer not null references schedule (ID),
    transaction_date timestamp default now(),
    row integer not null,
    place integer not null,
    unique(schedule_id,row,place)
);

INSERT INTO cinema (CinemaName, Address, Hall_count)
VALUES  ('CinemaScope1','Москва, Первая улица, д.1',1),
        ('CinemaScope2','Москва, Первая улица, д.2',2),
        ('CinemaScope3','Москва, Первая улица, д.3',3),
        ('CinemaScope4','Москва, Первая улица, д.4',4);

INSERT INTO movie (MovieName,Description,Genre,Duration,Producer,Country,Release_Year,Age_group,PathToAvatar)
VALUES  ('Гренландия','Greenland description','Tragedy',112,'Tarantino','America',2016,16,'/media/greenland.jpg'),
        ('Антибеллум','Антибеллум description','Comedy',118,'Tarantino','America',2012,12,'/media/antibellum.jpg'),
        ('Довод','Довод description','Thriller',160,'Nolan','America',2020,18,'/media/dovod.jpg'),
        ('Гнездо','Гнездо description','Drama',180,'No name','Canada',2006,10,'/media/gnezdo.jpg'),
        ('Сделано в Италии','Италиан description','Comedy',100,'Zarukko','Italy',2020,12,'/media/italian.jpg'),
        ('Мулан','Мулан description','Tragedy',132,'Zue che ke','China',2020,18,'/media/mulan.jpg'),
        ('Никогда всегда всегда никогда','Никогда description','Fantastic',130,'Васильев','Russia',2018,18,'/media/nikogda.jpg'),
        ('После','После description','Fantastic',180,'Rukko','Spain',2020,18,'/media/posle.jpg'),
        ('Стрельцов','Стрельцов description','Drama',120,'Васильев','Russia',2008,18,'/media/strelcov.jpg');

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
       (7,1,'2020-09-03','2020-11-29');
INSERT INTO cinema_hall (Place_amount,Hall_params)
VALUES (15,'{"levels":[{"place":1,"row":1},{"place":2,"row":1}]}'),
       (10,'{"levels":[{"place":1,"row":2}]}');

INSERT INTO schedule(Movie_ID, Cinema_ID, Hall_ID, Premiere_time)
VALUES (1,2,3,now()),
       (1,3,1,now()),
       (2,2,1,now()),
       (4,3,2,now()),
       (3,4,1,now()),
       (5,1,5,now()),
       (8,2,2,now()),
       (7,3,3,now()),
       (3,1,1,now()),
       (2,1,4,now()),
       (6,1,3,now()),
       (7,1,2,now());
