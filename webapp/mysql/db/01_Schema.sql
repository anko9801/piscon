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
    geom        POINT AS (POINT(latitude, longitude)) STORED NOT NULL,
    rent        INTEGER             NOT NULL,
    door_height TINYINT UNSIGNED    NOT NULL,
    door_width  TINYINT UNSIGNED    NOT NULL,
    features    VARCHAR(64)         NOT NULL,
    popularity  INTEGER             NOT NULL,
    INDEX all_door_index (`door_height`, `door_width`, `rent`),
    INDEX door_width (`door_width`, `rent`),
    INDEX door_height (`door_height`, `rent`),
    INDEX rent_index (`rent`),
    INDEX rent_popularity (`rent`, `popularity` DESC, `id`),
    INDEX popularity_id (`popularity` DESC, `id`),
    INDEX latitude_longitude (`latitude`, `longitude`)
    SPATIAL INDEX (geom)
) ENGINE=MyISAM;