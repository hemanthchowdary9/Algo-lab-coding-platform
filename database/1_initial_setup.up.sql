CREATE TABLE IF NOT EXISTS challenges (
    id INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    title VARCHAR(50),
    difficulty VARCHAR(10),
    category varchar(20),
    challenge JSONB
    );

CREATE TABLE IF NOT EXISTS coder (
    username VARCHAR(20) PRIMARY KEY,
    email VARCHAR(50),
    password VARCHAR(50),
    role VARCHAR(20) default 'NORMAL'
    );


CREATE TABLE IF NOT EXISTS user_submission (
    username VARCHAR(20) PRIMARY KEY,
    submissions JSONB
    );