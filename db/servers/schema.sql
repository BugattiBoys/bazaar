CREATE TABLE servers
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    ip_address cidr,
    api_key uuid,
    created_at TIMESTAMP,
    last_ping TIMESTAMP
);
