package aclManager

import (
	"context"
	"github.com/b52-unofficial/TXhero-backend/config"
	"github.com/b52-unofficial/TXhero-backend/dashboard/stateManager/common/logger"
	"github.com/b52-unofficial/TXhero-backend/dashboard/stateManager/data"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TriggerAclManager ACLManager 트리거 호출을 위한 함수
func TriggerAclManager(bid *data.BidInfo) {
	conf := config.GetConfig()
	var r bool

	client, err := ethclient.Dial(conf.Endpoints.AclManager)
	if err != nil {
		logger.Log.Debugf("Failed to connect to the Ethereum client: %v", err)
		return
	}

	// Remove prev peer
	client.Client().CallContext(context.Background(), &r, "admin_removeTrustedPeer", "")

	// Add new peer
	client.Client().CallContext(context.Background(), &r, "admin_addTrustedPeer", "")

	//TODO : implement

}
