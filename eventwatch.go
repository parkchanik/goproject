package main

import (
    "context"
    "fmt"
    "log"
    "strings"
    "math/big"
    "database/sql"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"

    sendabox "./contract" 
    "./db"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	//client, err := ethclient.Dial("ws://localhost:8545/ws")
    checkErr(err)

    //contractAddress := common.HexToAddress("0x868eF2CfBd938ca9Ae5ddFa2b6CE2cAdd73b3c36") // local
    //contractAddress := common.HexToAddress("0x9e51806f5d074cff7dcea42997860d1176def99b") // local office
	contractAddress := common.HexToAddress("0x8bb30f3e1e5c63f30f3e41cdfaedabe1a45827b9") // ropsten
    

    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }
	
    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.TODO(), query, logs)
    checkErr(err)
    
    
    contractAbi, err := abi.JSON(strings.NewReader(string(sendabox.SendaboxABI)))
    checkErr(err)
    
    for {
        select {
        case err := <-sub.Err():
            fmt.Println("sdd")
			log.Fatal(err)
        case vLog := <-logs:
            /*
            fmt.Println("vlog")
            fmt.Println(vLog) // pointer to event log
            fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
            fmt.Println(vLog.BlockNumber)     // 2394201
            fmt.Println(vLog.TxHash.Hex()) 
            */
            fmt.Printf("BlockHash : %s\n" , vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
            fmt.Printf("BlockNumber : %d\n" , vLog.BlockNumber) // 2394201
            fmt.Printf("TxHash : %s\n" , vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
    
            fmt.Printf("TxIndex  : %d\n" , vLog.TxIndex)    // 
            fmt.Printf("Index  : %d\n" , vLog.Index)    //  i_logIndex
            
            event := struct {
                //BoxIdx  *big.Int
                //Sender  common.Address
                Value   *big.Int
                Token   *big.Int
                Message string
                Raw     types.Log // Blockchain specific contextual infos
            }{}
            err := contractAbi.Unpack(&event, "ev_SendABoxEvent", vLog.Data)
            checkErr(err)
            
            
        
            fmt.Printf("Value : %d\n" , event.Value)
            fmt.Printf("Token : %d\n" , event.Token)
            fmt.Printf("Message : %s\n" , event.Message)
            
        /*
        var params = [box_idx , sender
        ,  message , wei ,token
      , logIndex , transactionIndex , transactionHash , blockHash , blockNumber , id];

      var stmt = 'CALL SP_SEND_BOX(?,?,?,?,?,?,?,?,?,?,?); ';
        */
        //fmt.Printf("Topics[0] : %s\n " , vLog.Topics[0].Hex()) //topic[0] event Signature
        //fmt.Printf("Topics[1] : %s\n " , vLog.Topics[1].Hex()) //topic[1] 내가 첫번째 index 로 잡았던 event  _box_idx
        //fmt.Printf("Topics[2] : %s\n " , vLog.Topics[2]) //topic[2] 내가 두번째 index 로 잡았던 event  _sender
       

            event_boxidx := vLog.Topics[1].Hex()
            event_sender := vLog.Topics[2].Hex()

            fmt.Printf("boxid : %d\n" , event_boxidx)
            fmt.Printf("sender : %s\n" , event_sender)

            stmt, err := db.Prepare("insert into user (username, password, first_name, middle_name, last_name, email, mobile_phone, login_attempt, remote_address, active_status) values(?,?,?,?,?,?,?,?,?,?);")
            checkErr(err)

            _, err = stmt.Exec(Username, string(HashedPassword), FirstName, MiddleName, LastName, Email, MobilePhone, LoginAttempt, RemoteAddress, ActiveStatus)
            checkErr(err)

            stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
            checkErr(err)

            res, err := stmt.Exec("yundream", "software", "2017-06-10")
            checkErr(err)
            tx, err := db.Begin()
            if err != nil {
                log.Fatal(err)
            }
            defer tx.Rollback()
            stmt, err := tx.Prepare("INSERT INTO foo VALUES (?)")
            if err != nil {
                log.Fatal(err)
            }
            defer stmt.Close() // danger!
            for i := 0; i < 10; i++ {
                _, err = stmt.Exec(i)
                if err != nil {
                    log.Fatal(err)
                }
            }
            err = tx.Commit()
            if err != nil {
                log.Fatal(err)
            }
            // stmt.Close() runs her

            defer stmt.Close()
        
        
        }
    }
}