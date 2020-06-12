package authcache

type authRules struct {
	Rules []authRule `json:"rules"`
}

type authRule struct {
	PathRegExp string       `json:"pathRegExp"`
	Methods    []authMethod `json:"methods"`
}

type authMethod struct {
	Name  string `json:"name"`
	Roles []int  `json:"roles"`
}
