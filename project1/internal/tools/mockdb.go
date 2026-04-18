package tools

import {
	"time"
}

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Jotaro": {
		AuthToken: "123456",
		Username: "Jotaro".
	},
	"Dio": {
		AuthToken: "666FFF",
		Username: "Dio",
	},
	"Zeppelin": {
		AuthToken: "ZEPZEP",
		Username: "Zeppelin",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Jotaro": {
		Coins: 10,
		Username: "Jotaro",
	},
	"Dio": {
		Coins: 42069,
		Username: "Dio",
	},
	"Zeppelin": {
		Coins: 9999999,
		Username: "Zeppelin",
	},
}

func (d, *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d, *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}