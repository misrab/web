package web

import (
	// "log"
	"testing"

	// "net/url"
)

func TestRequests(t *testing.T) {
	var err error
	data := []byte(``)
	_, err = Get("http://google.com", "", "")
	if err != nil { t.Errorf(err.Error()) }
	_, err = Post("http://google.com", "", "", data)
	if err != nil { t.Errorf(err.Error()) }
	_, err = Patch("http://google.com", "", "", data)
	if err != nil { t.Errorf(err.Error()) }
	_, err = Delete("http://google.com", "", "")
	if err != nil { t.Errorf(err.Error()) }
}