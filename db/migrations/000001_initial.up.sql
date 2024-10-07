CREATE TABLE IF NOT EXISTS todo_items
(
    id          uuid PRIMARY KEY,
    description VARCHAR(256) not null,
    due_date    time not null,
    file_id     varchar(256) not null
);
