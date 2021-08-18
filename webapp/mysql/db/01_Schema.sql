DROP DATABASE IF EXISTS isuumo;
CREATE DATABASE isuumo;

DROP TABLE IF EXISTS isuumo.estate;

CREATE TABLE isuumo.estate
(
    id          INTEGER             NOT NULL PRIMARY KEY,
    name        VARCHAR(16)         NOT NULL,
    description VARCHAR(100)        NOT NULL,
    thumbnail   VARCHAR(100)        NOT NULL,
    address     VARCHAR(64)         NOT NULL,
    latitude    DOUBLE PRECISION    NOT NULL,
    longitude   DOUBLE PRECISION    NOT NULL,
    rent        INTEGER             NOT NULL,
    door_height TINYINT UNSIGNED    NOT NULL,
    door_width  TINYINT UNSIGNED    NOT NULL,
    features    VARCHAR(64)         NOT NULL,
    popularity  INTEGER             NOT NULL,
    INDEX all_index (`door_height`, `door_width`, `rent`, `id`),
    INDEX nazotte_index (`latitude`, `longitude`, `popularity` DESC, `id` ASC),
    INDEX union_index (`door_width`, `door_height`, `popularity` DESC, `id` ASC),
    INDEX rent_index (`rent`, `id`),
    INDEX rent_popularity (`rent`, `popularity` DESC, `id`)
) ENGINE=MyISAM;