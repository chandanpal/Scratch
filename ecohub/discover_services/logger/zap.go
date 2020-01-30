package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"github.com/discover_services/config"
	"time"
)

type Logger struct {
	sugaredLogger *zap.SugaredLogger
}

var Logs *Logger

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(TimestampFormat))
}

// func getEncoder(isJSON bool) zapcore.Encoder {
// 	encoderConfig := zap.NewProductionEncoderConfig()
// 	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
// 	if isJSON {
// 		return zapcore.NewJSONEncoder(encoderConfig)
// 	}
// 	return zapcore.NewConsoleEncoder(encoderConfig)
// }


func getEncoder(isJSON bool) zapcore.Encoder {
	if isJSON {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = CustomTimeEncoder
		return zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = CustomTimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Debug:
		return zapcore.DebugLevel
	case Error:
		return zapcore.ErrorLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func newZapLogger() (*Logger, error) {
	cores := []zapcore.Core{}

	if config.CONF.Logger.EnableConsole {
		level := getZapLevel(config.CONF.Logger.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(config.CONF.Logger.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	if config.CONF.Logger.EnableFile {
		level := getZapLevel(config.CONF.Logger.FileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: config.CONF.Logger.FileLocation,
			MaxSize:  config.CONF.Logger.MaxSize,
			Compress: config.CONF.Logger.Compress,
			MaxAge:   config.CONF.Logger.MaxAge,
		})
		core := zapcore.NewCore(getEncoder(config.CONF.Logger.FileJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

  // AddCallerSkip skips 2 number of callers, this is important else the file that gets 
  // logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(1),
		zap.AddCaller(),
	).Sugar()

	return &Logger{
		sugaredLogger: logger,
	}, nil
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *Logger) GetLoggerWithFields(fields Fields) *Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &Logger{newLogger}
}

func (l *Logger) GetLogger(name string) *Logger {
	newLogger := l.sugaredLogger.Named(name)
	return &Logger{newLogger}
}




func Initialize( defaultLog bool) ( logger *Logger, err error) {
	
	logger, err = newZapLogger()
	if err != nil {
		logger.Fatalf("Could not instantiate log %s", err.Error())
		
	}
	if defaultLog {
		Logs = logger
	}
	return 
}
