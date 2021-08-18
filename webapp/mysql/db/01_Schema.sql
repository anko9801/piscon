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
    latitude    DOUBLE              NOT NULL,
    longitude   DOUBLE              NOT NULL,
    geom        GEOMETRY            NOT NULL,
    rent        INTEGER             NOT NULL,
    door_height TINYINT UNSIGNED    NOT NULL,
    door_width  TINYINT UNSIGNED    NOT NULL,
    features    VARCHAR(64)         NOT NULL,
    popularity  INTEGER             NOT NULL,
    INDEX all_index (`door_height`, `door_width`, `rent`, `id`),
    INDEX union_index (`door_width`, `door_height`, `popularity` DESC, `id` ASC),
    INDEX rent_index (`rent`, `id`),
    INDEX rent_popularity (`rent`, `popularity` DESC, `id`),
    SPATIAL INDEX (geom)
) ENGINE=MyISAM;

CREATE TRIGGER insert_trigger BEFORE INSERT ON estate FOR EACH ROW UPDATE estate SET geom=POINT(latitude, longitude);