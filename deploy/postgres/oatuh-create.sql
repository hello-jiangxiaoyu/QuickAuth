CREATE TABLE IF NOT EXISTS codes (
    id           BIGSERIAL PRIMARY KEY,
    user_id      BIGSERIAL NOT NULL,
    app_id       CHAR(32) NOT NULL,
    code         CHAR(32) NOT NULL,
    scope        VARCHAR(255) NOT NULL,
    state        CHAR(63) NOT NULL,
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_codes_tenant_user_id ON codes(code, app_id);


CREATE TABLE IF NOT EXISTS providers (
    id              BIGSERIAL PRIMARY KEY,
    tenant_id       BIGSERIAL NOT NULL,
    app_id          CHAR(32) NOT NULL,
    type            VARCHAR(32) NOT NULL,
    client_id       VARCHAR(255) NOT NULL,
    client_secret   VARCHAR(255) NOT NULL,
    agent_id        VARCHAR(255) NOT NULL,
    create_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX idx_provider_tenant_type_id ON providers(tenant_id, app_id);


CREATE TABLE IF NOT EXISTS user_pools (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(127) NOT NULL,
    describe    VARCHAR(127) NOT NULL,
    is_disabled INTEGER NOT NULL DEFAULT 0,
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS users (
    id           CHAR(32) PRIMARY KEY,
    user_pool_id BIGSERIAL NOT NULL,
    username     VARCHAR(127) NOT NULL,
    password     VARCHAR(127) NOT NULL,
    display_name VARCHAR(127) NOT NULL,
    email        VARCHAR(127) NOT NULL,
    phone        VARCHAR(20) NOT NULL,
    type         INTEGER NOT NULL,
    is_disabled  INTEGER NOT NULL DEFAULT 0,
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_users_user_pool_username ON users(user_pool_id, username);
CREATE UNIQUE INDEX idx_users_email ON users(email);
CREATE UNIQUE INDEX idx_users_phone ON users(phone);
