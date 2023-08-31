package worker

import (
	"github.com/b52-unofficial/TXhero-backend/config"
	"github.com/b52-unofficial/TXhero-backend/stateManager/common/constant"
	"github.com/b52-unofficial/TXhero-backend/stateManager/common/logger"
	"github.com/b52-unofficial/TXhero-backend/stateManager/data"
	"github.com/b52-unofficial/TXhero-backend/stateManager/service/aclManager"
	"github.com/b52-unofficial/TXhero-backend/stateManager/service/eth"
	"github.com/go-co-op/gocron"
	"math/rand"
	"time"
)

func anotherTask() {
	logger.Log.Debug("I am another task. I will execute regardless of the other task's failure.")
}

// Unconfirmed Tx를 확인해 confirm된 Tx정보를 DB에 update하는 Job
func syncTxConfirmed() {
	logger.Log.Debug("syncTxConfirmed Job Start")

	//DB에서 unconfirmed인 tx들을 가져와서 확인
	unconfirmedTxs, err := data.GetTransactionDataByStatus(constant.TX_STATUS_UNCONFIRMED)
	if err != nil {
		logger.Log.Error(err)
		return
	}

	if unconfirmedTxs == nil {
		logger.Log.Debug("No unconfirmed txs")
		return
	}

	//unconfirmedTxs를 순회하면서 처리
	logger.Log.Debug("unconfirmed txs: ", len(unconfirmedTxs))
	for _, tx := range unconfirmedTxs {
		//eth에서 tx가 confirm되었는지 확인
		isConfirmed, gasFee := eth.CheckTxConfirmed(tx.TxHash)
		if isConfirmed {
			//confirm되었다면 DB에 update
			data.UpdateTxConfirmed(tx.TxHash, gasFee)
		}
	}
}

// TODO 이건 Mocking임
// 다음 라운드 bid 셋업
func setNextRoundWinningBid() {
	logger.Log.Debug("setNextRoundWinningBid Job Start")
	//현재 라운드의 winning bid를 계산하기 위해 컨트랙트를 뒤짐
	eth.GetBidInfo()

	//가져온 정보로 쿵짝쿵짝해서 계산 (Mocking)
	//대충 0.05~1 사이의 랜덤한 숫자를 뽑아서 topBid으로 설정
	rand.Seed(time.Now().UnixNano())
	topBidInfo := &data.BidInfo{TopBid: 0.05 + rand.Float64()*(1-0.05), BuilderId: 1}

	//다음 라운드의 winning bid 정보를 DB에 Insert 일단 Happy case만 가정
	data.InsertNextBidInfo(topBidInfo)

	//ACL Manager 트리거
	aclManager.TriggerAclManager(topBidInfo)
}

// 이번 라운드 보상 정산하는 Job
func distributeRoundRewards() {

}

// TODO 추후 삭제
func forTest() {
	logger.Log.Debug("for test")
	eth.CheckTxConfirmed("0xd0f15e4eaef7d2c7b9de48baec93dfa91cfd449852b916b9e3e8908b5c495064")
}

func RegisterCron() {
	conf := config.GetConfig()
	//TEST
	//forTest()

	scheduler := gocron.NewScheduler(time.UTC)

	//syncTxConfirmed job 등록
	scheduler.Cron(conf.Job.SyncTx).Do(syncTxConfirmed)
	scheduler.Cron(conf.Job.NextRoundWinningBid).Do(setNextRoundWinningBid)
	scheduler.Every(5).Seconds().Do(anotherTask)

	scheduler.StartBlocking()
}
