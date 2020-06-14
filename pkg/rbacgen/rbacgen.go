package rbacgen

import (
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/omc-college/management-system/pkg/rbac"
)

func GetRoleTmpl(specs []string) (rbac.RoleTmpl, error) {
	var roleTmpl roleTmpl
	var tmplFeature featureEntry
	var tmplEndpoint endpoint

	var rbacFeaturesWithDetails map[string]apiFeature
	var rbacFeatures []string
	var isFeatureExisting bool

	roleTmpl.Entries = make(map[string]featureEntry)

	for _, spec := range specs {
		swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(spec)
		if err != nil {
			return rbac.RoleTmpl{}, err
		}

		err = json.Unmarshal(swagger.Components.ExtensionProps.Extensions["x-rbac-features"].(json.RawMessage), &rbacFeaturesWithDetails)
		if err != nil {
			return rbac.RoleTmpl{}, err
		}

		for path, pathItem := range swagger.Paths {
			for operationName, operation := range pathItem.Operations() {
				err := json.Unmarshal(operation.ExtensionProps.Extensions["x-rbac-features"].(json.RawMessage), &rbacFeatures)
				if err != nil {
					return rbac.RoleTmpl{}, err
				}

				// check whether current endpoint is a part of any feature
				if len(rbacFeatures) == 0 {
					err := fmt.Errorf("endpoint is not a part of any feature")
					return rbac.RoleTmpl{}, err
				}

				tmplEndpoint.Name = operation.OperationID
				tmplEndpoint.Path = path
				tmplEndpoint.Method = operationName

				for _, featureName := range rbacFeatures {
					_, isFeatureExisting = roleTmpl.Entries[featureName]
					if !isFeatureExisting {
						tmplFeature.Name = featureName
						tmplFeature.Description = rbacFeaturesWithDetails[featureName].Description
						tmplFeature.Endpoints = make(map[string]endpoint)

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

func toRoleTmpl(roleTmpl roleTmpl) (genericRoleTmpl rbac.RoleTmpl) {
	var genericFeatures []rbac.FeatureEntry
	var genericEndpoints []rbac.Endpoint

	genericFeatures = []rbac.FeatureEntry{}
	for _, feature := range roleTmpl.Entries {
		genericEndpoints = []rbac.Endpoint{}
		for _, endpoint := range feature.Endpoints {
			genericEndpoint := rbac.Endpoint{Name: endpoint.Name, Path: endpoint.Path, Method: endpoint.Method}
			genericEndpoints = append(genericEndpoints, genericEndpoint)
		}
		genericFeature := rbac.FeatureEntry{Name: feature.Name, Description: feature.Description, Endpoints: genericEndpoints}
		genericFeatures = append(genericFeatures, genericFeature)
	}
	genericRoleTmpl = rbac.RoleTmpl{Entries: genericFeatures}

	return genericRoleTmpl
}
