
CREATE TABLE IF NOT EXISTS users (
    id           uuid PRIMARY KEY,
    user_pool_id uuid NOT NULL,
    username     VARCHAR(127) NOT NULL,
    password     VARCHAR(127),
    display_name VARCHAR(127),
    email        VARCHAR(127),
    phone        VARCHAR(20),
    is_disabled  BOOLEAN NOT NULL DEFAULT FALSE,
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_users_user_pool_username ON users(user_pool_id, username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);

CREATE TABLE IF NOT EXISTS user_pools (
    id          uuid PRIMARY KEY,
    name        VARCHAR(127) NOT NULL,
    describe    VARCHAR(127),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);



CREATE TABLE IF NOT EXISTS tenants (
    id           uuid PRIMARY KEY,
    client_id    uuid NOT NULL,
    user_pool_id uuid NOT NULL,
    type         INTEGER NOT NULL,
    name         VARCHAR(127) NOT NULL,
    host         VARCHAR(127) NOT NULL,
    company      VARCHAR(127) NOT NULL,
    describe     VARCHAR(127),
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_tenants_host ON tenants(host);
CREATE INDEX idx_tenants_client_user_pool_id ON tenants(client_id, user_pool_id);



CREATE TABLE IF NOT EXISTS clients (
    id             uuid PRIMARY KEY,
    name           VARCHAR(127) NOT NULL,
    describe       VARCHAR(127),
    grant_types    VARCHAR(127) ARRAY,
    token_expire   INTEGER NOT NULL,
    refresh_expire INTEGER NOT NULL,
    create_time    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS client_secrets (
    id           SERIAL PRIMARY KEY,
    client_id    uuid NOT NULL,
    secret       CHARACTER(63) NOT NULL,
    describe     VARCHAR(127),
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX idx_client_secrets_client_id ON client_secrets(client_id, secret);

CREATE TABLE IF NOT EXISTS redirect_uris (
    id           SERIAL PRIMARY KEY,
    client_id    uuid NOT NULL,
    uri          CHARACTER(63) NOT NULL,
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX idx_redirect_uri_client_id_uri ON redirect_uris(client_id, uri);



