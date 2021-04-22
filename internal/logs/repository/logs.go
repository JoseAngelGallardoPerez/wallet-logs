package repository

import (
	"github.com/Confialink/wallet-pkg-list_params"
	"github.com/Confialink/wallet-pkg-list_params/adapters"
	"github.com/jinzhu/gorm"

	"github.com/Confialink/wallet-logs/internal/logs/db"
	"github.com/Confialink/wallet-logs/internal/logs/models"
)

type LogsRepository struct {
	db *gorm.DB
}

func newLogsRepository() *LogsRepository {
	return &LogsRepository{db.GetConnection()}
}

func (self *LogsRepository) Create(log *models.Log) error {
	return self.db.Create(log).Error
}

func (self *LogsRepository) GetList(params *list_params.ListParams) (
	[]*models.Log, error,
) {
	var logs []*models.Log
	adapter := adapters.NewGorm(self.db)
	return logs, adapter.LoadList(&logs, params, "logs")
}

func (self *LogsRepository) GetListCount(params *list_params.ListParams) (
	uint64, error,
) {
	var count uint64
	str, arguments := params.GetWhereCondition()
	query := self.db.Where(str, arguments...)

	query = query.Joins(params.GetJoinCondition())

	if err := query.Model(&models.Log{}).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

func (self *LogsRepository) Get(id uint64, includes *list_params.Includes) (
	log *models.Log, err error,
) {
	log = &models.Log{}
	query := self.db

	err = query.Where("id = ?", id).First(log).Error
	if err != nil {
		return
	}

	if includes != nil {
		interfaceLogs := []interface{}{log}
		for _, customIncludesFunc := range includes.GetCustomIncludesFunctions() {
			if err := customIncludesFunc(interfaceLogs); err != nil {
				return nil, err
			}
		}
	}

	return
}
