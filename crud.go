package sessions

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func SessionSave(w http.ResponseWriter, r *http.Request, cookieID string, userID string) (string, error) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	id := uuid.New()
	cookie := http.Cookie{Name: cookieID, Value: id.String(), Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)

	c, err := r.Cookie("user_session")
	if err != nil {
		log.Panic(err)
		return "", errors.New("C")
	}
	fmt.Println(c)

	return id.String(), nil
}

func SessionGet(w http.ResponseWriter, r *http.Request, cookieID string) (string, error) {
	coo, err := r.Cookie(cookieID)
	if err != nil {
		log.Panic(err)
	}

	return coo.Value, err
}

func SessionExpire(w http.ResponseWriter, r *http.Request, cookieID string) (bool, error) {
	expiration := time.Now().Add(-1 * time.Hour)
	cookie := http.Cookie{Name: "user_session", Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)

	c, err := r.Cookie("user_session")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(c)

	return false, nil
}
