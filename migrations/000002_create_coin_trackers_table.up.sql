CREATE TABLE coin_trackers
(
    id INTEGER PRIMARY KEY NOT NULL,
    user_id INTEGER NOT NULL,
    name varchar(200) NOT NULL,
    rank varchar(200) NOT NULL,
    symbol varchar(200) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);