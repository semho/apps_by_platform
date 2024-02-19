package authorization

import (
	"context"
	"platform/authorization/identity"
	"platform/services"
	"platform/sessions"
)

const USER_SESSION_KEY string = "USER"

type SessionSignInMgr struct {
	context.Context
}

func (mgr *SessionSignInMgr) SignIn(user identity.User) error {
	session, err := mgr.getSession()
	if err == nil {
		session.SetValue(USER_SESSION_KEY, user.GetID())
	}
	return err
}

func (mgr *SessionSignInMgr) SignOut(user identity.User) error {
	session, err := mgr.getSession()
	if err == nil {
		session.SetValue(USER_SESSION_KEY, nil)
	}
	return err
}

func (mgr *SessionSignInMgr) getSession() (s sessions.Session, err error) {
	err = services.GetServiceForContext(mgr.Context, &s)
	return
}

func RegisterDefaultSignInService() {
	err := services.AddScoped(func(c context.Context) identity.SignInManager {
		return &SessionSignInMgr{Context: c}
	})
	if err != nil {
		panic(err)
	}

}
