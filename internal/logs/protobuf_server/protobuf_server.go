package protobuf_server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Confialink/wallet-logs/internal/logs/config"
	"github.com/Confialink/wallet-logs/internal/logs/models"
	"github.com/Confialink/wallet-logs/internal/logs/repository"
	pb "github.com/Confialink/wallet-logs/rpc/logs"
)

type protobufServer struct {
	logsRepository *repository.LogsRepository
}

func newProrobufServer() *protobufServer {
	return &protobufServer{
		repository.Logs(),
	}
}

func StartProtobufServer() {
	cfg := config.GetConfig()
	twirpHandler := pb.NewLogsServiceServer(newProrobufServer(), nil)
	mux := http.NewServeMux()
	mux.Handle(pb.LogsServicePathPrefix, twirpHandler)
	go http.ListenAndServe(fmt.Sprintf(":%s", cfg.ProtobufPort), mux)
}

func (self *protobufServer) CreateLog(ctx context.Context, req *pb.CreateLogReq) (*pb.CreateLogResp, error) {
	parsedTime, err := time.Parse(time.RFC3339, req.LogTime)
	if err != nil {
		return &pb.CreateLogResp{
			Error: &pb.Error{
				Title:   "Wrong field value",
				Details: "Time should be in RFC3339 format",
			},
		}, nil
	}

	rawJson := json.RawMessage{}
	err = rawJson.UnmarshalJSON(req.DataFields)
	if err != nil {
		return &pb.CreateLogResp{
			Error: &pb.Error{
				Title:   "Invalid format",
				Details: "Can't unmarshal bytes to json",
			},
		}, nil
	}

	log := models.Log{
		Subject:    &req.Subject,
		UserId:     req.UserId,
		LoggedAt:   &parsedTime,
		DataTitle:  &req.DataTitle,
		DataFields: rawJson,
	}

	err = self.logsRepository.Create(&log)
	if err != nil {
		return &pb.CreateLogResp{
			Error: &pb.Error{
				Title: err.Error(),
			},
		}, nil
	}
	return &pb.CreateLogResp{}, nil
}
