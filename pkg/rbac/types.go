package rbac

import (
	"fmt"
)

const RoleType = "role"
const RoleOperationCreate = "role.create"
const RoleOperationUpdate = "role.update"
const RoleOperationDelete = "role.delete"
const RolesTopicName = "roles"

var ErrNotAuthorized = fmt.Errorf("cannot authorize")
var ErrInvalidType = fmt.Errorf("cannot recognize type")
var ErrInvalidOperation = fmt.Errorf("cannot recognize operation")
var ErrInvalidPayload = fmt.Errorf("invalid payload passed")
var ErrCreateExistingRole = fmt.Errorf("already existing role should not be created")
var ErrDeleteNotExistingRole = fmt.Errorf("unexisting role should not be deleted")

type Role struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Entries []FeatureEntry `json:"entries"`
}

type FeatureEntry struct {
	ID          int        `json:"id,omitempty" yaml:"entries,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Endpoints   []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	ID     int    `json:"id,omitempty" yaml:"entries,omitempty"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type RoleTmpl struct {
	Entries []FeatureEntry `json:"entries"`
}

type Rule struct {
	PathRegExp string   `json:"pathRegExp"`
	Methods    []Method `json:"methods"`
}

type Method struct {
	Name  string `json:"name"`
	Roles []int  `json:"roles"`
}

type Input struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Token  string `json:"token"`
	Cache  *Cache `json:"cache"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
