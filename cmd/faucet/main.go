package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/mitchellh/mapstructure"
	"github.com/planetmint/faucet/config"
	"github.com/planetmint/faucet/types"
)

func sendFunds(query url.Values) (err error) {
	var result types.SendParams
	err = mapstructure.Decode(query, &result)
	if err != nil {
		return
	}
	// Did someone set address in request multiple times?
	if len(result.Address) != 1 {
		err = errors.New("wrong number of receiving addresses")
		return
	}
	// Is address in configuration file bech32?
	sendingAddress := config.GetConfig().Address
	addr0, err := sdk.AccAddressFromBech32(sendingAddress)
	if err != nil {
		return
	}
	// Is address in request bech32?
	receivingAddress := result.Address[0]
	addr1, err := sdk.AccAddressFromBech32(receivingAddress)
	if err != nil {
		return
	}
	// Don't do that!
	if receivingAddress == sendingAddress {
		err = errors.New("sending and receiving address are the same")
		return
	}
	// Create 'bank send' message
	coin := sdk.NewCoins(sdk.NewInt64Coin("plmnt", 100))
	_ = banktypes.NewMsgSend(addr0, addr1, coin)
	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Are we supposed to send funds?
	query := r.URL.Query()
	var err error
	if len(query) != 0 {
		err = sendFunds(query)
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	// Show index.html
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, config.GetConfig())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
}

func main() {
	// Initialize Cosmos SDK
	accountAddressPrefix := "plmnt"
	accountPubKeyPrefix := accountAddressPrefix + "pub"
	sdkconfig := sdk.GetConfig()
	sdkconfig.SetBech32PrefixForAccount(accountAddressPrefix, accountPubKeyPrefix)
	// Load our configuration file
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("fatal error loading config file: %s", err)
	}
	serviceBind := config.GetString("service-bind")
	servicePort := config.GetInt("service-port")
	// Start our service
	log.Printf("Listening on '%s:%d' ...", serviceBind, servicePort)
	http.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", serviceBind, servicePort), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
