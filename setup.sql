drop database if exists dictionary;
create database dictionary;
use dictionary;

create table vocabularies (
    id int auto_increment,
    writing text not null,
    reading text not null,
    primary key (id)
);

create table meanings (
    id int auto_increment,
    vocabulary_id int,
    text text not null,
    primary key (id),
    foreign key (vocabulary_id) references vocabularies(id)
);
