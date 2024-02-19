package authorization

import "platform/authorization/identity"

type roleCondition struct {
	allowedRoles []string
}

func (c *roleCondition) Validate(user identity.User) bool {
	for _, allowedRole := range c.allowedRoles {
		if user.InRole(allowedRole) {
			return true
		}
	}
	return false
}

func NewRoleCondition(roles ...string) identity.AuthorizationCondition {
	return &roleCondition{allowedRoles: roles}
}
