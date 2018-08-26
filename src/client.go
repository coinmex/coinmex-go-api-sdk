package coinmex_api

/*
 http client, request, response
 @author Bitcoin Light
 @date 2018-08-10
 @version 0.9.1
*/

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
 Get a http client
*/
func NewClient(config Config) *Client {
	var client Client
	client.Config = config
	timeout := config.TimeoutSecond
	if timeout <= 0 {
		timeout = 30
	}
	client.HttpClient = &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	return &client
}

/*
 Send a http request to remote server and get a response data
*/
func (client *Client) Request(method string, requestPath string,
	params, result interface{}) (response *http.Response, err error) {
	config := client.Config

	// url
	url := config.Endpoint + requestPath

	// get json and bin styles request body
	var jsonBody string
	var binBody = bytes.NewReader(make([]byte, 0))
	if params != nil {
		jsonBody, binBody, err = ParseRequestParams(params)
		if err != nil {
			return response, err
		}
	}

	// get a http request
	request, err := http.NewRequest(method, url, binBody)
	if err != nil {
		return response, err
	}

	// Sign and set request headers
	timestamp := IsoTime()
	preHash := PreHashString(timestamp, method, requestPath, jsonBody)
	sign, err := Sign(preHash, config.SecretKey)
	if err != nil {
		return response, err
	}
	Headers(request, config, timestamp, sign)

	if config.IsPrint {
		printRequest(config, request, jsonBody, preHash)
	}

	// send a request to remote server, and get a response
	response, err = client.HttpClient.Do(request)
	if err != nil {
		return response, err
	}
	defer response.Body.Close()

	// get a response results and parse
	status := response.StatusCode
	message := response.Status
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response, err
	}

	if config.IsPrint {
		printResponse(status, message, body)
	}

	response.Header.Add(ResultJsonString, string(body))

	if status >= 200 && status < 300 {
		if body != nil && result != nil {
			err := JsonBytes2Struct(body, result)
			if err != nil {
				return response, err
			}
		}
		return response, nil
	} else if status == 400 || status == 401 || status == 500 {
		if body != nil {
			var responseMessage ResponseMessage
			err := JsonBytes2Struct(body, &responseMessage)
			if err != nil {
				return response, err
			}
			message = strconv.Itoa(status) + " " + responseMessage.Message
		}
		return response, errors.New(message)
	} else {
		return response, errors.New(message)
	}
	return response, nil
}

func printRequest(config Config, request *http.Request, body string, preHash string) {
	if config.SecretKey != "" {
		fmt.Println("  SecretKey: " + config.SecretKey)
	}
	fmt.Println("  Request(" + IsoTime() + "):")
	fmt.Println("\tUrl: " + request.URL.String())
	fmt.Println("\tMethod: " + strings.ToUpper(request.Method))
	if len(request.Header) > 0 {
		fmt.Println("\tHeaders: ")
		for k, v := range request.Header {
			fmt.Println("\t\t" + k + ": " + v[0])
		}
	}
	fmt.Println("\tBody: " + body)
	if preHash != "" {
		fmt.Println("  PreHash: " + preHash)
	}
}

func printResponse(status int, message string, body []byte) {
	fmt.Println("  Response(" + IsoTime() + "):")
	statusString := strconv.Itoa(status)
	message = strings.Replace(message, statusString, "", -1)
	message = strings.Trim(message, " ")
	fmt.Println("\tStatus: " + statusString)
	fmt.Println("\tMessage: " + message)
	fmt.Println("\tBody: " + string(body))
}
