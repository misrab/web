/*
	Convenience methods for dealing with a web api, possible restful

	- Responses are assumed to be JSON
	- Basic auth headers built in, if username provided (not nil)
*/


package web

import (
	"bytes"

	"io/ioutil"
	"net/http"
	"net/url"
)

func Get(url, username, password string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil { return nil, err }

	if username != "" { req.SetBasicAuth(username, password) }

	client := &http.Client{}
	resp, err := client.Do(req)
	// resp, err := http.Get(url)
	if err != nil { return nil, err }
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// data e.g. url.Values{"key": {"Value"}, "id": {"123"}}
func Post(url, username, password string, data url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
	if err != nil { return nil, err }

	if username != "" { req.SetBasicAuth(username, password) }

	client := &http.Client{}
	resp, err := client.Do(req)
	// resp, err := http.PostForm("http://example.com/form", *data)
	if err != nil { return nil, err }
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Patch(url, username, password string, data url.Values) ([]byte, error) {
	req, err := http.NewRequest("PATCH", url, bytes.NewBufferString(data.Encode()))
	if err != nil { return nil, err }

	if username != "" { req.SetBasicAuth(username, password) }

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil { return nil, err }
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Delete(url, username, password string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil { return nil, err }

	if username != "" { req.SetBasicAuth(username, password) }

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil { return nil, err }
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

