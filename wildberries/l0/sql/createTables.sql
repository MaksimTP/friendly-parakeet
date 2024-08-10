-- DROP TABLE "order";
-- DROP TABLE item;
-- DROP TABLE delivery;
-- DROP TABLE payment;

CREATE TABLE IF NOT EXISTS "order" (
    order_uid VARCHAR not NULL,
    track_number VARCHAR not NULL,
    entry VARCHAR not NULL,
    delivery_id BIGINT not NULL,
    payment_id BIGINT not NULL,
    items_ids BIGINT[],
    locale VARCHAR(2) not NULL,
    internal_signature VARCHAR,
    customer_id VARCHAR,
    delivery_service VARCHAR,
    shardkey VARCHAR,
    sm_id int,
    date_created timestamptz DEFAULT CURRENT_TIMESTAMP not NULL,
    oof_shard VARCHAR
);

alter table "order" ADD CONSTRAINT uk_order unique (order_uid, delivery_id, payment_id);
alter table "order" ADD CONSTRAINT ch_locale check ( locale in ('en', 'ru', 'kz', 'ua', 'kg', 'fr', 'de') );
alter table "order" ADD CONSTRAINT fk_delivery_id FOREIGN KEY (delivery_id) REFERENCES delivery(id);
alter table "order" ADD CONSTRAINT fk_payment_id FOREIGN KEY (payment_id) REFERENCES payment(id);
alter table "order" ADD CONSTRAINT fk_item_id FOREIGN KEY (items_ids) REFERENCES item(id);

CREATE TABLE IF NOT EXISTS item (
    id BIGINT not NULL,
    chrt_id BIGINT not NULL,
    track_number VARCHAR not NULL,
    price BIGINT,
    rid VARCHAR,
    sale INT,
    size VARCHAR(3),
    total_price BIGINT,
    nm_id BIGINT,
    brand VARCHAR(50),
    status VARCHAR(3)
);

alter table item ADD CONSTRAINT uk_item unique (id);


CREATE TABLE IF NOT EXISTS payment (
    id BIGINT not NULL,
    transaction VARCHAR not NULL,
    request_id VARCHAR,
    currency VARCHAR(3) not NULL,
    provider VARCHAR not NULL,
    amount BIGINT,
    payment_dt BIGINT,
    bank VARCHAR(30),
    delivery_cost INT,
    goods_total BIGINT,
    custom_fee INT
)

alter table payment ADD CONSTRAINT uk_payment unique (id);


CREATE TABLE IF NOT EXISTS delivery (
    id BIGINT not NULL,
    name VARCHAR not NULL,
    phone VARCHAR(20) not NULL,
    zip VARCHAR(12) not NULL,
    city VARCHAR(40) not NULL,
    address VARCHAR(60) not NULL,
    region VARCHAR(40) not NULL,
    email VARCHAR not NULL
);

alter table delivery ADD CONSTRAINT uk_delivery unique (id);

SELECT * FROM "order" as o
JOIN delivery as d on o.delivery_id = d.id
JOIN payment as p on o.payment_id = p.id
JOIN item as i on i.id in o.items_ids 