CREATE TABLE IF NOT EXISTS codes (
    id           SERIAL PRIMARY KEY,
    user_id      uuid NOT NULL,
    client_id    uuid NOT NULL,
    code         CHARACTER(31) NOT NULL,
    scope        VARCHAR(255) NOT NULL,
    state        CHARACTER(31) NOT NULL,
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_codes_tenant_user_id ON codes(client_id, code);


CREATE TABLE IF NOT EXISTS providers (
    tenant_id       uuid NOT NULL,
    type            CHARACTER(31) NOT NULL,
    client_id       uuid NOT NULL,
    client_secret   uuid NOT NULL,
    create_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
ALTER TABLE providers ADD PRIMARY KEY (tenant_id, type);
