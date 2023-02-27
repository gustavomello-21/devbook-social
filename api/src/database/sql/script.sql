CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS  users;

CREATE TABLE users(
  id int auto_increment primary key,
  fullname varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  password varchar(100) not null,
  createdAt timestamp default current_timestamp()
);

CREATE TABLE followers(
  user_id int not null,
  FOREIGN KEY (user_id),
  REFERENCES users(id)
  ON DELETE CASCADE,

  follower_id int not null,
  FOREIGN KEY (user_id),
  REFERENCES users(id)
  ON DELETE CASCADE,

  PRIMARY KEY(user_id, follower_id)
)
