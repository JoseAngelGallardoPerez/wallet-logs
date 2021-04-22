package repository

var logs *LogsRepository

func Logs() *LogsRepository {
	if logs == nil {
		return newLogsRepository()
	}

	return logs
}
