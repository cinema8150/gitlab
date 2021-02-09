package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab/config"
	"net/http"
)

// Method request类型
type Method string

const (
	// MethodGET Get请求
	MethodGET Method = "GET"
	// MethodPOST Post请求
	MethodPOST Method = "POST"
)

// POST 以post方式发送请求
func POST(path string, params map[string]interface{}, dataType interface{}) error {
	return request("POST", path, params, dataType)
}

// GET 以get方式发送请求
func GET(path string, params map[string]interface{}, dataType interface{}) error {
	return request("GET", path, params, dataType)
}

func request(method string, path string, params map[string]interface{}, dataType interface{}) error {

	var config = config.Shared()
	if len(config.Host) == 0 {
		return errors.New("Config gitlab host before using gitlab function, use `gitlab config --host xxx`.")
	}
	if len(config.Token) == 0 {
		println(config.Host)
		msg := "Config gitlab token before using gitlab function, use `gitlab config --token xxx`.\n" +
			"Or you can create a new personal access_token from :" +
			" https://" + config.Host + "/profile/personal_access_tokens"
		return errors.New(msg)
	}

	var query string
	for k, v := range params {
		if len(query) == 0 {
			query = "?"
		} else {
			query = query + "&"
		}
		query = fmt.Sprintf("%s%s=%v", query, k, v)
	}

	url := fmt.Sprintf("https://%s%s%s", config.Host, path, query)
	println(url)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("user-agent", "vscode-restclient")
	req.Header.Add("private-token", config.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	jsonErr := json.NewDecoder(res.Body).Decode(dataType)

	// fmt.Printf("api result :%s\n", dataType)

	return jsonErr
}
