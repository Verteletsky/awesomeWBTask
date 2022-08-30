CREATE TABLE Orders
(
    id                 serial primary key,
    order_uid          varchar,
    track_number       varchar,
    entry              varchar,
    locale             varchar,
    internal_signature varchar,
    customer_id        varchar not null,
    delivery_service   varchar not null,
    shardkey           varchar,
    sm_id              integer,
    date_created       date,
    oof_shard          varchar
);

CREATE TABLE Delivery
(
    orderId integer references Orders (id),
    name    varchar               NOT NULL,
    phone   varchar               not null,
    zip     varchar               not null,
    city    varchar               not null,
    address varchar               not null,
    region  varchar               not null,
    email   character varying(30) not null
);

CREATE TABLE Payment
(
    orderId       integer references Orders (id),
    transaction   varchar not null,
    request_id    varchar,
    currency      varchar(3),
    provider      varchar not null,
    amount        integer not null,
    payment_dt    integer not null,
    bank          varchar not null,
    delivery_cost integer not null,
    goods_total   integer not null,
    custom_fee    integer not null
);

CREATE table Items
(
    orderId      integer references Orders (id),
    chrt_id      integer not null,
    track_number varchar not null,
    price        integer not null,
    rid          varchar not null,
    name         varchar not null,
    sale         integer not null,
    size         varchar not null,
    total_price  real    not null,
    nm_id        integer not null,
    brand        varchar not null,
    status       integer not null
);

