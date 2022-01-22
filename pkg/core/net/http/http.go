/*
Copyright 2021 The NitroCI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package http

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	extendedJson "nitroci/pkg/core/encoding/json"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
	return nil
}

type HttpResult struct {
	StatusCode int
	Body       *string
}

func (httpResult HttpResult) ToJson(target interface{}) error {
	_, err := extendedJson.IsJSON(*httpResult.Body)
	if err != nil {
		return err
	}
	body := httpResult.Body
	return json.NewDecoder(strings.NewReader(*body)).Decode(target)
}

func HttpGet(url string, username string, password string) (httpResult *HttpResult, err error) {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(bodyBytes)
	httpResult = &HttpResult{
		StatusCode: resp.StatusCode,
		Body:       &body,
	}
	return httpResult, nil
}
