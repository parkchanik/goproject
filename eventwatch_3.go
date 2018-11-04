package main

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    //"github.com/miguelmota/go-web3-example/greeter"
    "log"
)

func main() {
    //client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	client, err := ethclient.Dial("ws://localhost:7545")
    if err != nil {
        log.Fatal(err)
    }

    greeterAddress := "0x868eF2CfBd938ca9Ae5ddFa2b6CE2cAdd73b3c36"
    priv := "abcdbcf6bdc3a8e57f311a2b4f513c25b20e3ad4606486d7a927d8074872cefg"

    key, err := crypto.HexToECDSA(priv)

    contractAddress := common.HexToAddress(greeterAddress)
    greeterClient, err := greeter.NewGreeter(contractAddress, client)

    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(key)

    // not sure why I have to set this when using testrpc
    // var nonce int64 = 0
    // auth.Nonce = big.NewInt(nonce)

    tx, err := greeterClient.Greet(auth, "hello")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Pending TX: 0x%x\n", tx.Hash())

    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    var ch = make(chan types.Log)
    ctx := context.Background()

    sub, err := client.SubscribeFilterLogs(ctx, query, ch)

    if err != nil {
        log.Println("Subscribe:", err)
        return
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case log := <-ch:
            fmt.Println("Log:", log)
        }
    }

}