DROP DATABASE if exists clear;
CREATE DATABASE clear;
USE clear;

DROP USER if exists user;
CREATE USER user@"%" IDENTIFIED BY 'pass';
CREATE USER user@localhost IDENTIFIED BY 'pass';
GRANT ALL ON clear.* TO user@"%";
GRANT ALL ON clear.* TO user@localhost;

DROP TABLE if exists tasks;

CREATE TABLE tasks (
  id INT NOT NULL,
  name VARCHAR(256) NOT NULL,
  status BIT(1) NOT NULL,
  `order` INT NOT NULL,
  timestamp TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);

INSERT INTO tasks (id, name, status, `order`, timestamp) values (1, '1. Flick down to add a new task', 1, 1, now());
INSERT INTO tasks (id, name, status, `order`, timestamp) values (2, '2. Flick right to complete the task', 1, 2, now());
INSERT INTO tasks (id, name, status, `order`, timestamp) values (3, '3. Flick left to delete a task', 1, 3, now());
INSERT INTO tasks (id, name, status, `order`, timestamp) values (4, '4. Flick this task to right to complete', 1, 4, now());
INSERT INTO tasks (id, name, status, `order`, timestamp) values (5, '5. Flick up to delete all tasks', 1, 5, now());

DROP TABLE if exists users;

CREATE TABLE users (
  email VARCHAR(256) NOT NULL,
  `password` VARCHAR(256) NOT NULL,
  timestamp TIMESTAMP NOT NULL,
  PRIMARY KEY(email)
);

INSERT INTO users (email, `password`, timestamp) values ('hoge@example.com', 'passwd', now());
