
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE orders
    ADD CONSTRAINT orders_product_id_fkey FOREIGN KEY (product_id)
    REFERENCES products (id)
    ON UPDATE RESTRICT
    ON DELETE RESTRICT
;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP CONSTRAINT orders_product_id_fkey
