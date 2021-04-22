package logs

import "github.com/inconshreveable/log15"

var Logger log15.Logger

func init() {
	Logger = log15.New("Microservice", "Logs")
}
