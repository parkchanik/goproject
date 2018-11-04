package main

import (
    "context"
    "fmt"
    "log"
    //"math/big"
    "strings"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"

    sendabox "./contract" // for demo
)

func main() {
	//client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	client, err := ethclient.Dial("ws://localhost:7545")
	//client, err := ethclient.Dial("ws://localhost:7545")
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("111111")
    contractAddress := common.HexToAddress("0x8bb30f3e1e5c63f30f3e41cdfaedabe1a45827b9")
    query := ethereum.FilterQuery{
  
        Addresses: []common.Address{
            contractAddress,
        },
    }
	fmt.Println("222222")
    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
		fmt.Println("333333")
		log.Fatal(err)
		
    }
	fmt.Println("aaaaaa")
    contractAbi, err := abi.JSON(strings.NewReader(string(sendabox.SendaboxABI)))
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("bbbbbb")
    for _, vLog := range logs {
        fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
        fmt.Println(vLog.BlockNumber)     // 2394201
        fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
		fmt.Println("cccccc")
        event := struct {
            _box_idx   int
			_sender [32]byte
			_value   int
			_token   int
			_message string
        }{}
        err := contractAbi.Unpack(&event, "ev_SendABoxEvent", vLog.Data)
        if err != nil {
            log.Fatal(err)
        }
		fmt.Println("dddddd")
        fmt.Println(event._message[:])   // foo
        fmt.Println(event._sender[:]) // bar

        var topics [4]string
        for i := range vLog.Topics {
            topics[i] = vLog.Topics[i].Hex()
        }

        fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
    }

    eventSignature := []byte("ItemSet(bytes32,bytes32)")
    hash := crypto.Keccak256Hash(eventSignature)
    fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}