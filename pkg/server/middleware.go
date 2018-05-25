package server

import (
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthMiddleware struct {
	SessionStore *sessions.CookieStore `inject:""`
	Logger       logrus.FieldLogger    `inject:"auth logger"`
	Handler      http.Handler
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	session, err := m.SessionStore.Get(r, sessionName)
	if err != nil {
		m.Logger.Error(err)
		http.Error(w, "error in getting session", http.StatusInternalServerError)
		return
	}
	if session.IsNew {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, found := session.Values["auth"]
	if !found {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	m.Handler.ServeHTTP(w, r)
}
