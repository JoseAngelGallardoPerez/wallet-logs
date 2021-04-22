package models

import (
	"encoding/json"
	"time"
)

const (
	SubjectManualTransaction = "Manual transaction"
	SubjectRevenueDeduction  = "Revenue deduction"
)

type Log struct {
	Id         *uint64         `json:"id"`
	Subject    *string         `json:"subject"`
	UserId     string          `json:"user_id"`
	User       *LogUser        `json:"user"`
	LoggedAt   *time.Time      `json:"loggedAt"`
	DataTitle  *string         `json:"dataTitle"`
	DataFields json.RawMessage `json:"dataFields"`
}

type LogUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
