CREATE TABLE trades (
                        id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
                        status varchar,
                        type varchar,
                        initiating_server uuid,
                        item varchar,
                        trade_quantity int,
                        price int,
                        fulfilled_quantity int,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        fulfilled_at TIMESTAMP,
                        cancelled_at TIMESTAMP,
                        updated_at TIMESTAMP
);

CREATE INDEX idx_initiating_server ON trades (initiating_server);
CREATE INDEX idx_type ON trades (type);
CREATE INDEX idx_item ON trades (item);

CREATE TABLE trade_fulfillments (
                                    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
                                    trade_id uuid,
                                    fulfilling_server uuid,
                                    fulfilled_quantity int,
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_trade_id ON trade_fulfillments (trade_id);
CREATE INDEX idx_fulfilling_server ON trade_fulfillments (fulfilling_server);
