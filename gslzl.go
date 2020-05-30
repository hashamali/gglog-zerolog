package gslzl

import (
	"github.com/hashamali/gsl"
	"github.com/rs/zerolog"
)

// New will create a new Zerolog.
func New(logger zerolog.Logger) *Zerolog {
	return &Zerolog{
		Logger: logger,
	}
}

// SetLevel will set the global Zerolog level.
func SetLevel(level int) {
	zerolog.SetGlobalLevel(zerolog.Level(level))
}

// Zerolog is a Zerolog implementation of gsl.Log.
type Zerolog struct {
	Logger zerolog.Logger
}

// With adds metadata to the logger's context.
func (l Zerolog) With(metadata interface{}) gsl.Log {
	return &Zerolog{
		Logger: l.Logger.With().EmbedObject(metadata.(zerolog.LogObjectMarshaler)).Logger(),
	}
}

// Info logs info.
func (l Zerolog) Info(message string) {
	l.Logger.Info().Msg(message)
}

// Infof logs info for a format.
func (l Zerolog) Infof(format string, v ...interface{}) {
	l.Logger.Info().Msgf(format, v...)
}

// Infow logs info with metadata.
func (l Zerolog) Infow(metadata interface{}, message string) {
	l.Logger.Info().EmbedObject(metadata.(zerolog.LogObjectMarshaler)).Msg(message)
}

// Infofw logs info with metadata for a format.
func (l Zerolog) Infofw(metadata interface{}, format string, v ...interface{}) {
	l.Logger.Info().EmbedObject(metadata.(zerolog.LogObjectMarshaler)).Msgf(format, v...)
}

// Error logs an error.
func (l Zerolog) Error(message string) {
	l.Logger.Error().Msg(message)
}

// Errorf logs an error for a format.
func (l Zerolog) Errorf(format string, v ...interface{}) {
	l.Logger.Error().Msgf(format, v...)
}

// Errorw logs an error with metadata.
func (l Zerolog) Errorw(metadata interface{}, message string) {
	l.Logger.Error().EmbedObject(metadata.(zerolog.LogObjectMarshaler)).Msg(message)
}

// Errorfw logs an error with metadata for a format.
func (l Zerolog) Errorfw(metadata interface{}, format string, v ...interface{}) {
	l.Logger.Error().EmbedObject(metadata.(zerolog.LogObjectMarshaler)).Msgf(format, v...)
}

// V handles verbosity.
func (l Zerolog) V(verbosity int) bool {
	switch l.Logger.GetLevel() {
	case zerolog.PanicLevel:
		return verbosity > 3
	case zerolog.FatalLevel:
		return verbosity == 3
	case zerolog.ErrorLevel:
		return verbosity == 2
	case zerolog.WarnLevel:
		return verbosity == 1
	case zerolog.InfoLevel:
		return verbosity == 0
	case zerolog.DebugLevel:
		return true
	case zerolog.TraceLevel:
		return true
	default:
		return false
	}
}
