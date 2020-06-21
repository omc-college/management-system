package rbac_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/omc-college/management-system/pkg/rbac"
	"reflect"
	"testing"
)

type testEnvelope struct {
	entityType string
	operation string
	payload json.RawMessage
	isFaling error
}

func (te *testEnvelope) EntityType() string {
	return te.entityType
}

func (te *testEnvelope) Operation() string {
	return te.operation
}

func (te *testEnvelope) Payload() json.RawMessage {
	return te.payload
}

func (te *testEnvelope) IsFaling() error {
	return te.isFaling
}

func (te *testEnvelope) Ack() error {
	return nil
}

func TestUpdate_ErrorShouldBeReturnedWhenInvalidTypePassed(t *testing.T) {
	cache := rbac.Cache{}

	envelopeTypes := []testEnvelope{
		{
			"deadbeef",
			rbac.RoleOperationCreate,
			[]byte{},
			nil,
		},
	}

	for _, envelope := range envelopeTypes {
		t.Run(envelope.EntityType(), func(t *testing.T) {
			err := cache.Update(&envelope)
			if err != nil {
				if !errors.Is(err, rbac.ErrInvalidType) {
					t.Fatalf("Expected: %v; got: %v", rbac.ErrInvalidType, err)
				}
			}
			if err == nil {
				t.Fatalf("Expected: %v; got: %v", rbac.ErrInvalidType, err)
			}
		})
	}
}

func TestUpdate_OkWhenValidTypePassed(t *testing.T) {
	cache := rbac.Cache {}

	envelopeTypes := []testEnvelope{
		{
			rbac.RoleType,
			rbac.RoleOperationCreate,
			[]byte{},
			nil,
		},
	}

	for _, envelope := range envelopeTypes {
		t.Run(envelope.EntityType(), func(t *testing.T) {
			err := cache.Update(&envelope)
			if err != nil {
				if errors.Is(err, rbac.ErrInvalidType) {
					t.Fatalf("got invalid type error: %s", err.Error())
				}
			}
		})
	}
}

func TestUpdate_ErrorShouldBeReturnedWhenInvalidOperationPassed(t *testing.T) {
	cache := rbac.Cache{}

	envelopeOperations := []testEnvelope{
		{
			rbac.RoleType,
			"deadbeef",
			[]byte{},
			nil,
		},
	}

	for _, envelope := range envelopeOperations {
		t.Run(envelope.Operation(), func(t *testing.T) {
			err := cache.Update(&envelope)
			if err != nil {
				if !errors.Is(err, rbac.ErrInvalidOperation) {
					t.Fatalf("Expected: %v; got: %v", rbac.ErrInvalidOperation, err)
				}
			}
			if err == nil {
				t.Fatalf("Expected: %v; got: %v", rbac.ErrInvalidType, err)
			}
		})
	}
}

func TestUpdate_OkWhenValidOperationPassed(t *testing.T) {
	cache := rbac.Cache{}

	envelopeOperations := []testEnvelope{
		{
			rbac.RoleType,
			rbac.RoleOperationCreate,
			[]byte{},
			nil,
		},
		{
			rbac.RoleType,
			rbac.RoleOperationUpdate,
			[]byte{},
			nil,
		},
		{
			rbac.RoleType,
			rbac.RoleOperationDelete,
			[]byte{},
			nil,
		},
	}

	for _, envelope := range envelopeOperations {
		t.Run(envelope.Operation(), func(t *testing.T) {
			err := cache.Update(&envelope)
			if err != nil {
				if errors.Is(err, rbac.ErrInvalidOperation) {
					t.Fatalf("got invalid type error: %s", err.Error())
				}
			}
		})
	}
}

func TestUpdate_ErrorShouldBeReturnedWhenInvalidPayloadPassed(t *testing.T) {
	cache := rbac.Cache{}

	invalidPayload := []byte("S32}aita{26ma{")

	envelopePayload := []testEnvelope{
		{
			rbac.RoleType,
			rbac.RoleOperationCreate,
			invalidPayload,
			nil,
		},
		{
			rbac.RoleType,
			rbac.RoleOperationUpdate,
			invalidPayload,
			nil,
		},
		{
			rbac.RoleType,
			rbac.RoleOperationDelete,
			invalidPayload,
			nil,
		},
	}

	for _, envelope := range envelopePayload {
		t.Run(envelope.Operation(), func(t *testing.T) {
			err := cache.Update(&envelope)
			if err != nil {
				if !errors.Is(err, rbac.ErrInvalidPayload) {
					t.Fatalf("Expected: %v; got: %v", rbac.ErrInvalidPayload, err)
				}
			}
			if err == nil {
				t.Fatalf("Expected: %v; got: %v", rbac.ErrInvalidType, err)
			}
		})
	}
}

