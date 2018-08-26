package coinmex_api

/*
 examples
 @author Bitcoin Light
 @date 2018-08-10
 @version 0.9.1
*/

import (
	"testing"
)

const (
	code = "ct_usdt"
	currencyCode = "ct"
	withdrawCode = "ddd"
	amount = "15"
	address = "0xbfeb8b30a947749dcd4da32b9372a65b37042060"
	orderId = 21042365534240
	status = "open"
	atype = "limit"
	price = "10"
	side = "sell"
	size = "1"
	start = "1534132800000"
	end = "1534150800000"
	klineType = "1hour"
	before = 916300
	after = 916320
	limit = 15
)

/*
 =============================== Public API ===============================
*/

func TestClient_GetCode(t *testing.T) {
	response, err := NewTestClient().GetCode()
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+" CoinMex's all codes: ", response)
}

func TestClient_GetTicker(t *testing.T) {
	response, err := NewTestClient().GetTicker(code)
	result := GetResponseJsonString(response)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+code+" ticker: ", result)
}

func TestClient_GetDepth(t *testing.T) {
	response, err := NewTestClient().GetDepth(code)
	result := GetResponseJsonString(response)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+code+" orderbook: ", result)
}

func TestClient_GetDeal(t *testing.T) {
	var paginationParams PaginationParams
	paginationParams.Code = code
	paginationParams.Before = before
	paginationParams.After = after
	paginationParams.Limit = limit
	response, err := NewTestClient().GetDeal(paginationParams)
	result := GetResponseJsonString(response)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+code+" deals: ", result)
}

func TestClient_GetKline(t *testing.T) {
	var getKlineParams GetKlineParams
	getKlineParams.Code = code
	getKlineParams.Start = start
	getKlineParams.End = end
	getKlineParams.Atype = klineType
	response, err := NewTestClient().GetKline(getKlineParams)
	result := GetResponseJsonString(response)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+getKlineParams.Code+" candles: ", result)
}

func TestGetServerTime(t *testing.T) {
	response, err := NewTestClient().GetServerTime()
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+" CoinMex's server time: ", response)
}

/*
 =============================== Private API ===============================
*/

func TestClient_GetAccount(t *testing.T) {
	response, err := NewTestClient().GetAccount()
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+" CoinMex Spot Account: ", response)
}

func TestClient_GetLedger(t *testing.T) {
	response, err := NewTestClient().GetLedger(currencyCode)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+currencyCode+" ledger: ", response)
}

func TestClient_Withdraw(t *testing.T) {
	var withdrawParams WithdrawParams
	withdrawParams.CurrencyCode = withdrawCode
	withdrawParams.Amount = amount
	withdrawParams.Address = address
	response, err := NewTestClient().Withdraw(withdrawParams)
	result := GetResponseJsonString(response)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+withdrawParams.CurrencyCode+" withdraw: ", result)
}

func TestClient_Order(t *testing.T) {
	var takeOrderParams TakeOrderParams
	takeOrderParams.Code = code
	takeOrderParams.Atype = atype
	takeOrderParams.Price = price
	takeOrderParams.Side = side
	takeOrderParams.Size = size
	response, err := NewTestClient().Order(takeOrderParams)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+" take order: ", response)
}

func TestClient_GetOrderInfo(t *testing.T) {
	var getOrderParams GetOrderParams
	getOrderParams.Code = code
	getOrderParams.OrderId = orderId
	response, err := NewTestClient().GetOrderInfo(getOrderParams)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+" order info: ", response)
}

func TestClient_GetOrdersInfo(t *testing.T) {
	var getOrdersParams GetOrdersParams
	getOrdersParams.Code = code
	getOrdersParams.Status = status
	response, err := NewTestClient().GetOrdersInfo(getOrdersParams)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+" orders info: ", response)
}

func TestClient_RevokeOrder(t *testing.T) {
	var getOrderParams GetOrderParams
	getOrderParams.Code = code
	getOrderParams.OrderId = orderId
	response, err := NewTestClient().RevokeOrder(getOrderParams)
	result := GetResponseJsonString(response)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+" revoke order: ", result)
}

func TestClient_RevokeOrders(t *testing.T) {
	var getOrdersParams GetOrdersParams
	getOrdersParams.Code = code
	getOrdersParams.Status = status
	response, err := NewTestClient().RevokeOrders(getOrdersParams)
	result := GetResponseJsonString(response)
	if err != nil {
		t.Error(err)
	}
	FlagPrintln(GUnitTest+" revoke orders: ", result)
}