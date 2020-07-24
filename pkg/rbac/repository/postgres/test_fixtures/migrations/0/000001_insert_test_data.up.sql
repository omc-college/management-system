BEGIN;

INSERT INTO roles(id, name) VALUES (2, 'admin');
INSERT INTO roles(id, name) VALUES (3, 'user');

INSERT INTO features(id, name, description) 
VALUES (1, 'rolesManagement', 'description of rolesManagement');
INSERT INTO features(id, name, description) 
VALUES (2, 'rolesReading', 'description of rolesReading');

INSERT INTO endpoints(id, name, path, method)
VALUES (1, 'GetAllRoles', '/roles', 'GET');
INSERT INTO endpoints(id, name, path, method)
VALUES (2, 'CreateRole', '/roles', 'POST');

INSERT INTO roles_to_features(role_id, feature_id) VALUES (2, 1);
INSERT INTO roles_to_features(role_id, feature_id) VALUES (2, 2);
INSERT INTO roles_to_features(role_id, feature_id) VALUES (3, 2);

INSERT INTO features_to_endpoints(feature_id, endpoint_id) VALUES (1, 1);
INSERT INTO features_to_endpoints(feature_id, endpoint_id) VALUES (1, 2);
INSERT INTO features_to_endpoints(feature_id, endpoint_id) VALUES (2, 1);

COMMIT;