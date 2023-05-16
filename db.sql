CREATE DATABASE pern;

CREATE TABLE profile(
    id SERIAL PRIMARY KEY, -- AUTOMAITALLY +1   
    name VARCHAR(150),  --VARVHAR = TEXT IWTH CONSTRAINT
    username VARCHAR(150),
    email VARCHAR(150),
    PASSWORD TEXT, --TEXT WILL BE HASHED
    rgb INT[] DEFAULT '{120, 45, 245}',
    avatar INT DEFAULT 0,
    goalTime INT DEFAULT 0,
    UNIQUE (username,email)
);
--getting data
SELECT * FROM profile;
SELECT * FROM profile WHERE id = 1;
--updating data
UPDATE profile SET name = 'hiroshi' WHERE id = 2;
UPDATE goaltracker SET jan[2] = 3 WHERE user_id = 1;
UPDATE profile SET goaltime = 2 WHERE user_id = 1;

UPDATE profile SET rgb = '{127, 255, 212}' WHERE id=1;


--deleting data
DELETE FROM profile WHERE id = 3;
DELETE FROM habit WHERE id = 3;
-----------
CREATE TABLE post(
    id SERIAL PRIMARY KEY,
    name VARCHAR(250),
    content TEXT,
    user_id INT , --NOT JUST NORMAL INTEGER, ITS A FOREIGN KEY
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES profile(id)
    UNIQUE(user_id)
);

INSERT INTO post (name, content, user_id) VALUES ('coding', 'learning sql', 1); --insert user_id 
INSERT INTO post (name, content, user_id) VALUES ('basketball', 'playing basketball', 1);

SELECT * FROM profile JOIN post ON post.user_id = profile.id; --getting all post/users with same user id
SELECT * FROM post WHERE user_id = 1; --getting all post with same user id
--- goal tracker
INSERT INTO goaltracker (user_id) VALUES (3);

CREATE TABLE docsdata(
    id SERIAL PRIMARY KEY,
    docname VARCHAR(100),
    dacdata TEXT DEFAULT 'example',
    user_id INT,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES profile(id)
);

CREATE TABLE habit(
    id SERIAL PRIMARY KEY,
    habitname VARCHAR(50),
    habitdata TEXT DEFAULT 'description',
    user_id INT,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES profile(id)
);
DELETE FROM habit WHERE habitdata = $1 AND user_id = $2;

DELETE FROM docsdata WHERE docname = 'golang' AND user_id = 1;

ALTER TABLE profile ADD avatar INT DEFAULT 0;
UPDATE profile SET avatar = $1 WHERE user_id = $2;