func TestUpdate_OkWhenValidPayloadPassed(t *testing.T) {
	cache := rbac.Cache{}

	payload := rbac.Role {
		ID: 1,
		Name: "roleName",
		Entries: []rbac.FeatureEntry {
			{
				ID: 1,
				Name: "featureName",
				Description: "featureDescr",
				Endpoints: []rbac.Endpoint {
					{
						ID: 1,
						Name: "endpointName",
						Path: "endpointPath",
						Method: "endpointMethod",
					},
				},
			},
		},
	}

	payloadRole, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("did not marshal payloadRole")
	}

	IDPayload := 1

	payloadID, err := json.Marshal(IDPayload)
	if err != nil {
		t.Fatalf("did not marshal payloadID")
	}

	envelopePayload := []testEnvelope{
		{
			rbac.RoleType,
			rbac.RoleOperationCreate,
			payloadRole,
			nil,
		},
		{
			rbac.RoleType,
			rbac.RoleOperationUpdate,
			payloadRole,
			nil,
		},
		{
			rbac.RoleType,
			rbac.RoleOperationDelete,
			payloadID,
			nil,
		},
	}

	for _, envelope := range envelopePayload {
		t.Run(envelope.Operation(), func(t *testing.T) {
			err := cache.Update(&envelope)
			if err != nil {
				if errors.Is(err, rbac.ErrInvalidPayload) {
					t.Fatalf("got invalid type error: %s", err.Error())
				}
			}
		})
	}
}

func TestUpdate_OkWhenNotExistingRoleCreated(t *testing.T) {
	 testCases := []struct{
		Path string
		Method string
		Role int
		ExpectedCache rbac.Cache
	}{
		{"/path", "POST", 2, rbac.Cache {
			[]rbac.Rule {
				{
					"^/path$",
					[]rbac.Method {
						{
							"GET",
							[]int{1},
						},
						{
							"POST",
							[]int{2},
						},
					},
				},
			},
		} },
		{"/overPath", "GET", 2, rbac.Cache {
		[]rbac.Rule {
				{
					"^/path$",
					[]rbac.Method {
						{
							"GET",
							[]int{1},
						},
					},
				},
				{
					"^/overPath$",
					[]rbac.Method {
						{
							"GET",
							[]int{2},
						},
					},
				},
			},
		} },
		{"/path", "GET", 2, rbac.Cache {
			[]rbac.Rule {
				{
					"^/path$",
					[]rbac.Method {
						{
							"GET",
							[]int{1,2},
						},
					},
				},
			},
		}},
		{"/overPath", "POST", 2, rbac.Cache {
			[]rbac.Rule {
				{
					"^/path$",
					[]rbac.Method {
						{
							"GET",
							[]int{1},
						},
					},
				},
				{
					"^/overPath$",
					[]rbac.Method {
						{
							"POST",
							[]int{2},
						},
					},
				},
			},
		}},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v, %s", testCaseIndex, testCase.Path), func(t *testing.T) {
			cache := rbac.Cache {
				[]rbac.Rule {
					{
						"^/path$",
						[]rbac.Method {
							{
								"GET",
								[]int{1},
							},
						},
					},
				},
			}

			payload := rbac.Role {
				ID: testCase.Role,
				Name: "roleName",
				Entries: []rbac.FeatureEntry {
					{
						ID: 1,
						Name: "featureName",
						Description: "featureDescr",
						Endpoints: []rbac.Endpoint {
							{
								ID: 1,
								Name: "endpointName",
								Path: testCase.Path,
								Method: testCase.Method,
							},
						},
					},
				},
			}

			payloadRole, err := json.Marshal(payload)
			if err != nil {
				t.Fatal("did not marshal payloadRole")
			}

			envelope := testEnvelope{
				rbac.RoleType,
				rbac.RoleOperationCreate,
				payloadRole,
				nil,
			}

			err = cache.Update(&envelope)
			if err != nil {
				t.Fatalf("got update error: %v", err)
			}

			if !reflect.DeepEqual(cache, testCase.ExpectedCache) {
				t.Fatalf("Expected: %v; got: %v", testCase.ExpectedCache, cache)
			}
		})
	}
}

func TestUpdate_ErrorShouldBeReturnedWhenExistingRoleCreated(t *testing.T) {
	testCases := []struct{
		Path string
		Method string
		Role int
		PathRegExp string
	}{
		{"/path", "POST", 1, "^/path$"},
		{"/overPath", "GET", 1, "^/overPath$"},
		{"/path", "GET", 1, "^/path$"},
		{"/overPath", "POST", 1, "^/overPath$"},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v, %s", testCaseIndex, testCase.Path), func(t *testing.T) {
			cache := rbac.Cache {
				[]rbac.Rule {
					{
						"^/path$",
						[]rbac.Method {
							{
								"GET",
								[]int{1},
							},
						},
					},
				},
			}

			payload := rbac.Role {
				ID: testCase.Role,
				Name: "roleName",
				Entries: []rbac.FeatureEntry {
					{
						ID: 1,
						Name: "featureName",
						Description: "featureDescr",
						Endpoints: []rbac.Endpoint {
							{
								ID: 1,
								Name: "endpointName",
								Path: testCase.Path,
								Method: testCase.Method,
							},
						},
					},
				},
			}

			payloadRole, err := json.Marshal(payload)
			if err != nil {
				t.Fatal("did not marshal payloadRole")
			}

			envelope := testEnvelope{
				rbac.RoleType,
				rbac.RoleOperationCreate,
				payloadRole,
				nil,
			}

			err = cache.Update(&envelope)
			if err != nil {
				if !errors.Is(err, rbac.ErrCreateExistingRole) {
					t.Fatalf("Expected: %v; got: %v", rbac.ErrCreateExistingRole, err)
				}
			}
			if err == nil {
				t.Fatalf("Expected: %v; got: %v", rbac.ErrCreateExistingRole, err)
			}
		})
	}
}

