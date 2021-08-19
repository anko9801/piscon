DROP DATABASE IF EXISTS isuumo;
CREATE DATABASE isuumo;

DROP TABLE IF EXISTS isuumo.chair;

CREATE TABLE isuumo.chair
(
    id          INTEGER          NOT NULL PRIMARY KEY,
    name        VARCHAR(32)      NOT NULL,
    description VARCHAR(100)     NOT NULL,
    thumbnail   VARCHAR(100)     NOT NULL,
    price       INTEGER          NOT NULL,
    height      TINYINT UNSIGNED NOT NULL,
    width       TINYINT UNSIGNED NOT NULL,
    depth       TINYINT UNSIGNED NOT NULL,
    color       VARCHAR(4)       NOT NULL,
    features    VARCHAR(32)      NOT NULL,
    kind        VARCHAR(16)      NOT NULL,
    popularity  INTEGER          NOT NULL,
    stock       TINYINT UNSIGNED NOT NULL,
    INDEX all_index (`price`, `height`, `width`, `depth`, `kind`, `color`, `features`, `stock`, `popularity` DESC, `id`),
    INDEX stock_price (`stock`, `price`, `id`)
) ENGINE=InnoDB;
