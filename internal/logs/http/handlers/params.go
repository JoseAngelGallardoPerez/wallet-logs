package handlers

import (
	"github.com/Confialink/wallet-logs/internal/logs/models"
	"github.com/Confialink/wallet-logs/internal/logs/services"
	"github.com/Confialink/wallet-pkg-list_params"
)

type handlerParams struct {
}

func newHandlerParams() *handlerParams {
	return new(handlerParams)
}

var showOutputFields = []interface{}{"Id", "LoggedAt", "Subject",
	"DataTitle", "DataFields", "UserId",
	map[string][]interface{}{"User": {"Username", "Email", "FirstName", "LastName"}},
}

var showListOutputFields = []interface{}{"Id", "LoggedAt", "Subject", "UserId",
	map[string][]interface{}{"User": {"Username", "Email", "FirstName", "LastName"}},
}

func (p *handlerParams) get(query string) *list_params.Includes {
	includes := list_params.NewIncludes(query)
	includes.AllowSelectFields(showOutputFields)
	allowIncludes(includes)
	addIncludes(includes)
	return includes
}

func (p *handlerParams) transactionsLogCsv(query string) *list_params.ListParams {
	params := p.list(query)
	params.Includes.AddIncludes("user")
	params.AddFilter("subject", []string{models.SubjectManualTransaction, models.SubjectRevenueDeduction}, list_params.OperatorIn)
	params.Pagination.PageSize = 0
	return params
}

func (p *handlerParams) informationLogCsv(query string) *list_params.ListParams {
	params := p.list(query)
	params.Includes.AddIncludes("user")
	params.AllowFilters([]string{
		"loggedAtFrom",
		"loggedAtTo",
		"subject",
		"subject:nin",
		list_params.FilterIn("subject"),
	})
	params.AddFilter("subject", []string{models.SubjectManualTransaction, models.SubjectRevenueDeduction}, list_params.OperatorNin)
	params.Pagination.PageSize = 0
	return params
}

func (p *handlerParams) list(query string) *list_params.ListParams {
	params := list_params.NewListParamsFromQuery(query, models.Log{})
	params.AllowSelectFields(showListOutputFields)
	params.AllowPagination()
	allowIncludes(params.Includes)
	addIncludes(params.Includes)
	allowFilters(params)
	addFilters(params)
	addSortings(params)
	return params
}

func allowIncludes(params *list_params.Includes) {
	params.Allow([]string{"user"})
}

func addIncludes(params *list_params.Includes) {
	params.AddCustomIncludes("user", services.GetLogsFiller().FillUsers)
}

func allowFilters(params *list_params.ListParams) {
	params.AllowFilters([]string{
		"loggedAtFrom",
		"loggedAtTo",
		"subject",
		list_params.FilterIn("subject"),
	})
}

func addFilters(params *list_params.ListParams) {
	params.AddCustomFilter("loggedAtFrom", list_params.DateFromFilter("logged_at"))
	params.AddCustomFilter("loggedAtTo", list_params.DateToFilter("logged_at"))
}

func addSortings(params *list_params.ListParams) {
	params.AllowSortings([]string{"loggedAt"})
	params.Sortings = append(params.Sortings, list_params.SortingListParameter{Field: "loggedAt", Direction: list_params.DescDirection})
}
