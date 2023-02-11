DROP SCHEMA IF EXISTS suga;
CREATE SCHEMA suga;
USE suga;

create table `users`(
    `id` serial not null primary key,
    `name` varchar(128),
    `data` json
);
-- insert into users values (1, 'Yamada', JSON_OBJECT(1,'val1', 2, 'val2'));