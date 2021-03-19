CREATE DATABASE IF NOT EXISTS ca_tech;
use ca_tech;

CREATE TABLE app_users
(
    id bigint unsigned AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE,
    disabled boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE tokens
(
    id bigint unsigned AUTO_INCREMENT,
    user_id bigint unsigned NOT NULL UNIQUE,
    token   text NOT NULL,
    created_at       timestamp NOT NULL DEFAULT current_timestamp,
    updated_at       timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (user_id)
     REFERENCES app_users(id),
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
