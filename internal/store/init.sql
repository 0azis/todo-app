CREATE TABLE users (
    userid serial not null,
    email varchar(50),
    password varchar(60),
    primary key (userid)
)

CREATE TABLE notes (
    noteid varchar(50),
    userid integer,
    text varchar(50),
    date timestamp not null default CURRENT_TIMESTAMP
)