BEGIN;

ALTER TABLE endpoints
ADD name varchar(255) UNIQUE NOT NULL;

ALTER TABLE features
ADD description varchar(255) NOT NULL,
ADD CONSTRAINT features_name_unique
    UNIQUE (name);

ALTER TABLE roles
ADD CONSTRAINT roles_name_unique
    UNIQUE (name);

ALTER TABLE features_to_endpoints
DROP CONSTRAINT features_to_endpoints_feature_id_fkey,
ADD CONSTRAINT features_to_endpoints_feature_id_fkey
    FOREIGN KEY (feature_id)
    REFERENCES features(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,

DROP CONSTRAINT features_to_endpoints_endpoint_id_fkey,
ADD CONSTRAINT features_to_endpoints_endpoint_id_fkey
    FOREIGN KEY (endpoint_id)
    REFERENCES endpoints(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

COMMIT;