package coinmex_api

/*
 utils
 @author Bitcoin Light
 @date 2018-08-10
 @version 0.9.1
*/

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	"fmt"
	"net/url"
)

/*
 signing a message
 using: hmac sha256 + base64
  eg:
    message = Pre_hash function comment
    secretKey = E65791902180E9EF4510DB6A77F6EBAE

  return signed string = TO6uwdqz+31SIPkd4I+9NiZGmVH74dXi+Fd5X0EzzSQ=
*/
func Sign(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

/*
 the pre hash string
  eg:
    timestamp = 2018-03-08T10:59:25.789Z
    method  = DELETE
    request_path = /orders
    body = {"code":"btc_usdt","orderId":"377454671037440"}

  return pre hash string = 2018-03-08T10:59:25.789ZDELETE/orders{"code":"btc_usdt","orderId":"377454671037440"}
*/
func PreHashString(timestamp string, method string, requestPath string, body string) string {
	return timestamp + strings.ToUpper(method) + requestPath + body
}

/*
  json string convert struct
*/
func JsonString2Struct(jsonString string, result interface{}) error {
	jsonBytes := []byte(jsonString)
	err := json.Unmarshal(jsonBytes, result)
	return err
}

/*
  json byte array convert struct
*/
func JsonBytes2Struct(jsonBytes []byte, result interface{}) error {
	err := json.Unmarshal(jsonBytes, result)
	return err
}

/*
 struct convert json string
 */
func Struct2JsonString(structt interface{}) (jsonString string, err error) {
	data, err := json.Marshal(structt)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
 Get a iso time
  eg: 2018-03-16T18:02:48.284Z
*/
func IsoTime() string {
	utcTime := time.Now().UTC()
	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	return iso
}

/*
 Get a http request body is a json string and a byte array.
*/
func ParseRequestParams(params interface{}) (string, *bytes.Reader, error) {
	if params == nil {
		return "", nil, errors.New("illegal parameter")
	}
	data, err := json.Marshal(params)
	if err != nil {
		return "", nil, errors.New("json convert string error")
	}
	jsonBody := string(data)
	binBody := bytes.NewReader(data)
	return jsonBody, binBody, nil
}

/*
 Set http request headers:
   Accept: application/json
   Content-Type: application/json; charset=UTF-8  (default)
   Cookie: locale=en_US        (English)
   ACCESS-KEY: (Your setting)
   ACCESS-SIGN: (Use your setting, auto sign and add)
   ACCESS-TIMESTAMP: (Auto add)
   ACCESS-PASSPHRASE: Your setting
*/
func Headers(request *http.Request, config Config, timestamp string, sign string) {
	request.Header.Add(ACCEPT, APPLICATION_JSON)
	request.Header.Add(CONTENT_TYPE, APPLICATION_JSON_UTF8)
	request.Header.Add(ACCESS_KEY, config.ApiKey)
	request.Header.Add(ACCESS_SIGN, sign)
	request.Header.Add(ACCESS_TIMESTAMP, timestamp)
	request.Header.Add(ACCESS_PASSPHRASE, config.Passphrase)
}

/*
 Get a new map.eg: {string:string}
 */
func NewParams() map[string]string {
	return make(map[string]string)
}

/*
 Get api requestPath + requestParams
	params := NewParams()
	params["depth"] = "200"
	params["conflated"] = "0"
	url := BuildParams("/api/futures/v3/products/BTC-USD-0310/book", params)
 return eg:/api/futures/v3/products/BTC-USD-0310/book?conflated=0&depth=200
 */
func BuildParams(requestPath string, params map[string]string) string {
	urlParams := url.Values{}
	for k := range params {
		urlParams.Add(k, params[k])
	}
	return requestPath + "?" + urlParams.Encode()
}

func GetResponseJsonString(response *http.Response) string {
	return response.Header.Get(ResultJsonString)
}

func Int64ToString(arg int64) string {
	return strconv.FormatInt(arg, 10)
}

func Int2String(arg int) string {
	return strconv.Itoa(arg)
}

func FlagPrintln(flag string, info interface{}) {
	fmt.Print(flag)
	if info != nil {
		jsonString, err := Struct2JsonString(info)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(jsonString)
	} else {
		fmt.Println("{}")
	}
}
