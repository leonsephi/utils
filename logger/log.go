package logger

import (
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// log levels
const (
	LogDebug = "debug"
	LogInfo  = "info"
	LogWarn  = "warn"
	LogError = "error"
)

// log configurations
type LogConfig struct {
	Level    string // log level
	FilePath string // log file path
	MaxSize  int    // max size in MB before rotating
	MaxAge   int    // max age in days to keep backups
	Compress bool   // whether to compress rotated files
	Output   string // output type: "console" or "file"
}

func InitLogger(configs LogConfig) error {
	err := nil
	// set log level
	var level = new(slog.LevelVar)
	switch configs.Level {
	case LogDebug:
		// set to debug level
		level.Set(slog.LevelDebug)
	case LogInfo:
		// set to info level
		level.Set(slog.LevelInfo)
	case LogWarn:
		// set to warn level
		level.Set(slog.LevelWarn)
	case LogError:
		// set to error level
		level.Set(slog.LevelError)
	default:
		// set to info level
		level.Set(slog.LevelInfo)
	}
	opts := &slog.HandlerOptions{Level: level}
	// file logging
	lumberjackLogger := &lumberjack.Logger{
		Filename: configs.FilePath,
		MaxSize:  configs.MaxSize,  // megabytes
		MaxAge:   configs.MaxAge,   // days
		Compress: configs.Compress, // compress backups
	}
	if configs.Output != "file" {
		consoleHandler := slog.NewTextHandler(os.Stdout, opts)
		slog.SetDefault(slog.New(consoleHandler))
	} else {
		fileHandler := slog.NewJSONHandler(lumberjackLogger, opts)
		slog.SetDefault(slog.New(fileHandler))
	}
	return err
}
