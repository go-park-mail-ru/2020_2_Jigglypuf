SET timezone ='+3';
CREATE DATABASE BackendCinemaInterface;
/* users table */
CREATE TABLE users
(
    ID serial NOT NULL PRIMARY KEY,
    Login VARCHAR(32) NOT NULL UNIQUE ,
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
    PathToAvatar varchar(64) default '/media/cinemaDefault.jpg'
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
    Actors VARCHAR(64) default '',
    Rating_count INTEGER DEFAULT 0,
    PathToAvatar VARCHAR(64),
    pathToSliderAvatar VARCHAR(64) default ''
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
    Cost INTEGER NOT NULL,
    UNIQUE(Cinema_ID,Hall_ID,Premiere_time)
);

/* tickets */
CREATE TABLE ticket
(
    ID serial not null unique primary key,
    User_login VARCHAR(32) not null references users (Login),
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

INSERT INTO movie (MovieName,Actors,Description,Genre,Duration,Producer,Country,Release_Year,Age_group,PathToAvatar,pathToSliderAvatar)
VALUES  ('Гренландия','Билл Кроун','Greenland description','Tragedy',112,'Tarantino','America',2016,16,'/media/greenland.jpg',''),
        ('Антибеллум','Джейк Келлуа','Антибеллум description','Comedy',118,'Tarantino','America',2012,12,'/media/antibellum.jpg',''),
        ('Довод','Роберт Паттинсон','Главный герой — секретный агент, который проходит жестокий тест на надежность и присоединяется к невероятной миссии. От ее выполнения зависит судьба мира, а для успеха необходимо отбросить все прежние представления о пространстве и времени.','Боевик',160,'Кристофер Нолан','America',2020,18,'/media/tenet_poster.jpg','/media/tenet_slider.png'),
        ('Гнездо','Сара Коннор','Гнездо description','Drama',180,'No name','Canada',2006,10,'/media/gnezdo.jpg',''),
        ('Сделано в Италии','Джордж Клуни','Италиан description','Comedy',100,'Zarukko','Italy',2020,12,'/media/italian.jpg',''),
        ('Мулан','Лю Ифэй','История о бесстрашной молодой девушке, которая выдаёт себя за мужчину, чтобы вступить в ряды армии, противостоящей Северным захватчикам, надвигающимся на Китай. Старшая дочь храброго воина Хуа, Мулан — энергичная и решительная девушка. Когда Император издаёт указ о том, что один мужчина из каждой семьи должен вступить в ряды Имперской армии, Мулан занимает место своего больного отца, еще не зная о том, что ей предстоит прославиться как один из самых величайших воинов в истории Китая.','Tragedy',132,'Zue che ke','China',2020,18,'/media/mulan.jpg','/media/mulan_slider.png'),
        ('Никогда всегда всегда никогда','Евгения Смолина','Никогда description','Fantastic',130,'Васильев','Russia',2018,18,'/media/nikogda.jpg',''),
        ('После','Сьюзэн МакМартин','После description','Fantastic',180,'Rukko','Spain',2020,18,'/media/posle.jpg',''),
        ('Стрельцов','Виталий Хлев','К 20 годам у кумира миллионов Эдуарда Стрельцова есть все, о чем только можно мечтать: талант, слава и любовь. Вся страна с замиранием сердца ждет от сборной и ее восходящей звезды победы на предстоящем Чемпионате мира по футболу. Но за два дня до отъезда команды против Стрельцова выдвигается обвинение, которое вмиг все перечеркивает. Вместо дуэли с гениальным бразильцем Пеле, которая могла стать самой зрелищной в истории футбола, Стрельцова ждет тюрьма. Неужели он и правда преступник? Сможет ли он после 5 лет лагерей вновь выйти на поле и доказать, что он — настоящий чемпион, достойный всенародной любви?','Drama',120,'Илья Алексеевич Учитель','Russia',2008,18,'/media/strelcov.jpg','/media/streltsov_slider.png'),
        ('Ловец Снов','Рада Митчел', 'Совсем немного времени прошло после убийства жены Люка соседским мальчишкой в отдалённом лесном домике, но мужчина привозит туда свою новую пассию Гейл и сына Джоша. Ребёнка мучают страшные сны, в которых ему является мёртвая мама, а Гейл — детский психилог со стажем — изо всех сил пытается помочь мальчику. Однажды, наслушавшись рассказов соседки про ловцы снов, Джош крадёт у неё из сундука полезную, как он думал, для избавления от кошмаров вещь, но после этого его сны становятся ещё более реалистичными и пугающими.','Триллер',85,'Керри Харрис','США',2020,18,'/media/dream_catcher_poster.jpg','/media/dream_catcher_slider.png'),
        ('Однажды в… Голливуде','Леонардно Ди Каприо','Фильм повествует о череде событий, произошедших в Голливуде в 1969 году, на закате его «золотого века». По сюжету, известный ТВ актер Рик Далтон и его дублер Клифф Бут пытаются найти свое место в стремительно меняющемся мире киноиндустрии.','Комедия',160,'Квентин Тарантино','США',2019,16,'/media/once_upon_a_time_in_hollywood_poster.jpg','');

INSERT INTO cinema_hall (Place_amount,Hall_params)
VALUES (15,'{"levels":[{"place":1,"row":1},{"place":2,"row":1}]}'),
       (10,'{"levels":[{"place":1,"row":2}]}');

INSERT INTO schedule(Movie_ID, Cinema_ID, Hall_ID, Premiere_time, Cost)
VALUES (1,2,2,now() + interval '1 hour', 400),
       (3,3,1,now() + interval '2 days', 350),
       (2,2,1,now() + interval '30 days', 840),
       (4,3,2,now() + interval '3 days', 200),
       (6,4,1,now() + interval '1 day 2 hours', 150),
       (5,4,2,now() + interval '2 hours 30 minutes', 899),
       (3,1,2,now() + interval '1 hour', 540),
       (2,1,1,now() + interval '1 day', 630),
       (5,1,1,now() + interval '3 days',320),
       (1,2,2,now() + interval '1 month',777),
       (3,3,1,now() + interval '2 hours',322),
       (2,2,1,now() + interval '20 days',228),
       (4,3,2,now() + interval '1 day', 666),
       (6,4,1,now() + interval '3 days 2 hours',133),
       (5,4,2,now() + interval '3 hours 30 minutes',1701),
       (3,1,2,now() + interval '10 hours',324),
       (7,1,1,now() + interval '5 days',764),
       (10,2,2,now() + interval '6 days', 549),
       (10,3,1,now() + interval '3 days', 359),
       (10,2,1,now() + interval '15 days',349),
       (10,3,2,now() + interval '9 days',339),
       (10,4,1,now() + interval '1 day 15 hours',777),
       (10,4,2,now() + interval '10 hours',322),
       (10,1,2,now() + interval '5 hours',228),
       (10,1,1,now() + interval '11 days',666);
