package gooseneck

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog/pkgerrors"
)

func InitLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	level := strings.ToLower(os.Getenv(LOG_LEVEL))
	var setLevel zerolog.Level

	switch level {
	case "trace", "-1":
		setLevel = zerolog.TraceLevel
	case "debug", "0":
		setLevel = zerolog.DebugLevel
	case "info", "1":
		setLevel = zerolog.InfoLevel
	case "warn", "2":
		setLevel = zerolog.WarnLevel
	case "error", "3":
		setLevel = zerolog.ErrorLevel
	case "fatal", "4":
		setLevel = zerolog.FatalLevel
	case "panic", "5":
		setLevel = zerolog.PanicLevel
	default:
		setLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(setLevel)
}

var (
	Trace = log.Trace
	Debug = log.Debug
	Info  = log.Info
	Warn  = log.Warn
	Error = log.Error
	Fatal = log.Fatal
	Panic = log.Panic
)
