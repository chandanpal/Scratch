package logger

// import (
// 	"errors"
// 	"github.com/cisco-runner/config"
// )

// A global variable so that log functions can be directly accessed
// type Logs struct {
// 	log *Logger 
// }

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	TimestampFormat = "02/01/2006 15:04:05.999 PST"
)

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

// const (
// 	InstanceZapLogger int = iota
// 	InstanceLogrusLogger
// )

// var (
// 	errInvalidLoggerInstance = errors.New("Invalid logger instance")
// )


// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type LoggerConfiguration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}


// //Logger is our contract for the logger
// type Logger interface {
// 	Debugf(format string, args ...interface{})

// 	Infof(format string, args ...interface{})

// 	Warnf(format string, args ...interface{})

// 	Errorf(format string, args ...interface{})

// 	Fatalf(format string, args ...interface{})

// 	Panicf(format string, args ...interface{})

// 	GetLoggerWithFields(keyValues Fields) Logger
// 	GetLogger(name string) Logger

// }


// //NewLogger returns an instance of logger
// func NewLogger(config config.Config, loggerInstance int) error {
// 	switch loggerInstance {
// 	case InstanceZapLogger:
// 		logger, err := newZapLogger(config)
// 		if err != nil {
// 			return err
// 		}
// 		log = logger
// 		return nil

// 	// case InstanceLogrusLogger:
// 	// 	logger, err := newLogrusLogger(config)
// 	// 	if err != nil {
// 	// 		return err
// 	// 	}
// 	// 	log = logger
// 	// 	return nil

// 	default:
// 		return errInvalidLoggerInstance
// 	}
// }

// func(l *Logs)Debugf(format string, args ...interface{}) {
// 	l.Debugf(format, args...)
// }

// func (l *Logs)Infof(format string, args ...interface{}) {
// 	l.Infof(format, args...)
// }

// func (l *Logs)Warnf(format string, args ...interface{}) {
// 	l.Warnf(format, args...)
// }

// func (l *Logs)Errorf(format string, args ...interface{}) {
// 	l.Errorf(format, args...)
// }

// func (l *Logs)Fatalf(format string, args ...interface{}) {
// 	l.Fatalf(format, args...)
// }

// func (l *Logs)Panicf(format string, args ...interface{}) {
// 	l.Panicf(format, args...)
// }

// func (l *Logs)GetLoggerWithFields(keyValues Fields) Logger {
// 	return l.GetLoggerWithFields(keyValues)
// }

// func  (l *Logs)GetLogger(name string) Logger {
// 	return l.GetLogger(name)
// }


// func Initialize(config config.Config) {
	
// 	err := NewLogger(config, InstanceZapLogger)
// 	if err != nil {
// 		log.Fatalf("Could not instantiate log %s", err.Error())
// 	}
	
// }


// func Initialize(config config.Config) (logger Logger,  err error) {
	
// 	logger, err = newZapLogger(config)
// 		if err != nil {
// 			logger.Fatalf("Could not instantiate log %s", err.Error())
			
// 		}
// 		return
// }

