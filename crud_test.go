package sessions

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionGet(t *testing.T) {
	cV := "expected"
	httpRecorder := httptest.NewRecorder()
	http.SetCookie(httpRecorder, &http.Cookie{Name: cookieID, Value: cV})
	request := &http.Request{Header: http.Header{"Cookie": httpRecorder.HeaderMap["Set-Cookie"]}}
	cr, err := SessionGet(httpRecorder, request, cookieID)
	if err != nil {
		t.Errorf("Getting client failed: %+v", err)
	}
	if cr != cV {
		t.Error(cr)
	}

}

func TestSessionSave(t *testing.T) {
	httpRecorder := httptest.NewRecorder()
	request := &http.Request{}
	s, err := SessionSave(httpRecorder, request, cookieID)
	if err != nil {
		t.Errorf("Getting client failed: %+v", err)
	}
	cookie := httpRecorder.Result().Cookies()[0].Value
	if cookie != s.Value {
		t.Errorf("Getting client failed: %+v", err)
		fmt.Println(cookie)
		fmt.Println(s.Value)
	}
}
