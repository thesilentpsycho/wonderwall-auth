package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"bitbucket.org/libertywireless/wonderwall-auth/config"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

func BuildRequestUrl(targetUrl string, params map[string]interface{}) (string, error) {
	requestURL, err := url.ParseRequestURI(targetUrl)

	if err != nil {
		return "", err
	}

	query := requestURL.Query()

	for key, val := range params {
		param, ok := val.(string)

		if ok {
			query.Add(key, param)
		} else {
			tokens := val.([]string)
			for index := 0; index < len(tokens); index++ {
				query.Add(key, tokens[index])
			}
		}
	}

	requestURL.RawQuery = query.Encode()
	return requestURL.String(), nil
}

func Get(url string, headers map[string]interface{}, params map[string]interface{}) (*http.Response, error) {
	requestURL, urlError := BuildRequestUrl(url, params)

	if urlError != nil {
		return nil, urlError
	}

	client := http.Client{
		Timeout: config.GetConfig().DefaultTimeout,
	}

	req, err := http.NewRequest("GET", requestURL, nil)

	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		param, _ := val.(string)
		req.Header.Add(key, param)
	}

	return client.Do(req)

}

func Put(url string, headers map[string]interface{}, data interface{}) (*http.Response, error) {
	targetURL := url

	requestBody, bodyErr := json.Marshal(data)
	if bodyErr != nil {
		return nil, bodyErr
	}

	client := http.Client{
		Timeout: config.GetConfig().DefaultTimeout,
	}

	req, err := http.NewRequest("PUT", targetURL, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		param, _ := val.(string)
		req.Header.Add(key, param)
	}

	return client.Do(req)
}

func Unmarshal(r io.ReadCloser, s interface{}) error {
	err := json.NewDecoder(r).Decode(&s)

	if err != nil {
		return err
	}
	return nil
}
