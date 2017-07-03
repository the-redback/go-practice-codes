
-- +migrate Up
Create table if not exists people(id int);

-- +migrate Down
Drop Table people;
