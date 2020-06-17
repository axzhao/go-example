CREATE TABLE Article(
    id serial,
    title varchar(100) NOT NULL,
    createddate timestamptz DEFAULT NULL,
    body text,
    userid int DEFAULT NULL,
    PRIMARY KEY(id)
);

INSERT INTO Article(title) VALUES('first article');

