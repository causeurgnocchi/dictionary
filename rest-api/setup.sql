drop database if exists dictionary;
create database dictionary;
use dictionary;

create table vocabularies (
    id int auto_increment,
    writing text not null,
    reading text not null,
    meaning text not null,
    primary key (id)
);

create table groups (
    id int auto_increment,
    name text not null,
    primary key (id)
);