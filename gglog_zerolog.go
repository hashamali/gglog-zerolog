package gglogzerolog

import (
	"github.com/hashamali/gglog"
	"github.com/rs/zerolog"
)

type ggLogZeroLog struct {
	Logger zerolog.Logger
}

func (l ggLogZeroLog) With(metadata map[string]interface{}) gglog.GGLog {
	return &ggLogZeroLog{
		Logger: l.Logger.With().Fields(metadata).Logger(),
	}
}

func (l ggLogZeroLog) Info(message string) {
	l.Logger.Info().Msg(message)
}

func (l ggLogZeroLog) Infof(format string, v ...interface{}) {
	l.Logger.Info().Msgf(format, v...)
}

func (l ggLogZeroLog) InfoWith(metadata map[string]interface{}, message string) {
	l.Logger.Info().Fields(metadata).Msg(message)
}

func (l ggLogZeroLog) InfofWith(metadata map[string]interface{}, format string, v ...interface{}) {
	l.Logger.Info().Fields(metadata).Msgf(format, v...)
}

func (l ggLogZeroLog) Error(message string) {
	l.Logger.Error().Msg(message)
}

func (l ggLogZeroLog) Errorf(format string, v ...interface{}) {
	l.Logger.Error().Msgf(format, v...)
}

func (l ggLogZeroLog) ErrorWith(metadata map[string]interface{}, message string) {
	l.Logger.Error().Fields(metadata).Msg(message)
}

func (l ggLogZeroLog) ErrorfWith(metadata map[string]interface{}, format string, v ...interface{}) {
	l.Logger.Error().Fields(metadata).Msgf(format, v...)
}
