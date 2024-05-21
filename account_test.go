package deribit

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"

	asserts "github.com/stretchr/testify/assert"
)

var client *Client

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientId := os.Getenv("ClientID")
	secret := os.Getenv("Secret")
	proxy := os.Getenv("HTTP_PROXY")
	client = NewWithProxy(EndpointTest, clientId, secret, proxy)
	m.Run()
}

func TestGetPaymentResult(t *testing.T) {
	assert := asserts.New(t)
	resp, err := client.GetAccountSummary("BTC", true)
	fmt.Printf("%+v", resp)
	assert.Nil(err)
	assert.NotEqual(0, resp.MaintenanceMargin)
}

func TestSell(t *testing.T) {
	assert := asserts.New(t)
	resp, err := client.Sell("BTC-12JUN22-28000", "0.1", "market")
	assert.Nil(err)
	assert.NotNil(resp)
}

func TestGetTxLogs(t *testing.T) {
	startts := time.Date(2024, 5, 17, 8, 0, 0, 0, time.UTC).UnixMilli()
	txlogs, err := client.GetTransactionLogResult(TxLogReq{
		Currency:     "BTC",
		Query:        "delivery",
		StartTs:      startts,
		EndTs:        startts+3600000,
		Count:        100,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(txlogs)
}
