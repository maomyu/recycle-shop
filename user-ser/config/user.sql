DROP TABLE
IF EXISTS USER;

CREATE TABLE
IF NOT EXISTS USER (
	id VARCHAR (20) PRIMARY KEY,
	nickname VARCHAR (20) NOT NULL,
	passwd VARCHAR (20) NOT NULL,
	truename VARCHAR (20) DEFAULT '',
	sex INT (20) NOT NULL DEFAULT 3,
	email VARCHAR (30) NOT NULL UNIQUE,
	headerimage VARCHAR (100) DEFAULT '',
	school VARCHAR (20) DEFAULT '',
	signature VARCHAR (100) DEFAULT '设置个性签名，让更多的人认识你吧',
	birthday DATE NOT NULL,
	studentid VARCHAR (16) DEFAULT '',
	role INT (10) NOT NULL DEFAULT 0
) ENGINE = INNODB DEFAULT CHARSET = utf8;

INSERT INTO USER (
	id,
	nickname,
	passwd,
	email,
	birthday
)
VALUES
	(
		'wBnNA7YRFepIlBEuA7bP',
		'2',
		'2',
		'wuyazi@163.com',
		'2000-09-08'
	);