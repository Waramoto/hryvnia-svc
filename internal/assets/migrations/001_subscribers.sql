-- +migrate Up

create table if not exists subscribers
(
    id        bigserial primary key,
    email     text unique not null,
    last_send timestamp   not null,
    status    int8        not null
);

-- +migrate Down

drop table subscribers;
