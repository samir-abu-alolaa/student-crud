CREATE DATABASE IF NOT EXISTS myproject;
USE myproject;

CREATE TABLE IF NOT EXISTS students (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(100) UNIQUE
);
INSERT INTO students (name, age, email) VALUES ('Alice', 22, 'alice@example.com');
INSERT INTO students (name, age, email) VALUES ('Bob', 25, 'bob@example.com');
