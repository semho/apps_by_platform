package sessions

import (
	"context"
	gorilla "github.com/gorilla/sessions"
	"platform/services"
)

const SESSION__CONTEXT_KEY string = "platform_session"

type Session interface {
	GetValue(key string) any
	GetValueDefault(key string, defVal any) any
	SetValue(key string, val any)
}

type SessionAdaptor struct {
	gSession *gorilla.Session
}

func (s SessionAdaptor) GetValue(key string) any {
	return s.gSession.Values[key]
}

func (s SessionAdaptor) GetValueDefault(key string, defVal any) any {
	if val, ok := s.gSession.Values[key]; ok {
		return val
	}
	return defVal
}

func (s SessionAdaptor) SetValue(key string, val any) {
	switch typedVal := val.(type) {
	case nil:
		s.gSession.Values[key] = nil
	case int, float64, bool, string:
		s.gSession.Values[key] = typedVal
	default:
		panic("Sessions only support int, float64, bool, and string values")
	}
}

func RegisterSessionService() {
	err := services.AddScoped(func(c context.Context) Session {
		val := c.Value(SESSION__CONTEXT_KEY)
		s, ok := val.(*gorilla.Session)
		if !ok {
			panic("Cannot get session from context")
		}
		return &SessionAdaptor{gSession: s}
	})
	if err != nil {
		panic(err)
	}
}
