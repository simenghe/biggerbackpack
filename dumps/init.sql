DROP DATABASE IF EXISTS backpack;
CREATE DATABASE backpack;
USE backpack;
CREATE TABLE csgoteams
(
  id INT
  AUTO_INCREMENT,
  team_name VARCHAR
  (100),
  ranking INT,
  url VARCHAR
  (100),
  fetch_date DATE,
  PRIMARY KEY
  (id)
);
  -- Astralis Example
  INSERT INTO csgoteams
    (team_name, ranking, url, fetch_date)
  values
    ('Astralis', 1, 'hltv', now());