package services

import (
	"fmt"
	"time"

	"github.com/Confialink/wallet-pkg-list_params"
	"github.com/Confialink/wallet-pkg-utils/csv"
	"github.com/Confialink/wallet-pkg-utils/timefmt"

	"github.com/Confialink/wallet-logs/internal/logs/repository"
	"github.com/Confialink/wallet-logs/internal/logs/services/syssettings"
)

const (
	emptyValue = "-"
)

// Csv service for generating csv files
type Csv struct {
	repository *repository.LogsRepository
}

// NewCsv returns new Csv service
func NewCsv() *Csv {
	return &Csv{repository.Logs()}
}

// GetFile returns new csv file or error
func (s *Csv) GetFile(params *list_params.ListParams, filePrefix string) (*csv.File, error) {
	logs, err := s.repository.GetList(params)
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()
	timeSettings, err := syssettings.GetTimeSettings()
	if err != nil {
		return nil, err
	}

	file := csv.NewFile()
	formattedCurrentTime := timefmt.FormatFilenameWithTime(currentTime, timeSettings.Timezone)
	file.Name = fmt.Sprintf("%s-%s.csv", filePrefix, formattedCurrentTime)

	header := []string{"Date", "From", "Email", "FirstName", "LastName", "Subject"}
	file.WriteRow(header)

	for _, v := range logs {
		formattedCreatedAt := timefmt.Format(*v.LoggedAt, timeSettings.DateTimeFormat, timeSettings.Timezone)
		username := emptyValue
		email := emptyValue
		lastName := emptyValue
		firstName := emptyValue
		if v.User != nil {
			username = v.User.Username
			email = v.User.Email
			firstName = v.User.FirstName
			lastName = v.User.LastName
		}
		record := []string{
			formattedCreatedAt,
			username,
			email,
			firstName,
			lastName,
			*v.Subject,
		}
		file.WriteRow(record)
	}

	return file, nil
}
