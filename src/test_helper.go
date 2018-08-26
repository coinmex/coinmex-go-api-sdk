package coinmex_api

/*
 Get a http client
 @author Bitcoin Light
 @date 2018-08-10
 @version 0.9.1
*/

func NewTestClient() *Client {
	var config Config
	config.Endpoint = "https://www.coinmex.com"
	config.ApiKey = ""
	config.SecretKey = ""
	config.Passphrase = ""
	config.TimeoutSecond = 45
	config.IsPrint = false

	client := NewClient(config)
	return client
}
