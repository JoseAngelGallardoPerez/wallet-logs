package services

import "github.com/Confialink/wallet-logs/internal/logs/config/logs"

var logsFiller *LogsFiller
var rpcUsers *RpcUsers

func GetLogsFiller() *LogsFiller {
	if logsFiller == nil {
		logsFiller = &LogsFiller{logs.Logger.New("Service", "LogsFiller")}
	}

	return logsFiller
}

func GetRpcUsers() *RpcUsers {
	if rpcUsers == nil {
		rpcUsers = &RpcUsers{}
	}

	return rpcUsers
}
