-- CREATE SCHEMA IF NOT EXISTS categories;

CREATE TABLE categories (
  id   varchar(36)  NOT NULL PRIMARY KEY,
  name text    NOT NULL,
  description  text
);

CREATE TABLE courses (
  id   varchar(36)  NOT NULL PRIMARY KEY,
  category_id   varchar(36)  NOT NULL,
  name text    NOT NULL,
  description  text,
  price  DOUBLE PRECISION  NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);