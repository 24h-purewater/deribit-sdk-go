package deribit

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	asserts "github.com/stretchr/testify/assert"
)

func loadTestClient() *Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientId := os.Getenv("ClientID")
	secret := os.Getenv("Secret")
	return New(EndpointTest,clientId,secret)
}

func TestGetPaymentResult(t *testing.T) {
	assert := asserts.New(t)
	client := loadTestClient()
	resp, err := client.GetAccountSummary("BTC", true)
	assert.Nil(err)
	assert.NotEqual(0,resp.MaintenanceMargin)
}


func TestSell(t *testing.T) {
	assert := asserts.New(t)
	client := loadTestClient()
	resp, err := client.Sell("BTC-12JUN22-28000","0.1", "market")
	assert.Nil(err)
	assert.NotNil(resp)
}