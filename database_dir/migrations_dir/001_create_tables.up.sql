create type roleType as enum ('admin','sub-admin','user');
create table if not exists
    person
    (
        id serial primary key ,
        name varchar,
        email varchar,
        password varchar,
        createdBy int
    );

create table if not exists
    location
    (
        sNum serial primary key,
        pId int,
        address point,
        foreign key(pId) references person(id)
    );

create table if not exists
    roles
    (
        sNum serial primary key,
        pId int,
        role roleType,
        foreign key(pId) references person(id)
    );

create table if not exists
    restaurant
    (
        id serial primary key,
        name varchar,
        address point,
        createdBy int
    );

create table if not exists
    dishes
    (
        sNum serial primary key,
        rId int,
        name varchar,
        price int,
        createdBy int,
        foreign key(rId) references restaurant(id)
    );

