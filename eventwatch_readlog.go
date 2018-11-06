package main

import (
    "context"
    "fmt"
    "log"
    "strings"
    "math/big"
    //"strconv"
    "database/sql"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"

    //"github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    _ "github.com/go-sql-driver/mysql"
    sendabox "./contract" 
)

func main() {

    
    var db, err = sql.Open("mysql", "dbchanik:1q2w3e$r@tcp(localhost:3306)/aboxdb")
    if err != nil {
        fmt.Println("DB ERROR")
		log.Fatal(err)
		
    }

	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	//client, err := ethclient.Dial("ws://localhost:7545/")
    if err != nil {
        fmt.Println("err1")
		log.Fatal(err)
		
    }

    //contractAddress := common.HexToAddress("0x868eF2CfBd938ca9Ae5ddFa2b6CE2cAdd73b3c36") // local house
    //contractAddress := common.HexToAddress("0x9e51806f5D074Cff7dCea42997860D1176dEF99B") // local office
	contractAddress := common.HexToAddress("0x8bb30f3e1e5c63f30f3e41cdfaedabe1a45827b9") // ropsten
	
    query := ethereum.FilterQuery{
        //FromBlock: big.NewInt(4364359),
       // ToBlock:   big.NewInt(4364484),
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
    
    // go-ethereum/core/types/log.go
	fmt.Println("3333")
	for _, vLog := range logs {
              
		event := struct {
			//BoxIdx  *big.Int
			//Sender  common.Address
			Value   *big.Int
			Token   *big.Int
			Message string
			//Raw     types.Log // Blockchain specific contextual infos
		}{}
		
        err := contractAbi.Unpack(&event, "ev_SendABoxEvent", vLog.Data)
        if err != nil {
            log.Fatal(err)
        }

        //fmt.Println("BoxIdx")
		//fmt.Println(event.BoxIdx)
		//fmt.Println("Sender")
        //fmt.Println(event.Sender.Hex())
           /*
        fmt.Printf("Topics[0] : %s\n " , vLog.Topics[0].Hex()) //topic[0] event Signature
        fmt.Printf("Topics[1] : %s\n " , vLog.Topics[1].Hex()) //topic[1] 내가 첫번째 index 로 잡았던 event  _box_idx
        fmt.Printf("Topics[2] : %s\n " , vLog.Topics[2].Hex()) //topic[2] 내가 두번째 index 로 잡았던 event  _sender
        */
        message := string(event.Message)
        wei := event.Value.Int64()
        token := event.Token.Int64() // Int64()로 컨버팅 하지 않으면 에러 - converting argument $5 type: unsupported type big.Int, a struct
        blockhash := vLog.BlockHash.Hex()
        blocknumber := vLog.BlockNumber
        txhash := vLog.TxHash.Hex()
        txindex := vLog.TxIndex
        logindex := vLog.Index


        boxidx_big := vLog.Topics[1].Big()
        boxidx := boxidx_big.Int64()

        sender := vLog.Topics[2].String()

        fmt.Printf("BlockHash : %s\n" , blockhash) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
        fmt.Printf("BlockNumber : %d\n" , blocknumber) // 2394201
        fmt.Printf("TxHash : %s\n" , txhash)    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
        fmt.Printf("TxIndex  : %d\n" , txindex)    // 
        fmt.Printf("Index  : %d\n" , logindex)    //  i_logIndex
        fmt.Printf("Value : %d\n" , wei)
        fmt.Printf("Token : %d\n" , token)
        fmt.Printf("Message : %s\n" , message)

        fmt.Printf("boxid : %d\n" , boxidx)
        fmt.Printf("sender : %s\n" , sender)
        
        /*
        var params = [box_idx , sender
        ,  message , wei ,token
      , logIndex , transactionIndex , transactionHash , blockHash , blockNumber , id];

      var stmt = 'CALL SP_SEND_BOX(?,?,?,?,?,?,?,?,?,?,?); ';
        */
        type OutReturn struct {
            O_return         int
        }
        var outreturn OutReturn
        row := db.QueryRow("CALL SP_SEND_BOX(?,?,?,?,?,?,?,?,?,?,?);" , boxidx , sender , message , wei , token , logindex , txindex , txhash , blockhash , blocknumber , 0  )

        err = row.Scan(&outreturn.O_return)
        if err != nil {
            fmt.Println("err Scan")
            log.Fatal(err)
            
        }
    
        /*
        var topics [4]string
        for i := range vLog.Topics {
            topics[i] = vLog.Topics[i].Hex()
            fmt.Println(topics[i])
            
        }
        */  
        //fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}
	

}