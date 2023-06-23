
create database quick_auth;


CREATE TABLE IF NOT EXISTS users (
    id           uuid PRIMARY KEY,
    user_pool_id uuid NOT NULL,
    username     VARCHAR(127) NOT NULL,
    password     VARCHAR(127),
    display_name VARCHAR(127),
    email        VARCHAR(127),
    phone        VARCHAR(20),
    is_disabled  BOOLEAN,
    create_time  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE DEFAULT now()
);
CREATE UNIQUE INDEX idx_users_user_pool_username ON users(user_pool_id, username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);


CREATE TABLE IF NOT EXISTS user_pools (
    id          uuid PRIMARY KEY,
    name        VARCHAR(127),
    describe    VARCHAR(127),
    create_time TIMESTAMP WITH TIME ZONE DEFAULT now(),
    update_time TIMESTAMP WITH TIME ZONE DEFAULT now()
);


CREATE TABLE IF NOT EXISTS tenants (
    id           uuid PRIMARY KEY,
    client_id       uuid NOT NULL,
    user_pool_id uuid NOT NULL,
    type         INTEGER,
    name         VARCHAR(127),
    host         VARCHAR(127) NOT NULL,
    company      VARCHAR(127) NOT NULL,
    describe     VARCHAR(127),
    create_time  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE DEFAULT now()
);
CREATE UNIQUE INDEX idx_tenants_host ON tenants(host);
CREATE INDEX idx_tenants_client_user_pool_id ON tenants(client_id, user_pool_id);


CREATE TABLE IF NOT EXISTS clients (
    id             uuid PRIMARY KEY,
    name           VARCHAR(127),
    describe       VARCHAR(127),
    grant_type     TEXT,
    create_time    TIMESTAMP WITH TIME ZONE DEFAULT now(),
    update_time    TIMESTAMP WITH TIME ZONE DEFAULT now()
);


CREATE TABLE IF NOT EXISTS client_secrets (
    id           SERIAL PRIMARY KEY,
    client_id    uuid,
    secret       CHARACTER(63),
    describe     VARCHAR(127),
    create_time  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE DEFAULT now()
);
CREATE INDEX idx_client_secrets_client_id ON client_secrets(client_id);


CREATE TABLE IF NOT EXISTS codes (
    id           SERIAL PRIMARY KEY,
    client_id    uuid,
    code         CHARACTER(31),
    scope        VARCHAR(255),
    create_time  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE DEFAULT now()
);
CREATE INDEX idx_code_tenant_user_id ON codes(client_id, code);
