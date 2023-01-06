INSERT INTO roles(name, access_level) VALUES ('root', 1), ('admin', 2), ('supervisor', 3), ('cashier', 4), ('waiter', 5), ('cook', 5);

INSERT INTO users(username, password, verified) VALUES ('root', 'root', true), ('luisflahan', '4051', true), ('mark', 'mark', true), ('ann', '123', false);

INSERT INTO inherit_user_roles(user_id, role_id) VALUES ('1','1'),('2','1');
