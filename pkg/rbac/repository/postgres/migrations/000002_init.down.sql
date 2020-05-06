BEGIN;

ALTER TABLE endpoints
DROP name;

ALTER TABLE features
DROP description,
DROP CONSTRAINT features_name_unique;

ALTER TABLE roles
DROP CONSTRAINT roles_name_unique;

ALTER TABLE features_to_endpoints
DROP CONSTRAINT features_to_endpoints_feature_id_fkey,
DROP CONSTRAINT features_to_endpoints_endpoint_id_fkey;

COMMIT;