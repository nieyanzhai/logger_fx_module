package logger_fx_module

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
)

var (
	logPath = path.Join("./config", "log.json")
)

func New() *zap.SugaredLogger {
	return newLogger()
}

func newLogger() *zap.SugaredLogger {
	cfg, err := loadConfig(logPath)
	if err != nil {
		panic(err)
	}
	lumberjackLogger := newLumberjackLogger(cfg)
	return newZapLogger(lumberjackLogger)
}

func newZapLogger(w io.Writer) *zap.SugaredLogger {
	c := zap.NewProductionEncoderConfig()
	c.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(c),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(w)),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)
	return zap.New(core, zap.AddCaller()).Sugar()
}

func newLumberjackLogger(cfg Log) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   path.Join(cfg.Path, cfg.Name), // Specify the log file path
		MaxSize:    cfg.MaxSize,                   // Maximum size in megabytes of each log file
		MaxBackups: cfg.MaxBackups,                // Maximum number of log files to keep
		MaxAge:     cfg.MaxAge,                    // Maximum number of days to retain log files
		Compress:   cfg.Compress,                  // Whether to compress old log files
		LocalTime:  cfg.LocalTime,                 // Use the local time zone
	}
}

func loadConfig(path string) (Log, error) {
	var t Log
	bytes, err := os.ReadFile(path)
	if err != nil {
		return t, err
	}

	if err := json.Unmarshal(bytes, &t); err != nil {
		return t, err
	}

	return t, nil
}
