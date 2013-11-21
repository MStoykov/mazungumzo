package translator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Response struct {
	action      string
	src         string
	dest        string
	Text        string
	Translation string
}

const (
	apiUrl             = "http://syslang.com?src=%s&dest=%s&text=%s&email=%s&password=%s&outformat=json"
	apiEmail    string = "fmi@golang.bg"
	apiPassword string = "g0lang3xample"
)

func fetch(url string) *http.Response {
	transport := http.Transport{
		ResponseHeaderTimeout: time.Duration(8 * time.Second),
	}

	client := http.Client{
		Transport: &transport,
	}

	response, err := client.Get(url)
	if err != nil {
		return nil
	}
	return response
}

func Translate(src, dest, text string) (string, error) {
	var (
		response Response
		result   *http.Response
	)

	result = fetch(fmt.Sprintf(apiUrl, src, dest, url.QueryEscape(text), apiEmail, apiPassword))
	if result == nil {
		return text, errors.New("Timed out")
	}

	res, err := ioutil.ReadAll(result.Body)
	result.Body.Close()
	if err != nil {
		return text, err
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return text, nil
	}

	return response.Translation, nil
}
