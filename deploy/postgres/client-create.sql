
CREATE TABLE IF NOT EXISTS users (
    id           BIGSERIAL PRIMARY KEY,
    user_pool_id BIGSERIAL NOT NULL,
    open_id      uuid NOT NULL,
    username     VARCHAR(127) NOT NULL,
    password     VARCHAR(127) NOT NULL DEFAULT '',
    display_name VARCHAR(127),
    email        VARCHAR(127),
    phone        VARCHAR(20),
    type         INTEGER NOT NULL,
    is_disabled  INTEGER NOT NULL DEFAULT 0,
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_users_user_pool_username ON users(user_pool_id, username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);

CREATE TABLE IF NOT EXISTS user_pools (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(127) NOT NULL,
    describe    VARCHAR(127),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);



CREATE TABLE IF NOT EXISTS tenants (
    id              BIGSERIAL PRIMARY KEY,
    app_id          CHAR(32) NOT NULL,
    user_pool_id    BIGSERIAL NOT NULL,
    type            INTEGER NOT NULL,
    name            VARCHAR(127) NOT NULL,
    host            VARCHAR(127) NOT NULL,
    company         VARCHAR(127) NOT NULL,
    grant_types     VARCHAR(127) ARRAY,
    redirect_uris   VARCHAR(127) ARRAY,
    token_expire    INTEGER NOT NULL,
    refresh_expire  INTEGER NOT NULL,
    code_expire     INTEGER NOT NULL,
    describe        VARCHAR(127),
    create_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_tenants_host ON tenants(host);
CREATE INDEX idx_tenants_client_user_pool_id ON tenants(app_id, user_pool_id);



CREATE TABLE IF NOT EXISTS apps (
    id             CHAR(32) PRIMARY KEY,
    name           VARCHAR(127) NOT NULL,
    describe       VARCHAR(127),
    create_time    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS app_secrets (
    app_id       CHAR(32) NOT NULL,
    secret       CHARACTER(63) NOT NULL,
    describe     VARCHAR(127),
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
ALTER TABLE app_secrets ADD PRIMARY KEY (app_id, secret);



