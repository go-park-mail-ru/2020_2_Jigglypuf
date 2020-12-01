SET timezone ='+3';

/* profile table */

CREATE TABLE if not exists profile
(
    user_id integer NOT NULL PRIMARY KEY,
    ProfileName VARCHAR(32),
    ProfileSurname VARCHAR(32),
    AvatarPath VARCHAR(64)
);