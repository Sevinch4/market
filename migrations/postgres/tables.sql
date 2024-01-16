create type user_type_enum as enum ('admin', 'customer');

create table users(
    id uuid primary key ,
    full_name text ,
    phone varchar(30) unique not null ,
    password varchar(30) not null ,
    cash integer default 0,
    user_type user_type_enum not null
);

create table baskets(
    id uuid primary key ,
    customer_id uuid references users(id) not null ,
    total_sum integer default 0
);