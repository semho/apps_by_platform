package placeholder

import (
	"platform/authorization/identity"
	"platform/services"
	"strings"
)

var users = map[int]identity.User{
	1: identity.NewBasicUser(1, "Alice", "Administrator"),
	2: identity.NewBasicUser(2, "Bob"),
}

type PlaceholderUserStore struct{}

func (store *PlaceholderUserStore) GetUserByID(id int) (user identity.User, found bool) {
	user, found = users[id]
	return user, found
}

func (store *PlaceholderUserStore) GetUserByName(name string) (user identity.User, found bool) {
	for _, user = range users {
		if strings.EqualFold(user.GetDisplayName(), name) {
			return user, true
		}
	}
	return
}

func RegisterPlaceholderUserStore() {
	err := services.AddSingleton(func() identity.UserStore {
		return &PlaceholderUserStore{}
	})
	if err != nil {
		panic(err)
	}

}
