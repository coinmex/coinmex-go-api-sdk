package coinmex_api

/*
 CoinMex Spot Api
 @author Bitcoin Light
 @date 2018-08-10
 @version 0.9.1
*/

import (
	"net/http"
)

/*
 =============================== Public API ===============================
*/

/*
  query code
*/
func (client *Client) GetCode() ([]CodeItems, error) {
	var codeItems []CodeItems
	_, err := client.Request(GET, SPOT_CODE, nil, &codeItems)
	return codeItems, err
}

/*
 query ticker
*/
func (client *Client) GetTicker(code string) (*http.Response, error) {
	response, err := client.Request(GET, SPOT_TICKER+code+"/ticker", nil, nil)
	return response, err
}

/*
 query depth
*/
func (client *Client) GetDepth(code string) (*http.Response, error) {
	response, err := client.Request(GET, SPOT_DEPTH+code+"/orderbook", nil, nil)
	return response, err
}

/*
 query deal
*/
func (client *Client) GetDeal(paginationParams PaginationParams) (*http.Response, error) {
	params := NewParams()
	params["before"] = Int2String(paginationParams.Before)
	params["after"] = Int2String(paginationParams.After)
	params["limit"] = Int2String(paginationParams.Limit)
	requestPath := BuildParams(SPOT_DEAL+paginationParams.Code+"/fills", params)
	response, err := client.Request(GET, requestPath, nil, nil)
	return response, err
}

/*
 query k-line
*/
func (client *Client) GetKline(getKlineParams GetKlineParams) (*http.Response, error) {
	params := NewParams()
	params["start"] = getKlineParams.Start
	params["end"] = getKlineParams.End
	params["type"] = getKlineParams.Atype
	requestPath := BuildParams(SPOT_KLINE+getKlineParams.Code+"/candles", params)
	response, err := client.Request(GET, requestPath, nil, nil)
	return response, err
}

/*
 query server-time
*/
func (client *Client) GetServerTime() (ServerTime, error) {
	var serverTime ServerTime
	_, err := client.Request(GET, SERVER_TIMESTAMP, nil, &serverTime)
	return serverTime, err
}

/*
 =============================== Private API ===============================
*/

/*
 query account
*/
func (client *Client) GetAccount() ([]CurrencyCodeItems, error) {
	var currencyCodeItems []CurrencyCodeItems
	_, err := client.Request(GET, SPOT_ACCOUNT, nil, &currencyCodeItems)
	return currencyCodeItems, err
}

/*
 query ledger
*/
func (client *Client) GetLedger(currencyCode string) ([]LedgerItems, error) {
	var ledgerItems []LedgerItems
	_, err := client.Request(GET, SPOT_LEDGER+currencyCode+"/ledger", nil, &ledgerItems)
	return ledgerItems, err
}

/*
 withdraw
*/
func (client *Client) Withdraw(withdrawItem WithdrawParams) (*http.Response, error) {
	params := NewParams()
	params["currencyCode"] = withdrawItem.CurrencyCode
	params["amount"] = withdrawItem.Amount
	params["address"] = withdrawItem.Address
	response, err := client.Request(POST, WITHDRAW, params, nil)
	return response, err
}

/*
 order
*/
func (client *Client) Order(takeOrderParams TakeOrderParams) (ResultItems, error) {
	var resultItems ResultItems
	_, err := client.Request(POST, SPOT_ORDER, takeOrderParams, &resultItems)
	return resultItems, err
}

/*
 revoke order by ID
*/
func (client *Client) RevokeOrder(getOrderParams GetOrderParams) (*http.Response, error) {
	params := NewParams()
	params["code"] = getOrderParams.Code
	requestPath := BuildParams(SPOT_REVOKE_ORDER+Int64ToString(getOrderParams.OrderId), params)
	response, err := client.Request(DELETE, requestPath, params, nil)
	return response, err
}

/*
 revoke orders by code
*/
func (client *Client) RevokeOrders(getOrdersParams GetOrdersParams) (*http.Response, error) {
	params := NewParams()
	params["code"] = getOrdersParams.Code
	requestPath := BuildParams(SPOT_REVOKE_ORDERS, params)
	response, err := client.Request(DELETE, requestPath, params, nil)
	return response, err
}

/*
 query order by ID
*/
func (client *Client) GetOrderInfo(getOrderParams GetOrderParams) (OrderItems, error) {
	params := NewParams()
	params["code"] = getOrderParams.Code
	requestPath := BuildParams(SPOT_ORDER_INFO+Int64ToString(getOrderParams.OrderId), params)
	var orderItems OrderItems
	_, err := client.Request(GET, requestPath, params, &orderItems)
	return orderItems, err
}

/*
 query orders by code
*/
func (client *Client) GetOrdersInfo(getOrdersParams GetOrdersParams) ([]OrderItems, error) {
	params := NewParams()
	params["code"] = getOrdersParams.Code
	params["status"] = getOrdersParams.Status
	requestPath := BuildParams(SPOT_ORDERS_INFO, params)
	var orderItems []OrderItems
	_, err := client.Request(GET, requestPath, nil, &orderItems)
	return orderItems, err
}