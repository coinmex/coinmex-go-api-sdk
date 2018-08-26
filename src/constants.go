package coinmex_api

/*
 constants
 @author Bitcoin Light
 @date 2018-08-10
 @version 0.9.1
*/

const (
	/*
	  http headers
	*/
	ACCESS_KEY        = "ACCESS-KEY"
	ACCESS_SIGN       = "ACCESS-SIGN"
	ACCESS_TIMESTAMP  = "ACCESS-TIMESTAMP"
	ACCESS_PASSPHRASE = "ACCESS-PASSPHRASE"

	BEFORE = "CB_BEFORE"
	AFTER  = "CB_AFTER"
	LIMIT  = "LIMIT"

	CONTENT_TYPE = "Content-Type"
	ACCEPT       = "Accept"

	//COOKIE       = "Cookie"
	//LOCALE       = "locale="

	APPLICATION_JSON      = "application/json"
	APPLICATION_JSON_UTF8 = "application/json; charset=UTF-8"

	/*
	 http methods
	*/
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"

	/*
	 others
	*/
	ResultJsonString = "Resultjsonstring"

	/*
	 unit-testing result flag
	*/
	GUnitTest = "[UnitTest]"

	/*
	 private api
	*/
	SPOT_ACCOUNT = "/api/v1/spot/ccex/account/assets"
	SPOT_LEDGER = "/api/v1/spot/ccex/account/"
	WITHDRAW = "/api/v1/spot/ccex/account/withdraw"

	SPOT_ORDER = "/api/v1/spot/ccex/orders"
	SPOT_REVOKE_ORDER = "/api/v1/spot/ccex/orders/"
	SPOT_REVOKE_ORDERS = "/api/v1/spot/ccex/orders"
	SPOT_ORDER_INFO = "/api/v1/spot/ccex/orders/"
	SPOT_ORDERS_INFO = "/api/v1/spot/ccex/orders"

	/*
	 public api
	*/
	SPOT_CODE = "/api/v1/spot/public/products"
	SPOT_TICKER = "/api/v1/spot/public/products/"
	SPOT_DEPTH = "/api/v1/spot/public/products/"
	SPOT_DEAL = "/api/v1/spot/public/products/"
	SPOT_KLINE = "/api/v1/spot/public/products/"

	SERVER_TIMESTAMP = "/api/v1/spot/public/time"

)
