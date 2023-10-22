package common

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type LoggerModel struct {
	Level       string `json:"level"`
	Timestamp   string `json:"timestamp"`
	IP          string `json:"ip"`
	Category    string `json:"category"`
	PID         int    `json:"pid"`
	Thread      string `json:"thread"`
	RequestID   string `json:"request_id"`
	Source      string `json:"source"`
	AccessToken string `json:"access_token"`
	ClientID    string `json:"client_id"`
	UserID      string `json:"user_id" `
	Resource    string `json:"resource"`
	Application string `json:"application"`
	Version     string `json:"version"`
	Class       string `json:"class"`
	Time        int64  `json:"time"`
	ByteIn      int    `json:"byte_in"`
	ByteOut     int    `json:"byte_out"`
	Status      int    `json:"status" `
	Code        string `json:"code"`
	Message     string `json:"message"`
}

func GenerateLogModel(version string, application string) (output LoggerModel) {
	output.IP = "-"
	output.Category = "-"
	output.PID = os.Getpid()
	output.Thread = "-"
	output.RequestID = "-"
	output.Source = "-"
	output.AccessToken = "-"
	output.Resource = "-"
	output.Application = application
	output.Version = version
	output.Class = "-"
	output.Code = "-"
	output.Message = "-"
	return output
}

func (object LoggerModel) String() string {
	b, err := json.Marshal(object)
	if err != nil {
		fmt.Println("error coy ", err)
		return ""
	}
	return string(b)
}

func (object LoggerModel) ToLoggerObject() (output []zap.Field) {
	output = append(output, zap.String("level", object.Level))
	output = append(output, zap.String("timestamp", object.Timestamp))
	output = append(output, zap.String("ip", object.IP))
	output = append(output, zap.String("category", object.Category))
	output = append(output, zap.Int("pid", object.PID))
	output = append(output, zap.String("class", object.Class))
	output = append(output, zap.String("thread", object.Thread))
	output = append(output, zap.String("request_id", object.RequestID))
	output = append(output, zap.String("source", object.Source))
	output = append(output, zap.String("access_token", object.AccessToken))
	output = append(output, zap.String("client_id", object.ClientID))
	output = append(output, zap.String("user_id", object.UserID))
	output = append(output, zap.String("resource", object.Resource))
	output = append(output, zap.String("application", object.Application))
	output = append(output, zap.String("version", object.Version))
	output = append(output, zap.Int64("time", object.Time))
	output = append(output, zap.Int("byte_in", object.ByteIn))
	output = append(output, zap.Int("byte_out", object.ByteOut))
	output = append(output, zap.Int("status", object.Status))
	output = append(output, zap.String("code", object.Code))
	output = append(output, zap.String("message", object.Message))

	return output
}

type PanicLogger struct {
	FileName     string
	FunctionName string
	Input        interface{}
	ErrorMessage string
}

func PrintLogWithLevel(level int, logModel LoggerModel) {
	logModel.Timestamp = time.Now().Format(time.RFC3339Nano)

	switch level {
	case 0:
		logModel.Level = "NOTSET"
		LoggingNexchief.LoggerDebugLevel.Info("", logModel.ToLoggerObject()...)
	case 10:
		logModel.Level = "DEBUG"
		LoggingNexchief.LoggerDebugLevel.Debug("", logModel.ToLoggerObject()...)
	case 20:
		logModel.Level = "INFO"
		LoggingNexchief.LoggerDebugLevel.Info("", logModel.ToLoggerObject()...)
	case 30:
		logModel.Level = "WARN"
		LoggingNexchief.LoggerDebugLevel.Warn("", logModel.ToLoggerObject()...)
	case 40:
		logModel.Level = "ERROR"
		LoggingNexchief.LoggerDebugLevel.Error("", logModel.ToLoggerObject()...)
	case 50:
		logModel.Level = "CRITICAL"
		LoggingNexchief.LoggerDebugLevel.Fatal("", logModel.ToLoggerObject()...)
	default:
		logModel.Level = "INFO"
		LoggingNexchief.LoggerDebugLevel.Info("", logModel.ToLoggerObject()...)
	}

}

func ConfigZapNexchief(level zapcore.Level, output []string) (logger *zap.Logger) {
	cfg := zap.Config{
		Encoding:    "json",
		OutputPaths: output,
	}
	cfg.Level = zap.NewAtomicLevelAt(level)

	logger, err := cfg.Build()
	if err != nil {
		os.Exit(9)
	}
	return
}

func SetLoggerServer(logFileName []string) {
	LoggingNexchief.LoggerDebugLevel = ConfigZapNexchief(zap.DebugLevel, logFileName)
}

type logger struct {
	LoggerDebugLevel *zap.Logger
}

var LoggingNexchief logger
