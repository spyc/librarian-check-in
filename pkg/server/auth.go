package server

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/pquerna/otp/totp"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

const sessionName = "library.attendance"

type authRequest struct {
	Passcode string `json:"passcode"`
}

type AuthHandler struct {
	SessionStore *sessions.CookieStore `inject:""`
	Logger       logrus.FieldLogger    `inject:"auth logger"`
}

func (h *AuthHandler) getSecret() string {
	pwd, _ := os.Getwd()
	filepath := path.Join(pwd, "data", "token")
	secret, _ := ioutil.ReadFile(filepath)
	return string(secret)
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, err := h.SessionStore.Get(r, sessionName)
	if err != nil {
		session, _ = h.SessionStore.New(r, sessionName)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when reading body", http.StatusUnprocessableEntity)
		return
	}

	var req authRequest
	if err := json.Unmarshal(body, &req); err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when reading body", http.StatusUnprocessableEntity)
		return
	}

	if totp.Validate(req.Passcode, h.getSecret()) {
		session.Values["auth"] = true
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
