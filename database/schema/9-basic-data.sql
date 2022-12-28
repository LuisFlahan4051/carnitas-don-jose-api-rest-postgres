INSERT INTO roles(name, access_level) VALUES ('root', 1), ('admin', 2);

INSERT INTO users(username, password, verified) VALUES ('root', '123', true), ('luisflahan', '4051', true);

INSERT INTO inherit_user_roles(user_id, role_id) VALUES ('1','1'),('2','1');
