package eth

import (
	"context"
	"fmt"
	"github.com/b52-unofficial/TXhero-backend/config"
	"github.com/b52-unofficial/TXhero-backend/dashboard/stateManager/common/logger"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*type TxInfo struct {
	Hash        string
	IsConfirmed bool
	gasCost     *big.Int
}*/

// ConfirmedTx 체크해서 confirm여부와 gasCost를 리턴
func CheckTxConfirmed(txHex string) (bool, float64) {
	conf := config.GetConfig()

	// Ethereum client 연결 - 이거 싱글톤으로 하는게 좋을지?
	client, err := ethclient.Dial(conf.Eth.Rpc)
	if err != nil {
		logger.Log.Debugf("Failed to connect to the Ethereum client: %v", err)
		return false, 0
	}

	// Transaction Hash로 Transaction 조회
	txHash := common.HexToHash(txHex)
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)

	//해시를 못찾는 Case. Public으로 flush되면 pending으로 뜰거임.
	if err != nil {
		logger.Log.Debugf("Failed to retrieve transaction: %v", err)
		return false, 0
	}

	// Pending인 경우는 Public으로 flush되어 멤풀에 있지만 아직 블록에 담기지 않은 경우. 추후 이를 통해 정상 Flush여부를 확인할 수도 있음.
	// TODO flush관련 처리
	if isPending {
		logger.Log.Debugf("Transaction is still pending")
		return false, 0
	} else {
		// Pending이 아닌 경우는 블록에 담긴 경우.
		txReceipt, err := client.TransactionReceipt(context.Background(), txHash)

		// TODO Finality 관련 처리

		if err != nil {
			logger.Log.Errorf("Failed to retrieve transaction: %v", err)
			return false, 0
		}
		fmt.Println("Transaction is confirmed")
		totalCost := new(big.Int).Mul(new(big.Int).SetUint64(txReceipt.GasUsed), tx.GasPrice())

		// Wei에서 Ether로 변환
		totalCostEther := new(big.Float).SetInt(totalCost)
		totalCostEther = new(big.Float).Quo(totalCostEther, big.NewFloat(1e18))
		//to float64
		totalCostEtherFloat, _ := totalCostEther.Float64()

		logger.Log.Debugf("Total transaction cost: %f Ether\n", totalCostEtherFloat)
		return true, totalCostEtherFloat
		// Additional logic to check the number of confirmations can also be added here.
	}
}

// TODO 이건 Mocking임
func GetBidInfo() {
	//다음 라운드의 winning bid를 계산하기 위해 컨트랙트를 뒤짐
}
