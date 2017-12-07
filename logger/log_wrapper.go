package logger


type LogData struct {
	Logs []*MyLog
}


func (lw *LogData) LogItem(logType string) *MyLog {
	for _, v := range lw.Logs {
		if(v.Name == logType) {
			return v
		}
	}

	return nil
}
