package opa

import (
	"context"

	"github.com/omc-college/management-system/pkg/rbac"
	"github.com/open-policy-agent/opa/rego"
)

type PolicyAgent struct {
	policyPath string
}

func NewPolicyAgent(policyPath string) *PolicyAgent {
	return &PolicyAgent{
		policyPath: policyPath,
	}
}

func (pa PolicyAgent)GetDecision(ctx context.Context, requestRegoInput rbac.Input) error {
	authorizationRego := rego.New(
		rego.Query("data.authorization.isAccessGranted"),
		rego.Input(requestRegoInput),
		rego.Load([]string{pa.policyPath}, nil))

	regoResult, err := authorizationRego.Eval(ctx)
	if err != nil {
		return err
	}

	if !regoResult[0].Expressions[0].Value.(bool) {
		return rbac.ErrNotAuthorized
	}

	return nil
}
