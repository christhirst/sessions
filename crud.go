package sessions

import (
	"fmt"
	"net/http"
	"time"
)

func SessionSave(w http.ResponseWriter, r *http.Request, cookieID string) (bool, error) {

	expiration := time.Now().Add(365 * 24 * time.Hour)

	cookie := http.Cookie{Name: "goID", Value: "testing", Expires: expiration, HttpOnly: true}

	fmt.Println("iiiiiiii")
	http.SetCookie(w, &cookie)

	c, err := r.Cookie("goID")
	fmt.Println(c)

	return false, nil
}
