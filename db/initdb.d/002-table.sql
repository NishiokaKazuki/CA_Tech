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

CREATE TABLE characters
(
    id bigint unsigned AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    disabled boolean NOT NULL DEFAULT false,
    created_at       timestamp NOT NULL DEFAULT current_timestamp,
    updated_at       timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE characters_is_in_possessions
(
    user_id bigint unsigned NOT NULL,
    character_id bigint unsigned NOT NULL,
    quantity tinyint unsigned NOT NULL,
    FOREIGN KEY (user_id)
     REFERENCES app_users(id),
    FOREIGN KEY (character_id)
     REFERENCES characters(id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE gacha_groups
(
    id bigint unsigned AUTO_INCREMENT,
    percentage int unsigned NOT NULL,
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE characters_in_gacha_groups
(
    gacha_group_id bigint unsigned NOT NULL,
    character_id bigint unsigned NOT NULL,
    FOREIGN KEY (gacha_group_id)
     REFERENCES gacha_groups(id),
    FOREIGN KEY (character_id)
     REFERENCES characters(id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
