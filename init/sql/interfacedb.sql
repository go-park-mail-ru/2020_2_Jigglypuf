SET timezone ='+3';

/* cinema table */
CREATE TABLE if not exists cinema
(
    ID serial NOT NULL PRIMARY KEY,
    CinemaName VARCHAR(32) NOT NULL,
    Address VARCHAR(64) NOT NULL,
    Hall_count integer NOT NULL,
    PathToAvatar varchar(64) default '/media/cinemaDefault.jpg'
);

/* movie table */

CREATE TABLE if not exists genre
(
    ID serial not null primary key,
    Genre_Name varchar(64)
);

CREATE TABLE if not exists movie
(
    ID serial NOT NULL PRIMARY KEY,
    MovieName TEXT NOT NULL UNIQUE,
    Description TEXT,
    Duration integer,
    Producer VARCHAR(64),
    Country VARCHAR(64),
    Release_Year integer,
    Age_group integer,
    Rating FLOAT DEFAULT 0.0,
    Rating_count INTEGER DEFAULT 0,
    PathToAvatar VARCHAR(64),
    pathToSliderAvatar VARCHAR(64) default ''
);

CREATE TABLE if not exists movie_reply
(
    ID serial not null primary key,
    MovieID integer not null references movie(ID),
    UserID integer not null,
    replyText text not null,
    UNIQUE(UserID, MovieID)
);

CREATE TABLE if not exists movie_genre
(
    ID serial not null primary key,
    movie_id integer not null references movie(ID),
    genre_id integer not null references genre(ID),
    UNIQUE (movie_id,genre_id)
);

CREATE TABLE if not exists actor
(
    ID serial not null primary key,
    Name varchar(64) default '',
    Surname varchar(64) default '',
    Patronymic varchar(64) default '',
    Description text default ''
);

CREATE TABLE if not exists movie_actors
(
    ID serial not null primary key,
    movie_id integer not null references movie(id),
    actor_id integer not null references actor(id),
    unique(movie_id, actor_id)
);

/* rating table */

CREATE TABLE if not exists rating_history
(
    ID serial NOT NULL PRIMARY KEY ,
    user_id integer,
    movie_id integer references movie (ID),
    movie_rating integer NOT NULL,
    UNIQUE(user_id, movie_id)
);

/* cinema halls structure */
CREATE TABLE if not exists cinema_hall
(
    ID serial NOT NULL UNIQUE PRIMARY KEY,
    Place_amount integer not null,
    Hall_params jsonb not null
);

/* schedule table */

CREATE TABLE if not exists schedule
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
CREATE TABLE if not exists ticket
(
    ID serial not null unique primary key,
    User_login VARCHAR(32) not null,
    schedule_id integer not null references schedule (ID),
    transaction_date timestamp default now(),
    row integer not null,
    place integer not null,
    transaction varchar(128) not null unique,
    unique(schedule_id,row,place)
);

INSERT INTO cinema (CinemaName, Address, Hall_count, pathtoavatar)
VALUES  ('Пять Звёзд','Москва, Большой Овчинниковский пер., 16',1, '/media/pyat_zvezd.jpg'),
        ('Формула Кино','Москва, Хорошевское ш., 27',2, '/media/formula_kino.jpg'),
        ('Кронверк Синема','Москва, Семеновская пл., 1',3, '/media/kronverk_cinema.jpg'),
        ('Центрфильм','Москва, Щёлковское ш., 100',4, '/media/cinemaDefault.jpg'),
        ('Пионер','Митьковский пр-д, 10', 3, '/media/Pioner_cinema.jpeg'),
        ('Москино Звезда','ул. Земляной Вал, 18/22 , стр.2', 3, '/media/Moskino_zvezda.jpg');


INSERT into genre(Genre_Name)
values ('Трагедия'),
       ('Комедия'),
       ('Приключения'),
       ('Боевик'),
       ('Фантастика'),
       ('История'),
       ('Мелодрама'),
       ('Аниме'),
       ('Триллер'),
       ('Драма'),
       ('Фэнтези'),
       ('Криминал'),
       ('Мультфильм');



