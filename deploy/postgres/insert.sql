

insert into apps (id, name, describe)
values ('16a83dfccebc4267b58b53ee2a4f3590', 'default', 'test client');


insert into user_pools (name, describe) values ('test', 'default client test user pool');
insert into user_pools (name, describe) values ('test2', 'default client test user pool2');

insert into tenants (app_id, type, name, host, company, describe, redirect_uris, grant_types)
 VALUES ('16a83dfccebc4267b58b53ee2a4f3590', 1, 'default', 'localhost', 'local', 'localhost test tenant', '{localhost}', '{read_user}');



