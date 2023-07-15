

insert into apps (id, name, describe)
values ('8767EE8483AD8BD6FA7962A9F77E228A', 'default', 'test client');


insert into user_pools (name, describe) values ('test', 'default client test user pool');
insert into user_pools (name, describe) values ('test2', 'default client test user pool2');

insert into tenants (app_id, type, name, host, company, describe, redirect_uris, grant_types)
 VALUES ('8767EE8483AD8BD6FA7962A9F77E228A', 1, 'default', 'localhost', 'local', 'localhost test tenant', '{localhost}', '{read_user}');