INSERT INTO actor(Name, Surname)
VALUES ('Эдвард', 'Нортон'),
       ('Брэд','Питт'),
       ('Виктор', 'Хориняк'),
       ('Мила', 'Сивацкая'),
       ('Галь', 'Гадот'),
       ('Крис', 'Пайн'),
       ('Харрис', 'Дикинсон'),
       ('Даниэль', 'Брюль'),
       ('Мэл', 'Гибсон'),
       ('Уолтон', 'Гоггинс'),
       ('Дмитрий', 'Высоцкий'),
       ('Сергей', 'Маковецкий'),
       ('Майкл','Дж.Фокс'),
       ('Кристофер', 'Ллойд'),
       ('Мэттью', 'МакКонахи'),
       ('Чарли', 'Ханнэм'),
       ('Леонардо', 'ДиКаприо'),
       ('Мэтт', 'Дэймон'),
       ('Федор', 'Федотов'),
       ('Софья', 'Присс'),
       ('Минами', 'Такаяма'),
       ('Рэи', 'Сакума'),
       ('Скарлетт', 'Йоханссон'),
       ('Флоренс', 'Пью'),
       ('Николас', 'Кейдж'),
       ('Эмма', 'Стоун'),
       ('Леонид', 'Барац'),
       ('Ирина', 'Горбачева');

