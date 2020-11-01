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
    Premiere_time timestamptz NOT NULL,
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

INSERT INTO cinema_hall (Place_amount,Hall_params)
VALUES (15,'{"levels":[{"place":1,"row":1},{"place":2,"row":1}]}'),
       (10,'{"levels":[{"place":1,"row":2}]}');

INSERT INTO schedule(Movie_ID, Cinema_ID, Hall_ID, Premiere_time)
VALUES (1,2,2,now() + interval '1 hour'),
       (3,3,1,now() + interval '2 days'),
       (2,2,1,now() + interval '30 days'),
       (4,3,2,now() + interval '3 days'),
       (6,4,1,now() + interval '1 day 2 hours'),
       (5,4,2,now() + interval '2 hours 30 minutes'),
       (3,1,2,now() + interval '1 hour'),
       (2,1,1,now() + interval '1 day'),
       (5,1,1,now() + interval '3 days'),
       (1,2,2,now() + interval '1 month'),
       (3,3,1,now() + interval '2 hours'),
       (2,2,1,now() + interval '20 days'),
       (4,3,2,now() + interval '1 day'),
       (6,4,1,now() + interval '3 days 2 hours'),
       (5,4,2,now() + interval '3 hours 30 minutes'),
       (3,1,2,now() + interval '10 hours'),
       (7,1,1,now() + interval '5 days');
