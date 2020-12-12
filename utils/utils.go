package utils

import "net/http"

func GetData(c http.Cookie, URL string) (http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	req.AddCookie(&c)

	var resp = new(http.Response)
	resp, err = client.Do(req)

	return *resp, err
}
