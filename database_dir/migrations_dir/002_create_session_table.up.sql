create table if not exists
    sessions
    (
        sessionToken varchar primary key,
        email varchar,
        expiry timestamp
    );