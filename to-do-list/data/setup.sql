drop table if exists tasks;
drop table if exists sessions;
drop table if exists ingredients;
drop table if exists receipts;
drop table if exists users;

create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  username   varchar(255) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  avatar     varchar(255),
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

create table tasks(
    id          serial primary key,
    user_id     integer references users(id),
    title       text,
    deadline    date,
    description text,
    isImportant boolean default false,
    isFinished  boolean default false,
    created_at  timestamp
);
create table receipts(
    id          serial primary key,
    user_id     integer references users(id),
    name        varchar(255),
    photo       varchar(255),
    duration    integer,
    instruction text,
    created_at timestamp
);
create table ingredients(
    id          serial primary key,
    name        varchar(255),
    receipt_id  integer references receipts(id),
    amount      integer,
    unit        varchar(10)
);
alter table users alter "avatar" set default 'private/avatar/default-avatar.jpg';
