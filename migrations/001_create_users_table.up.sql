-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS users (
                                     id INT AUTO_INCREMENT PRIMARY KEY,
                                     dni CHAR(8),
    name VARCHAR(30),
    last_name VARCHAR(30),
    second_last_name VARCHAR(30),
    birthdate DATE,
    phone CHAR(9),
    email VARCHAR(50)
    );

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS users;
