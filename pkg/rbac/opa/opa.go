package opa

import (
	"context"

	"github.com/open-policy-agent/opa/rego"
)

func GetDecision(ctx context.Context, requestRegoInput RegoInput) error {
	authorizationRego := rego.New(
		rego.Query("data.authorization.isAccessGranted"),
		rego.Input(requestRegoInput),
		rego.Load([]string{"../../pkg/rbac/opa/authorization.rego"}, nil))

	regoResult, err := authorizationRego.Eval(ctx)
	if err != nil {
		return err
	}

	if !regoResult[0].Expressions[0].Value.(bool) {
		return ErrNotAuthorized
	}

	return nil
}
