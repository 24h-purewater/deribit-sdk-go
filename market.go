package deribit

const (
	urlGetInstruments = "/api/v2/public/get_instruments"
	urlGetInstrument  = "/api/v2/public/get_instrument"
	urlGetOrderBook   = "/api/v2/public/get_order_book"
	urlGetIndexPrice  = "/api/v2/public/get_index_price"
)

// GetInstruments https://docs.deribit.com/?shell#public-get_instrument
func (c *Client) GetInstruments(currency, kind string, expired bool) (instrumentList []Instrument, err error) {
	var instrumentListResp InstrumentListResp
	if err = c.GetAndUnmarshalJson(urlGetInstruments, map[string]interface{}{
		"currency": currency,
		"kind":     kind,
		"expired":  expired,
	}, &instrumentListResp); err != nil {
		return instrumentList, err
	}
	return instrumentListResp.Result, err
}

func (c *Client) GetInstrument(instrumentName string) (o Instrument, err error) {
	var instrumentResp InstrumentResp
	if err = c.GetAndUnmarshalJson(urlGetInstrument, map[string]interface{}{
		"instrument_name": instrumentName,
	}, &instrumentResp); err != nil {
		return o, err
	}
	return instrumentResp.Result, err
}

func (c *Client) GetOrderBook(instrumentName string, depth int) (o OrderBook, err error) {
	var orderBookResp OrderBookResp
	if err = c.GetAndUnmarshalJson(urlGetOrderBook, map[string]interface{}{
		"instrument_name": instrumentName,
		"depth":           depth,
	}, &orderBookResp); err != nil {
		return o, err
	}
	return orderBookResp.Result, err
}

func (c *Client) GetIndexPrice(indexName string) (o IndexPriceResult, err error) {
	var indexPriceResp IndexPriceResp
	if err = c.GetAndUnmarshalJson(urlGetIndexPrice, map[string]interface{}{
		"index_name": indexName,
	}, &indexPriceResp); err != nil {
		return o, err
	}
	return indexPriceResp.Result, err
}
