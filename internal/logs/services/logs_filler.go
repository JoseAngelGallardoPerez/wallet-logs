package services

import (
	pb "github.com/Confialink/wallet-users/rpc/proto/users"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-logs/internal/logs/models"
)

type LogsFiller struct {
	logger log15.Logger
}

var exists = struct{}{}

func (self *LogsFiller) FillUsers(records []interface{}) error {
	userIds := make(map[string]struct{})
	logs := make([]*models.Log, len(records))
	for i, v := range records {
		logs[i] = v.(*models.Log)
		userIds[logs[i].UserId] = exists
	}

	users, err := getUsers(userIds)
	if err != nil {
		return err
	}

	for _, log := range logs {
		user := getUser(users, log.UserId)
		if user == nil {
			self.logger.Error("Failed to fill log by user", "userId", log.UserId)
			continue
		}

		fillLog(log, user)
	}

	return nil
}

func fillLog(log *models.Log, user *pb.User) {
	log.User = &models.LogUser{
		Username:  user.Username,
		Email:     user.Email,
		LastName:  user.LastName,
		FirstName: user.FirstName,
	}
}

func getUser(users []*pb.User, id string) *pb.User {
	for _, v := range users {
		if v.UID == id {
			return v
		}
	}
	return nil
}

func getUsers(userIds map[string]struct{}) ([]*pb.User, error) {
	var sliceIds []string
	for v := range userIds {
		sliceIds = append(sliceIds, v)
	}

	return GetRpcUsers().GetByUIDs(sliceIds)
}
