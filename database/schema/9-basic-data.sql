INSERT INTO roles(name, access_level) VALUES ('root', 1), ('admin', 2), ('supervisor', 3), ('cashier', 4), ('waiter', 5), ('cook', 5);

INSERT INTO users(username, password, verified) VALUES ('root', 'root', true), ('luisflahan', '4051', true), ('mark', 'mark', true), ('ann', '123', false);

INSERT INTO inherit_user_roles(user_id, role_id) VALUES ('1','1'),('2','1');

/*SELECT inherit_role.id, inherit_role.created_at, inherit_role.updated_at, inherit_role.deleted_at, role.name, role.access_level, inherit_role.user_id FROM roles role, inherit_user_roles inherit_role WHERE inherit_role.role_id = role.id;
*/

/*SELECT inherit_role.id, inherit_role.created_at, inherit_role.updated_at, inherit_role.deleted_at, role.name, role.access_level, inherit_role.user_id, user.id, user.created_at, user.updated_at, user.deleted_at, user.name, user.lastname, user.username, user.password, user.photo, user.verified, user.warning, user.darktheme, user.active_contract, user.address, user.born, user.degree_study, user.relation_ship, user.curp, user.rfc, user.citizen_id, user.credential_id, user.origin_state, user.score, user.qualities, user.defects, user.branch_id, user.origin_branch_id FROM roles role, inherit_user_roles inherit_role, users user WHERE inherit_role.role_id = role.id AND inherit_role.user_id = user.id;
*/

INSERT INTO notifications(type, solved, description) VALUES ('DONE', true, 'Database created; Basic data created');
INSERT INTO notification_images(image, notification_id) VALUES ('http//localhost:8080/notification/1/images/image_1.webp', 1);