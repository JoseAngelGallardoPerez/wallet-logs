package services

import (
	"github.com/Confialink/wallet-logs/internal/srvdiscovery"
	"context"
	"net/http"

	pb "github.com/Confialink/wallet-users/rpc/proto/users"
)

type RpcUsers struct {
}

func (rpc *RpcUsers) GetByUIDs(uids []string) ([]*pb.User, error) {
	client, err := rpc.getClient()
	if err != nil {
		return nil, err
	}

	req := pb.Request{UIDs: uids}
	resp, err := client.GetByUIDs(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}

func (rpc *RpcUsers) getClient() (pb.UserHandler, error) {
	usersUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameUsers)
	if nil != err {
		return nil, err
	}
	return pb.NewUserHandlerProtobufClient(usersUrl.String(), http.DefaultClient), nil
}