INSERT INTO movie (MovieName,Description,Duration,Producer,Country,Release_Year,Age_group,PathToAvatar,pathToSliderAvatar)
VALUES  ('Бойцовский клуб','Сотрудник страховой компании страдает хронической бессонницей и отчаянно пытается вырваться из мучительно скучной жизни. Однажды в очередной командировке он встречает некоего Тайлера Дёрдена — харизматического торговца мылом с извращенной философией. Тайлер уверен, что самосовершенствование — удел слабых, а единственное, ради чего стоит жить — саморазрушение. Проходит немного времени, и вот уже новые друзья лупят друг друга почем зря на стоянке перед баром, и очищающий мордобой доставляет им высшее блаженство. Приобщая других мужчин к простым радостям физической жестокости, они основывают тайный Бойцовский клуб, который начинает пользоваться невероятной популярностью.',139,'Дэвид Финчер','США',1999,18,'/media/fight_club.png','/media/fight_club_slider.png'),
        ('Последний богатырь: Корень зла','Во второй части зрители узнают об истоках древнего зла, с которым героям пришлось столкнуться в первом фильме, увидят новые уголки сказочного Белогорья, и станут свидетелями захватывающих схваток с участием былинных богатырей.',121,'Дмитрий Дьяченко','Россия',2020,6,'/media/bogatir.png','/media/bogatir_slider.png'),
        ('Чудо-женщина: 1984','1984 год. Диана всё ещё грустит по погибшему Стиву, борется с мелким криминалом и работает в музее Смитсоновского института. Однажды она знакомится с новой коллегой Барбарой, специалисткой широкого профиля, которой поручено изучить новые артефакты. Среди древних предметов оказывается загадочный кристалл, который исполняет желания. Так к Диане внезапно возвращается Стив, а застенчивая и неуклюжая Барбара обретает невероятную силу, сноровку и привлекательность. Но за волшебным кристаллом охотится бизнесмен Максвелл Лорд, и у него имеются очень коварные планы на этот артефакт.',151,'Пэтти Дженкинс','США',2021,12,'/media/wonder_woman.png','/media/wonder_woman_slider.png'),
        ('King’s Man: Начало','«Kingsman» — организация супершпионов, действующая на благо человечества вдали от любопытных глаз. И один из первых и самых талантливых оперативников в истории организации — Конрад, молодой и наглый сын герцога Оксфордского. Как и многие его друзья он мечтал служить на благо Англии, но в итоге оказался втянутым в тайный мир шпионов и убийц.',128,'Мэттью Вон','Великобритания',2021,16,'/media/kingsman.png','/media/kingsman_slider.png'),
        ('Охота на Санту','Обиженный на Санту мальчик нанимает киллера, чтобы отомстить за плохой подарок. Но он не подозревает, что Санта не так прост и за долгие годы службы приобрел много необычных навыков.',100,'Эшом Нелмс','Великобритания',2020,18,'/media/fatman.png','/media/fatman_slider.png'),
        ('Конь Юлий и большие скачки','Дождались: говорящий конь Юлий влюбился! И на этот раз все серьезно – он просит руки, то есть копыта королевской кобылы по имени Звезда Востока у султана Рашида. Но сватать кобылу королевских кровей может только особа царских кровей. Юлий обращается к князю Киевскому за помощью, но получает отказ. Но наш жених не намерен отказываться от любви. Он похищает князя и насильно везет его к султану, на сватовство. И у богатырей тоже есть дело в восточной стороне. Все сойдутся на больших скачках, где победитель получит всё.',86,'Дарина Шмидт','Россия',2020,6,'/media/konyuliy.png','/media/konyuliy_slider.png'),
        ('Назад в будущее','Подросток Марти с помощью машины времени, сооружённой его другом-профессором доком Брауном, попадает из 80-х в далекие 50-е. Там он встречается со своими будущими родителями, ещё подростками, и другом-профессором, совсем молодым.',116,'Роберт Земекис','США',1985,12,'/media/backtothefuture.png','/media/backtothefuture_slider.png'),
        ('Джентльмены','Один ушлый американец ещё со студенческих лет приторговывал наркотиками, а теперь придумал схему нелегального обогащения с использованием поместий обедневшей английской аристократии и очень неплохо на этом разбогател. Другой пронырливый журналист приходит к Рэю, правой руке американца, и предлагает тому купить киносценарий, в котором подробно описаны преступления его босса при участии других представителей лондонского криминального мира - партнёра-еврея, китайской диаспоры, чернокожих спортсменов и даже русского олигарха.',113,'Гай Ричи','Великобритания',2019,18,'/media/gentelmen.png','/media/gentelmen_slider.png'),
        ('Отступники','Два лучших выпускника полицейской академии оказались по разные стороны баррикады: один из них – агент мафии в рядах правоохранительных органов, другой – «крот», внедрённый в мафию. Каждый считает своим долгом обнаружить и уничтожить противника, но постоянная жизнь в искажённых реалиях меняет внутренний мир героев.',151,'Мартин Скорсезе','США',2006,16,'/media/departed.png','/media/departed_slider.png'),
        ('Серебряные коньки', '1899 год, рождественский Петербург. Яркая праздничная жизнь бурлит на скованных льдом реках и каналах столицы. Накануне нового столетия судьба сводит тех, кому, казалось бы, не суждено было встретиться. Люди из совершенно разных миров, Матвей - сын фонарщика, его единственное богатство - доставшиеся по наследству посеребренные коньки; Алиса - дочь крупного сановника, грезящая о науке. У каждого - своя непростая история, но, однажды столкнувшись, они устремляются к мечте вместе.',136,'Михаил Локшин','Россия',2020,6,'/media/silver.png','/media/silver_slider.png'),
        ('Ведьмина служба доставки','Молодая ведьма Кики по достижении 13 лет должна прожить среди людей определённое время. Вместе с котом Дзидзи она отправляется в город, где знакомится с добрым пекарем, который помогает ей начать собственное дело - экстренную службу доставки. Новая работа знакомит Кики со множеством различных людей и предоставляет возможность обрести новых друзей и совершить массу всевозможных проделок.',103,'Хаяо Миядзаки','Япония',1989, 0,'/media/majo.png','/media/majo_slider.png'),
        ('Чёрная Вдова','Наташе Романофф предстоит лицом к лицу встретиться со своим прошлым. Чёрной Вдове придется вспомнить о том, что было в её жизни задолго до присоединения к команде Мстителей, и узнать об опасном заговоре, в который оказываются втянуты её старые знакомые - Елена, Алексей, также известный как Красный Страж, и Мелина.',136,'Кейт Шортланд','США',2021, 16,'/media/blackwidow.png','/media/blackwidow_slider.png'),
        ('Семейка Крудс: Новоселье','Такие харизматичные герои как члены клана Крудс просто не в силах усидеть на месте. Они смело идут навстречу самым головокружительным приключениям и готовы ответить непредсказуемостью и находчивостью на любой вызов судьбы.',95,'Джоэль Кроуфорд','США',2020, 6,'/media/cruds.png','/media/cruds_slider.png'),
        ('Обратная связь','31 декабря семеро друзей вновь собираются в загородном доме, чтобы вместе встретить Новый год, но предновогодний вечер полон сюрпризов.',97,'Алексей Нужный','России',2020, 16,'/media/feedback.png','/media/feedback_slider.png');

