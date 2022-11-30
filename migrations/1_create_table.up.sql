create table "users"(
    "id" serial not null primary key,
    "first_name" varchar(50) not null,
    "last_name" varchar(50) not null,
    "username" varchar(50) unique,
    "phone_number" varchar(30) unique,
    "email" varchar(100) not null unique,
    "password" varchar not null,
    "image_url" varchar,
    "created_at" timestamp default current_timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp 
);

create table "notes" (
    id serial not null primary key,
    user_id int not null references users(id),
    title varchar(100) not null,
    description varchar not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);