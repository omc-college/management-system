package opa

import "github.com/omc-college/management-system/pkg/rbac/authcache"

type RegoInput struct {
	Path   string          `json:"path"`
	Method string          `json:"method"`
	Token  string          `json:"token"`
	Cache  authcache.Cache `json:"cache"`
}
