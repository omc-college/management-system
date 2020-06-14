package rbac

import (
	"fmt"
)

const RoleType = "role"
const RoleOperationCreate = "role.create"
const RoleOperationUpdate = "role.update"
const RoleOperationDelete = "role.delete"
const RolesTopicName = "roles"

const authCacheFilename = "authCache.json"

var ErrNotAuthorized = fmt.Errorf("cannot authorize")

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

type rules struct {
	Rules []rule `json:"rules"`
}

type rule struct {
	PathRegExp string   `json:"pathRegExp"`
	Methods    []method `json:"methods"`
}

type method struct {
	Name  string `json:"name"`
	Roles []int  `json:"roles"`
}

type Input struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Token  string `json:"token"`
	Cache  Cache  `json:"cache"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
