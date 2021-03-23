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

INSERT INTO characters(id, name) VALUES
    ( 1, '金木 研'),
    ( 2, '佐々木 琲世'),
    ( 3, '霧嶋 董香'),
    ( 4, '四方 蓮示'),
    ( 5, '西尾 錦'),
    ( 6, '笛口 雛実'),
    ( 7, '月山 習'),
    ( 8, '霧嶋 絢都'),
    ( 9, 'エト'),
    (10, '有馬 貴将'),
    (11, '旧多 二福'),
    (12, '亜門 鋼太朗');

INSERT INTO characters_is_in_possessions(user_id, character_id, quantity) VALUES
    ( 1,  1, 1),
    ( 1,  2, 1),
    ( 1,  3, 1),
    ( 1,  4, 1),
    ( 1,  5, 1),
    ( 1,  6, 1),
    ( 1,  7, 1),
    ( 1,  8, 1),
    ( 1,  9, 1),
    ( 1, 10, 1),
    ( 1, 11, 1),
    ( 1, 12, 1),
    ( 2,  1, 2),
    ( 2,  2, 2),
    ( 2,  3, 2),
    ( 2,  4, 2),
    ( 2,  5, 2),
    ( 2,  6, 2),
    ( 2,  7, 2),
    ( 2,  8, 2),
    ( 2,  9, 2),
    ( 2, 10, 2),
    ( 2, 11, 2),
    ( 2, 12, 2),
    ( 3,  1, 1),
    ( 3,  2, 2),
    ( 3,  3, 3),
    ( 3,  4, 4),
    ( 3,  5, 5),
    ( 3,  6, 6),
    ( 3,  7, 7),
    ( 3,  8, 8),
    ( 3,  9, 9),
    ( 3, 10, 9),
    ( 3, 11, 9),
    ( 3, 12, 9);

INSERT INTO gacha_groups(id, probability) VALUES
    ( 1, 1),
    ( 2, 4),
    ( 3, 95);

INSERT INTO characters_in_gacha_groups(gacha_group_id, character_id) VALUES
    ( 1,  1),
    ( 1,  2),
    ( 1,  3),
    ( 3,  4),
    ( 3,  5),
    ( 3,  6),
    ( 3,  7),
    ( 3,  8),
    ( 3,  9),
    ( 2, 10),
    ( 2, 11),
    ( 2, 12);