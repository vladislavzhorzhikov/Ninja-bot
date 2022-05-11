create table if not exists randoms (
    id integer primary key,
    user_name varchar (64) not null,
    rand_count integer not null);

INSERT INTO randoms VALUES (0, 'test', 0);

