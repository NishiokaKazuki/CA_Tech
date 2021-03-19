CREATE DATABASE IF NOT EXISTS ca_tech;
use ca_tech;

INSERT INTO app_users(id, name) VALUES
    (1, 'test_01'),
    (2, 'test_02'),
    (3, 'test_03'),
    (4, 'test_04'),
    (5, 'test_05'),
    (6, 'test_06');

INSERT INTO tokens(id, user_id, token) VALUES
    (1, 1, 'AAAAA'),
    (2, 2, 'BBBBB'),
    (3, 3, 'CCCCC'),
    (4, 4, 'DDDDD'),
    (5, 5, 'RRRRR'),
    (6, 6, 'FFFFF');
