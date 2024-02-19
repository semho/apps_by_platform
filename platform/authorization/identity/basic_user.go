package identity

import "strings"

type basicUser struct {
	Id            int
	Name          string
	Roles         []string
	Authenticated bool
}

func (bu *basicUser) GetID() int {
	return bu.Id
}

func (bu *basicUser) GetDisplayName() string {
	return bu.Name
}

func (bu *basicUser) InRole(role string) bool {
	for _, r := range bu.Roles {
		if strings.EqualFold(r, role) {
			return true
		}
	}
	return false
}

func (bu *basicUser) IsAuthenticated() bool {
	return bu.Authenticated
}

var UnauthenticatedUser User = &basicUser{}

func NewBasicUser(id int, name string, roles ...string) User {
	return &basicUser{
		Id:            id,
		Name:          name,
		Roles:         roles,
		Authenticated: true,
	}
}
