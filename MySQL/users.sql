CREATE DATABASE IF NOT EXISTS person;

CREATE TABLE IF NOT EXIST `person`.`users` (
  `id` INT NOT NULL,
  `username` VARCHAR(45) NULL,
  `date_of_birth` DATE NULL,
  PRIMARY KEY (`id`))
COMMENT = 'Users in a person database';


INSERT INTO `person`.`users` (`id`, `username`, `date_of_birth`) VALUES ('1', 'mohan', '1990-05-08');
INSERT INTO `person`.`users` (`id`, `username`, `date_of_birth`) VALUES ('2', 'namrata', '1989-04-24');
INSERT INTO `person`.`users` (`id`, `username`, `date_of_birth`) VALUES ('3', 'swapnil', '1985-02-28');
INSERT INTO `person`.`users` (`id`, `username`, `date_of_birth`) VALUES ('4', 'Gita', '1992-08-14');