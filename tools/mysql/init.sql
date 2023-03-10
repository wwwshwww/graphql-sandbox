CREATE DATABASE IF NOT EXISTS test_db;
USE test_db;

CREATE TABLE IF NOT EXISTS users(
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    name VARCHAR (127) NOT NULL UNIQUE,
    password VARCHAR (127) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS links(
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    title VARCHAR (255) ,
    address VARCHAR (255) ,
    user_id INT ,
    FOREIGN KEY (user_id) REFERENCES users(id) ,
    PRIMARY KEY (id)
);