INSERT INTO movie_actors(movie_id, actor_id)
VALUES (1,1),
       (1,2),
       (2,3),
       (2,4),
       (3,5),
       (3,6),
       (4,7),
       (4,8),
       (5,9),
       (5,10),
       (6,11),
       (6,12),
       (7,13),
       (7,14),
       (8,15),
       (8,16),
       (9,17),
       (9,18),
       (10,19),
       (10,20),
       (11,21),
       (11,22),
       (12,23),
       (12,24),
       (13,25),
       (13,26),
       (14,27),
       (14,28);

INSERT INTO movie_genre (movie_id, genre_id)
VALUES (1,9),
       (1,10),
       (2,2),
       (2,3),
       (3,11),
       (3,4),
       (4,4),
       (4,9),
       (5,9),
       (5,11),
       (6,13),
       (6,3),
       (7,5),
       (7,2),
       (8,2),
       (8,12),
       (9,10),
       (9,9),
       (10,6),
       (10,7),
       (11,8),
       (11,13),
       (12,5),
       (12,4),
       (13,13),
       (13,11),
       (14,2);


INSERT INTO cinema_hall (Place_amount,Hall_params)
VALUES (15,'{"levels":[{"place":1,"row":1},{"place":2,"row":1},{"place":3,"row":1},{"place":4,"row":1},{"place":5,"row":1},{"place":6,"row":1},{"place":7,"row":1},{"place":8,"row":1},
  {"place":1,"row":2},{"place":2,"row":2},{"place":3,"row":2},{"place":4,"row":2},{"place":5,"row":2},{"place":6,"row":2},{"place":7,"row":2},{"place":8,"row":2},
  {"place":1,"row":3},{"place":2,"row":3},{"place":3,"row":3},{"place":4,"row":3},{"place":5,"row":3},{"place":6,"row":3},{"place":7,"row":3},{"place":8,"row":3},
  {"place":1,"row":4},{"place":2,"row":4},{"place":3,"row":4},{"place":4,"row":4},{"place":5,"row":4},{"place":6,"row":4},{"place":7,"row":4},{"place":8,"row":4},
  {"place":1,"row":5},{"place":2,"row":5},{"place":3,"row":5},{"place":4,"row":5},{"place":5,"row":5},{"place":6,"row":5},{"place":7,"row":5},{"place":8,"row":5}]}'),
       (10,'{"levels":[{"place":1,"row":1},{"place":2,"row":1},{"place":3,"row":1},{"place":4,"row":1},{"place":5,"row":1},{"place":6,"row":1},{"place":7,"row":1},{"place":8,"row":1},
         {"place":1,"row":2},{"place":2,"row":2},{"place":3,"row":2},{"place":4,"row":2},{"place":5,"row":2},{"place":6,"row":2},{"place":7,"row":2},{"place":8,"row":2},
         {"place":1,"row":3},{"place":2,"row":3},{"place":3,"row":3},{"place":4,"row":3},{"place":5,"row":3},{"place":6,"row":3},{"place":7,"row":3},{"place":8,"row":3},
         {"place":1,"row":4},{"place":2,"row":4},{"place":3,"row":4},{"place":4,"row":4},{"place":5,"row":4},{"place":6,"row":4},{"place":7,"row":4},{"place":8,"row":4},
         {"place":1,"row":5},{"place":2,"row":5},{"place":3,"row":5},{"place":4,"row":5},{"place":5,"row":5},{"place":6,"row":5},{"place":7,"row":5},{"place":8,"row":5}]}');

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
       (8,2,2,now() + interval '19 hours 2 minutes', 549),
       (9,3,1,now() + interval '20 hours 12 minutes', 359),
       (8,2,1,now() + interval '19 hours 25 minutes',349),
       (9,3,2,now() + interval '1 day 22 minutes',339),
       (9,4,1,now() + interval '23 hours 10 minutes',777),
       (10,2,2,now() + interval '7 days', 549),
       (10,3,1,now() + interval '3 days', 359),
       (10,2,1,now() + interval '15 days',349),
       (10,3,2,now() + interval '9 days',339),
       (10,4,1,now() + interval '1 day 15 hours',777),
       (10,4,2,now() + interval '10 hours',322),
       (10,1,2,now() + interval '5 hours',228),
       (10,2,2,now() + interval '7 hours',228),
       (10,3,2,now() + interval '9 hours',228),
       (10,1,1,now() + interval '11 days',666),
       (11,2,2,now() + interval '6 days', 549),
       (11,3,1,now() + interval '4 days', 359),
       (11,2,1,now() + interval '16 days',349),
       (11,3,2,now() + interval '10 days',339),
       (11,4,1,now() + interval '1 day 16 hours',777),
       (11,4,2,now() + interval '11 hours',322),
       (11,1,2,now() + interval '6 hours',228),
       (11,2,2,now() + interval '5 hours',228),
       (11,3,2,now() + interval '5 hours',228),
       (11,5,2,now() + interval '1 day',1337),
       (12,3,1,now() + interval '2 days 31 minutes', 359),
       (13,2,1,now() + interval '1 day 1 minute',349),
       (14,3,2,now() + interval '20 hours 38 minutes',339),
       (12,3,1,now() + interval '26 hours 12 minutes',777),
       (13,2,2,now() + interval '14 hours 10 minutes',322),
       (14,1,2,now() + interval '8 hours 18 minutes',228),
       (12,4,2,now() + interval '18 hours 10 minutes',228),
       (13,6,2,now() + interval '17 hours 18 minutes',228),
       (14,5,2,now() + interval '19 hours 12 minutes',1337),
       (6,5,2,now() + interval '2 days 29 minutes',337),
       (4,6,1,now() + interval '1 day 5 hours',358),
       (8,6,2,now() + interval '1 day 7 hours',767);




