
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE orders (
    id SERIAL,
    status text,
    product_id integer,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    PRIMARY KEY(id)
);
CREATE INDEX products_product_id ON orders (product_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX products_product_id;
DROP TABLE orders;
