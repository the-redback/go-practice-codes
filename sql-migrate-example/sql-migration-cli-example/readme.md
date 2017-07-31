### drivers for sql-migrate
> go get github.com/rubenv/sql-migrate/...



### development section
> sql-migrate new -env=development

a sql file will be generated in migrations/postgres1
add sql to that generated file, example

```
-- +migrate Up
Create table if not exists people(id int);

-- +migrate Down
Drop Table people;
```


### create another file by
> sql-migrate new -env=development

### add folowing sql to new file, example
```
-- +migrate Up
Insert into people values(1);
Insert into people values(2);
Insert into people values(3);

-- +migrate Down
Delete from people where id=1;
Delete from people where id=2;
Delete from people where id=3;
```


### See status 
> sql-migrate status -env=development

### execute sql file UP. 
> sql-migrate up -env=development

this will execute all the sqls of all the files

### Sql execution Down. 
> sql-migrate down -env=development

Down executions will execute down parts of one file each time. The last generated file will be executed 1st.

> sql-migrate status -env=development


### Similar with -env=production
