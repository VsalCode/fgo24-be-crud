-- Active: 1750752675497@@127.0.0.1@5432@dbuserscrud
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(100) UNIQUE,
    password VARCHAR(255) NOT NULL
);

INSERT INTO users (username, email, phone, password)
VALUES 
('budistyle', 'budi@mail.co', 08129356789, md5('maxstyle123'));

DROP TABLE users;

INSERT INTO users (username, email, phone, password)
VALUES 
('Yassirmaxstyle', 'yassirmax@mail.co', 08123456789, md5('maxstyle123')),
('bruno', 'bruno@mail.co', 08123456788, md5('bruno123')),
('james', 'james@mail.co', 08123456787, md5('james123')),
('john', 'john@mail.co', 08123456786, md5('john123')),
('jane', 'jane@mail.co', 08123456785, md5('jane123')),
('doe', 'doe@mail.co', 08123456784, md5('doe123'));

SELECT * FROM users;