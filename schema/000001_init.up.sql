CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   username VARCHAR(255) UNIQUE NOT NULL,
   password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE todo_lists (
   id SERIAL PRIMARY KEY,
   title VARCHAR(255) NOT NULL,
   description VARCHAR(255)
);

CREATE TABLE users_lists (
    id SERIAL PRIMARY KEY,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done boolean not null default false
);

CREATE TABLE lists_items (
    id SERIAL PRIMARY KEY,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);

