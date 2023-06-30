

insert into clients (id, name, describe, grant_types, redirect_uris, token_expire, refresh_expire, code_expire)
values ('8767EE84-83AD-8BD6-FA79-62A9F77E228A', 'default', 'test client', '{code}', '{localhost}', 10, 100, 1);


insert into user_pools (name, describe) values ('test', 'default client test user pool');
insert into user_pools (name, describe) values ('test2', 'default client test user pool2');

insert into tenants (client_id, type, name, host, company, describe)
 VALUES ('8767EE84-83AD-8BD6-FA79-62A9F77E228A', 1, 'default', 'localhost', 'local', 'localhost test tenant');
