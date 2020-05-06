package openapi

import (
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/omc-college/management-system/pkg/rbac/models"
)

func GetRoleTmpl(specs []string) (models.RoleTmpl, error) {
	var roleTmpl roleTmpl
	var tmplFeature featureEntry
	var tmplEndpoint endpoint

	var rbacFeaturesWithDetails map[string]apiFeature
	var rbacFeatures []string
	var isFeatureExisting bool

	roleTmpl.Entries = make(map[string]featureEntry)

	// iterate input specs
	for _, spec := range specs {
		// unmarshal current openapi spec
		swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(spec)
		if err != nil {
			return models.RoleTmpl{}, err
		}

		// unmarshal value of "x-rbac-features" from components
		err = json.Unmarshal(swagger.Components.ExtensionProps.Extensions["x-rbac-features"].(json.RawMessage), &rbacFeaturesWithDetails)
		if err != nil {
			return models.RoleTmpl{}, err
		}

		// iterate paths
		for path, pathItem := range swagger.Paths {
			// iterate operations
			for operationName, operation := range pathItem.Operations() {
				// unmarshal value of "x-rbac-features" extension property
				err := json.Unmarshal(operation.ExtensionProps.Extensions["x-rbac-features"].(json.RawMessage), &rbacFeatures)
				if err != nil {
					return models.RoleTmpl{}, err
				}

				// check whether current endpoint is a part of any feature
				if len(rbacFeatures) == 0 {
					err := fmt.Errorf("endpoint is not a part of any feature")
					return models.RoleTmpl{}, err
				}

				// form endpoint
				tmplEndpoint.Name = operation.OperationID
				tmplEndpoint.Path = path
				tmplEndpoint.Method = operationName

				// iterate features
				for _, featureName := range rbacFeatures {
					// check whether feature exists in roleTmpl
					_, isFeatureExisting = roleTmpl.Entries[featureName]
					if !isFeatureExisting {
						// init feature
						tmplFeature.Name = featureName
						tmplFeature.Description = rbacFeaturesWithDetails[featureName].Description
						tmplFeature.Endpoints = make(map[string]endpoint)

						// save feature to roleTmpl's features map
						roleTmpl.Entries[featureName] = tmplFeature
					}

					tmplFeature = roleTmpl.Entries[featureName]

					tmplFeature.Endpoints[tmplEndpoint.Name] = tmplEndpoint

					roleTmpl.Entries[tmplFeature.Name] = tmplFeature
				}
			}
		}
	}

	return toRoleTmpl(roleTmpl), nil
}

func toRoleTmpl(roleTmpl roleTmpl) (genericRoleTmpl models.RoleTmpl) {
	var genericFeatures []models.FeatureEntry
	var genericEndpoints []models.Endpoint

	genericFeatures = []models.FeatureEntry{}
	for _, feature := range roleTmpl.Entries {
		genericEndpoints = []models.Endpoint{}
		for _, endpoint := range feature.Endpoints {
			genericEndpoint := models.Endpoint{Name: endpoint.Name, Path: endpoint.Path, Method: endpoint.Method}
			genericEndpoints = append(genericEndpoints, genericEndpoint)
		}
		genericFeature := models.FeatureEntry{Name: feature.Name, Description: feature.Description, Endpoints: genericEndpoints}
		genericFeatures = append(genericFeatures, genericFeature)
	}
	genericRoleTmpl = models.RoleTmpl{Entries: genericFeatures}

	return genericRoleTmpl
}
