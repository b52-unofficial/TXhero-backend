package contractApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/b52-unofficial/TXhero-backend/config"
	"github.com/b52-unofficial/TXhero-backend/stateManager/common/logger"
	"github.com/b52-unofficial/TXhero-backend/stateManager/data"
	"net/http"
)

type ReqData struct {
	UserRewardInfo []*data.RewardInfo `json:"userRewardInfo"`
}

func RequestUpdateRewardData(rewardData []*data.RewardInfo) error {
	logger.Log.Debugf("RequestUpdateRewardData: %v", len(rewardData))
	conf := config.GetConfig()

	data := ReqData{
		UserRewardInfo: rewardData,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		logger.Log.Errorf("Failed to marshal rewardData: %v", err)
		return err
	}

	logger.Log.Debug("RequestUpdateRewardData: ", string(payloadBytes))

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", conf.Endpoints.ContractApi+"/reward", body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	http.DefaultClient.Do(req)

	return nil

}

func ApiServerHealthCheck() bool {
	conf := config.GetConfig()

	resp, err := http.Get(conf.Endpoints.ContractApi + "/health")
	if err != nil {
		logger.Log.Errorf("Failed to connect to the Ethereum client: %v", err)
		return false
	}

	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}
