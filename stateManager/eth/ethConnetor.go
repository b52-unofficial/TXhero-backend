package eth

import (
	"context"
	"fmt"
	"github.com/b52-unofficial/TXhero-backend/config"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckTxConfirmed(txHex string) bool {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Rpc)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	txHash := common.HexToHash(txHex)
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatalf("Failed to retrieve transaction: %v", err)
	}

	if isPending {
		fmt.Println("Transaction is still pending")
		return false
	} else {
		fmt.Println("Transaction is confirmed")
		fmt.Println(tx.Gas())
		// pring block number
		fmt.Println(tx.GasPrice())
		return true
		// Additional logic to check the number of confirmations can also be added here.
	}
}
