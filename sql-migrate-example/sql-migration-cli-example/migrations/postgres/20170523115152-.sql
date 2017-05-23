
-- +migrate Up
Insert into student values(11);
Insert into student values(22);
Insert into student values(33);
-- +migrate Down
Delete from student where id=11;
Delete from student where id=22;
Delete from student where id=33;
