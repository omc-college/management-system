package rbac

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/omc-college/management-system/pkg/pubsub"
	"github.com/sirupsen/logrus"
)

type Cache struct {
	Rules []Rule `json:"rules"`
}

// NewCache inits Cache based on full history from MQ
func NewCache() *Cache {
	return &Cache{}
}

func (cache *Cache) ListenUpdates(rolesChannel <-chan pubsub.Envelope) {
	for envelope := range rolesChannel {
		err := cache.Update(envelope)
		if err != nil {
			logrus.Fatalf("cannot update cache: %s", err.Error())
		}
	}
}

func (cache *Cache) Update(envelope pubsub.Envelope) error {
	if envelope.GetEntityType() != RoleType {
		return ErrInvalidType
	}

	switch envelope.GetOperation() {
	case RoleOperationCreate:
		return cache.createRole(envelope.GetPayload())
	case RoleOperationUpdate:
		return cache.updateRole(envelope.GetPayload())
	case RoleOperationDelete:
		return cache.deleteRole(envelope.GetPayload())
	default:
		return ErrInvalidOperation
	}
}

func (cache *Cache) createRole(rawNewRole json.RawMessage) error {
	var newRole Role

	err := json.Unmarshal(rawNewRole, &newRole)
	if err != nil {
		return ErrInvalidPayload
	}

	paramRegExp, err := regexp.Compile(`{\w+}`)
	if err != nil {
		return err
	}

	for _, cacheRule := range cache.Rules {
		for _, cacheMethod := range cacheRule.Methods {
			for _, cacheRoleID := range cacheMethod.Roles {
				if newRole.ID == cacheRoleID {
					return ErrCreateExistingRole
				}
			}
		}
	}

	for _, newFeature := range newRole.Entries {
		for _, newEndpoint := range newFeature.Endpoints {
			var newPathRegExp = fmt.Sprintf("^%s$", paramRegExp.ReplaceAll([]byte(newEndpoint.Path), []byte("\\w+")))
			var existingCacheRuleIndex int
			var isPathRegExpExisting bool

			for cacheRuleIndex, cacheRule := range cache.Rules {
				if newPathRegExp == cacheRule.PathRegExp {
					isPathRegExpExisting = true
					existingCacheRuleIndex = cacheRuleIndex

					break
				}
			}

			if !isPathRegExpExisting {

				newAuthMethod := Method{
					Name:  newEndpoint.Method,
					Roles: []int{newRole.ID},
				}

				newAuthRule := Rule{
					PathRegExp: newPathRegExp,
					Methods:    []Method{newAuthMethod},
				}

				cache.Rules = append(cache.Rules, newAuthRule)

				continue
			}

			var existingCacheAuthMethodIndex int
			var isMethodExisting bool

			for cacheAuthMethodID, cacheAuthMethod := range cache.Rules[existingCacheRuleIndex].Methods {
				if newEndpoint.Method == cacheAuthMethod.Name {
					isMethodExisting = true
					existingCacheAuthMethodIndex = cacheAuthMethodID
					break
				}
			}

			if !isMethodExisting {
				newAuthMethod := Method{
					Name:  newEndpoint.Method,
					Roles: []int{newRole.ID},
				}

				cache.Rules[existingCacheRuleIndex].Methods = append(cache.Rules[existingCacheRuleIndex].Methods, newAuthMethod)

				continue
			}

			cache.Rules[existingCacheRuleIndex].Methods[existingCacheAuthMethodIndex].Roles = append(cache.Rules[existingCacheRuleIndex].Methods[existingCacheAuthMethodIndex].Roles, newRole.ID)
		}
	}

	sort.SliceStable(cache.Rules, func(i, j int) bool {
		iLength := len(strings.Split(cache.Rules[i].PathRegExp, "/"))
		jLength := len(strings.Split(cache.Rules[j].PathRegExp, "/"))
		return iLength > jLength
	})

	return nil
}

func (cache *Cache) updateRole(rawNewRole json.RawMessage) error {
	var newRole Role

	err := json.Unmarshal(rawNewRole, &newRole)
	if err != nil {
		return ErrInvalidPayload
	}

	RawNewRoleID, err := json.Marshal(newRole.ID)
	if err != nil {
		return ErrInvalidPayload
	}

	err = cache.deleteRole(RawNewRoleID)
	if err != nil {
		return err
	}

	err = cache.createRole(rawNewRole)
	if err != nil {
		return err
	}

	return nil
}

func (cache *Cache) deleteRole(rawRoleID json.RawMessage) error {
	var roleID int

	err := json.Unmarshal(rawRoleID, &roleID)
	if err != nil {
		return ErrInvalidPayload
	}

	var isRuleDeleted bool
	var isRoleDeleted bool

	for cacheRuleIndex, cacheRule := range cache.Rules {
		for cacheAuthMethodIndex, cacheAuthMethod := range cacheRule.Methods {
			for cacheRoleIDIndex, cacheRoleID := range cacheAuthMethod.Roles {
				if roleID == cacheRoleID {
					cache.Rules[cacheRuleIndex].Methods[cacheAuthMethodIndex].Roles = append(cacheAuthMethod.Roles[:cacheRoleIDIndex], cacheAuthMethod.Roles[cacheRoleIDIndex+1:]...)

					if len(cache.Rules[cacheRuleIndex].Methods[cacheAuthMethodIndex].Roles) == 0 {
						cache.Rules[cacheRuleIndex].Methods = append(cacheRule.Methods[:cacheAuthMethodIndex], cacheRule.Methods[cacheAuthMethodIndex+1:]...)

						if len(cache.Rules[cacheRuleIndex].Methods) == 0 {
							isRuleDeleted = true

							cache.Rules = append(cache.Rules[:cacheRuleIndex], cache.Rules[cacheRuleIndex+1:]...)
						}
					}
					isRoleDeleted = true
				}
			}
		}
	}

	if !isRoleDeleted {
		return ErrDeleteNotExistingRole
	}

	if isRuleDeleted {
		sort.SliceStable(cache.Rules, func(i, j int) bool {
			iLength := len(strings.Split(cache.Rules[i].PathRegExp, "/"))
			jLength := len(strings.Split(cache.Rules[j].PathRegExp, "/"))
			return iLength > jLength
		})
	}

	return nil
}
