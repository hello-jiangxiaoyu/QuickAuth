
CREATE TABLE IF NOT EXISTS users (
    id           CHAR(32) PRIMARY KEY,
    user_pool_id BIGSERIAL NOT NULL,
    username     VARCHAR(127) NOT NULL,
    password     VARCHAR(127) NOT NULL DEFAULT '',
    display_name VARCHAR(127) NOT NULL,
    email        VARCHAR(127) NOT NULL,
    phone        VARCHAR(20) NOT NULL,
    type         INTEGER NOT NULL,
    is_disabled  INTEGER NOT NULL DEFAULT 0,
    create_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_users_user_pool_username ON users(username, user_pool_id);
CREATE UNIQUE INDEX idx_users_email ON users(email);
CREATE UNIQUE INDEX idx_users_phone ON users(phone);


CREATE TABLE IF NOT EXISTS tenants (
    id              BIGSERIAL PRIMARY KEY,
    app_id          CHAR(32) NOT NULL,
    user_pool_id    BIGSERIAL NOT NULL,
    type            INTEGER NOT NULL,
    name            VARCHAR(127) NOT NULL,
    host            VARCHAR(127) NOT NULL,
    company         VARCHAR(127) NOT NULL,
    grant_types     VARCHAR(127) ARRAY NOT NULL,
    redirect_uris   VARCHAR(127) ARRAY NOT NULL,
    code_expire     INTEGER NOT NULL DEFAULT 120,
    id_expire       INTEGER NOT NULL DEFAULT 604800,  -- id_token过期时间，默认7天
    access_expire   INTEGER NOT NULL DEFAULT 604800,  -- access_token过期时间，默认7天
    refresh_expire  INTEGER NOT NULL DEFAULT 2592000, -- refresh_token过期时间，默认30天
    is_code         INTEGER NOT NULL DEFAULT 1,  -- 是否开启authorization_code
    is_refresh      INTEGER NOT NULL DEFAULT 1,  -- 是否返回refresh_token
    is_password     INTEGER NOT NULL DEFAULT 0,  -- 是否开启password授权模式
    is_credential   INTEGER NOT NULL DEFAULT 1,  -- 是否开启client_credential
    is_device_flow  INTEGER NOT NULL DEFAULT 0,  -- 是否开启device_flow
    describe        VARCHAR(127) NOT NULL,
    is_disabled     INTEGER NOT NULL DEFAULT 0,
    create_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_tenants_host ON tenants(host);
CREATE INDEX idx_tenants_client_user_pool_id ON tenants(app_id, user_pool_id);



CREATE TABLE IF NOT EXISTS apps (
    id             CHAR(32) PRIMARY KEY,
    name           VARCHAR(127) NOT NULL,
    describe       VARCHAR(127) NOT NULL,
    create_time    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS app_secrets (
    id              BIGSERIAL PRIMARY KEY,
    app_id          CHAR(32) NOT NULL,
    secret          CHARACTER(63) NOT NULL,    -- 客户端凭证密钥
    scope           VARCHAR(255) NOT NULL,     -- 客户端凭证权限范围
    access_expire   INTEGER NOT NULL DEFAULT 604800,
    refresh_expire  INTEGER NOT NULL DEFAULT 2592000,
    describe        VARCHAR(127) NOT NULL,
    create_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    update_time     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_app_secret_id ON app_secrets (app_id, secret);



