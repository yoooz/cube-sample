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
