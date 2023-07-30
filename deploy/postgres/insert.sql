


insert into user_pools (name, describe) values ('test', 'default client test user pool');
insert into user_pools (name, describe) values ('test2', 'default client test user pool2');

insert into tenants (app_id, type, name, host, company, describe, redirect_uris, grant_types)
 VALUES ('764e2e6b5f9b4ac983d2d18ec845b923', 1, 'default', 'localhost:3000', 'local', 'localhost test tenant', '{http://localhost}', '{read_user}');


DELETE FROM user_pools WHERE ID IN (SELECT ID FROM user_pools WHERE ID NOT IN (SELECT ID FROM tenants));
