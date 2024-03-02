CREATE TABLE trade_fulfillments (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    trade_id uuid,
    fulfilling_server uuid,
    fulfilled_quantity int,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_trade_id ON trade_fulfillments (trade_id);
CREATE INDEX idx_fulfilling_server ON trade_fulfillments (fulfilling_server);
