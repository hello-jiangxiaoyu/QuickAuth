
create database quick_auth;


CREATE TABLE IF NOT EXISTS users (
    id           uuid PRIMARY KEY,
    user_pool_id uuid NOT NULL,
    username     VARCHAR(255) NOT NULL,
    password     VARCHAR(255),
    display_name VARCHAR(255),
    email        VARCHAR(255),
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
    name        VARCHAR(255),
    describe    VARCHAR(255),
    create_time TIMESTAMP WITH TIME ZONE DEFAULT now(),
    update_time TIMESTAMP WITH TIME ZONE DEFAULT now()
);


CREATE TABLE IF NOT EXISTS tenants (
    id           uuid PRIMARY KEY,
    app_id       uuid NOT NULL,
    user_pool_id uuid NOT NULL,
    type         INTEGER,
    name         VARCHAR(255),
    host         VARCHAR(255) NOT NULL,
    company      VARCHAR(255) NOT NULL,
    describe     VARCHAR(255),
    create_time  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE DEFAULT now()
);
CREATE UNIQUE INDEX idx_tenants_host ON tenants(host);
CREATE INDEX idx_tenants_app_user_pool ON tenants(app_id, user_pool_id);



CREATE TABLE IF NOT EXISTS apps (
    id           uuid PRIMARY KEY,
    name         VARCHAR(255),
    describe     VARCHAR(255),
    create_time  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    update_time  TIMESTAMP WITH TIME ZONE DEFAULT now()
);


