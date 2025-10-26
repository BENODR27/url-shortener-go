package logger

import (
    "os"
    "github.com/rs/zerolog"
)

func New(debug bool) zerolog.Logger {
    level := zerolog.InfoLevel
    if debug {
        level = zerolog.DebugLevel
    }

    return zerolog.New(os.Stdout).
        Level(level).
        With().
        Timestamp().
        Logger()
}
