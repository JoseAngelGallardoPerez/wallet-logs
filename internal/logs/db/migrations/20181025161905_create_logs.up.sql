CREATE TABLE logs (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  subject varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  logged_at timestamp NULL DEFAULT NULL,
  data_title varchar(255) NOT NULL,
  data_fields JSON,
  PRIMARY KEY (id)
) DEFAULT CHARSET=utf8;
