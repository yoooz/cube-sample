set global local_infile = 1;

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'root';

CREATE DATABASE IF NOT EXISTS sample;

CREATE TABLE IF NOT EXISTS sample.alphabet(
    `id` INT(10) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS sample.task_status_history(
    `date` DATE NOT NULL,
    `user` VARCHAR(64) NOT NULL,
    `task` VARCHAR(64) NOT NULL,
    `status_type` INT(10) NOT NULL,
    `time` DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS sample.user_company(
    `id` INT(10) NOT NULL AUTO_INCREMENT,
    `paying` CHAR(1) NOT NULL,
    `city` VARCHAR(64) NOT NULL,
    `company_name` VARCHAR(64) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS sample.sample(
    `a` INT(10),
    `b` INT(10),
    `c` INT(10),
    `d` INT(10),
    `e` INT(10),
    `f` INT(10),
    `g` INT(10),
    `h` INT(10),
    `i` INT(10),
    `j` INT(10)
);
