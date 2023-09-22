DROP TABLE IF EXISTS blogs

CREATE TABLE blogs (
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(45) NOT NULL,
  body longtext,
  created_at datetime NULL,
  updated_at datetime NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO blogs (title, body, created_at, updated_at)
VALUES
  ("title1", "tech content", "2023-09-19", "2023-09-19"),
  ("title2", "muscle content", "2023-09-19", "2023-09-19"),
  ("title3", "mac content", "2023-09-19", "2023-09-19"),
  ("title4", "network content", "2023-09-19", "2023-09-19")
