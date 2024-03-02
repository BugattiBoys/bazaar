CREATE TABLE trades (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    status varchar,
    type varchar,
    -- the ID of the server that opened the trade
    initiating_server uuid,
    -- the item that is being sold / bought
    item varchar,
    -- the total number of items to fulfill the trade
    trade_quantity int,
    -- the price for buying / selling the item
    price int,
    -- the number of units for this trade already fulfilled
    fulfilled_quantity int,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fulfilled_at TIMESTAMP,
    cancelled_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX idx_initiating_server ON trades (initiating_server);
CREATE INDEX idx_type ON trades (type);
CREATE INDEX idx_item ON trades (item);
