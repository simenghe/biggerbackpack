DROP DATABASE backpack;
CREATE DATABASE backpack;

USE backpack;

CREATE TABLE users(
id INT AUTO_INCREMENT,
   team_name VARCHAR(100),
   ranking INT,
   url VARCHAR(100),
   fetch_date DATETIME,
   PRIMARY KEY(id)
);
