create type transaction as enum ('STORAGE', 'SALE');

create table if not exists robots
(
    id                        bigserial primary key,
    type                      text unique,
    count_of_robots           bigint,
    manufacturing_cost        bigint,
    storage_cost              bigint,
    selling_price             bigint,
    manufacturing_rate        bigint,
    last_update_number_robots timestamp with time zone,
    last_update_storage_cost  timestamp with time zone
);

create table if not exists transaction_histories
(
    id                 bigserial primary key,
    transaction        transaction,
    robot_id           bigint,
    count_robots       bigint,
    amount             bigint,
    manufacturing_cost bigint,
    time               timestamp with time zone
);