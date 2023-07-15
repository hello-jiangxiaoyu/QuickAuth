

insert into apps (id, name, describe)
values ('8767EE8483AD8BD6FA7962A9F77E228A', 'default', 'test client');


insert into user_pools (name, describe) values ('test', 'default client test user pool');
insert into user_pools (name, describe) values ('test2', 'default client test user pool2');

insert into tenants (app_id, type, name, host, company, describe, redirect_uris, grant_types)
 VALUES ('8767EE8483AD8BD6FA7962A9F77E228A', 1, 'default', 'localhost', 'local', 'localhost test tenant', '{localhost}', '{read_user}');


INSERT INTO "app_secrets"
    ("app_id","secret","scope","access_expire","refresh_expire","describe")
VALUES ('8767EE8483AD8BD6FA7962A9F77E228A','zSXqDyajvlNqw0I1RwJil5PiB8BN3R1','user_profile',604800,2592000,'new')
RETURNING "id","create_time","update_time"
