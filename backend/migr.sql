CREATE TABLE `Users` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL,
	`role` VARCHAR(255) NOT NULL,
	`balance` INT,
	`password` VARCHAR(255) NOT NULL,
	PRIMARY KEY (`id`)
);


CREATE TABLE `Transactions` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`type` INT NOT NULL,
	`userId` INT NOT NULL,
	`value` INT NOT NULL,			
	`refund` INT NOT NULL,
	`to` INT,
	PRIMARY KEY (`id`)
);

CREATE TABLE `Sessions` (
	`userId` INT NOT NULL,
	`status` VARCHAR(255) NOT NULL,
	`refreshToken` VARCHAR(255) NOT NULL,
	`timeClose` INT NOT NULL
);

ALTER TABLE `Transactions` ADD CONSTRAINT `Transactions_fk0` FOREIGN KEY (`userId`) REFERENCES `Users`(`id`);



INSERT INTO Users VALUES(1, "bob12312", "admin", 0, "qwerty123");
INSERT INTO Users VALUES(2, "Jack_1989", "admin", 0,"bobqewsax");
INSERT INTO Users VALUES(3, "viggoMarvin", "admin", 0,"asdcxvsa");
INSERT INTO Users VALUES(4, "grishaMalkevich", "user", 0,"qedsvrev");


INSERT INTO Transactions VALUES(1, 100, 4, 0, 0, 1);








