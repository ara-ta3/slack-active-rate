
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE actives (
    id int(11) NOT NULL AUTO_INCREMENT,
    presence varchar(11) NOT NULL,
    user varchar(11) NOT NULL,
    hour timestamp NOT NULL,
    PRIMARY KEY id
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE actives;
