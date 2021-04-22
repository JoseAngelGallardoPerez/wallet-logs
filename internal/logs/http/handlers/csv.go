package handlers

import (
	"github.com/Confialink/wallet-pkg-errors"
	"github.com/Confialink/wallet-pkg-list_params"
	"github.com/Confialink/wallet-pkg-utils/csv"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-logs/internal/logs/config/logs"
	"github.com/Confialink/wallet-logs/internal/logs/errcodes"
	"github.com/Confialink/wallet-logs/internal/logs/services"
)

//CsvLogsHandler contains handlers to donwload csvs
type CsvLogsHandler struct {
	csvService *services.Csv
	params     *handlerParams
	logger     log15.Logger
}

// NewCsvLogsHandler returns new CsvLogsHandler
func NewCsvLogsHandler() *CsvLogsHandler {
	return &CsvLogsHandler{
		services.NewCsv(),
		newHandlerParams(),
		logs.Logger.New("handler", "CsvLogsHandler"),
	}
}

// DownloadTransactions handler for downloading transactions logs csv
func (h *CsvLogsHandler) DownloadTransactions(c *gin.Context) {
	listParams := h.params.transactionsLogCsv(c.Request.URL.RawQuery)
	h.downloadHandler(c, "transactions-log", listParams)
}

// DownloadInformation handler for downloading information logs csv
func (h *CsvLogsHandler) DownloadInformation(c *gin.Context) {
	listParams := h.params.informationLogCsv(c.Request.URL.RawQuery)
	h.downloadHandler(c, "information-log", listParams)
}

func (h *CsvLogsHandler) downloadHandler(c *gin.Context, filePrefix string, listParams *list_params.ListParams) {
	if ok, paramsErrors := listParams.Validate(); !ok {
		errcodes.AddErrorMeta(c, errcodes.BadCollectionParams, paramsErrors)
		return
	}

	file, err := h.csvService.GetFile(listParams, "transactions-log")
	if err != nil {
		privateError := errors.PrivateError{Message: "Can not get csv file"}
		privateError.AddLogPair("error", err.Error())
		errors.AddErrors(c, &privateError)
		return
	}

	if err = csv.Send(file, c.Writer); err != nil {
		privateError := errors.PrivateError{Message: "Can not send csv file"}
		privateError.AddLogPair("error", err.Error())
		errors.AddErrors(c, &privateError)
	}
}