-- INSERT INTO schedule(Movie_ID, Cinema_ID, Hall_ID, Premiere_time, Cost)
-- VALUES (1,1,1, now() + interval '7 hour', 322),
--        (8,1,2, now() + interval '8 hour', 599),
--        (9,1,1, now() + interval '10 hours 30 minutes', 299),
--        (8,2,1, now() + interval '9 hour', 400),
--        (9,2,1, now() + interval '7 hour 30 minutes', 399),
--        (9,2,2, now() + interval '1 day 4 hours', 599),
--        (1,2,2, now() + interval '8 hour', 228),
--        (2,2,1, now() + interval '11 hours', 500),
--        (4,2,2, now() + interval '2 days', 300),
--        (4,3,2, now() + interval '10 hour', 399),
--        (2,3,2, now() + interval '3 days', 500),
--        (1,3,1, now() + interval '1 day 10 hour', 300),
--        (2,3,1, now() + interval '10 hour', 399),
--        (3,3,2, now() + interval '1 day', 599),
--        (6,4,1, now() + interval '8 hour 30 minutes',500),
--        (2,4,2, now() + interval '8 hours', 699),
--        (1,4,1, now() + interval '18 hour 30 minutes',500),
--        (2,4,2, now() + interval '35 hours', 699),
--        (8,4,1, now() + interval '24 hour 30 minutes',500),
--        (6,4,2, now() + interval '19 hours', 699);

