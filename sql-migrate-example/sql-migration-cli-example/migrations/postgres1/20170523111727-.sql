
-- +migrate Up
Insert into people values(1);
Insert into people values(2);
Insert into people values(3);

-- +migrate Down
Delete from people where id=1;
Delete from people where id=2;
Delete from people where id=3;
