package deribit

const (
	urlGetOpenOrdersByCurrency = "/api/v2/private/get_open_orders_by_currency"
	urlGetUserTradesByCurrency = "/api/v2/private/get_user_trades_by_currency"
	urlBuy                     = "/api/v2/private/buy"
	urlSell                    = "/api/v2/private/sell"
	urlCancel                  = "/api/v2/private/cancel"
)

//Buy https://docs.deribit.com/#private-buy
func (c *Client) Buy(instrumentName, amount, orderType string) (o BuyResult, err error) {
	var buyResult BuyResp
	if err = c.GetAndUnmarshalJson(urlBuy, map[string]interface{}{
		"instrument_name": instrumentName,
		"amount":          amount,
		"type":            orderType,
	}, &buyResult); err != nil {
		return o, err
	}
	return buyResult.Result, err
}

//Sell https://docs.deribit.com/#private-sell
func (c *Client) Sell(instrumentName, amount, orderType string) (o SellResult, err error) {
	var sellResp SellResp
	if err = c.GetAndUnmarshalJson(urlSell, map[string]interface{}{
		"instrument_name": instrumentName,
		"amount":          amount,
		"type":            orderType,
		//"price":           price,
	}, &sellResp); err != nil {
		return o, err
	}
	return sellResp.Result, err
}

//Cancel https://docs.deribit.com/#private-cancel
func (c *Client) Cancel(orderID string) (o CancelResult, err error) {
	var cancelResp CancelResp
	if err = c.GetAndUnmarshalJson(urlCancel, map[string]interface{}{
		"order_id": orderID,
	}, &cancelResp); err != nil {
		return o, err
	}
	return cancelResp.Result, err
}

//GetOpenOrdersByCurrency https://docs.deribit.com/#private-get_open_orders_by_currency
func (c *Client) GetOpenOrdersByCurrency(currency, kind, orderType string) (orderList []Order, err error) {
	var orderListResp OrderListResp
	if err = c.GetAndUnmarshalJson(urlGetOpenOrdersByCurrency, map[string]interface{}{
		"currency": currency,
		"kind":     kind,
		"type":     orderType,
	}, &orderListResp); err != nil {
		return nil, err
	}
	return orderListResp.Result, err
}

//GetUserTradesByCurrency https://docs.deribit.com/#private-get_user_trades_by_currency
func (c *Client) GetUserTradesByCurrency(currency, kind, startId, endId, sorting string, count int64, includeOld bool) (tradeList []Trade, err error) {
	var tradeListResp TradeListResp
	if err = c.GetAndUnmarshalJson(urlGetUserTradesByCurrency, map[string]interface{}{
		"currency":    currency,
		"kind":        kind,
		"start_id":    startId,
		"end_id":      endId,
		"count":       count,
		"include_old": includeOld,
		"sorting":     sorting,
	}, &tradeListResp); err != nil {
		return nil, err
	}
	return tradeListResp.Result.Trades, err
}
