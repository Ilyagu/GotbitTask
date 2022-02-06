create extension if not exists citext;

create table if not exists tasks (
  id serial unique not null primary key,
  creator citext collate "ucs_basic" not null,
  title citext not null,
  description text default '',
  executor citext,
  created   timestamp with time zone default now(),
  completed boolean default false
);