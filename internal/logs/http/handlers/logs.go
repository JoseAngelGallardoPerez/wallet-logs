package handlers

import (
	"net/http"
	"strconv"

	"github.com/Confialink/wallet-pkg-errors"
	"github.com/Confialink/wallet-pkg-json_response"
	"github.com/Confialink/wallet-pkg-model_serializer"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-logs/internal/logs/config/logs"
	"github.com/Confialink/wallet-logs/internal/logs/errcodes"
	"github.com/Confialink/wallet-logs/internal/logs/repository"
)

// LogsHandler contains handler functions for logs
type LogsHandler struct {
	logsRepository *repository.LogsRepository
	params         *handlerParams
	logger         log15.Logger
}

// NewLogsHandler returns new LogsHandler
func NewLogsHandler() *LogsHandler {
	return &LogsHandler{
		repository.Logs(),
		newHandlerParams(),
		logs.Logger.New("Handler", "LogsHandler"),
	}
}

// List handler for list of logs
func (h *LogsHandler) List(c *gin.Context) {
	listParams := h.params.list(c.Request.URL.RawQuery)
	if ok, validationErrors := listParams.Validate(); !ok {
		errcodes.AddErrorMeta(c, errcodes.BadCollectionParams, validationErrors)
		return
	}

	list, err := h.logsRepository.GetList(listParams)
	if err != nil {
		privateError := errors.PrivateError{Message: "Can not get logs"}
		privateError.AddLogPair("error", err.Error())
		errors.AddErrors(c, &privateError)
		return
	}
	count, err := h.logsRepository.GetListCount(listParams)
	if err != nil {
		privateError := errors.PrivateError{Message: "Can not get logs count"}
		privateError.AddLogPair("error", err.Error())
		errors.AddErrors(c, &privateError)
		return
	}

	serialized := model_serializer.SerializeList(list, listParams.GetOutputFields())
	resp, err := json_response.NewListResponseAndPageLinks(serialized,
		c.Request.URL.RequestURI(), count, uint64(listParams.Pagination.PageNumber),
		uint64(listParams.Pagination.PageSize),
	)
	if err != nil {
		privateError := errors.PrivateError{Message: "Can not build response"}
		privateError.AddLogPair("error", err.Error())
		errors.AddErrors(c, &privateError)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Get handler for single log
func (h *LogsHandler) Get(c *gin.Context) {
	id := c.Params.ByName("id")
	id64, _ := strconv.ParseUint(id, 10, 64)

	includes := h.params.get(c.Request.URL.RawQuery)
	if ok, validationErrors := includes.Validate(); !ok {
		errcodes.AddErrorMeta(c, errcodes.BadCollectionParams, validationErrors)
		return
	}

	log, err := h.logsRepository.Get(id64, includes)
	if err != nil {
		privateError := errors.PrivateError{Message: "Can not get log"}
		privateError.AddLogPair("error", err.Error())
		errors.AddErrors(c, &privateError)
		return
	}

	serialized := model_serializer.Serialize(log, includes.GetOutputFields())
	c.JSON(http.StatusOK, json_response.NewResponse(serialized))
}
