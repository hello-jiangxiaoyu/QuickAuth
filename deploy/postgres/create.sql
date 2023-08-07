
CREATE TABLE IF NOT EXISTS apps (
    id             CHAR(32) PRIMARY KEY,
    tag            VARCHAR(127) NOT NULL,
    name           VARCHAR(127) NOT NULL UNIQUE,
    describe       VARCHAR(255) NOT NULL,
    icon           VARCHAR(127) NOT NULL,
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);



CREATE TABLE IF NOT EXISTS user_pools (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(127) NOT NULL,
    describe    VARCHAR(127) NOT NULL,
    is_disabled BOOLEAN NOT NULL DEFAULT false,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS users (
    id              CHAR(32) PRIMARY KEY,
    user_pool_id    BIGSERIAL NOT NULL,
    username        VARCHAR(127) NOT NULL,
    password        VARCHAR(127) NOT NULL,
    nick_name       VARCHAR(127) NOT NULL,
    display_name    VARCHAR(127) NOT NULL,
    gender          CHAR NOT NULL DEFAULT 0,  -- M:男性 F:女性 O:其他
    birthdate       DATE,                     -- 出生日期
    email           VARCHAR(127) NULL,
    email_verified  BOOLEAN NOT NULL DEFAULT false,
    phone           VARCHAR(20) NULL,
    phone_verified  BOOLEAN NOT NULL DEFAULT false,
    addr            VARCHAR(255) NOT NULL,
    avatar          VARCHAR(255) NOT NULL,
    type            INTEGER NOT NULL,
    is_disabled     BOOLEAN NOT NULL DEFAULT false,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_user_pool_username ON users(user_pool_id, username);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_phone ON users(phone);




CREATE TABLE IF NOT EXISTS tenants (
    id              BIGSERIAL PRIMARY KEY,
    app_id          CHAR(32) NOT NULL REFERENCES apps(id),
    user_pool_id    BIGSERIAL NOT NULL REFERENCES user_pools(id),
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
    is_code         BOOLEAN NOT NULL DEFAULT true,  -- 是否开启authorization_code
    is_refresh      BOOLEAN NOT NULL DEFAULT true,  -- 是否返回refresh_token
    is_password     BOOLEAN NOT NULL DEFAULT false,  -- 是否开启password授权模式
    is_credential   BOOLEAN NOT NULL DEFAULT true,  -- 是否开启client_credential
    is_device_flow  BOOLEAN NOT NULL DEFAULT false,  -- 是否开启device_flow
    config          JSONB NOT NULL,
    describe        VARCHAR(127) NOT NULL,
    is_disabled     BOOLEAN NOT NULL DEFAULT false,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_tenants_host ON tenants(host);
CREATE UNIQUE INDEX IF NOT EXISTS idx_tenants_name ON tenants(app_id, name);
CREATE INDEX IF NOT EXISTS idx_tenants_client_user_pool_id ON tenants(user_pool_id);




CREATE TABLE IF NOT EXISTS app_secrets (
    id              BIGSERIAL PRIMARY KEY,
    app_id          CHAR(32) NOT NULL,
    secret          CHAR(63) NOT NULL,    -- 客户端凭证密钥
    scope           VARCHAR(127) ARRAY NOT NULL,     -- 客户端凭证权限范围
    access_expire   INTEGER NOT NULL DEFAULT 604800,
    refresh_expire  INTEGER NOT NULL DEFAULT 2592000,
    describe        VARCHAR(127) NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_app_secret_id ON app_secrets (app_id, secret);

CREATE TABLE IF NOT EXISTS codes (
    id           BIGSERIAL PRIMARY KEY,
    user_id      BIGSERIAL NOT NULL,
    app_id       CHAR(32) NOT NULL,
    code         CHAR(32) NOT NULL,
    scope        VARCHAR(255) NOT NULL,
    state        CHAR(63) NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_codes_tenant_id ON codes(code, app_id);


CREATE TABLE IF NOT EXISTS providers (
    id              BIGSERIAL PRIMARY KEY,
    tenant_id       BIGSERIAL NOT NULL,
    app_id          CHAR(32) NOT NULL,
    type            VARCHAR(32) NOT NULL,
    client_id       VARCHAR(255) NOT NULL,
    client_secret   VARCHAR(255) NOT NULL,
    agent_id        VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_provider_tenant_type_id ON providers(tenant_id, app_id);
CREATE INDEX IF NOT EXISTS idx_provider_type ON providers(type);

