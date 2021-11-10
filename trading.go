package deribit

import (
	"encoding/json"
)

const (
	urlGetOpenOrdersByCurrency = "/api/v2/private/get_open_orders_by_currency"
	urlGetUserTradesByCurrency = "/api/v2/private/get_user_trades_by_currency"
	urlBuy                     = "/api/v2/private/buy"
	urlSell                    = "/api/v2/private/sell"
	urlCancel                  = "/api/v2/private/cancel"
)

//https://docs.deribit.com/#private-buy
func (c *Client) Buy(instrumentName, amount, orderType string) (o BuyResult, err error) {
	res, err := c.Get(urlBuy, map[string]interface{}{
		"instrument_name": instrumentName,
		"amount":          amount,
		"type":            orderType,
	})
	if err != nil {
		return o, err
	}
	err = json.Unmarshal(res, &o)
	return o, err
}

//https://docs.deribit.com/#private-sell
func (c *Client) Sell(instrumentName, amount, orderType string) (o SellResult, err error) {
	res, err := c.Get(urlSell, map[string]interface{}{
		"instrument_name": instrumentName,
		"amount":          amount,
		"type":            orderType,
		//"price":           price,
	})
	if err != nil {
		return o, err
	}
	err = json.Unmarshal(res, &o)
	return o, err
}

//https://docs.deribit.com/#private-cancel
func (c *Client) Cancel(orderID string) (o CancelResult, err error) {
	res, err := c.Get(urlCancel, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return o, err
	}
	err = json.Unmarshal(res, &o)
	return o, err
}

//https://docs.deribit.com/#private-get_open_orders_by_currency
func (c *Client) GetOpenOrdersByCurrency(currency, kind, orderType string) (o OrderListResult, err error) {
	res, err := c.Get(urlGetOpenOrdersByCurrency, map[string]interface{}{
		"currency": currency,
		"kind":     kind,
		"type":     orderType,
	})
	if err != nil {
		return o, err
	}
	err = json.Unmarshal(res, &o)
	return o, err
}

//https://docs.deribit.com/#private-get_user_trades_by_currency
func (c *Client) GetUserTradesByCurrency(currency, kind, startId, endId, sorting string, count int64, includeOld bool) (o TradeListResult, err error) {
	res, err := c.Get(urlGetUserTradesByCurrency, map[string]interface{}{
		"currency":    currency,
		"kind":        kind,
		"start_id":    startId,
		"end_id":      endId,
		"count":       count,
		"include_old": includeOld,
		"sorting":     sorting,
	})
	if err != nil {
		return o, err
	}
	err = json.Unmarshal(res, &o)
	return o, err
}
