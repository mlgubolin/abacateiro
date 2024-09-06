-- This is a sample migration.

create table if not exists users(
  id serial primary key,
  name varchar not null,
  email varchar not null,
  password varchar not null, --- DO NOT DO THIS IN REAL LIFE ---
  document varchar
);

---- create above / drop below ----

-- drop table people;