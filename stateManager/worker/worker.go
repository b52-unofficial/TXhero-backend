package worker

import (
	"fmt"
	"github.com/b52-unofficial/TXhero-backend/stateManager/eth"
	"github.com/go-co-op/gocron"
	"time"
)

func task1() {
	fmt.Println("Task 1:", time.Now())
}

func task2() {
	fmt.Println("Task 2:", time.Now())
}

func syncTxConfirmed() {
	//
	eth.CheckTxConfirmed("0xefea99a327c18308d44fccba0fc43f5e8e7fa20267b808f44e192459e7de369a")
}

// RegisterCron 함수는 주기적으로 실행할 작업을 등록합니다.
func RegisterCron() {
	syncTxConfirmed()

	// 새로운 스케쥴러 생성
	s := gocron.NewScheduler(time.UTC)

	// 첫 번째 작업: 매 5초마다 실행
	s.Every(5).Second().Do(task1)

	// 두 번째 작업: 매 30초마다 실행
	s.Every(10).Second().Do(task2)

	fmt.Println("Cron registered")

	// 모든 작업 실행 시작
	s.StartBlocking()
}
