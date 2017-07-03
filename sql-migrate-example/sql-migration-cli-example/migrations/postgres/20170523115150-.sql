
-- +migrate Up
Create table if not exists student(id int);
-- +migrate Down
Drop table student;
