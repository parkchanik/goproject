package main

import (
    "context"
    "fmt"
    "log"
    "strings"
    "math/big"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"

    sendabox "./contract" 
)

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	//client, err := ethclient.Dial("ws://localhost:7545")
    if err != nil {
		log.Fatal(err)
		fmt.Println("err1")
    }

	//contractAddress := common.HexToAddress("0x868eF2CfBd938ca9Ae5ddFa2b6CE2cAdd73b3c36") // local
	contractAddress := common.HexToAddress("0x8bb30f3e1e5c63f30f3e41cdfaedabe1a45827b9") // ropsten
	
    query := ethereum.FilterQuery{
  
        Addresses: []common.Address{
            contractAddress,
        },
    }

	fmt.Println("1111")
	logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
        log.Fatal(err)
	}
	
	fmt.Println("2222")
	contractAbi, err := abi.JSON(strings.NewReader(string(sendabox.SendaboxABI)))
    if err != nil {
        log.Fatal(err)
	}
	
	fmt.Println("3333")
	for _, vLog := range logs {
        fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
        fmt.Println(vLog.BlockNumber)     // 2394201
        fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		event := struct {
			BoxIdx  *big.Int
			Sender  common.Address
			Value   *big.Int
			Token   *big.Int
			Message string
			Raw     types.Log // Blockchain specific contextual infos
		}{}
		
        err := contractAbi.Unpack(&event, "ev_SendABoxEvent", vLog.Data)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("event1")
		fmt.Println(event.BoxIdx)
		fmt.Println("event2")
		fmt.Println(event.Sender)
		fmt.Println(event.Value)
		fmt.Println(event.Token)
		fmt.Println(event.Message)
			
        var topics [4]string
        for i := range vLog.Topics {
            topics[i] = vLog.Topics[i].Hex()
        }

        fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}
	

}