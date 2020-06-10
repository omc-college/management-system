package opa

import (
	"context"

	"github.com/open-policy-agent/opa/rego"
	"github.com/sirupsen/logrus"
)

func GetDecision(ctx context.Context, requestRegoInput RegoInput) bool {
	var err error

	authorizationRego := rego.New(
		rego.Query("data.authorization.isAccessGranted"),
		rego.Input(requestRegoInput),
		rego.Load([]string{"../../pkg/rbac/opa/authorization.rego",
			"../../pkg/rbac/api/middleware/authorizationCache.json"}, nil))

	regoResult, err := authorizationRego.Eval(ctx)
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	return regoResult[0].Expressions[0].Value.(bool)
}
