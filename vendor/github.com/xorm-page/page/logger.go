package page

//Logger logger interface
type CLogger interface {
	Debug(v ...interface{})
	Debugf(formater string, v ...interface{})
}

var logger CLogger

type simpleLogger struct {
}

//Debug log debug level
func (sl simpleLogger) Debug(v ...interface{}) {

}

//Debugf  log debug level with formater
func (sl simpleLogger) Debugf(formater string, v ...interface{}) {

}

//Logger get logger
func Logger() CLogger {

	if logger == nil {
		logger = simpleLogger{}
	}
	return logger
}

//SetLogger set looger
func SetLogger(clogger CLogger) {
	logger = clogger
}
