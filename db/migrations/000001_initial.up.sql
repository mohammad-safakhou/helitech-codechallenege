CREATE TABLE IF NOT EXISTS todo_items
(
    id          uuid PRIMARY KEY,
    description VARCHAR(256),
    due_date    time,
    file_id     varchar(256)
);
