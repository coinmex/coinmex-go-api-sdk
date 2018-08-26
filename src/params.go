package coinmex_api

/*
 params
 @author Bitcoin Light
 @date 2018-08-10
 @version 0.9.1
*/

import (
	"net/http"
)

type Client struct {
	Config     Config
	HttpClient *http.Client
}

type Config struct {
	// Endpoint url.
	Endpoint string
	// Api-Key provided by CoinMex.
	ApiKey string
	// Secret-Key provided by CoinMex.
	SecretKey string
	// The Passphrase provided by yourself.
	Passphrase string
	// The second of http request timeout.
	TimeoutSecond int
	// Print response message or not.
	IsPrint bool
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type GetKlineParams struct {
	Code  string `json:"code"`
	Atype string `json:"type"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type TakeOrderParams struct {
	Code  string `json:"code"`
	Atype string `json:"type"`
	Side  string `json:"side"`
	Size  string `json:"size"`
	Price string `json:"price"`
	Funds string `json:"funds"`
}

type CurrencyCodeItems struct {
	Id              int    `json:"id"`
	CurrencyCode    string `json:"currencyCode"`
	Balance         string `json:"balance"`
	Available       string `json:"available"`
	Hold            string `json:"hold"`
}

type LedgerItems struct {
	Id              int64  `json:"id"`
	Atype           string `json:"type"`
	Amount          string `json:"amount"`
	Balance         string `json:"balance"`
	CreatedDate     int64  `json:"createdDate"`
	Details         LedgerDetailsItems `json:"details"`
}

type LedgerDetailsItems struct {
	OrderId         int64  `json:"orderId"`
	ProductId       string `json:"productId"`
}

type CodeItems struct {
	Code            string `json:"code"`
	BaseCurrency    string `json:"baseCurrency"`
	BaseIncrement   string `json:"baseIncrement"`
	BaseMaxSize     string `json:"baseMaxSize"`
	BaseMinSize     string `json:"baseMinSize"`
	QuoteCurrency   string `json:"quoteCurrency"`
	QuoteIncrement  string `json:"quoteIncrement"`
	QuotePrecision  string `json:"quotePrecision"`
	VolumeIncrement string `json:"volumeIncrement"`
}

type OrderItems struct {
	AveragePrice    string `json:"averagePrice"`
	Code            string `json:"code"`
	CreatedDate     int64  `json:"createdDate"`
	FilledVolume    string `json:"filledVolume"`
	Funds           string `json:"funds"`
	OrderId         int64  `json:"orderId"`
	OrderType       string `json:"orderType"`
	Price           string `json:"price"`
	Side            string `json:"side"`
	Status          string `json:"status"`
	Volume          string `json:"volume"`
}

type ResultItems struct {
	OrderId         int64  `json:"orderId"`
	Result          bool    `json:"result"`
}


type WithdrawParams struct {
	CurrencyCode string
	Amount string
	Address string
}

type GetOrdersParams struct {
	Code string
	Status string
	CursorPage
}

type GetOrderParams struct {
	Code string
	OrderId int64
}

type PaginationParams struct {
	Code string
	CursorPage
}

type KlineParams struct {
	Code  string
	Atype string `json:"type"`
	Start string
	End   string
}

type CursorPage struct {
	// Request page before (newer) this pagination id.
	Before int
	// Request page after (older) this pagination id.
	After int
	// Number of results per request. Maximum 100. Default 100.
	Limit int
}

type ServerTime struct {
	Iso   string `json:"iso"`
	Epoch string `json:"epoch"`
}