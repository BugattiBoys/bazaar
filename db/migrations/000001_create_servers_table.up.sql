CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE servers
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    ip_address cidr,
    api_key uuid DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_ping TIMESTAMP
);
-- CREATE UNIQUE INDEX idx_api_keys ON servers api_key;