
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE products (
    id SERIAL,
    code text,
    price integer,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE products;
