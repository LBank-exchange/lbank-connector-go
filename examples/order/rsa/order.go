package main

const (
	apiKey    = ""
	secretKey = ""
)

var client = sve.NewClient(apiKey, secretKey)

func TestCreateOrder() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{
		"symbol": "lbk_usdt",
		"size":   "1",
	}
	client.NewOrderService().CreateOrder(data)
}

func main() {
	TestCreateOrder()
}