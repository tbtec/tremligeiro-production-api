create table if not exists tremligeiro_production.order
(
    order_id uuid not null primary key,
    status varchar(200) not null,
    created_at timestamp
);
