package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gmgale/AdventofCode2020/secrets"
)

// getData sends a GET request to server
// and returns the response body as a string.
func getData(URL string) (string, error) {

	// Errors are propegated to caller for brevity
	// (task specific), change for production code.
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	req.AddCookie(&secrets.MyCookie)

	resp := new(http.Response)
	body := new([]byte)

	resp, err = client.Do(req)
	*body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return string(*body), err
}
