package session

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type session struct {
	username string
	expiry   time.Time
}

var sessions = map[string]session{}

func CreateSession(writer http.ResponseWriter, user string) {
	token := uuid.NewString()
	expiry := time.Now().Add(10 * time.Minute)
	sessions[token] = session{
		username: user,
		expiry:   expiry,
	}
	http.SetCookie(writer, &http.Cookie{
		Name:     "ssauth",
		Value:    token,
		Expires:  expiry,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
}

func RemoveSession(writer http.ResponseWriter, request *http.Request) error {
	cookie, err := request.Cookie("ssauth")
	if err != nil {
		return err
	}
	delete(sessions, cookie.Value)
	// Client deletes cookie if expiry is past
	cookie.Expires = time.Now().AddDate(-1, 0, 0)
	http.SetCookie(writer, cookie)
	return nil
}

func IsSessionValid(request *http.Request) bool {
	cookie, err := request.Cookie("ssauth")
	if err != nil {
		return false
	}
	session, ok := sessions[cookie.Value]
	if !ok {
		return false
	}
	if session.expiry.Before(time.Now()) {
		return false
	}
	return true
}
