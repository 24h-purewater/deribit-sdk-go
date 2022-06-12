package deribit

import (
	"strconv"
)

const (
	urlGetPositions      = "/api/v2/private/get_positions"
	urlGetPosition       = "/api/v2/private/get_position"
	urlGetTransactionLog = "/api/v2/private/get_transaction_log"
	urlGetAccountSummary = "/api/v2/private/get_account_summary"
)

// GetPositions https://docs.deribit.com/#private-get_positions
func (c *Client) GetPositions(currency, kind string) (positionList []PositionListResult, err error) {
	var positionResp PositionListResp
	if err = c.GetAndUnmarshalJson(urlGetPositions, map[string]interface{}{
		"currency": currency,
		"kind":     kind,
	}, &positionResp); err != nil {
		return positionList, err
	}
	return positionResp.Result, err
}

func (c *Client) GetPosition(instrumentName string) (o GetPositionResult, err error) {
	var getPositionResp GetPositionResp
	if err = c.GetAndUnmarshalJson(urlGetPosition, map[string]interface{}{
		"instrument_name": instrumentName,
	}, &getPositionResp); err != nil {
		return o, err
	}
	return getPositionResp.Result, err
}

func (c *Client) GetTransactionLogResult(currency, query string, startTs, endTs int64) (logList []TransactionLog, err error) {
	var getTxLogResult GetTransactionLogResp
	if err = c.GetAndUnmarshalJson(urlGetTransactionLog, map[string]interface{}{
		"currency":        currency,
		"start_timestamp": strconv.Itoa(int(startTs)),
		"end_timestamp":   strconv.Itoa(int(endTs)),
		"query":           query,
	}, &getTxLogResult); err != nil {
		return logList, err
	}
	return getTxLogResult.Result.Logs, err
}

// GetAccountSummary https://docs.deribit.com/#private-get_account_summary
func (c *Client) GetAccountSummary(currency string, extended bool) (o AccountSummary, err error) {
	var getAccountSummaryResp GetAccountSummaryResp
	if err = c.GetAndUnmarshalJson(urlGetAccountSummary, map[string]interface{}{
		"currency": currency,
		"extended": extended,
	}, &getAccountSummaryResp); err != nil {
		return o, err
	}
	return getAccountSummaryResp.Result, err
}