func TestUpdate_OkWhenExistingRoleDeleted(t *testing.T) {
	testCases := []struct{
		Role int
		ExpectedCache rbac.Cache
	}{
		{ 1, rbac.Cache {
			[]rbac.Rule {},
		}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.Role), func(t *testing.T) {
			cache := rbac.Cache {
				[]rbac.Rule {
					{
						"^/path$",
						[]rbac.Method {
							{
								"GET",
								[]int{1},
							},
						},
					},
				},
			}

			payload := testCase.Role

			payloadRole, err := json.Marshal(payload)
			if err != nil {
				t.Fatal("did not marshal payloadRole")
			}

			envelope := testEnvelope{
				rbac.RoleType,
				rbac.RoleOperationDelete,
				payloadRole,
				nil,
			}

			err = cache.Update(&envelope)
			if err != nil {
				t.Fatalf("got update error: %v", err)
			}

			fmt.Print(cache.Rules)

			if !reflect.DeepEqual(cache, testCase.ExpectedCache) {
				t.Fatalf("Expected: %v; got: %v", testCase.ExpectedCache, cache)
			}
		})
	}
}

func TestUpdate_ErrorShouldBeReturnedWhenNotExistingRoleDeleted(t *testing.T) {
	testCases := []struct{
		Role int
	}{
		{ 2},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.Role), func(t *testing.T) {
			cache := rbac.Cache {
				[]rbac.Rule {
					{
						"^/path$",
						[]rbac.Method {
							{
								"GET",
								[]int{1},
							},
						},
					},
				},
			}

			payload := testCase.Role

			payloadRole, err := json.Marshal(payload)
			if err != nil {
				t.Fatal("did not marshal payloadRole")
			}

			envelope := testEnvelope {
				rbac.RoleType,
				rbac.RoleOperationDelete,
				payloadRole,
				nil,
			}

			err = cache.Update(&envelope)
			if err != nil {
				if !errors.Is(err, rbac.ErrDeleteNotExistingRole) {
					t.Fatalf("Expected: %v; got: %v", rbac.ErrDeleteNotExistingRole, err)
				}
			}
			if err == nil {
				t.Fatalf("Expected: %v; got: %v", rbac.ErrDeleteNotExistingRole, err)
			}
		})
	}
}


func TestUpdate_OkWhenExistingRoleUpdated(t *testing.T) {
	testCases := []struct{
		Path string
		Method string
		Role int
		ExpectedCache rbac.Cache
	}{
		{"/path", "POST", 1, rbac.Cache{
			[]rbac.Rule{
				{
					"^/path$",
					[]rbac.Method{
						{
							"POST",
							[]int{1},
						},
					},
				},
			},
		}},
		{"/overPath", "GET", 1, rbac.Cache{
			[]rbac.Rule{
				{
					"^/overPath$",
					[]rbac.Method{
						{
							"GET",
							[]int{1},
						},
					},
				},
			},
		}},
		{"/path", "GET", 1, rbac.Cache{
			[]rbac.Rule{
				{
					"^/path$",
					[]rbac.Method{
						{
							"GET",
							[]int{1},
						},
					},
				},
			},
		}},
		{"/overPath", "POST", 1, rbac.Cache{
			[]rbac.Rule{
				{
					"^/overPath$",
					[]rbac.Method{
						{
							"POST",
							[]int{1},
						},
					},
				},
			},
		}},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v, %s", testCaseIndex, testCase.Path), func(t *testing.T) {
			cache := rbac.Cache {
				[]rbac.Rule {
					{
						"^/path$",
						[]rbac.Method {
							{
								"GET",
								[]int{1},
							},
						},
					},
				},
			}

			payload := rbac.Role {
				ID: testCase.Role,
				Name: "roleName",
				Entries: []rbac.FeatureEntry {
					{
						ID: 1,
						Name: "featureName",
						Description: "featureDescr",
						Endpoints: []rbac.Endpoint {
							{
								ID: 1,
								Name: "endpointName",
								Path: testCase.Path,
								Method: testCase.Method,
							},
						},
					},
				},
			}

			payloadRole, err := json.Marshal(payload)
			if err != nil {
				t.Fatal("did not marshal payloadRole")
			}

			envelope := testEnvelope{
				rbac.RoleType,
				rbac.RoleOperationUpdate,
				payloadRole,
				nil,
			}

			err = cache.Update(&envelope)
			if err != nil {
				t.Fatalf("got update error: %v", err)
			}

			if !reflect.DeepEqual(cache, testCase.ExpectedCache) {
				t.Fatalf("Expected: %v; got: %v", testCase.ExpectedCache, cache)
			}
		})
	}
}






