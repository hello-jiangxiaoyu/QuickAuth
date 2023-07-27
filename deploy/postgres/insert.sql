


insert into user_pools (name, describe) values ('test', 'default client test user pool');
insert into user_pools (name, describe) values ('test2', 'default client test user pool2');

insert into tenants (app_id, type, name, host, company, describe, redirect_uris, grant_types)
 VALUES ('de8de8e1769c4f3ca17b089c89bbcf50', 1, 'default', 'localhost', 'local', 'localhost test tenant', '{localhost}', '{read_user}');



