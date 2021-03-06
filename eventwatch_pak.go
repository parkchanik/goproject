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
	

// NewSendaboxFilterer creates a new log filterer instance of Sendabox, bound to a specific deployed contract.
	/*
	func NewSendaboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SendaboxFilterer, error) {
		contract, err := bindSendabox(address, nil, nil, filterer)
		if err != nil {
			return nil, err
		}
		return &SendaboxFilterer{contract: contract}, nil
	}
	*/
	query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
	}
	
	logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.TODO(), query, logs)
    if err != nil {
		fmt.Println("err-sub")
		log.Fatal(err)
		
	}
	
	sendboxfilter , err := sendabox.NewSendaboxFilterer(contractAddress , sub)
	if err != nil {
		fmt.Println("err-sendboxfilter")
		log.Fatal(err)
		
	}

	
	sendboxevent := make(chan sendabox.SendaboxEvSendABoxEvent)

	eventSubcription , err := sendboxfilter.WatchEvSendABoxEvent( , sendboxevent , nil , nil)
	if err != nil {
		fmt.Println("err-eventSubcription")
		log.Fatal(err)
		
	}

	
	//func (_Sendabox *SendaboxFilterer) WatchEvSendABoxEvent(opts *bind.WatchOpts, sink chan<- *SendaboxEvSendABoxEvent, _box_idx []*big.Int, _sender []common.Address) (event.Subscription, error) {

   /*
    contractAbi, err := abi.JSON(strings.NewReader(string(sendabox.SendaboxABI)))
    if err != nil {
        log.Fatal(err)
    }
    */
    for {
        select {
        case err := <-sub.Err():
            fmt.Println("sdd")
			log.Fatal(err)
        case vLog := <-logs:
            fmt.Println("vlog")
            fmt.Println(vLog) // pointer to event log
            fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
            fmt.Println(vLog.BlockNumber)     // 2394201
            fmt.Println(vLog.TxHash.Hex()) 

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
                fmt.Println("abierr")
                log.Fatal(err)
            }
            
            fmt.Println("event1")
            fmt.Println(event.BoxIdx)
            fmt.Println("event2")
            fmt.Println(event.Sender)
            fmt.Println(event.Value)
            fmt.Println(event.Token)
            fmt.Println(event.Message)
          
            fmt.Println("event3")
        }
    }
}